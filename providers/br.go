// Package providers implements holiday providers for various countries.
// This file contains the provider for Brazil (BR).
package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/calc"
)

// BrazilProvider provides public holidays for Brazil
type BrazilProvider struct{}

func (BrazilProvider) RegisterHolidays(year int) []calc.Holiday {
	holiday := baseHolidaysBR(year)
	if h := blackAwarenessDayBR(year); h != nil {
		holiday = append(holiday, *h)
	}

	return holiday
}

func baseHolidaysBR(year int) []calc.Holiday {
	easterSunday := calc.CatholicEasterSunday(year)

	return []calc.Holiday{
		calc.NewHoliday(year, time.January, 1, "New Year's Day", nil),
		calc.NewHoliday(year, time.April, 21, "Tiradentes", nil),
		calc.NewHoliday(year, time.May, 1, "Labour Day", nil),
		calc.NewHoliday(year, time.July, 9, "Constitutionalist Revolution of 1932", []string{"BR-SP"}),
		calc.NewHoliday(year, time.September, 7, "Independence Day", nil),
		calc.NewHoliday(year, time.October, 12, "Our Lady of Aparecida", nil),
		calc.NewHoliday(year, time.November, 2, "All Souls' Day", nil),
		calc.NewHoliday(year, time.November, 15, "Republic Proclamation Day", nil),
		calc.NewHoliday(year, time.December, 25, "Christmas Day", nil),
		calc.NewHolidayFromTime(easterSunday.AddDate(0, 0, -47), "Carnival", nil),
		calc.NewHolidayFromTime(easterSunday.AddDate(0, 0, -48), "Carnival", nil),
		calc.NewHolidayFromTime(easterSunday, "Easter Sunday", nil),
		calc.NewHolidayFromTime(calc.CatholicGoodFriday(year), "Good Friday", nil),
		calc.NewHolidayFromTime(calc.CatholicCorpusChristi(year), "Corpus Christi", nil),
	}
}

func blackAwarenessDayBR(year int) *calc.Holiday {
	if year < 2024 {
		return nil
	}
	holiday := calc.NewHoliday(year, time.November, 20, "Black Awareness Day", nil)
	return &holiday
}
