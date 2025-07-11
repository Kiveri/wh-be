package model

import "time"

type EmployeeRole uint8

const (
	EmployeeRole_UNKNOWN    EmployeeRole = 0
	EmployeeRole_BASIC      EmployeeRole = 1
	EmployeeRole_SPECIALIST EmployeeRole = 2
	EmployeeRole_ADMIN      EmployeeRole = 3
)

type Employee struct {
	ID         int64
	FirstName  string
	LastName   string
	Patronymic *string
	Email      string
	Phone      string
	Role       EmployeeRole
	IsActive   bool
	HireDate   time.Time
	FireDate   *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewEmployee(
	firstName, lastname string,
	patronymic *string,
	email, phone string,
	role EmployeeRole,
	hireDate time.Time,
) *Employee {
	return &Employee{
		FirstName:  firstName,
		LastName:   lastname,
		Patronymic: patronymic,
		Email:      email,
		Phone:      phone,
		Role:       role,
		IsActive:   true,
		HireDate:   hireDate,
	}
}
