package model

import "time"

type Company struct {
	ID           int64
	Name         string
	Inn          int64
	EmployeesIDs []int64
	LegalAddress string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewCompany(
	name string,
	inn int64,
	employeesIDs []int64,
) *Company {
	return &Company{
		Name:         name,
		Inn:          inn,
		EmployeesIDs: employeesIDs,
		IsActive:     true,
	}
}
