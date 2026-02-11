// Package providers implements holiday providers for various countries.
// This file contains the provider for France (FR).
package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/calc"
)

// FranceProvider provides public holidays for France
type FranceProvider struct{}

func (FranceProvider) RegisterHolidays(year int) []calc.Holiday {
	return baseHolidaysFR(year)
}

func baseHolidaysFR(year int) []calc.Holiday {
	return []calc.Holiday{
		calc.NewHoliday(year, time.January, 1, "New Year's Day", nil),
		calc.NewHoliday(year, time.May, 1, "Labour Day", nil),
		calc.NewHoliday(year, time.May, 8, "Victory in Europe Day", nil),
		calc.NewHoliday(year, time.July, 14, "Bastille Day", nil),
		calc.NewHoliday(year, time.August, 15, "Assumption Day", nil),
		calc.NewHoliday(year, time.November, 1, "All Saints' Day", nil),
		calc.NewHoliday(year, time.November, 11, "Armistice Day", nil),
		calc.NewHoliday(year, time.December, 25, "Christmas Day", nil),
		calc.NewHolidayFromTime(calc.CatholicEasterMonday(year), "Easter Monday", nil),
		calc.NewHolidayFromTime(calc.CatholicAscensionDay(year), "Ascension Day", nil),
		calc.NewHolidayFromTime(calc.CatholicWhitMonday(year), "Whit Monday", nil),
	}
}
