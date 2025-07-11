package employee_usecase

import (
	"context"
	"github.com/Kiveri/wh-be/internal/domain/model"
)

type (
	UseCase struct {
		employeeRepo employeeRepo
	}

	employeeRepo interface {
		CreateEmployee(ctx context.Context, employee *model.Employee) error
	}
)

func NewUseCase(employeeRepo employeeRepo) *UseCase {
	return &UseCase{
		employeeRepo: employeeRepo,
	}
}
