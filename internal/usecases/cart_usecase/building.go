package cart_usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/dto"
	"github.com/Kiveri/wh-be/internal/domain/model/internal_entities"
	"github.com/samber/lo"
)

var (
	errNoAvailablePosition = errors.New("no more positions available")
)

type BuildingReq struct {
	ClientID             int64
	ExternalPositionsIDs []int64
}

func (u *UseCase) Building(ctx context.Context, req BuildingReq) error {
	carts, err := u.cartRepo.FindAllByFilter(ctx, dto.FindCartsFilter{
		ClientID: lo.ToPtr(req.ClientID),
		Status:   lo.ToPtr(internal_entities.CartStatus_BUILDING),
		IsActive: lo.ToPtr(true),
	})
	if err != nil {
		return fmt.Errorf("cartRepo.FindAllByFilter: %w", err)
	}

	var (
		cart      *internal_entities.Cart
		updCart   *internal_entities.Cart
		positions []*internal_entities.Position
	)

	if len(carts) == 0 {
		cart = internal_entities.NewCart(req.ClientID)

		updCart, err = u.addPositionsToExistCart(ctx, cart, positions, req)
		if err != nil {
			return fmt.Errorf("u.addPositionsToExistCart: %w", err)
		}
		err = u.cartRepo.CreateCart(ctx, updCart)
		if err != nil {
			return fmt.Errorf("cartRepo.CreateCart: %w", err)
		}
	} else {
		updCart, err = u.addPositionsToExistCart(ctx, cart, positions, req)
		if err != nil {
			return fmt.Errorf("u.addPositionsToExistCart: %w", err)
		}
		err = u.cartRepo.UpdateCart(ctx, cart, dto.UpdateCartFilter{
			PositionsIDs: updCart.PositionsIDs,
			TotalPrice:   lo.ToPtr(updCart.TotalPrice),
		})
	}

	return nil
}

func (u *UseCase) addPositionsToExistCart(ctx context.Context, cart *internal_entities.Cart, positions []*internal_entities.Position, req BuildingReq) (*internal_entities.Cart, error) {
	var (
		err                            error
		now                            = u.timer.NowMoscow()
		allExpiredByExternalPositionID = true
	)

	for _, externalID := range req.ExternalPositionsIDs {
		positions, err = u.positionRepo.FindAllByFilter(ctx, dto.FindPositionsFilter{
			ExternalID: lo.ToPtr(externalID),
			IsHasOrder: lo.ToPtr(false),
			IsActive:   lo.ToPtr(true),
		})
		if err != nil {
			return nil, fmt.Errorf("positionRepo.FindAllByFilter: %w", err)
		}
		if len(positions) == 0 {
			return nil, fmt.Errorf("%w with externalID %d", errNoAvailablePosition, externalID)
		}
		for _, position := range positions {
			if position.ExpirationDate != nil {
				checkNotExpired := position.IsPositionNotExpired(now, lo.FromPtr(position.ExpirationDate))
				if !checkNotExpired {
					err = u.positionRepo.Update(ctx, position, dto.UpdatePositionFilter{
						IsActive: lo.ToPtr(false),
					})
					if err != nil {
						return nil, fmt.Errorf("positionRepo.Update: %w", err)
					}
				} else {
					cart.AddPosition(position.ID)
					cart.IncTotalPrice(position.Price)
					allExpiredByExternalPositionID = false

					err = u.positionRepo.Update(ctx, position, dto.UpdatePositionFilter{
						IsHasOrder: lo.ToPtr(true),
					})
					if err != nil {
						return nil, fmt.Errorf("positionRepo.Update: %w", err)
					}

					break
				}

				if allExpiredByExternalPositionID {
					return nil, fmt.Errorf("%w with externalID %d", errNoAvailablePosition, externalID)
				}
			} else {
				cart.AddPosition(position.ID)
				cart.IncTotalPrice(position.Price)

				err = u.positionRepo.Update(ctx, position, dto.UpdatePositionFilter{
					IsHasOrder: lo.ToPtr(true),
				})
				if err != nil {
					return nil, fmt.Errorf("positionRepo.Update: %w", err)
				}

				break
			}
		}
	}

	return cart, nil
}
