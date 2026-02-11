package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/internal"
)

type BrazilProvider struct{}

func (_ BrazilProvider) RegisterHolidays(year int) []internal.Holiday {
	holiday := baseHolidaysBR(year)
	if h := blackAwarenessDayBR(year); h != nil {
		holiday = append(holiday, *h)
	}

	return holiday
}

func baseHolidaysBR(year int) []internal.Holiday {
	easterSunday := internal.CatholicEasterSunday(year)

	return []internal.Holiday{
		internal.NewHoliday(year, time.January, 1, "New Year's Day", nil),
		internal.NewHoliday(year, time.April, 21, "Tiradentes", nil),
		internal.NewHoliday(year, time.May, 1, "Labour Day", nil),
		internal.NewHoliday(year, time.July, 9, "Constitutionalist Revolution of 1932", []string{"BR-SP"}),
		internal.NewHoliday(year, time.September, 7, "Independence Day", nil),
		internal.NewHoliday(year, time.October, 12, "Our Lady of Aparecida", nil),
		internal.NewHoliday(year, time.November, 2, "All Souls' Day", nil),
		internal.NewHoliday(year, time.November, 15, "Republic Proclamation Day", nil),
		internal.NewHoliday(year, time.December, 25, "Christmas Day", nil),
		internal.NewHolidayFromTime(easterSunday.AddDate(0, 0, -47), "Carnival", nil),
		internal.NewHolidayFromTime(easterSunday.AddDate(0, 0, -48), "Carnival", nil),
		internal.NewHolidayFromTime(easterSunday, "Easter Sunday", nil),
		internal.NewHolidayFromTime(internal.CatholicGoodFriday(year), "Good Friday", nil),
		internal.NewHolidayFromTime(internal.CatholicCorpusChristi(year), "Corpus Christi", nil),
	}
}

func blackAwarenessDayBR(year int) *internal.Holiday {
	if year < 2024 {
		return nil
	}
	holiday := internal.NewHoliday(year, time.November, 20, "Black Awareness Day", nil)
	return &holiday
}
