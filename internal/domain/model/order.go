package model

import "time"

type (
	OrderStatus       uint8
	OrderDeliveryType uint8
)

const (
	OrderStatus_UNKNOWN    OrderStatus = 0
	OrderStatus_CREATED    OrderStatus = 1
	OrderStatus_BUILDING   OrderStatus = 2
	OrderStatus_BUILT      OrderStatus = 3
	OrderStatus_DELIVERING OrderStatus = 4
	OrderStatus_DELIVERED  OrderStatus = 5

	OrderDeliveryType_UNKNOWN      OrderDeliveryType = 0
	OrderDeliveryType_SELF_PICK_UP OrderDeliveryType = 1
	OrderDeliveryType_COURIER      OrderDeliveryType = 2
	OrderDeliveryType_PVZ          OrderDeliveryType = 3
)

type Order struct {
	ID                int64
	PostingsIDs       []int64
	OrderStatus       OrderStatus
	OrderDeliveryType OrderDeliveryType
	IsActive          bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func NewOrder(postingsIDs []int64, status OrderStatus, deliveryType OrderDeliveryType) *Order {
	return &Order{
		PostingsIDs:       postingsIDs,
		OrderStatus:       status,
		OrderDeliveryType: deliveryType,
		IsActive:          true,
	}
}
