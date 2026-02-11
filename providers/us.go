// Package providers implements holiday providers for various countries.
// This file contains the provider for the United States (US).
package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/calc"
)

// UnitedStatesProvider provides public holidays for the United States
type UnitedStatesProvider struct{}

func (UnitedStatesProvider) RegisterHolidays(year int) []calc.Holiday {
	holiday := baseHolidaysUS(year)
	if h := juneteenthUS(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := goodFridayUS(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := columbusDayUS(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := trumanDayUS(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := lincolnsBirthdayUS(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := indigenousPeoplesDayUS(year); h != nil {
		holiday = append(holiday, *h)
	}

	return holiday
}

func baseHolidaysUS(year int) []calc.Holiday {
	thirdMondayInJanuary := calc.FindDay(year, time.January, time.Monday, 3)
	thirdMondayInFebruary := calc.FindDay(year, time.February, time.Monday, 3)
	lastMondayInMay := calc.FindLastDay(year, time.May, time.Monday)
	firstMondayInSeptember := calc.FindDay(year, time.September, time.Monday, 1)
	fourthThursdayInNovember := calc.FindDay(year, time.November, time.Thursday, 4)

	return []calc.Holiday{
		{Date: time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC), Name: "New Year's Day", Subdivisions: nil, Fixed: true},
		{Date: thirdMondayInJanuary, Name: "Martin Luther King, Jr. Day", Subdivisions: nil, Fixed: false},
		{Date: thirdMondayInFebruary, Name: "Washington's Birthday", Subdivisions: nil, Fixed: false},
		{Date: lastMondayInMay, Name: "Memorial Day", Subdivisions: nil, Fixed: false},
		{Date: time.Date(year, time.July, 4, 0, 0, 0, 0, time.UTC), Name: "Independence Day", Subdivisions: nil, Fixed: true},
		{Date: firstMondayInSeptember, Name: "Labor Day", Subdivisions: nil, Fixed: false},
		{Date: fourthThursdayInNovember, Name: "Thanksgiving Day", Subdivisions: nil, Fixed: false},
		{Date: time.Date(year, time.November, 11, 0, 0, 0, 0, time.UTC), Name: "Veterans Day", Subdivisions: nil, Fixed: true},
		{Date: time.Date(year, time.December, 25, 0, 0, 0, 0, time.UTC), Name: "Christmas Day", Subdivisions: nil, Fixed: true},
	}
}

func juneteenthUS(year int) *calc.Holiday {
	if year < 2021 {
		return nil
	}
	holiday := calc.Holiday{
		Date:         time.Date(year, time.June, 19, 0, 0, 0, 0, time.UTC),
		Name:         "Juneteenth National Independence Day",
		Subdivisions: nil,
		Fixed:        true,
	}
	return &holiday
}

func goodFridayUS(year int) *calc.Holiday {
	holiday := calc.Holiday{
		Date:         calc.CatholicGoodFriday(year),
		Name:         "Good Friday",
		Subdivisions: []string{"US-CT", "US-DE", "US-HI", "US-IN", "US-KY", "US-LA", "US-NC", "US-ND", "US-NJ", "US-TN"},
		Fixed:        false,
	}
	return &holiday
}

func columbusDayUS(year int) *calc.Holiday {
	secondMondayInOctober := calc.FindDay(year, time.October, time.Monday, 2)
	holiday := calc.Holiday{
		Date: secondMondayInOctober,
		Name: "Columbus Day",
		Subdivisions: []string{
			"US-AL", "US-AZ", "US-CO", "US-CT", "US-GA", "US-ID", "US-IL", "US-IN", "US-IA", "US-KS",
			"US-KY", "US-LA", "US-ME", "US-MD", "US-MA", "US-MS", "US-MO", "US-MT", "US-NE", "US-NH",
			"US-NJ", "US-NM", "US-NY", "US-NC", "US-OH", "US-OK", "US-PA", "US-RI", "US-SC", "US-TN",
			"US-UT", "US-VA", "US-WV",
		},
		Fixed: false,
	}
	return &holiday
}

func trumanDayUS(year int) *calc.Holiday {
	holiday := calc.Holiday{
		Date:         time.Date(year, time.May, 8, 0, 0, 0, 0, time.UTC),
		Name:         "Truman Day",
		Subdivisions: []string{"US-MO"},
		Fixed:        true,
	}
	return &holiday
}

func lincolnsBirthdayUS(year int) *calc.Holiday {
	holiday := calc.Holiday{
		Date:         time.Date(year, time.February, 12, 0, 0, 0, 0, time.UTC),
		Name:         "Lincoln's Birthday",
		Subdivisions: []string{"US-CA", "US-CT", "US-IL", "US-IN", "US-KY", "US-MI", "US-NY", "US-MO", "US-OH"},
		Fixed:        true,
	}
	return &holiday
}

func indigenousPeoplesDayUS(year int) *calc.Holiday {
	if year < 1988 {
		return nil
	}

	var subdivisions []string
	switch {
	case year == 1988:
		subdivisions = []string{"US-HI"}
	case year >= 1989 && year < 2015:
		subdivisions = []string{"US-HI", "US-SD"}
	case year == 2015:
		subdivisions = []string{"US-AK", "US-HI", "US-SD"}
	case year >= 2016 && year < 2018:
		subdivisions = []string{"US-AK", "US-HI", "US-MN", "US-SD", "US-VT"}
	case year == 2018:
		subdivisions = []string{"US-AK", "US-HI", "US-IA", "US-MN", "US-NC", "US-SD", "US-VT"}
	case year == 2019:
		subdivisions = []string{"US-AK", "US-AL", "US-CA", "US-HI", "US-IA", "US-LA", "US-ME", "US-MI", "US-MN", "US-NC", "US-NM", "US-OK", "US-SD", "US-VT", "US-WI"}
	case year == 2020:
		subdivisions = []string{"US-AK", "US-AL", "US-CA", "US-HI", "US-IA", "US-LA", "US-ME", "US-MI", "US-MN", "US-NC", "US-NE", "US-NM", "US-OK", "US-SD", "US-VA", "US-VT", "US-WI"}
	case year >= 2021:
		subdivisions = []string{"US-AK", "US-AL", "US-CA", "US-HI", "US-IA", "US-LA", "US-ME", "US-MI", "US-MN", "US-NC", "US-NE", "US-NM", "US-OK", "US-OR", "US-SD", "US-TX", "US-VA", "US-VT", "US-WI"}
	}

	secondMondayInOctober := calc.FindDay(year, time.October, time.Monday, 2)
	holiday := calc.Holiday{
		Date:         secondMondayInOctober,
		Name:         "Indigenous Peoples' Day",
		Subdivisions: subdivisions,
		Fixed:        false,
	}
	return &holiday
}
