package order_usecase

import (
	"context"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/model/internal_entities"
)

type CreateReq struct {
	ClientID        int64
	PostingsIDs     []int64
	DeliveryType    internal_entities.OrderDeliveryType
	DeliveryAddress *string
}

func (u *UseCase) Create(ctx context.Context, req CreateReq) error {
	order := internal_entities.NewOrder(req.ClientID, req.DeliveryType)
	order.ChangeStatus(internal_entities.OrderStatus_BUILDING)
	order.AddPostings(req.PostingsIDs)

	err := u.orderRepo.CreateOrder(ctx, order)
	if err != nil {
		return fmt.Errorf("orderRepo.CreateOrder: %w", err)
	}

	return nil
}
