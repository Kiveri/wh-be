package postings

import (
	"context"
	"errors"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/dto"
	"github.com/Kiveri/wh-be/internal/domain/model/internal_entities"
	"github.com/samber/lo"
	"strings"
)

var errNoFiltersProvided = errors.New("at least one filter parameter must be provided")

func (r *Repo) FindAllByFilter(ctx context.Context, filter dto.FindPostingsFilter) ([]*internal_entities.Posting, error) {
	if filter.ID == nil && filter.CartID == nil &&
		filter.Status == nil && filter.IsActive == nil {
		return nil, fmt.Errorf("repo.FindAllByFilter: %w", errNoFiltersProvided)
	}

	var sb strings.Builder
	var args []interface{}
	paramCounter := 1

	sb.WriteString(`
        SELECT id, cart_id, positions_ids, status, is_active 
        FROM postings 
        WHERE 1=1
    `)

	if filter.ID != nil {
		fmt.Fprintf(&sb, " AND id = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.ID))
		paramCounter++
	}
	if filter.CartID != nil {
		fmt.Fprintf(&sb, " AND cart_id = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.CartID))
		paramCounter++
	}
	if filter.Status != nil {
		fmt.Fprintf(&sb, " AND status = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.Status))
		paramCounter++
	}
	if filter.IsActive != nil {
		fmt.Fprintf(&sb, " AND is_active = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.IsActive))
		paramCounter++
	}

	rows, err := r.cluster.Conn.Query(ctx, sb.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("cluster.Conn.Query: %w", err)
	}
	defer rows.Close()

	var postings []*internal_entities.Posting
	for rows.Next() {
		var posting internal_entities.Posting
		err = rows.Scan(
			&posting.ID,
			&posting.CartID,
			&posting.PositionsIDs,
			&posting.Status,
			&posting.IsActive,
		)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		postings = append(postings, &posting)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return postings, nil
}
