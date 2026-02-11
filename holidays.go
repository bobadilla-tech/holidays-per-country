package holidays

import "time"

func IsHoliday(date time.Time, countryCode string) bool {
	holidays := GetHolidays(countryCode, date.Year())

	// TODO(crydafan): Optimize with binary search
	for _, holiday := range holidays {
		if holiday.Date.Year() == date.Year() &&
			holiday.Date.Month() == date.Month() &&
			holiday.Date.Day() == date.Day() {
			return true
		}
	}

	return false
}

func GetHolidays(countryCode string, year int) []Holiday {
	if holidays, exists := holidaysCache[countryCode]; exists {
		return holidays
	}

	holidays := registry[countryCode].RegisterHolidays(year)
	holidaysCache[countryCode] = holidays
	return holidays
}

func GetHolidaysInRange(countryCode string, startDate, endDate time.Time) []Holiday {
	if startDate.After(endDate) {
		return []Holiday{}
	}

	var result []Holiday

	for year := startDate.Year(); year <= endDate.Year(); year++ {
		holidays := GetHolidays(countryCode, year)

		for _, holiday := range holidays {
			// Check if holiday falls within the date range
			if (holiday.Date.Equal(startDate) || holiday.Date.After(startDate)) &&
				(holiday.Date.Equal(endDate) || holiday.Date.Before(endDate)) {
				result = append(result, holiday)
			}
		}
	}

	return result
}
