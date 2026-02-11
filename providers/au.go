package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/internal"
)

type AustraliaProvider struct{}

func (_ AustraliaProvider) RegisterHolidays(year int) []internal.Holiday {
	// Labour Day variations by state
	firstMondayInMarch := internal.FindDay(year, time.March, time.Monday, 1)
	secondMondayInMarchLD := internal.FindDay(year, time.March, time.Monday, 2)
	// Monarch Birthday variations
	secondMondayInJune := internal.FindDay(year, time.June, time.Monday, 2)
	lastMondayInSeptember := internal.FindLastDay(year, time.September, time.Monday)

	secondMondayInMarch := internal.FindDay(year, time.March, time.Monday, 2)
	firstMondayInMay := internal.FindDay(year, time.May, time.Monday, 1)
	may27 := time.Date(year, time.May, 27, 0, 0, 0, 0, time.UTC)
	firstMondayAfterOr27May := internal.FindNextDay(may27, time.Monday)
	firstMondayInJune := internal.FindDay(year, time.June, time.Monday, 1)
	firstMondayInAugust := internal.FindDay(year, time.August, time.Monday, 1)
	firstMondayInOctober := internal.FindDay(year, time.October, time.Monday, 1)
	firstTuesdayInNovember := internal.FindDay(year, time.November, time.Tuesday, 1)

	holidays := []internal.Holiday{
		internal.NewHoliday(year, time.January, 1, "New Year's Day", nil),
		internal.NewHoliday(year, time.January, 26, "Australia Day", nil),
		internal.NewHolidayFromTime(secondMondayInMarch, "Canberra Day", []string{"AU-ACT"}),
		internal.NewHolidayFromTime(secondMondayInMarch, "Adelaide Cup Day", []string{"AU-SA"}),
		internal.NewHolidayFromTime(secondMondayInMarch, "Eight Hours Day", []string{"AU-TAS"}),
		internal.NewHoliday(year, time.April, 25, "Anzac Day", nil),
		internal.NewHolidayFromTime(firstMondayInMay, "May Day", []string{"AU-NT"}),
		internal.NewHolidayFromTime(firstMondayAfterOr27May, "Reconciliation Day", []string{"AU-ACT"}),
		internal.NewHolidayFromTime(firstMondayInJune, "Western Australia Day", []string{"AU-WA"}),
		internal.NewHolidayFromTime(firstMondayInAugust, "Picnic Day", []string{"AU-NT"}),
		internal.NewHolidayFromTime(firstTuesdayInNovember, "Melbourne Cup", []string{"AU-VIC"}),
		internal.NewHoliday(year, time.December, 25, "Christmas Day", nil),
		internal.NewHoliday(year, time.December, 26, "St. Stephen's Day", nil),
		internal.NewHolidayFromTime(internal.CatholicEasterSaturday(year), "Easter Eve", []string{"AU-ACT", "AU-NSW", "AU-NT", "AU-QLD", "AU-SA", "AU-VIC"}),
		internal.NewHolidayFromTime(internal.CatholicGoodFriday(year), "Good Friday", nil),
		internal.NewHolidayFromTime(internal.CatholicEasterMonday(year), "Easter Monday", nil),

		// Labour Day variations
		internal.NewHolidayFromTime(firstMondayInMarch, "Labour Day", []string{"AU-WA"}),
		internal.NewHolidayFromTime(secondMondayInMarchLD, "Labour Day", []string{"AU-VIC"}),
		internal.NewHolidayFromTime(firstMondayInMay, "Labour Day", []string{"AU-QLD"}),
		internal.NewHolidayFromTime(firstMondayInOctober, "Labour Day", []string{"AU-ACT", "AU-NSW", "AU-SA"}),
	}

	// Easter Sunday with year-based subdivisions
	var easterSubdivisions []string
	switch {
	case year >= 2024:
		easterSubdivisions = []string{"AU-ACT", "AU-NSW", "AU-NT", "AU-QLD", "AU-SA", "AU-VIC", "AU-WA"}
	case year == 2023:
		easterSubdivisions = []string{"AU-ACT", "AU-NSW", "AU-NT", "AU-QLD", "AU-VIC", "AU-WA"}
	case year == 2022:
		easterSubdivisions = []string{"AU-ACT", "AU-NSW", "AU-QLD", "AU-VIC", "AU-WA"}
	case year >= 2017 && year <= 2021:
		easterSubdivisions = []string{"AU-ACT", "AU-NSW", "AU-QLD", "AU-VIC"}
	case year == 2016:
		easterSubdivisions = []string{"AU-ACT", "AU-NSW", "AU-VIC"}
	case year >= 2010 && year <= 2015:
		easterSubdivisions = []string{"AU-NSW"}
	}
	if len(easterSubdivisions) > 0 {
		holidays = append(holidays, internal.NewHolidayFromTime(internal.CatholicEasterSunday(year), "Easter Sunday", easterSubdivisions))
	}

	// Monarch Birthday
	monarchName := "Queen's Birthday"
	if year >= 2023 {
		monarchName = "King's Birthday"
	}
	holidays = append(holidays,
		internal.NewHolidayFromTime(secondMondayInJune, monarchName, []string{"AU-ACT", "AU-NSW", "AU-NT", "AU-SA", "AU-TAS", "AU-VIC"}),
		internal.NewHolidayFromTime(lastMondayInSeptember, monarchName, []string{"AU-WA"}),
		internal.NewHolidayFromTime(firstMondayInOctober, monarchName, []string{"AU-QLD"}),
	)

	// National Day of Mourning for Queen Elizabeth II (2022 only)
	if year == 2022 {
		holidays = append(holidays, internal.NewHoliday(year, time.September, 22, "National Day of Mourning", nil))
	}

	// Friday before AFL Grand Final (Victoria)
	switch year {
	case 2016, 2017, 2018, 2019, 2021, 2023, 2024, 2025:
		lastFridayInSeptember := internal.FindLastDay(year, time.September, time.Friday)
		holidays = append(holidays, internal.NewHolidayFromTime(lastFridayInSeptember, "Friday before AFL Grand Final", []string{"AU-VIC"}))
	case 2020:
		holidays = append(holidays, internal.NewHoliday(year, time.October, 23, "Friday before AFL Grand Final", []string{"AU-VIC"}))
	case 2022:
		holidays = append(holidays, internal.NewHoliday(year, time.September, 23, "Friday before AFL Grand Final", []string{"AU-VIC"}))
	}

	return holidays
}
