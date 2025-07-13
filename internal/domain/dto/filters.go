package dto

import (
	"github.com/Kiveri/wh-be/internal/domain/model/internal_entities"
)

type (
	FindCartsFilter struct {
		ID       *int64
		ClientID *int64
		Status   *internal_entities.CartStatus
		IsPaid   *bool
		IsActive *bool
	}

	UpdateCartFilter struct {
		ClientID     *int64
		PositionsIDs []int64
		TotalPrice   *int64
		Status       *internal_entities.CartStatus
		IsPaid       *bool
		IsActive     *bool
	}

	FindPositionsFilter struct {
		ID           *int64
		ExternalID   *int64
		Barcode      *int64
		Name         *string
		Manufacturer *string
		Type         *internal_entities.PositionType
		IsHasOrder   *bool
		OrderID      *int64
		IsActive     *bool
	}

	UpdatePositionFilter struct {
		ExternalID   *int64
		Barcode      *int64
		Name         *string
		Manufacturer *string
		Type         *internal_entities.PositionType
		IsHasOrder   *bool
		OrderID      *int64
		IsActive     *bool
	}

	FindPostingsFilter struct {
		ID       *int64
		CartID   *int64
		Status   *internal_entities.PostingStatus
		IsActive *bool
	}
)
