package carts

import (
	"context"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/model/internal_entities"
)

func (r *Repo) CreateCart(ctx context.Context, cart *internal_entities.Cart) error {
	err := r.cluster.Conn.QueryRow(ctx,
		"insert into carts (id, client_id, positions_ids, total_price, is_paid, is_active, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8)",
		cart.ID,
		cart.ClientID,
		cart.PositionsIDs,
		cart.TotalPrice,
		cart.IsPaid,
		cart.IsActive,
		r.timer.NowMoscow(),
		r.timer.NowMoscow(),
	)
	if err != nil {
		return fmt.Errorf("cluster.Conn.QueryRow: %v", err)
	}

	return nil
}
