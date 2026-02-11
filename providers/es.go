package providers

import "github.com/bobadilla-tech/holidays-per-country/providers/internal"

type SpainProvider struct{}

func (_ SpainProvider) RegisterHolidays(year int) []internal.Holiday {
	return []internal.Holiday{}
}
