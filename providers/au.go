// Package providers implements holiday providers for various countries.
// This file contains the provider for Australia (AU).
package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/calc"
)

// AustraliaProvider provides public holidays for Australia and its states/territories
type AustraliaProvider struct{}

func (AustraliaProvider) RegisterHolidays(year int) []calc.Holiday {
	holiday := baseHolidaysAU(year)
	holiday = append(holiday, labourDayVariationsAU(year)...)
	holiday = append(holiday, monarchBirthdayAU(year)...)

	if h := easterSundayAU(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := nationalMourningAU(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := aflGrandFinalAU(year); h != nil {
		holiday = append(holiday, *h)
	}

	return holiday
}

func baseHolidaysAU(year int) []calc.Holiday {
	secondMondayInMarch := calc.FindDay(year, time.March, time.Monday, 2)
	firstMondayInMay := calc.FindDay(year, time.May, time.Monday, 1)
	may27 := time.Date(year, time.May, 27, 0, 0, 0, 0, time.UTC)
	firstMondayAfterOr27May := calc.FindNextDay(may27, time.Monday)
	firstMondayInJune := calc.FindDay(year, time.June, time.Monday, 1)
	firstMondayInAugust := calc.FindDay(year, time.August, time.Monday, 1)
	firstTuesdayInNovember := calc.FindDay(year, time.November, time.Tuesday, 1)

	return []calc.Holiday{
		calc.NewHoliday(year, time.January, 1, "New Year's Day", nil),
		calc.NewHoliday(year, time.January, 26, "Australia Day", nil),
		calc.NewHolidayFromTime(secondMondayInMarch, "Canberra Day", []string{"AU-ACT"}),
		calc.NewHolidayFromTime(secondMondayInMarch, "Adelaide Cup Day", []string{"AU-SA"}),
		calc.NewHolidayFromTime(secondMondayInMarch, "Eight Hours Day", []string{"AU-TAS"}),
		calc.NewHoliday(year, time.April, 25, "Anzac Day", nil),
		calc.NewHolidayFromTime(firstMondayInMay, "May Day", []string{"AU-NT"}),
		calc.NewHolidayFromTime(firstMondayAfterOr27May, "Reconciliation Day", []string{"AU-ACT"}),
		calc.NewHolidayFromTime(firstMondayInJune, "Western Australia Day", []string{"AU-WA"}),
		calc.NewHolidayFromTime(firstMondayInAugust, "Picnic Day", []string{"AU-NT"}),
		calc.NewHolidayFromTime(firstTuesdayInNovember, "Melbourne Cup", []string{"AU-VIC"}),
		calc.NewHoliday(year, time.December, 25, "Christmas Day", nil),
		calc.NewHoliday(year, time.December, 26, "St. Stephen's Day", nil),
		calc.NewHolidayFromTime(calc.CatholicEasterSaturday(year), "Easter Eve", []string{"AU-ACT", "AU-NSW", "AU-NT", "AU-QLD", "AU-SA", "AU-VIC"}),
		calc.NewHolidayFromTime(calc.CatholicGoodFriday(year), "Good Friday", nil),
		calc.NewHolidayFromTime(calc.CatholicEasterMonday(year), "Easter Monday", nil),
	}
}

func labourDayVariationsAU(year int) []calc.Holiday {
	firstMondayInMarch := calc.FindDay(year, time.March, time.Monday, 1)
	secondMondayInMarch := calc.FindDay(year, time.March, time.Monday, 2)
	firstMondayInMay := calc.FindDay(year, time.May, time.Monday, 1)
	firstMondayInOctober := calc.FindDay(year, time.October, time.Monday, 1)

	return []calc.Holiday{
		calc.NewHolidayFromTime(firstMondayInMarch, "Labour Day", []string{"AU-WA"}),
		calc.NewHolidayFromTime(secondMondayInMarch, "Labour Day", []string{"AU-VIC"}),
		calc.NewHolidayFromTime(firstMondayInMay, "Labour Day", []string{"AU-QLD"}),
		calc.NewHolidayFromTime(firstMondayInOctober, "Labour Day", []string{"AU-ACT", "AU-NSW", "AU-SA"}),
	}
}

func easterSundayAU(year int) *calc.Holiday {
	var subdivisions []string
	switch {
	case year >= 2024:
		subdivisions = []string{"AU-ACT", "AU-NSW", "AU-NT", "AU-QLD", "AU-SA", "AU-VIC", "AU-WA"}
	case year == 2023:
		subdivisions = []string{"AU-ACT", "AU-NSW", "AU-NT", "AU-QLD", "AU-VIC", "AU-WA"}
	case year == 2022:
		subdivisions = []string{"AU-ACT", "AU-NSW", "AU-QLD", "AU-VIC", "AU-WA"}
	case year >= 2017 && year <= 2021:
		subdivisions = []string{"AU-ACT", "AU-NSW", "AU-QLD", "AU-VIC"}
	case year == 2016:
		subdivisions = []string{"AU-ACT", "AU-NSW", "AU-VIC"}
	case year >= 2010 && year <= 2015:
		subdivisions = []string{"AU-NSW"}
	default:
		return nil
	}

	holiday := calc.NewHolidayFromTime(calc.CatholicEasterSunday(year), "Easter Sunday", subdivisions)
	return &holiday
}

func monarchBirthdayAU(year int) []calc.Holiday {
	secondMondayInJune := calc.FindDay(year, time.June, time.Monday, 2)
	lastMondayInSeptember := calc.FindLastDay(year, time.September, time.Monday)
	firstMondayInOctober := calc.FindDay(year, time.October, time.Monday, 1)

	name := "Queen's Birthday"
	if year >= 2023 {
		name = "King's Birthday"
	}

	return []calc.Holiday{
		calc.NewHolidayFromTime(secondMondayInJune, name, []string{"AU-ACT", "AU-NSW", "AU-NT", "AU-SA", "AU-TAS", "AU-VIC"}),
		calc.NewHolidayFromTime(lastMondayInSeptember, name, []string{"AU-WA"}),
		calc.NewHolidayFromTime(firstMondayInOctober, name, []string{"AU-QLD"}),
	}
}

func nationalMourningAU(year int) *calc.Holiday {
	if year != 2022 {
		return nil
	}
	holiday := calc.NewHoliday(year, time.September, 22, "National Day of Mourning", nil)
	return &holiday
}

func aflGrandFinalAU(year int) *calc.Holiday {
	switch year {
	case 2016, 2017, 2018, 2019, 2021, 2023, 2024, 2025:
		lastFridayInSeptember := calc.FindLastDay(year, time.September, time.Friday)
		holiday := calc.NewHolidayFromTime(lastFridayInSeptember, "Friday before AFL Grand Final", []string{"AU-VIC"})
		return &holiday
	case 2020:
		holiday := calc.NewHoliday(year, time.October, 23, "Friday before AFL Grand Final", []string{"AU-VIC"})
		return &holiday
	case 2022:
		holiday := calc.NewHoliday(year, time.September, 23, "Friday before AFL Grand Final", []string{"AU-VIC"})
		return &holiday
	default:
		return nil
	}
}
