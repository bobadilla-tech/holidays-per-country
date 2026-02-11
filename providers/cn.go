package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/internal"
)

type ChinaProvider struct{}

func (_ ChinaProvider) RegisterHolidays(year int) []internal.Holiday {
	holiday := []internal.Holiday{
		internal.NewHoliday(year, time.January, 1, "New Year's Day", nil),
		internal.NewHoliday(year, time.May, 1, "Labour Day", nil),
		internal.NewHoliday(year, time.October, 1, "National Day", nil),
	}

	if year > 1901 && year < 2100 {
		if springFestival, ok := internal.ChineseLunisolarToGregorian(year, 1, 1); ok {
			holiday = append(holiday, internal.NewHolidayFromTime(springFestival, "Chinese New Year (Spring Festival)", nil))
		}
		if dragonBoatFestival, ok := internal.ChineseLunisolarToGregorian(year, 5, 5); ok {
			holiday = append(holiday, internal.NewHolidayFromTime(dragonBoatFestival, "Dragon Boat Festival", nil))
		}
		if midAutumnFestival, ok := internal.ChineseLunisolarToGregorian(year, 8, 15); ok {
			holiday = append(holiday, internal.NewHolidayFromTime(midAutumnFestival, "Mid-Autumn Festival", nil))
		}
	}

	return holiday
}
