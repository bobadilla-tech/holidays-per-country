package providers

import "github.com/bobadilla-tech/holidays-per-country/providers/internal"

type CanadaProvider struct{}

func (_ CanadaProvider) RegisterHolidays(year int) []internal.Holiday {
	return []internal.Holiday{}
}
