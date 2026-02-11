// Package providers implements holiday providers for various countries.
// This file contains the provider for France (FR).
package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/internal"
)

// FranceProvider provides public holidays for France
type FranceProvider struct{}

func (FranceProvider) RegisterHolidays(year int) []internal.Holiday {
	return baseHolidaysFR(year)
}

func baseHolidaysFR(year int) []internal.Holiday {
	return []internal.Holiday{
		internal.NewHoliday(year, time.January, 1, "New Year's Day", nil),
		internal.NewHoliday(year, time.May, 1, "Labour Day", nil),
		internal.NewHoliday(year, time.May, 8, "Victory in Europe Day", nil),
		internal.NewHoliday(year, time.July, 14, "Bastille Day", nil),
		internal.NewHoliday(year, time.August, 15, "Assumption Day", nil),
		internal.NewHoliday(year, time.November, 1, "All Saints' Day", nil),
		internal.NewHoliday(year, time.November, 11, "Armistice Day", nil),
		internal.NewHoliday(year, time.December, 25, "Christmas Day", nil),
		internal.NewHolidayFromTime(internal.CatholicEasterMonday(year), "Easter Monday", nil),
		internal.NewHolidayFromTime(internal.CatholicAscensionDay(year), "Ascension Day", nil),
		internal.NewHolidayFromTime(internal.CatholicWhitMonday(year), "Whit Monday", nil),
	}
}
