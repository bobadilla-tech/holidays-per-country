package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/internal"
)

type UnitedStatesProvider struct {
}

func (_ UnitedStatesProvider) RegisterHolidays(year int) []internal.Holiday {
	holidays := []internal.Holiday{
		internal.NewHoliday(year, time.January, 1, "New Year's Day"),
		//internal.VariableHoliday(thirdMondayInJanuary, "Martin Luther King, Jr. Day", "Martin Luther King, Jr. Day"),
		//internal.VariableHoliday(thirdMondayInFebruary, "Washington's Birthday", "Presidents Day"),
		//internal.VariableHoliday(lastMondayInMay, "Memorial Day", "Memorial Day"),
		internal.NewHoliday(year, time.July, 4, "Independence Day"),
		//h.VariableHoliday(firstMondayInSeptember, "Labor Day", "Labour Day"),
		//h.VariableHoliday(fourthThursdayInNovember, "Thanksgiving Day", "Thanksgiving Day"),
		internal.NewHoliday(year, time.November, 11, "Veterans Day"),
		internal.NewHoliday(year, time.December, 25, "Christmas Day"),
	}

	return holidays
}
