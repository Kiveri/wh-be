package position_usecase

import (
	"context"
	"github.com/Kiveri/wh-be/internal/domain/model/internal_entities"
)

type (
	UseCase struct {
		positionRepo positionRepo
	}

	positionRepo interface {
		CreatePosition(ctx context.Context, position *internal_entities.Position) error
	}
)

func NewUseCase(positionRepo positionRepo) *UseCase {
	return &UseCase{
		positionRepo: positionRepo,
	}
}
