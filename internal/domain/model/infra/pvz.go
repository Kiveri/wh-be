package infra

import "time"

type Pvz struct {
	ID               int64
	OwnerFirstName   string
	OwnerLastName    string
	OwnerPatronymic  *string
	OwnerEmail       string
	OwnerPhone       string
	OwnerHomeAddress string
	Address          string
	OpeningDate      time.Time
	IsActive         bool
	OrdersInPvzIDs   []int64
}

func NewPvz(
	ownerFirstName string,
	ownerLastName string,
	ownerEmail string,
	ownerPhone string,
	ownerHomeAddress string,
	address string,
	openingDate time.Time) *Pvz {
	return &Pvz{
		OwnerFirstName:   ownerFirstName,
		OwnerLastName:    ownerLastName,
		OwnerEmail:       ownerEmail,
		OwnerPhone:       ownerPhone,
		OwnerHomeAddress: ownerHomeAddress,
		Address:          address,
		OpeningDate:      openingDate,
		IsActive:         true,
	}
}
