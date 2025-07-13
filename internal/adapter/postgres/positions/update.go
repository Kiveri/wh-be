package positions

import (
	"context"
	"errors"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/dto"
	"github.com/Kiveri/wh-be/internal/domain/model/internal_entities"
	"github.com/samber/lo"
	"strings"
)

var (
	errNoFilters        = errors.New("at least one filter parameter must be provided")
	errEmptyPosition    = errors.New("position is nil")
	errPositionMatching = errors.New("no positions matched the filter")
)

func (r *Repo) Update(ctx context.Context, position *internal_entities.Position, filter dto.UpdatePositionFilter) error {
	if position == nil {
		return fmt.Errorf("repo.Update: %w", errEmptyPosition)
	}

	if filter.ExternalID == nil && filter.Barcode == nil &&
		filter.Name == nil && filter.Manufacturer == nil &&
		filter.Type == nil && filter.IsHasOrder == nil && filter.OrderID == nil &&
		filter.IsActive == nil {
		return fmt.Errorf("repo.Update: %w", errNoFilters)
	}

	var sb strings.Builder
	var args []interface{}

	sb.WriteString("UPDATE positions SET updated_at = $1")
	args = append(args, r.timer.NowMoscow())
	paramCounter := 2

	if position.Name != "" {
		fmt.Fprintf(&sb, ", name = $%d", paramCounter)
		args = append(args, position.Name)
		paramCounter++
	}
	if position.Manufacturer != "" {
		fmt.Fprintf(&sb, ", manufacturer = $%d", paramCounter)
		args = append(args, position.Manufacturer)
		paramCounter++
	}
	if position.Price != 0 {
		fmt.Fprintf(&sb, ", price = $%d", paramCounter)
		args = append(args, position.Price)
		paramCounter++
	}
	if position.Type != 0 {
		fmt.Fprintf(&sb, ", type = $%d", paramCounter)
		args = append(args, position.Type)
		paramCounter++
	}
	if position.ProductionDate != nil {
		fmt.Fprintf(&sb, ", production_date = $%d", paramCounter)
		args = append(args, position.ProductionDate)
		paramCounter++
	}
	if position.ExpirationDate != nil {
		fmt.Fprintf(&sb, ", expiration_date = $%d", paramCounter)
		args = append(args, position.ExpirationDate)
		paramCounter++
	}

	fmt.Fprintf(&sb, ", is_has_order = $%d", paramCounter)
	args = append(args, position.IsHasOrder)
	paramCounter++

	if position.OrderID != nil {
		fmt.Fprintf(&sb, ", order_id = $%d", paramCounter)
		args = append(args, position.OrderID)
		paramCounter++
	}

	fmt.Fprintf(&sb, ", is_active = $%d", paramCounter)
	args = append(args, position.IsActive)
	paramCounter++

	sb.WriteString(" WHERE 1=1")

	if filter.ExternalID != nil {
		fmt.Fprintf(&sb, " AND external_id = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.ExternalID))
		paramCounter++
	}
	if filter.Barcode != nil {
		fmt.Fprintf(&sb, " AND barcode = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.Barcode))
		paramCounter++
	}
	if filter.Name != nil {
		fmt.Fprintf(&sb, " AND name = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.Name))
		paramCounter++
	}
	if filter.Manufacturer != nil {
		fmt.Fprintf(&sb, " AND manufacturer = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.Manufacturer))
		paramCounter++
	}
	if filter.Type != nil {
		fmt.Fprintf(&sb, " AND type = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.Type))
		paramCounter++
	}
	if filter.IsHasOrder != nil {
		fmt.Fprintf(&sb, " AND is_has_order = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.IsHasOrder))
		paramCounter++
	}
	if filter.OrderID != nil {
		fmt.Fprintf(&sb, " AND order_id = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.OrderID))
		paramCounter++
	}
	if filter.IsActive != nil {
		fmt.Fprintf(&sb, " AND is_active = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.IsActive))
		paramCounter++
	}

	result, err := r.cluster.Conn.Exec(ctx, sb.String(), args...)
	if err != nil {
		return fmt.Errorf("cluster.Conn.Exec: %w", err)
	}

	if rowsAffected := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("result.RowsAffected: %w", errPositionMatching)
	}

	return nil
}
