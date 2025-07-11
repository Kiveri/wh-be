package carts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/dto"
	"github.com/Kiveri/wh-be/internal/domain/model"
	"github.com/samber/lo"
)

var errNoFiltersProvided = errors.New("at least one filter parameter must be provided")

func (r *Repo) FindAllByFilter(ctx context.Context, filter dto.FindCartFilter) ([]*model.Cart, error) {
	if filter.ID == nil && filter.ClientID == nil && filter.IsPaid == nil && filter.IsActive == nil {
		return nil, fmt.Errorf("repo.FindAllByFilter: %w", errNoFiltersProvided)
	}

	query := `SELECT id, client_id, positions_ids, total_price, is_paid, is_active FROM carts WHERE 1=1`

	var args []interface{}
	var paramCounter = 1

	if filter.ID != nil {
		query += fmt.Sprintf(" AND id = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.ID))
		paramCounter++
	}
	if filter.ClientID != nil {
		query += fmt.Sprintf(" AND client_id = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.ClientID))
		paramCounter++
	}
	if filter.IsPaid != nil {
		query += fmt.Sprintf(" AND is_paid = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.IsPaid))
		paramCounter++
	}
	if filter.IsActive != nil {
		query += fmt.Sprintf(" AND is_active = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.IsActive))
		paramCounter++
	}

	rows, err := r.cluster.Conn.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("cluster.Conn.Query: %w", err)
	}
	defer rows.Close()

	var carts []*model.Cart
	for rows.Next() {
		var cart model.Cart
		err = rows.Scan(
			&cart.ID,
			&cart.ClientID,
			&cart.PositionsIDs,
			&cart.TotalPrice,
			&cart.IsPaid,
			&cart.IsActive,
		)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		carts = append(carts, &cart)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return carts, nil
}
