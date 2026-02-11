// Package providers implements holiday providers for various countries.
// This file contains the provider for Mexico (MX).
package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/calc"
)

// MexicoProvider provides public holidays for Mexico
type MexicoProvider struct{}

func (MexicoProvider) RegisterHolidays(year int) []calc.Holiday {
	firstMondayOfFebruary := calc.FindDay(year, time.February, time.Monday, 1)
	thirdMondayOfMarch := calc.FindDay(year, time.March, time.Monday, 3)
	thirdMondayOfNovember := calc.FindDay(year, time.November, time.Monday, 3)

	holiday := []calc.Holiday{
		calc.NewHoliday(year, time.January, 1, "New Year's Day", nil),
		calc.NewHolidayFromTime(firstMondayOfFebruary, "Constitution Day", nil),
		calc.NewHolidayFromTime(thirdMondayOfMarch, "Benito Juarez's Birthday", nil),
		calc.NewHoliday(year, time.May, 1, "Labour Day", nil),
		calc.NewHoliday(year, time.September, 16, "Independence Day", nil),
		calc.NewHolidayFromTime(thirdMondayOfNovember, "Revolution Day", nil),
		calc.NewHoliday(year, time.December, 25, "Christmas Day", nil),
		// non-public holiday in the source data
		calc.NewHolidayFromTime(calc.CatholicMaundyThursday(year), "Maundy Thursday", nil),
		// non-public holiday in the source data
		calc.NewHolidayFromTime(calc.CatholicGoodFriday(year), "Good Friday", nil),
	}

	if h := inaugurationDayMX(year); h != nil {
		holiday = append(holiday, *h)
	}

	return holiday
}

func inaugurationDayMX(year int) *calc.Holiday {
	name := "Inauguration Day"

	switch year {
	case 1934, 1940, 1946, 1952, 1958, 1964, 1970, 1976, 1982, 1988, 1994, 2000, 2006, 2012, 2018:
		holiday := calc.NewHoliday(year, time.December, 1, name, nil)
		return &holiday
	case 2024, 2030, 2036, 2042, 2048, 2054, 2060, 2066, 2072, 2078:
		holiday := calc.NewHoliday(year, time.October, 1, name, nil)
		return &holiday
	}

	return nil
}
