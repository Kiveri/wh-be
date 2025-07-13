package clients

import (
	"context"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/model/persons"
)

func (r *Repo) CreateClient(ctx context.Context, client *persons.Client) error {
	err := r.cluster.Conn.QueryRow(ctx,
		"insert into clients (id, first_name, last_name, patronymic, email, phone, home_address, company_id, is_active, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		client.ID,
		client.FirstName,
		client.LastName,
		client.Patronymic,
		client.Email,
		client.Phone,
		client.HomeAddress,
		client.CompanyID,
		client.IsActive,
		r.timer.NowMoscow(),
		r.timer.NowMoscow(),
	)
	if err != nil {
		return fmt.Errorf("cluster.Conn.QueryRow: %v", err)
	}

	return nil
}
