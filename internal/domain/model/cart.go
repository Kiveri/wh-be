package model

import "time"

type Cart struct {
	ID           int64
	ClientID     int64
	PositionsIDs []int64
	TotalPrice   int64
	IsPaid       bool
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewCart(
	clientID int64,
	positionsIDs []int64,
	totalPrice int64,
) *Cart {
	return &Cart{
		ClientID:     clientID,
		PositionsIDs: positionsIDs,
		TotalPrice:   totalPrice,
		IsPaid:       false,
		IsActive:     true,
	}
}
