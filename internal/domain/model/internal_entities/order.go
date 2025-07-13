package internal_entities

type (
	OrderStatus       uint8
	OrderDeliveryType uint8
)

const (
	OrderStatus_CREATED    OrderStatus = 1
	OrderStatus_BUILDING   OrderStatus = 2
	OrderStatus_BUILT      OrderStatus = 3
	OrderStatus_DELIVERING OrderStatus = 4
	OrderStatus_DELIVERED  OrderStatus = 5
	OrderStatus_COMPLETED  OrderStatus = 6

	OrderDeliveryType_UNKNOWN      OrderDeliveryType = 0
	OrderDeliveryType_SELF_PICK_UP OrderDeliveryType = 1
	OrderDeliveryType_COURIER      OrderDeliveryType = 2
	OrderDeliveryType_PVZ          OrderDeliveryType = 3
)

type Order struct {
	ID              int64
	ClientID        int64
	PostingsIDs     []int64
	Status          OrderStatus
	DeliveryType    OrderDeliveryType
	DeliveryAddress string
	IsActive        bool
}

func NewOrder(clientID int64, deliveryType OrderDeliveryType) *Order {
	return &Order{
		ClientID:     clientID,
		Status:       OrderStatus_CREATED,
		DeliveryType: deliveryType,
		IsActive:     true,
	}
}

func (o *Order) AddPostings(postings []int64) {
	o.PostingsIDs = append(o.PostingsIDs, postings...)
}

func (o *Order) ChangeStatus(status OrderStatus) {
	o.Status = status
}
