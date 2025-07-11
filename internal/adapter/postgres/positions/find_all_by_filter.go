package positions

import (
	"context"
	"errors"
	"fmt"
	"github.com/Kiveri/wh-be/internal/domain/dto"
	"github.com/Kiveri/wh-be/internal/domain/model"
	"github.com/samber/lo"
)

var errNoFiltersProvided = errors.New("at least one filter parameter must be provided")

func (r *Repo) FindAllByFilter(ctx context.Context, filter dto.FindPositionFilter) ([]*model.Position, error) {
	if filter.ID == nil && filter.ExternalID == nil && filter.Barcode == nil &&
		filter.Name == nil && filter.Manufacturer == nil &&
		filter.Type == nil && filter.IsHasOrder == nil &&
		filter.IsActive == nil {
		return nil, fmt.Errorf("repo.findAllByFilter: %w", errNoFiltersProvided)
	}

	query := `SELECT id, external_id, barcode, name, manufacturer, price, type, production_date, expiration_date, is_has_order, is_active FROM positions WHERE 1=1`

	var args []interface{}
	var paramCounter = 1

	if filter.ID != nil {
		query += fmt.Sprintf(" AND id = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.ID))
		paramCounter++
	}
	if filter.ExternalID != nil {
		query += fmt.Sprintf(" AND external_id = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.ExternalID))
		paramCounter++
	}
	if filter.Barcode != nil {
		query += fmt.Sprintf(" AND barcode = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.Barcode))
		paramCounter++
	}
	if filter.Name != nil {
		query += fmt.Sprintf(" AND name = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.Name))
		paramCounter++
	}
	if filter.Manufacturer != nil {
		query += fmt.Sprintf(" AND manufacturer = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.Manufacturer))
		paramCounter++
	}
	if filter.Type != nil {
		query += fmt.Sprintf(" AND type = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.Type))
		paramCounter++
	}
	if filter.IsHasOrder != nil {
		query += fmt.Sprintf(" AND is_has_order = $%d", paramCounter)
		args = append(args, lo.ToPtr(filter.IsHasOrder))
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

	var positions []*model.Position
	for rows.Next() {
		var position model.Position
		err = rows.Scan(
			&position.ID,
			&position.ExternalID,
			&position.Barcode,
			&position.Name,
			&position.Manufacturer,
			&position.Price,
			&position.Type,
			&position.ProductionDate,
			&position.ExpirationDate,
			&position.IsHasOrder,
			&position.IsActive,
		)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		positions = append(positions, &position)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return positions, nil
}
