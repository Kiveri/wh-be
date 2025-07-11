package pkg

type Date struct {
	Day   int
	Month int
	Year  int
}

func NewDate(day, month, year int) *Date {
	return &Date{Day: day, Month: month, Year: year}
}
