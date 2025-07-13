package companies

import (
	"context"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/model/infra"
)

func (r *Repo) CreateCompany(ctx context.Context, company *infra.Company) error {
	err := r.cluster.Conn.QueryRow(ctx,
		"insert into companies (id, name, inn, owners_ids, legal_address, is_active, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8)",
		company.ID,
		company.Name,
		company.Inn,
		company.OwnersIDs,
		company.LegalAddress,
		company.IsActive,
		r.timer.NowMoscow(),
		r.timer.NowMoscow(),
	)
	if err != nil {
		return fmt.Errorf("cluster.Conn.QueryRow: %v", err)
	}

	return nil
}
