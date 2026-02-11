package providers

import "github.com/bobadilla-tech/holidays-per-country/providers/internal"

type IndiaProvider struct{}

func (_ IndiaProvider) RegisterHolidays(year int) []internal.Holiday {
	return []internal.Holiday{}
}
