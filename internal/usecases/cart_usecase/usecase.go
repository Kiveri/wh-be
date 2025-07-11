package cart_usecase

import (
	"context"
	"github.com/Kiveri/wh-be/internal/domain/dto"
	"github.com/Kiveri/wh-be/internal/domain/model"
	"time"
)

type (
	UseCase struct {
		cartRepo     cartRepo
		positionRepo positionRepo
		timer        timer
	}

	cartRepo interface {
		CreateCart(ctx context.Context, cart *model.Cart) error
		FindAllByFilter(ctx context.Context, filter dto.FindCartFilter) ([]*model.Cart, error)
		UpdateCart(ctx context.Context, cart *model.Cart, filter dto.UpdateCartFilter) error
	}
	positionRepo interface {
		FindAllByFilter(ctx context.Context, filter dto.FindPositionFilter) ([]*model.Position, error)
		Update(ctx context.Context, position *model.Position, filter dto.FindPositionFilter) error
	}
	timer interface {
		NowMoscow() time.Time
	}
)

func NewUseCase(cartRepo cartRepo, positionRepo positionRepo, timer timer) *UseCase {
	return &UseCase{
		cartRepo:     cartRepo,
		positionRepo: positionRepo,
		timer:        timer,
	}
}
