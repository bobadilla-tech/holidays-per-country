package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/internal"
)

type UnitedStatesProvider struct {
}

func (_ UnitedStatesProvider) RegisterHolidays(year int) []internal.Holiday {
	thirdMondayInJanuary := internal.FindDay(year, time.January, time.Monday, 3)
	thirdMondayInFebruary := internal.FindDay(year, time.February, time.Monday, 3)
	lastMondayInMay := internal.FindLastDay(year, time.May, time.Monday)
	firstMondayInSeptember := internal.FindDay(year, time.September, time.Monday, 1)
	secondMondayInOctober := internal.FindDay(year, time.October, time.Monday, 2)
	fourthThursdayInNovember := internal.FindDay(year, time.November, time.Thursday, 4)

	holidays := []internal.Holiday{
		internal.NewHoliday(year, time.January, 1, "New Year's Day", nil),
		internal.NewHolidayFromTime(thirdMondayInJanuary, "Martin Luther King, Jr. Day", nil),
		internal.NewHolidayFromTime(thirdMondayInFebruary, "Washington's Birthday", nil),
		internal.NewHolidayFromTime(lastMondayInMay, "Memorial Day", nil),
		internal.NewHoliday(year, time.July, 4, "Independence Day", nil),
		internal.NewHolidayFromTime(firstMondayInSeptember, "Labor Day", nil),
		internal.NewHolidayFromTime(fourthThursdayInNovember, "Thanksgiving Day", nil),
		internal.NewHoliday(year, time.November, 11, "Veterans Day", nil),
		internal.NewHoliday(year, time.December, 25, "Christmas Day", nil),
	}

	// State-specific holidays
	if year >= 2021 {
		holidays = append(holidays, internal.NewHoliday(year, time.June, 19, "Juneteenth National Independence Day", nil))
	}

	// Good Friday (various states)
	holidays = append(holidays, internal.NewHolidayFromTime(internal.CatholicGoodFriday(year), "Good Friday", []string{
		"US-CT", "US-DE", "US-HI", "US-IN", "US-KY", "US-LA", "US-NC", "US-ND", "US-NJ", "US-TN"}))

	// Columbus Day (various states)
	holidays = append(holidays, internal.NewHolidayFromTime(secondMondayInOctober, "Columbus Day", []string{
		"US-AL", "US-AZ", "US-CO", "US-CT", "US-GA", "US-ID", "US-IL", "US-IN", "US-IA", "US-KS",
		"US-KY", "US-LA", "US-ME", "US-MD", "US-MA", "US-MS", "US-MO", "US-MT", "US-NE", "US-NH",
		"US-NJ", "US-NM", "US-NY", "US-NC", "US-OH", "US-OK", "US-PA", "US-RI", "US-SC", "US-TN",
		"US-UT", "US-VA", "US-WV"}))
	holidays = append(holidays, internal.NewHoliday(year, time.May, 8, "Truman Day", []string{"US-MO"}))

	// Lincoln's Birthday
	holidays = append(holidays, internal.NewHoliday(year, time.February, 12, "Lincoln's Birthday", []string{
		"US-CA", "US-CT", "US-IL", "US-IN", "US-KY", "US-MI", "US-NY", "US-MO", "US-OH"}))

	if year >= 1988 {
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
		holidays = append(holidays, internal.NewHolidayFromTime(secondMondayInOctober, "Indigenous Peoples' Day", subdivisions))
	}

	return holidays
}
