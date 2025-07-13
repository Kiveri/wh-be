package employees

import (
	"context"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/model/persons"
)

func (r *Repo) CreateEmployee(ctx context.Context, employee *persons.Employee) error {
	err := r.cluster.Conn.QueryRow(ctx,
		"insert into employees (id, first_name, last_name, patronymic, email, phone, home_address, role, is_active, hire_date, fire_date, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
		employee.ID,
		employee.FirstName,
		employee.LastName,
		employee.Patronymic,
		employee.Email,
		employee.Phone,
		employee.Role,
		employee.HomeAddress,
		employee.IsActive,
		employee.HireDate,
		employee.FireDate,
		r.timer.NowMoscow(),
		r.timer.NowMoscow(),
	)
	if err != nil {
		return fmt.Errorf("cluster.Conn.QueryRow: %v", err)
	}

	return nil
}
