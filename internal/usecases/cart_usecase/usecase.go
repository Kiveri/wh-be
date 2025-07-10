package cart_usecase

import (
	"context"
	"github.com/Kiveri/wh-be/internal/domain/model"
)

type (
	UseCase struct {
		cartRepo cartRepo
	}

	cartRepo interface {
		CreateCart(ctx context.Context, cart *model.Cart) error
	}
)

func NewUseCase(cartRepo cartRepo) *UseCase {
	return &UseCase{
		cartRepo: cartRepo,
	}
}
