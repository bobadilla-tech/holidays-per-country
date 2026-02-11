// Package providers implements holiday providers for various countries.
// This file contains the provider for Germany (DE).
package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/common"
	"github.com/bobadilla-tech/holidays-per-country/providers/internal"
)

// GermanyProvider provides public holidays for Germany and its federal states
type GermanyProvider struct{}

func (GermanyProvider) RegisterHolidays(year int) []common.Holiday {
	holiday := baseHolidaysDE(year)
	if h := internationalWomensDayDE(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := reformationDayDE(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := repentanceAndPrayerDayDE(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := liberationDayDE(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := worldChildrensDayDE(year); h != nil {
		holiday = append(holiday, *h)
	}

	return holiday
}

func baseHolidaysDE(year int) []common.Holiday {
	return []common.Holiday{
		internal.NewHoliday(year, time.January, 1, "New Year's Day", nil),
		internal.NewHoliday(year, time.January, 6, "Epiphany", []string{"DE-BW", "DE-BY", "DE-ST"}),
		internal.NewHoliday(year, time.May, 1, "Labour Day", []string{"DE-BW", "DE-BY", "DE-BE", "DE-BB", "DE-HB", "DE-HH", "DE-HE", "DE-MV", "DE-NI", "DE-NW", "DE-RP", "DE-SL", "DE-SN", "DE-ST", "DE-TH"}),
		internal.NewHoliday(year, time.August, 15, "Assumption Day", []string{"DE-SL"}),
		internal.NewHoliday(year, time.October, 3, "German Unity Day", []string{"DE-BW", "DE-BY", "DE-BE", "DE-BB", "DE-HB", "DE-HH", "DE-HE", "DE-MV", "DE-NI", "DE-NW", "DE-RP", "DE-SL", "DE-SN", "DE-ST", "DE-TH"}),
		internal.NewHoliday(year, time.November, 1, "All Saints' Day", []string{"DE-BW", "DE-BY", "DE-NW", "DE-RP", "DE-SL"}),
		internal.NewHoliday(year, time.December, 25, "Christmas Day", []string{"DE-BW", "DE-BY", "DE-BE", "DE-BB", "DE-HB", "DE-HH", "DE-HE", "DE-MV", "DE-NI", "DE-NW", "DE-RP", "DE-SL", "DE-SN", "DE-ST", "DE-TH"}),
		internal.NewHoliday(year, time.December, 26, "St. Stephen's Day", nil),
		internal.NewHolidayFromTime(internal.CatholicGoodFriday(year), "Good Friday", nil),
		internal.NewHolidayFromTime(internal.CatholicEasterSunday(year), "Easter Sunday", []string{"DE-BB"}),
		internal.NewHolidayFromTime(internal.CatholicEasterMonday(year), "Easter Monday", nil),
		internal.NewHolidayFromTime(internal.CatholicAscensionDay(year), "Ascension Day", nil),
		internal.NewHolidayFromTime(internal.CatholicPentecost(year), "Pentecost", []string{"DE-BB"}),
		internal.NewHolidayFromTime(internal.CatholicWhitMonday(year), "Whit Monday", nil),
		internal.NewHolidayFromTime(internal.CatholicCorpusChristi(year), "Corpus Christi", []string{"DE-BW", "DE-BY", "DE-HE", "DE-NW", "DE-RP", "DE-SL"}),
	}
}

func internationalWomensDayDE(year int) *common.Holiday {
	if year >= 2019 && year <= 2022 {
		holiday := internal.NewHoliday(year, time.March, 8, "International Women's Day", []string{"DE-BE"})
		return &holiday
	}
	if year >= 2023 {
		holiday := internal.NewHoliday(year, time.March, 8, "International Women's Day", []string{"DE-BE", "DE-MV"})
		return &holiday
	}

	return nil
}

func reformationDayDE(year int) *common.Holiday {
	if year == 2017 {
		holiday := internal.NewHoliday(year, time.October, 31, "Reformation Day", nil)
		return &holiday
	}
	if year >= 2018 {
		holiday := internal.NewHoliday(year, time.October, 31, "Reformation Day", []string{"DE-BB", "DE-MV", "DE-SN", "DE-ST", "DE-TH", "DE-HB", "DE-HH", "DE-NI", "DE-SH"})
		return &holiday
	}

	holiday := internal.NewHoliday(year, time.October, 31, "Reformation Day", []string{"DE-BB", "DE-MV", "DE-SN", "DE-ST", "DE-TH"})
	return &holiday
}

func repentanceAndPrayerDayDE(year int) *common.Holiday {
	christmas := time.Date(year, time.December, 25, 0, 0, 0, 0, time.UTC)
	sundayBeforeChristmas := internal.FindDayBefore(christmas, time.Sunday)
	adventSunday := sundayBeforeChristmas.AddDate(0, 0, -21)
	dayOfPrayer := adventSunday.AddDate(0, 0, -11)

	if year >= 1934 && year < 1939 {
		holiday := internal.NewHolidayFromTime(dayOfPrayer, "Repentance and Prayer Day", nil)
		return &holiday
	}
	if year >= 1945 && year <= 1980 {
		holiday := internal.NewHolidayFromTime(dayOfPrayer, "Repentance and Prayer Day", []string{"DE-BW", "DE-BE", "DE-HB", "DE-HH", "DE-HE", "DE-NI", "DE-NW", "DE-RP", "DE-SL", "DE-SH"})
		return &holiday
	}
	if year >= 1981 && year <= 1989 {
		holiday := internal.NewHolidayFromTime(dayOfPrayer, "Repentance and Prayer Day", []string{"DE-BW", "DE-BY", "DE-BE", "DE-HB", "DE-HH", "DE-HE", "DE-NI", "DE-NW", "DE-RP", "DE-SL", "DE-SH"})
		return &holiday
	}
	if year >= 1990 && year <= 1994 {
		holiday := internal.NewHolidayFromTime(dayOfPrayer, "Repentance and Prayer Day", nil)
		return &holiday
	}
	if year >= 1995 {
		holiday := internal.NewHolidayFromTime(dayOfPrayer, "Repentance and Prayer Day", []string{"DE-SN"})
		return &holiday
	}

	return nil
}

func liberationDayDE(year int) *common.Holiday {
	if year == 2020 || year == 2025 {
		holiday := internal.NewHoliday(year, time.May, 8, "Liberation Day", []string{"DE-BE"})
		return &holiday
	}

	return nil
}

func worldChildrensDayDE(year int) *common.Holiday {
	if year < 2019 {
		return nil
	}
	holiday := internal.NewHoliday(year, time.September, 20, "World Children's Day", []string{"DE-TH"})
	return &holiday
}
