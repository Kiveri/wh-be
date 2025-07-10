package posting_usecase

import (
	"context"
	"github.com/Kiveri/wh-be/internal/domain/model"
)

type (
	UseCase struct {
		postingRepo postingRepo
	}

	postingRepo interface {
		CreatePosting(ctx context.Context, posting *model.Posting) error
	}
)

func NewUseCase(postingRepo postingRepo) *UseCase {
	return &UseCase{
		postingRepo: postingRepo,
	}
}
