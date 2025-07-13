package internal_entities

import (
	"time"
)

type PositionType uint8

const (
	PositionType_BASIC            PositionType = 1
	PositionType_BasicConsumable  PositionType = 2
	PositionType_Liquid           PositionType = 3
	PositionType_LiquidConsumable PositionType = 4
	PositionType_OverSize         PositionType = 5
)

type Position struct {
	ID             int64
	ExternalID     int64
	Barcode        int64
	Name           string
	Manufacturer   string
	Price          int64
	Type           PositionType
	ProductionDate *time.Time
	ExpirationDate *time.Time
	IsHasOrder     bool
	OrderID        *int64
	IsActive       bool
}

func NewPosition(
	externalID, barcode int64,
	name, manufacturer string,
	price int64,
	positionType PositionType,
) *Position {
	return &Position{
		ExternalID:   externalID,
		Barcode:      barcode,
		Name:         name,
		Manufacturer: manufacturer,
		Type:         positionType,
		Price:        price,
		IsActive:     true,
	}
}

func (p *Position) IsPositionNotExpired(now, expDate time.Time) bool {
	if now.After(expDate) {
		return false
	}

	diff := expDate.Sub(now)

	if diff < 30*24*time.Hour {
		return false
	}

	return true
}
