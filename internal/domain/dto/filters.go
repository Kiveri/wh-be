package dto

import "github.com/Kiveri/wh-be/internal/domain/model"

type (
	FindCartFilter struct {
		ID       *int64
		ClientID *int64
		Status   *model.CartStatus
		IsPaid   *bool
		IsActive *bool
	}

	UpdateCartFilter struct {
		ClientID     *int64
		PositionsIDs []int64
		TotalPrice   *int64
		Status       *model.CartStatus
		IsPaid       *bool
		IsActive     *bool
	}

	FindPositionFilter struct {
		ID           *int64
		ExternalID   *int64
		Barcode      *int64
		Name         *string
		Manufacturer *string
		Type         *model.PositionType
		IsHasOrder   *bool
		IsActive     *bool
	}

	UpdatePositionFilter struct {
		ExternalID   *int64
		Barcode      *int64
		Name         *string
		Manufacturer *string
		Type         *model.PositionType
		IsHasOrder   *bool
		IsActive     *bool
	}
)
