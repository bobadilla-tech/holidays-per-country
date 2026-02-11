package providers

import "github.com/bobadilla-tech/holidays-per-country/providers/internal"

type MexicoProvider struct{}

func (_ MexicoProvider) RegisterHolidays(year int) []internal.Holiday {
	return []internal.Holiday{}
}
