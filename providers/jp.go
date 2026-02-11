package providers

import "github.com/bobadilla-tech/holidays-per-country/providers/internal"

type JapanProvider struct{}

func (_ JapanProvider) RegisterHolidays(year int) []internal.Holiday {
	return []internal.Holiday{}
}
