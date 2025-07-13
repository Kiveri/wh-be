package client_usecase

import (
	"context"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/model"
)

type CreateClientReq struct {
	FirstName   string
	LastName    string
	Patronymic  *string
	Email       string
	Phone       string
	HomeAddress string
	CompanyID   *int64
}

func (u *UseCase) CreateClient(ctx context.Context, req CreateClientReq) error {
	client := model.NewClient(req.FirstName, req.LastName, req.Email, req.Phone, req.HomeAddress)

	if req.Patronymic != nil {
		client.Patronymic = req.Patronymic
	}
	if req.CompanyID != nil {
		client.CompanyID = req.CompanyID
	}

	err := u.clientRepo.CreateClient(ctx, client)
	if err != nil {
		return fmt.Errorf("clientRepo.CreateClient: %w", err)
	}

	return nil
}
