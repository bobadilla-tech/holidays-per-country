// Package providers implements holiday providers for various countries.
// This file contains the provider for Canada (CA).
package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/calc"
)

// CanadaProvider provides public holidays for Canada and its provinces
type CanadaProvider struct{}

func (CanadaProvider) RegisterHolidays(year int) []calc.Holiday {
	thirdMondayInFebruary := calc.FindDay(year, time.February, time.Monday, 3)
	mondayOnOrBeforeMay25 := calc.FindDayBefore(time.Date(year, time.May, 25, 0, 0, 0, 0, time.UTC), time.Monday)
	firstMondayInAugust := calc.FindDay(year, time.August, time.Monday, 1)
	thirdMondayInAugust := calc.FindDay(year, time.August, time.Monday, 3)
	firstMondayInSeptember := calc.FindDay(year, time.September, time.Monday, 1)
	secondMondayInOctober := calc.FindDay(year, time.October, time.Monday, 2)

	holiday := []calc.Holiday{
		calc.NewHoliday(year, time.January, 1, "New Year's Day", nil),
		calc.NewHolidayFromTime(thirdMondayInFebruary, "Louis Riel Day", []string{"CA-MB"}),
		calc.NewHolidayFromTime(thirdMondayInFebruary, "Islander Day", []string{"CA-PE"}),
		calc.NewHolidayFromTime(thirdMondayInFebruary, "Heritage Day", []string{"CA-NS"}),
		calc.NewHoliday(year, time.March, 17, "Saint Patrick's Day", []string{"CA-NL"}),
		calc.NewHoliday(year, time.April, 23, "Saint George's Day", []string{"CA-NL"}),
		calc.NewHolidayFromTime(mondayOnOrBeforeMay25, "National Patriots' Day", []string{"CA-QC"}),
		calc.NewHolidayFromTime(mondayOnOrBeforeMay25, "Victoria Day", nil),
		calc.NewHoliday(year, time.June, 21, "National Aboriginal Day", []string{"CA-NT"}),
		calc.NewHoliday(year, time.June, 24, "Discovery Day", []string{"CA-NL"}),
		calc.NewHoliday(year, time.June, 24, "National Holiday", []string{"CA-QC"}),
		calc.NewHoliday(year, time.July, 12, "Orangemen's Day", []string{"CA-NL"}),
		calc.NewHolidayFromTime(firstMondayInAugust, "Civic Holiday", []string{"CA-MB", "CA-NL", "CA-NT", "CA-NU", "CA-ON"}),
		calc.NewHolidayFromTime(firstMondayInAugust, "British Columbia Day", []string{"CA-BC"}),
		calc.NewHolidayFromTime(firstMondayInAugust, "Heritage Day", []string{"CA-AB", "CA-YT"}),
		calc.NewHolidayFromTime(firstMondayInAugust, "New Brunswick Day", []string{"CA-NB"}),
		calc.NewHolidayFromTime(firstMondayInAugust, "Natal Day", []string{"CA-NS"}),
		calc.NewHolidayFromTime(firstMondayInAugust, "Saskatchewan Day", []string{"CA-SK"}),
		calc.NewHolidayFromTime(thirdMondayInAugust, "Gold Cup Parade Day", []string{"CA-PE"}),
		calc.NewHolidayFromTime(thirdMondayInAugust, "Discovery Day", []string{"CA-YT"}),
		calc.NewHolidayFromTime(firstMondayInSeptember, "Labour Day", nil),
		calc.NewHoliday(year, time.September, 30, "National Day for Truth and Reconciliation", nil),
		calc.NewHolidayFromTime(secondMondayInOctober, "Thanksgiving", nil),
		calc.NewHoliday(year, time.November, 11, "Armistice Day", []string{"CA-NL"}),
		calc.NewHoliday(year, time.November, 11, "Remembrance Day", []string{"CA-AB", "CA-BC", "CA-NB", "CA-NT", "CA-NS", "CA-NU", "CA-PE", "CA-SK", "CA-YT"}),
		calc.NewHoliday(year, time.December, 25, "Christmas Day", nil),
		calc.NewHoliday(year, time.December, 26, "St. Stephen's Day", []string{"CA-AB", "CA-NB", "CA-NS", "CA-ON", "CA-PE"}),
		calc.NewHolidayFromTime(calc.CatholicGoodFriday(year), "Good Friday", nil),
		calc.NewHolidayFromTime(calc.CatholicEasterMonday(year), "Easter Monday", []string{"CA-AB", "CA-PE"}),
	}

	holiday = append(holiday, canadaDay(year)...)
	holiday = append(holiday, familyDay(year)...)
	if special := funeralForQueenElizabeth(year); special != nil {
		holiday = append(holiday, *special)
	}

	return holiday
}

func funeralForQueenElizabeth(year int) *calc.Holiday {
	if year != 2022 {
		return nil
	}

	holiday := calc.NewHoliday(year, time.September, 19, "State Funeral of Queen Elizabeth II", nil)
	return &holiday
}

func familyDay(year int) []calc.Holiday {
	name := "Family Day"
	thirdMondayInFebruary := calc.FindDay(year, time.February, time.Monday, 3)

	if year < 2019 {
		secondMondayInFebruary := calc.FindDay(year, time.February, time.Monday, 2)
		return []calc.Holiday{
			calc.NewHolidayFromTime(secondMondayInFebruary, name, []string{"CA-BC"}),
			calc.NewHolidayFromTime(thirdMondayInFebruary, name, []string{"CA-AB", "CA-ON", "CA-SK"}),
		}
	}

	return []calc.Holiday{
		calc.NewHolidayFromTime(thirdMondayInFebruary, name, []string{"CA-AB", "CA-BC", "CA-NB", "CA-ON", "CA-SK"}),
	}
}

func canadaDay(year int) []calc.Holiday {
	name := "Canada Day"
	canadaDay := time.Date(year, time.July, 1, 0, 0, 0, 0, time.UTC)

	if canadaDay.Weekday() == time.Sunday {
		return []calc.Holiday{
			calc.NewHolidayFromTime(canadaDay, name, []string{"CA-BC", "CA-MB", "CA-NB", "CA-NL", "CA-NS", "CA-ON", "CA-PE", "CA-QC", "CA-SK", "CA-NT", "CA-NU", "CA-YT"}),
			calc.NewHolidayFromTime(canadaDay.AddDate(0, 0, 1), name, []string{"CA-AB"}),
		}
	}

	return []calc.Holiday{calc.NewHolidayFromTime(canadaDay, name, nil)}
}
