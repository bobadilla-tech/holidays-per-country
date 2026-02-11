// Package providers implements holiday providers for various countries.
// This file contains the provider for the United Kingdom (GB).
package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/calc"
)

// UnitedKingdomProvider provides public holidays for the United Kingdom
type UnitedKingdomProvider struct{}

func (UnitedKingdomProvider) RegisterHolidays(year int) []calc.Holiday {
	firstMondayInAugust := calc.FindDay(year, time.August, time.Monday, 1)
	lastMondayInAugust := calc.FindLastDay(year, time.August, time.Monday)

	holiday := []calc.Holiday{
		calc.NewHoliday(year, time.January, 1, "New Year's Day", []string{"GB-ENG", "GB-NIR", "GB-SCT", "GB-WLS"}),
		calc.NewHoliday(year, time.January, 2, "2 January", []string{"GB-SCT"}),
		calc.NewHoliday(year, time.March, 17, "Saint Patrick's Day", []string{"GB-NIR"}),
		calc.NewHoliday(year, time.July, 12, "Battle of the Boyne", []string{"GB-NIR"}),
		calc.NewHolidayFromTime(firstMondayInAugust, "Summer Bank Holiday", []string{"GB-SCT"}),
		calc.NewHolidayFromTime(lastMondayInAugust, "Summer Bank Holiday", []string{"GB-ENG", "GB-WLS", "GB-NIR"}),
		calc.NewHoliday(year, time.November, 30, "Saint Andrew's Day", []string{"GB-SCT"}),
		calc.NewHoliday(year, time.December, 25, "Christmas Day", nil),
		calc.NewHoliday(year, time.December, 26, "St. Stephen's Day", nil),
		calc.NewHolidayFromTime(calc.CatholicGoodFriday(year), "Good Friday", nil),
		calc.NewHolidayFromTime(calc.CatholicEasterMonday(year), "Easter Monday", []string{"GB-ENG", "GB-WLS", "GB-NIR"}),
	}

	if h := earlyMayBankHoliday(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := springBankHoliday(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := queensPlatinumJubilee(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := queensStateFuneral(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := coronationBankHoliday(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := worldCupBankHoliday(year); h != nil {
		holiday = append(holiday, *h)
	}

	return holiday
}

func springBankHoliday(year int) *calc.Holiday {
	if year == 2022 {
		holiday := calc.NewHoliday(year, time.June, 2, "Spring Bank Holiday", nil)
		return &holiday
	}

	lastMondayInMay := calc.FindLastDay(year, time.May, time.Monday)
	holiday := calc.NewHolidayFromTime(lastMondayInMay, "Spring Bank Holiday", nil)
	return &holiday
}

func queensPlatinumJubilee(year int) *calc.Holiday {
	if year != 2022 {
		return nil
	}
	holiday := calc.NewHoliday(year, time.June, 3, "Queen's Platinum Jubilee", nil)
	return &holiday
}

func queensStateFuneral(year int) *calc.Holiday {
	if year != 2022 {
		return nil
	}
	holiday := calc.NewHoliday(year, time.September, 19, "Queen's State Funeral", nil)
	return &holiday
}

func coronationBankHoliday(year int) *calc.Holiday {
	if year != 2023 {
		return nil
	}
	holiday := calc.NewHoliday(year, time.May, 8, "Coronation Bank Holiday", nil)
	return &holiday
}

func worldCupBankHoliday(year int) *calc.Holiday {
	if year != 2026 {
		return nil
	}
	holiday := calc.NewHoliday(year, time.June, 15, "World Cup Bank Holiday", []string{"GB-SCT"})
	return &holiday
}

func earlyMayBankHoliday(year int) *calc.Holiday {
	if year == 1995 {
		holiday := calc.NewHoliday(year, time.May, 8, "Early May Bank Holiday", nil)
		return &holiday
	}
	if year == 2020 {
		secondFridayInMay := calc.FindDay(year, time.May, time.Friday, 2)
		holiday := calc.NewHolidayFromTime(secondFridayInMay, "Early May Bank Holiday", nil)
		return &holiday
	}

	firstMondayInMay := calc.FindDay(year, time.May, time.Monday, 1)
	holiday := calc.NewHolidayFromTime(firstMondayInMay, "Early May Bank Holiday", nil)
	return &holiday
}
