// Package providers implements holiday providers for various countries.
// This file contains the provider for Canada (CA).
package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/internal"
)

// CanadaProvider provides public holidays for Canada and its provinces
type CanadaProvider struct{}

func (CanadaProvider) RegisterHolidays(year int) []internal.Holiday {
	thirdMondayInFebruary := internal.FindDay(year, time.February, time.Monday, 3)
	mondayOnOrBeforeMay25 := internal.FindDayBefore(time.Date(year, time.May, 25, 0, 0, 0, 0, time.UTC), time.Monday)
	firstMondayInAugust := internal.FindDay(year, time.August, time.Monday, 1)
	thirdMondayInAugust := internal.FindDay(year, time.August, time.Monday, 3)
	firstMondayInSeptember := internal.FindDay(year, time.September, time.Monday, 1)
	secondMondayInOctober := internal.FindDay(year, time.October, time.Monday, 2)

	holiday := []internal.Holiday{
		internal.NewHoliday(year, time.January, 1, "New Year's Day", nil),
		internal.NewHolidayFromTime(thirdMondayInFebruary, "Louis Riel Day", []string{"CA-MB"}),
		internal.NewHolidayFromTime(thirdMondayInFebruary, "Islander Day", []string{"CA-PE"}),
		internal.NewHolidayFromTime(thirdMondayInFebruary, "Heritage Day", []string{"CA-NS"}),
		internal.NewHoliday(year, time.March, 17, "Saint Patrick's Day", []string{"CA-NL"}),
		internal.NewHoliday(year, time.April, 23, "Saint George's Day", []string{"CA-NL"}),
		internal.NewHolidayFromTime(mondayOnOrBeforeMay25, "National Patriots' Day", []string{"CA-QC"}),
		internal.NewHolidayFromTime(mondayOnOrBeforeMay25, "Victoria Day", nil),
		internal.NewHoliday(year, time.June, 21, "National Aboriginal Day", []string{"CA-NT"}),
		internal.NewHoliday(year, time.June, 24, "Discovery Day", []string{"CA-NL"}),
		internal.NewHoliday(year, time.June, 24, "National Holiday", []string{"CA-QC"}),
		internal.NewHoliday(year, time.July, 12, "Orangemen's Day", []string{"CA-NL"}),
		internal.NewHolidayFromTime(firstMondayInAugust, "Civic Holiday", []string{"CA-MB", "CA-NL", "CA-NT", "CA-NU", "CA-ON"}),
		internal.NewHolidayFromTime(firstMondayInAugust, "British Columbia Day", []string{"CA-BC"}),
		internal.NewHolidayFromTime(firstMondayInAugust, "Heritage Day", []string{"CA-AB", "CA-YT"}),
		internal.NewHolidayFromTime(firstMondayInAugust, "New Brunswick Day", []string{"CA-NB"}),
		internal.NewHolidayFromTime(firstMondayInAugust, "Natal Day", []string{"CA-NS"}),
		internal.NewHolidayFromTime(firstMondayInAugust, "Saskatchewan Day", []string{"CA-SK"}),
		internal.NewHolidayFromTime(thirdMondayInAugust, "Gold Cup Parade Day", []string{"CA-PE"}),
		internal.NewHolidayFromTime(thirdMondayInAugust, "Discovery Day", []string{"CA-YT"}),
		internal.NewHolidayFromTime(firstMondayInSeptember, "Labour Day", nil),
		internal.NewHoliday(year, time.September, 30, "National Day for Truth and Reconciliation", nil),
		internal.NewHolidayFromTime(secondMondayInOctober, "Thanksgiving", nil),
		internal.NewHoliday(year, time.November, 11, "Armistice Day", []string{"CA-NL"}),
		internal.NewHoliday(year, time.November, 11, "Remembrance Day", []string{"CA-AB", "CA-BC", "CA-NB", "CA-NT", "CA-NS", "CA-NU", "CA-PE", "CA-SK", "CA-YT"}),
		internal.NewHoliday(year, time.December, 25, "Christmas Day", nil),
		internal.NewHoliday(year, time.December, 26, "St. Stephen's Day", []string{"CA-AB", "CA-NB", "CA-NS", "CA-ON", "CA-PE"}),
		internal.NewHolidayFromTime(internal.CatholicGoodFriday(year), "Good Friday", nil),
		internal.NewHolidayFromTime(internal.CatholicEasterMonday(year), "Easter Monday", []string{"CA-AB", "CA-PE"}),
	}

	holiday = append(holiday, canadaDay(year)...)
	holiday = append(holiday, familyDay(year)...)
	if special := funeralForQueenElizabeth(year); special != nil {
		holiday = append(holiday, *special)
	}

	return holiday
}

func funeralForQueenElizabeth(year int) *internal.Holiday {
	if year != 2022 {
		return nil
	}

	holiday := internal.NewHoliday(year, time.September, 19, "State Funeral of Queen Elizabeth II", nil)
	return &holiday
}

func familyDay(year int) []internal.Holiday {
	name := "Family Day"
	thirdMondayInFebruary := internal.FindDay(year, time.February, time.Monday, 3)

	if year < 2019 {
		secondMondayInFebruary := internal.FindDay(year, time.February, time.Monday, 2)
		return []internal.Holiday{
			internal.NewHolidayFromTime(secondMondayInFebruary, name, []string{"CA-BC"}),
			internal.NewHolidayFromTime(thirdMondayInFebruary, name, []string{"CA-AB", "CA-ON", "CA-SK"}),
		}
	}

	return []internal.Holiday{
		internal.NewHolidayFromTime(thirdMondayInFebruary, name, []string{"CA-AB", "CA-BC", "CA-NB", "CA-ON", "CA-SK"}),
	}
}

func canadaDay(year int) []internal.Holiday {
	name := "Canada Day"
	canadaDay := time.Date(year, time.July, 1, 0, 0, 0, 0, time.UTC)

	if canadaDay.Weekday() == time.Sunday {
		return []internal.Holiday{
			internal.NewHolidayFromTime(canadaDay, name, []string{"CA-BC", "CA-MB", "CA-NB", "CA-NL", "CA-NS", "CA-ON", "CA-PE", "CA-QC", "CA-SK", "CA-NT", "CA-NU", "CA-YT"}),
			internal.NewHolidayFromTime(canadaDay.AddDate(0, 0, 1), name, []string{"CA-AB"}),
		}
	}

	return []internal.Holiday{internal.NewHolidayFromTime(canadaDay, name, nil)}
}
