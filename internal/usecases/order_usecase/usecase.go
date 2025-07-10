package order_usecase

import (
	"context"
	"github.com/Kiveri/wh-be/internal/domain/model"
)

type (
	UseCase struct {
		orderRepo orderRepo
	}

	orderRepo interface {
		CreateOrder(ctx context.Context, order *model.Order) error
	}
)

func NewUseCase(orderRepo orderRepo) *UseCase {
	return &UseCase{
		orderRepo: orderRepo,
	}
}
