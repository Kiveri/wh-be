package employee_usecase

import (
	"context"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/model"
	"github.com/Kiveri/wh-be/internal/pkg"
	"time"
)

type CreateEmployeeReq struct {
	FirstName  string
	LastName   string
	Patronymic *string
	Email      string
	Phone      string
	Role       model.EmployeeRole
	HireDate   pkg.Date
}

func (u *UseCase) CreateEmployee(ctx context.Context, req CreateEmployeeReq) error {
	hireDate := time.Date(
		req.HireDate.Year, time.Month(req.HireDate.Month), req.HireDate.Day, 0, 0, 0, 0, nil,
	)

	employee := model.NewEmployee(req.FirstName, req.LastName, req.Email, req.Phone, req.Role, hireDate)

	if req.Patronymic != nil {
		employee.Patronymic = req.Patronymic
	}

	err := u.employeeRepo.CreateEmployee(ctx, employee)
	if err != nil {
		return fmt.Errorf("employeeRepo.CreateEmployee: %w", err)
	}

	return nil
}
