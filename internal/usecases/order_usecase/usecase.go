package order_usecase

import (
	"context"
	"github.com/Kiveri/wh-be/internal/domain/dto"
	"github.com/Kiveri/wh-be/internal/domain/model/internal_entities"
)

type (
	UseCase struct {
		orderRepo    orderRepo
		positionRepo positionRepo
		postingRepo  postingRepo
	}

	orderRepo interface {
		CreateOrder(ctx context.Context, order *internal_entities.Order) error
	}
	positionRepo interface {
		FindAllByFilter(ctx context.Context, filter dto.FindPositionsFilter) ([]*internal_entities.Position, error)
		Update(ctx context.Context, position *internal_entities.Position, filter dto.UpdatePositionFilter) error
	}
	postingRepo interface {
		FindAllByFilter(ctx context.Context, filter dto.FindPostingsFilter) ([]*internal_entities.Posting, error)
	}
)

func NewUseCase(orderRepo orderRepo, positionRepo positionRepo, postingRepo postingRepo) *UseCase {
	return &UseCase{
		orderRepo:    orderRepo,
		positionRepo: positionRepo,
		postingRepo:  postingRepo,
	}
}
