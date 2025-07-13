package company_usecase

import (
	"context"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/model/infra"
)

type CreateReq struct {
	Name         string
	Inn          int64
	EmployeesIDs []int64
	LegalAddress string
	Type         infra.CompanyType
}

func (u *UseCase) Create(ctx context.Context, req CreateReq) error {
	company := infra.NewCompany(req.Name, req.Inn, req.EmployeesIDs, req.LegalAddress, req.Type)
	err := u.companyRepo.CreateCompany(ctx, company)
	if err != nil {
		return fmt.Errorf("companyRepo.CreateCompany: %w", err)
	}

	return nil
}
