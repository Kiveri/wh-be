package carts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/dto"
	"github.com/Kiveri/wh-be/internal/domain/model/internal_entities"
	"strings"
)

var (
	errNoFilters    = errors.New("at least one filter parameter must be provided")
	errEmptyCart    = errors.New("cart is nil")
	errCartMatching = errors.New("no carts matched the filter")
)

func (r *Repo) UpdateCart(ctx context.Context, cart *internal_entities.Cart, filter dto.UpdateCartFilter) error {
	if cart == nil {
		return fmt.Errorf("repo.UpdateCart: %w", errEmptyCart)
	}

	if filter.ClientID == nil && len(filter.PositionsIDs) == 0 &&
		filter.TotalPrice == nil && filter.Status == nil &&
		filter.IsPaid == nil && filter.IsActive == nil {
		return fmt.Errorf("repo.UpdateCart: %w", errNoFilters)
	}

	var sb strings.Builder
	var args []interface{}

	sb.WriteString("UPDATE carts SET updated_at = $1")
	args = append(args, r.timer.NowMoscow())
	paramCounter := 2

	if cart.ClientID != 0 {
		fmt.Fprintf(&sb, ", client_id = $%d", paramCounter)
		args = append(args, cart.ClientID)
		paramCounter++
	}
	if len(cart.PositionsIDs) > 0 {
		fmt.Fprintf(&sb, ", positions_ids = $%d", paramCounter)
		args = append(args, cart.PositionsIDs)
		paramCounter++
	}
	if cart.TotalPrice != 0 {
		fmt.Fprintf(&sb, ", total_price = $%d", paramCounter)
		args = append(args, cart.TotalPrice)
		paramCounter++
	}
	if cart.Status != 0 {
		fmt.Fprintf(&sb, ", status = $%d", paramCounter)
		args = append(args, cart.Status)
		paramCounter++
	}

	fmt.Fprintf(&sb, ", is_paid = $%d", paramCounter)
	args = append(args, cart.IsPaid)
	paramCounter++

	fmt.Fprintf(&sb, ", is_active = $%d", paramCounter)
	args = append(args, cart.IsActive)
	paramCounter++

	sb.WriteString(" WHERE 1=1")

	if filter.ClientID != nil {
		fmt.Fprintf(&sb, " AND client_id = $%d", paramCounter)
		args = append(args, *filter.ClientID)
		paramCounter++
	}
	if len(filter.PositionsIDs) > 0 {
		fmt.Fprintf(&sb, " AND positions_ids @> $%d", paramCounter)
		args = append(args, filter.PositionsIDs)
		paramCounter++
	}
	if filter.TotalPrice != nil {
		fmt.Fprintf(&sb, " AND total_price = $%d", paramCounter)
		args = append(args, *filter.TotalPrice)
		paramCounter++
	}
	if filter.Status != nil {
		fmt.Fprintf(&sb, " AND status = $%d", paramCounter)
		args = append(args, *filter.Status)
		paramCounter++
	}
	if filter.IsPaid != nil {
		fmt.Fprintf(&sb, " AND is_paid = $%d", paramCounter)
		args = append(args, *filter.IsPaid)
		paramCounter++
	}
	if filter.IsActive != nil {
		fmt.Fprintf(&sb, " AND is_active = $%d", paramCounter)
		args = append(args, *filter.IsActive)
		paramCounter++
	}

	result, err := r.cluster.Conn.Exec(ctx, sb.String(), args...)
	if err != nil {
		return fmt.Errorf("cluster.Conn.Exec: %w", err)
	}

	if rowsAffected := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("result.RowsAffected: %w", errCartMatching)
	}

	return nil
}
