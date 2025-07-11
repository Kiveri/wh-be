package model

import "time"

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
	Barcode        int64
	Name           string
	Manufacturer   string
	Price          int64
	PositionType   PositionType
	ProductionDate *time.Time
	ExpirationDate *time.Time
	IsHasOrder     bool
	IsActive       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewPosition(
	barcode int64,
	name, manufacturer string,
	price int64,
	positionType PositionType,
) *Position {
	return &Position{
		Barcode:      barcode,
		Name:         name,
		Manufacturer: manufacturer,
		PositionType: positionType,
		Price:        price,
		IsHasOrder:   false,
		IsActive:     true,
	}
}
