package model

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
}

func NewClient(
	firstName, lastName, email, phone, homeAddress string,
) *Client {
	return &Client{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		Phone:       phone,
		HomeAddress: homeAddress,
		IsActive:    true,
	}
}
