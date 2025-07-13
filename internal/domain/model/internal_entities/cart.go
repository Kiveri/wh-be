package internal_entities

type CartStatus uint8

const (
	CartStatus_BUILDING CartStatus = 1
	CartStatus_RELEASED CartStatus = 2
	CartStatus_DELETED  CartStatus = 3
)

type Cart struct {
	ID           int64
	ClientID     int64
	PositionsIDs []int64
	TotalPrice   int64
	Status       CartStatus
	IsPaid       bool
	IsActive     bool
}

func NewCart(
	clientID int64,
) *Cart {
	return &Cart{
		ClientID: clientID,
		Status:   CartStatus_BUILDING,
		IsPaid:   false,
		IsActive: true,
	}
}

func (c *Cart) AddPosition(id int64) {
	c.PositionsIDs = append(c.PositionsIDs, id)
}

func (c *Cart) IncTotalPrice(price int64) {
	c.TotalPrice += price
}
