package postings

import (
	"context"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/model/internal_entities"
)

func (r *Repo) CreatePosting(ctx context.Context, posting *internal_entities.Posting) error {
	err := r.cluster.Conn.QueryRow(ctx,
		"insert into postings (id, cart_id, positions_ids, status, is_active, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7)",
		posting.ID,
		posting.CartID,
		posting.PositionsIDs,
		posting.Status,
		posting.IsActive,
		r.timer.NowMoscow(),
		r.timer.NowMoscow(),
	)
	if err != nil {
		return fmt.Errorf("cluster.Conn.QueryRow: %v", err)
	}

	return nil
}
