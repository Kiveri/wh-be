package positions

import (
	"context"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/model"
)

func (r *Repo) CreatePosition(ctx context.Context, position *model.Position) error {
	err := r.cluster.Conn.QueryRow(ctx,
		"insert into positions (id, barcode, name, manufacturer, price, type, production_date, expiration_date, is_has_order, is_active, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
		position.ID,
		position.Barcode,
		position.Name,
		position.Manufacturer,
		position.Price,
		position.Type,
		position.ProductionDate,
		position.ExpirationDate,
		position.IsHasOrder,
		position.IsActive,
		r.timer.NowMoscow(),
		r.timer.NowMoscow(),
	)
	if err != nil {
		return fmt.Errorf("cluster.Conn.QueryRow: %v", err)
	}

	return nil
}
