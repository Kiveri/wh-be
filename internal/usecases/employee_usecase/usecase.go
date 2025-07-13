package employee_usecase

import (
	"context"
	"github.com/Kiveri/wh-be/internal/domain/model/persons"
)

type (
	UseCase struct {
		employeeRepo employeeRepo
	}

	employeeRepo interface {
		CreateEmployee(ctx context.Context, employee *persons.Employee) error
	}
)

func NewUseCase(employeeRepo employeeRepo) *UseCase {
	return &UseCase{
		employeeRepo: employeeRepo,
	}
}
