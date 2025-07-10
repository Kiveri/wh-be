package position_usecase

import (
	"context"
	"github.com/Kiveri/wh-be/internal/domain/model"
)

type (
	UseCase struct {
		positionRepo positionRepo
	}

	positionRepo interface {
		CreatePosition(ctx context.Context, position *model.Position) error
	}
)

func NewUseCase(positionRepo positionRepo) *UseCase {
	return &UseCase{
		positionRepo: positionRepo,
	}
}
