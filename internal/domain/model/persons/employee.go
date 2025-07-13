package persons

import "time"

type EmployeeRole uint8

const (
	EmployeeRole_BASIC      EmployeeRole = 1
	EmployeeRole_SPECIALIST EmployeeRole = 2
	EmployeeRole_ADMIN      EmployeeRole = 3
)

type Employee struct {
	ID          int64
	FirstName   string
	LastName    string
	Patronymic  *string
	Email       string
	Phone       string
	HomeAddress string
	Role        EmployeeRole
	IsActive    bool
	HireDate    time.Time
	FireDate    *time.Time
}

func NewEmployee(
	firstName, lastname, email, phone, homeAddress string,
	role EmployeeRole,
	hireDate time.Time,
) *Employee {
	return &Employee{
		FirstName:   firstName,
		LastName:    lastname,
		Email:       email,
		Phone:       phone,
		HomeAddress: homeAddress,
		Role:        role,
		IsActive:    true,
		HireDate:    hireDate,
	}
}
