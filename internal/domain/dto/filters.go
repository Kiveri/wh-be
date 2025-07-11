package dto

type CartFilter struct {
	ID           *int64
	ClientID     *int64
	PositionsIDs []*int64
	TotalPrice   *int64
	IsPaid       *bool
	IsActive     *bool
}
