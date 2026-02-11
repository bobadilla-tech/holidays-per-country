// Package providers implements holiday providers for various countries.
// This file contains the provider for Mexico (MX).
package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/common"
	"github.com/bobadilla-tech/holidays-per-country/providers/internal"
)

// MexicoProvider provides public holidays for Mexico
type MexicoProvider struct{}

func (MexicoProvider) RegisterHolidays(year int) []common.Holiday {
	firstMondayOfFebruary := internal.FindDay(year, time.February, time.Monday, 1)
	thirdMondayOfMarch := internal.FindDay(year, time.March, time.Monday, 3)
	thirdMondayOfNovember := internal.FindDay(year, time.November, time.Monday, 3)

	holiday := []common.Holiday{
		internal.NewHoliday(year, time.January, 1, "New Year's Day", nil),
		internal.NewHolidayFromTime(firstMondayOfFebruary, "Constitution Day", nil),
		internal.NewHolidayFromTime(thirdMondayOfMarch, "Benito Juarez's Birthday", nil),
		internal.NewHoliday(year, time.May, 1, "Labour Day", nil),
		internal.NewHoliday(year, time.September, 16, "Independence Day", nil),
		internal.NewHolidayFromTime(thirdMondayOfNovember, "Revolution Day", nil),
		internal.NewHoliday(year, time.December, 25, "Christmas Day", nil),
		// non-public holiday in the source data
		internal.NewHolidayFromTime(internal.CatholicMaundyThursday(year), "Maundy Thursday", nil),
		// non-public holiday in the source data
		internal.NewHolidayFromTime(internal.CatholicGoodFriday(year), "Good Friday", nil),
	}

	if h := inaugurationDayMX(year); h != nil {
		holiday = append(holiday, *h)
	}

	return holiday
}

func inaugurationDayMX(year int) *common.Holiday {
	name := "Inauguration Day"

	switch year {
	case 1934, 1940, 1946, 1952, 1958, 1964, 1970, 1976, 1982, 1988, 1994, 2000, 2006, 2012, 2018:
		holiday := internal.NewHoliday(year, time.December, 1, name, nil)
		return &holiday
	case 2024, 2030, 2036, 2042, 2048, 2054, 2060, 2066, 2072, 2078:
		holiday := internal.NewHoliday(year, time.October, 1, name, nil)
		return &holiday
	}

	return nil
}
