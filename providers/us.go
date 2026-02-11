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
		internal.NewHoliday(year, time.January, 1, "New Year's Day"),
		internal.NewHolidayFromTime(thirdMondayInJanuary, "Martin Luther King, Jr. Day"),
		internal.NewHolidayFromTime(thirdMondayInFebruary, "Washington's Birthday"),
		internal.NewHolidayFromTime(lastMondayInMay, "Memorial Day"),
		internal.NewHoliday(year, time.July, 4, "Independence Day"),
		internal.NewHolidayFromTime(firstMondayInSeptember, "Labor Day"),
		internal.NewHolidayFromTime(fourthThursdayInNovember, "Thanksgiving Day"),
		internal.NewHoliday(year, time.November, 11, "Veterans Day"),
		internal.NewHoliday(year, time.December, 25, "Christmas Day"),
	}

	// State-specific holidays
	if year >= 2021 {
		holidays = append(holidays, internal.NewHoliday(year, time.June, 19, "Juneteenth National Independence Day"))
	}

	// Good Friday (various states)
	holidays = append(holidays, internal.NewHolidayFromTime(internal.CatholicGoodFriday(year), "Good Friday"))

	// Columbus Day (various states)
	holidays = append(holidays, internal.NewHolidayFromTime(secondMondayInOctober, "Columbus Day"))
	holidays = append(holidays, internal.NewHoliday(year, time.May, 8, "Truman Day"))

	// Lincoln's Birthday
	holidays = append(holidays, internal.NewHoliday(year, time.February, 12, "Lincoln's Birthday"))

	// TODO(crydafan): Implement state-specific holidays
	//if year >= 1988 {
	//	secondMondayInOctober := internal.FindDay(year, time.October, time.Monday, 2)
	//	holidays = append(holidays, internal.NewHolidayFromTime(secondMondayInOctober, "Indigenous Peoples' Day"))
	//}

	return holidays
}
