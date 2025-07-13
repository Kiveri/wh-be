package posting_usecase

import (
	"context"
	"github.com/Kiveri/wh-be/internal/domain/dto"
	"github.com/Kiveri/wh-be/internal/domain/model/internal_entities"
)

type (
	UseCase struct {
		postingRepo postingRepo
		cartRepo    cartRepo
	}

	postingRepo interface {
		CreatePosting(ctx context.Context, posting *internal_entities.Posting) error
	}
	cartRepo interface {
		FindAllByFilter(ctx context.Context, filter dto.FindCartsFilter) ([]*internal_entities.Cart, error)
	}
)

func NewUseCase(postingRepo postingRepo, cartRepo cartRepo) *UseCase {
	return &UseCase{
		postingRepo: postingRepo,
		cartRepo:    cartRepo,
	}
}
