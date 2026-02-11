package providers

import "github.com/bobadilla-tech/holidays-per-country/providers/internal"

type UnitedKingdomProvider struct{}

func (_ UnitedKingdomProvider) RegisterHolidays(year int) []internal.Holiday {
	return []internal.Holiday{}
}
