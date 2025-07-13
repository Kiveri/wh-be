package orders

import (
	"context"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/model/internal_entities"
)

func (r *Repo) CreateOrder(ctx context.Context, order *internal_entities.Order) error {
	err := r.cluster.Conn.QueryRow(ctx,
		"insert into orders (id, postings_ids, order_status, order_delivery_type, is_active, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7)",
		order.ID,
		order.PostingsIDs,
		order.OrderStatus,
		order.OrderDeliveryType,
		order.IsActive,
		r.timer.NowMoscow(),
		r.timer.NowMoscow(),
	)
	if err != nil {
		return fmt.Errorf("cluster.Conn.QueryRow: %v", err)
	}

	return nil
}
