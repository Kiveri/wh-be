package client_usecase

import (
	"context"
	"github.com/Kiveri/wh-be/internal/domain/model"
)

type (
	UseCase struct {
		clientRepo clientRepo
	}

	clientRepo interface {
		CreateClient(ctx context.Context, client *model.Client) error
	}
)

func NewUseCase(clientRepo clientRepo) *UseCase {
	return &UseCase{
		clientRepo: clientRepo,
	}
}
