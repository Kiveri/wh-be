package infra

type CompanyType uint8

const (
	CompanyType_INDIVIDUAL CompanyType = 1
	CompanyType_LEGAL      CompanyType = 2
)

type Company struct {
	ID           int64
	Name         string
	Inn          int64
	OwnersIDs    []int64
	LegalAddress string
	Type         CompanyType
	IsActive     bool
}

func NewCompany(
	name string,
	inn int64,
	ownersIDs []int64,
	legalAddress string,
	companyType CompanyType,
) *Company {
	return &Company{
		Name:         name,
		Inn:          inn,
		OwnersIDs:    ownersIDs,
		LegalAddress: legalAddress,
		Type:         companyType,
		IsActive:     true,
	}
}
