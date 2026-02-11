// Package providers implements holiday providers for various countries.
// This file contains the provider for Japan (JP).
package providers

import (
	"math"
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/internal"
)

// JapanProvider provides public holidays for Japan
type JapanProvider struct{}

func (_ JapanProvider) RegisterHolidays(year int) []internal.Holiday {
	secondMondayInJanuary := internal.FindDay(year, time.January, time.Monday, 2)
	thirdMondayInJuly := internal.FindDay(year, time.July, time.Monday, 3)
	thirdMondayInSeptember := internal.FindDay(year, time.September, time.Monday, 3)

	holiday := []internal.Holiday{
		internal.NewHoliday(year, time.January, 1, "New Year's Day", nil),
		internal.NewHolidayFromTime(secondMondayInJanuary, "Coming of Age Day", nil),
		internal.NewHoliday(year, time.February, 11, "Foundation Day", nil),
		internal.NewHoliday(year, time.April, 29, "Showa Day", nil),
		internal.NewHoliday(year, time.May, 3, "Constitution Memorial Day", nil),
		internal.NewHoliday(year, time.May, 4, "Greenery Day", nil),
		internal.NewHoliday(year, time.May, 5, "Children's Day", nil),
		internal.NewHolidayFromTime(thirdMondayInJuly, "Marine Day", nil),
		internal.NewHoliday(year, time.August, 11, "Mountain Day", nil),
		internal.NewHolidayFromTime(thirdMondayInSeptember, "Respect for the Aged Day", nil),
		internal.NewHoliday(year, time.November, 3, "Culture Day", nil),
		internal.NewHoliday(year, time.November, 23, "Labour Thanksgiving Day", nil),
	}

	if h := emperorsBirthdayJP(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := vernalEquinoxDayJP(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := autumnalEquinoxDayJP(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := sportsDayJP(year); h != nil {
		holiday = append(holiday, *h)
	}

	return holiday
}

func emperorsBirthdayJP(year int) *internal.Holiday {
	if year < 1868 {
		return nil
	}

	if year < 1873 {
		return nil
	}

	var date time.Time
	switch {
	case year < 1912:
		date = time.Date(year, time.November, 3, 0, 0, 0, 0, time.UTC)
	case year < 1913:
		date = time.Date(year, time.August, 31, 0, 0, 0, 0, time.UTC)
	case year < 1927:
		date = time.Date(year, time.October, 31, 0, 0, 0, 0, time.UTC)
	case year < 1989:
		date = time.Date(year, time.April, 29, 0, 0, 0, 0, time.UTC)
	case year < 2019:
		date = time.Date(year, time.December, 23, 0, 0, 0, 0, time.UTC)
	case year == 2019:
		return nil
	default:
		date = time.Date(year, time.February, 23, 0, 0, 0, 0, time.UTC)
	}

	holiday := internal.NewHolidayFromTime(date, "The Emperor's Birthday", nil)
	return &holiday
}

func sportsDayJP(year int) *internal.Holiday {
	if year <= 1965 {
		return nil
	}

	if year > 1965 && year < 2000 {
		holiday := internal.NewHoliday(year, time.October, 10, "Health and Sports Day", nil)
		return &holiday
	}

	if year >= 2000 && year < 2020 {
		secondMondayInOctober := internal.FindDay(year, time.October, time.Monday, 2)
		holiday := internal.NewHolidayFromTime(secondMondayInOctober, "Health and Sports Day", nil)
		return &holiday
	}

	if year == 2020 {
		holiday := internal.NewHoliday(year, time.July, 24, "Sports Day", nil)
		return &holiday
	}

	if year == 2021 {
		holiday := internal.NewHoliday(year, time.July, 23, "Sports Day", nil)
		return &holiday
	}

	if year >= 2022 {
		secondMondayInOctober := internal.FindDay(year, time.October, time.Monday, 2)
		holiday := internal.NewHolidayFromTime(secondMondayInOctober, "Sports Day", nil)
		return &holiday
	}

	return nil
}

func vernalEquinoxDayJP(year int) *internal.Holiday {
	day, ok := vernalEquinoxDay(year)
	if !ok {
		return nil
	}
	holiday := internal.NewHoliday(year, time.March, day, "Vernal Equinox Day", nil)
	return &holiday
}

func autumnalEquinoxDayJP(year int) *internal.Holiday {
	day, ok := autumnalEquinoxDay(year)
	if !ok {
		return nil
	}
	holiday := internal.NewHoliday(year, time.September, day, "Autumnal Equinox Day", nil)
	return &holiday
}

func vernalEquinoxDay(year int) (int, bool) {
	if year < 1850 || year > 2151 {
		return 0, false
	}

	differencePerYear := 0.242194
	var equinoxDay float64
	switch {
	case year >= 1851 && year <= 1899:
		equinoxDay = math.Trunc(19.8277 + differencePerYear*float64(year-1980) - math.Trunc(float64(year-1983)/4.0))
	case year >= 1900 && year <= 1979:
		equinoxDay = math.Trunc(20.8357 + differencePerYear*float64(year-1980) - math.Trunc(float64(year-1983)/4.0))
	case year >= 1980 && year <= 2099:
		equinoxDay = math.Trunc(20.8431 + differencePerYear*float64(year-1980) - math.Trunc(float64(year-1980)/4.0))
	case year >= 2100 && year <= 2150:
		equinoxDay = math.Trunc(21.8510 + differencePerYear*float64(year-1980) - math.Trunc(float64(year-1980)/4.0))
	}

	return int(equinoxDay), true
}

func autumnalEquinoxDay(year int) (int, bool) {
	if year < 1850 || year > 2151 {
		return 0, false
	}

	differencePerYear := 0.242194
	var equinoxDay float64
	switch {
	case year >= 1851 && year <= 1899:
		equinoxDay = math.Trunc(22.2588 + differencePerYear*float64(year-1980) - math.Trunc(float64(year-1983)/4.0))
	case year >= 1900 && year <= 1979:
		equinoxDay = math.Trunc(23.2588 + differencePerYear*float64(year-1980) - math.Trunc(float64(year-1983)/4.0))
	case year >= 1980 && year <= 2099:
		equinoxDay = math.Trunc(23.2488 + differencePerYear*float64(year-1980) - math.Trunc(float64(year-1980)/4.0))
	case year >= 2100 && year <= 2150:
		equinoxDay = math.Trunc(24.2488 + differencePerYear*float64(year-1980) - math.Trunc(float64(year-1980)/4.0))
	}

	return int(equinoxDay), true
}
