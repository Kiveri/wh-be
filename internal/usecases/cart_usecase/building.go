package cart_usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/dto"
	"github.com/Kiveri/wh-be/internal/domain/model"
	"github.com/samber/lo"
)

var errNoAvailablePosition = errors.New("no more positions available")

type BuildingReq struct {
	ClientID             int64
	ExternalPositionsIDs []int64
}

func (u *UseCase) Building(ctx context.Context, req BuildingReq) error {
	carts, err := u.cartRepo.FindAllByFilter(ctx, dto.FindCartFilter{
		ClientID: lo.ToPtr(req.ClientID),
		Status:   lo.ToPtr(model.CartStatus_BUILDING),
		IsActive: lo.ToPtr(true),
	})
	if err != nil {
		return fmt.Errorf("cartRepo.FindAllByFilter: %w", err)
	}

	var (
		cart      *model.Cart
		updCart   *model.Cart
		positions []*model.Position
	)

	if len(carts) == 0 {
		cart = model.NewCart(req.ClientID)

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

func (u *UseCase) addPositionsToExistCart(ctx context.Context, cart *model.Cart, positions []*model.Position, req BuildingReq) (*model.Cart, error) {
	var (
		err error
		now = u.timer.NowMoscow()
	)
	for _, externalID := range req.ExternalPositionsIDs {
		positions, err = u.positionRepo.FindAllByFilter(ctx, dto.FindPositionFilter{
			ExternalID: lo.ToPtr(externalID),
			IsHasOrder: lo.ToPtr(false),
			IsActive:   lo.ToPtr(true),
		})
		if err != nil {
			return nil, fmt.Errorf("positionRepo.FindAllByFilter: %w", err)
		}
		if len(positions) == 0 {
			return nil, fmt.Errorf("%w", errNoAvailablePosition)
		}
		for _, position := range positions {
			if position.ExpirationDate != nil {
				checkNotExpired := position.IsPositionNotExpired(now, lo.FromPtr(position.ExpirationDate))
				if !checkNotExpired {
					err = u.positionRepo.Update(ctx, position, dto.FindPositionFilter{
						IsActive: lo.ToPtr(false),
					})
					if err != nil {
						return nil, fmt.Errorf("positionRepo.Update: %w", err)
					}
				} else {
					cart.AddPositions(position.ID)
					cart.IncTotalPrice(position.Price)

					err = u.positionRepo.Update(ctx, position, dto.FindPositionFilter{
						IsHasOrder: lo.ToPtr(true),
					})
					if err != nil {
						return nil, fmt.Errorf("positionRepo.Update: %w", err)
					}

					break
				}
			}
		}
	}

	return cart, nil
}
