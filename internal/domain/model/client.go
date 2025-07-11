package model

import "time"

type Client struct {
	ID          int64
	FirstName   string
	LastName    string
	Patronymic  *string
	Email       string
	Phone       string
	HomeAddress string
	CompanyID   *int64
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewClient(
	firstName, lastName string,
	patronymic *string,
) *Client {
	return &Client{
		FirstName:  firstName,
		LastName:   lastName,
		Patronymic: patronymic,
		IsActive:   true,
	}
}
