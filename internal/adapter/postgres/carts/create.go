package carts

import (
	"context"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/model"
)

func (r *Repo) CreateCart(ctx context.Context, cart *model.Cart) error {
	err := r.cluster.Conn.QueryRow(ctx,
		"insert into carts (id, client_id, positions_ids, total_price, is_paid, is_active, created_at, updated_at) ",
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
