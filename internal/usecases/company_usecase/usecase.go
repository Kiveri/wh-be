package company_usecase

import (
	"context"
	"github.com/Kiveri/wh-be/internal/domain/model/infra"
)

type (
	UseCase struct {
		companyRepo companyRepo
	}

	companyRepo interface {
		CreateCompany(ctx context.Context, company *infra.Company) error
	}
)

func NewUseCase(companyRepo companyRepo) *UseCase {
	return &UseCase{
		companyRepo: companyRepo,
	}
}
