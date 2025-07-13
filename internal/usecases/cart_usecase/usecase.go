package cart_usecase

import (
	"context"
	"github.com/Kiveri/wh-be/internal/domain/dto"
	"github.com/Kiveri/wh-be/internal/domain/model/internal_entities"
	"time"
)

type (
	UseCase struct {
		cartRepo     cartRepo
		positionRepo positionRepo
		timer        timer
	}

	cartRepo interface {
		CreateCart(ctx context.Context, cart *internal_entities.Cart) error
		FindAllByFilter(ctx context.Context, filter dto.FindCartsFilter) ([]*internal_entities.Cart, error)
		UpdateCart(ctx context.Context, cart *internal_entities.Cart, filter dto.UpdateCartFilter) error
	}
	positionRepo interface {
		FindAllByFilter(ctx context.Context, filter dto.FindPositionsFilter) ([]*internal_entities.Position, error)
		Update(ctx context.Context, position *internal_entities.Position, filter dto.UpdatePositionFilter) error
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
