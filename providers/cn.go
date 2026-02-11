// Package providers implements holiday providers for various countries.
// This file contains the provider for China (CN).
package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/calc"
)

// ChinaProvider provides public holidays for China
type ChinaProvider struct{}

func (ChinaProvider) RegisterHolidays(year int) []calc.Holiday {
	holiday := []calc.Holiday{
		calc.NewHoliday(year, time.January, 1, "New Year's Day", nil),
		calc.NewHoliday(year, time.May, 1, "Labour Day", nil),
		calc.NewHoliday(year, time.October, 1, "National Day", nil),
	}

	if year > 1901 && year < 2100 {
		if springFestival, ok := calc.ChineseLunisolarToGregorian(year, 1, 1); ok {
			holiday = append(holiday, calc.NewHolidayFromTime(springFestival, "Chinese New Year (Spring Festival)", nil))
		}
		if dragonBoatFestival, ok := calc.ChineseLunisolarToGregorian(year, 5, 5); ok {
			holiday = append(holiday, calc.NewHolidayFromTime(dragonBoatFestival, "Dragon Boat Festival", nil))
		}
		if midAutumnFestival, ok := calc.ChineseLunisolarToGregorian(year, 8, 15); ok {
			holiday = append(holiday, calc.NewHolidayFromTime(midAutumnFestival, "Mid-Autumn Festival", nil))
		}
	}

	return holiday
}
