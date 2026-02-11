package providers

import "github.com/bobadilla-tech/holidays-per-country/providers/internal"

type ChinaProvider struct{}

func (_ ChinaProvider) RegisterHolidays(year int) []internal.Holiday {
	return []internal.Holiday{}
}
