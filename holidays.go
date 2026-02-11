// Package holidays provides functionality for determining public holidays across different countries.
// It supports subdivision-specific holidays, movable holidays (Easter-based, lunar calendar, etc.),
// and includes production-grade features like thread-safe caching and binary search optimization.
package holidays

import (
	"sort"
	"time"
)

// IsHoliday checks if a specific date is a public holiday in the given country.
// Uses binary search for O(log n) lookup performance.
func IsHoliday(date time.Time, countryCode string) bool {
	holidays := GetHolidays(countryCode, date.Year())

	normalizedDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)

	idx := sort.Search(len(holidays), func(i int) bool {
		return !holidays[i].Date.Before(normalizedDate)
	})

	if idx < len(holidays) {
		holidayDate := holidays[idx].Date

		return holidayDate.Year() == date.Year() &&
			holidayDate.Month() == date.Month() &&
			holidayDate.Day() == date.Day()
	}

	return false
}

// GetHolidays returns all public holidays for a specific country and year.
// Results are cached for performance. Thread-safe for concurrent use.
// Automatically loads the country provider on first use (lazy loading).
func GetHolidays(countryCode string, year int) []Holiday {
	key := cacheKey(countryCode, year)

	cacheMutex.RLock()
	if holidays, exists := holidaysCache[key]; exists {
		cacheMutex.RUnlock()
		return holidays
	}
	cacheMutex.RUnlock()

	ensureProviderLoaded(countryCode)

	registryMutex.RLock()
	provider, exists := registry[countryCode]
	registryMutex.RUnlock()

	if !exists {
		return []Holiday{}
	}

	holidays := provider.RegisterHolidays(year)

	sort.Slice(holidays, func(i, j int) bool {
		return holidays[i].Date.Before(holidays[j].Date)
	})

	cacheMutex.Lock()
	holidaysCache[key] = holidays
	cacheMutex.Unlock()

	return holidays
}

// GetHolidaysInRange returns all holidays within a date range for a specific country.
// Results are returned in chronological order.
func GetHolidaysInRange(countryCode string, startDate, endDate time.Time) []Holiday {
	if startDate.After(endDate) {
		return []Holiday{}
	}

	var result []Holiday

	for year := startDate.Year(); year <= endDate.Year(); year++ {
		holidays := GetHolidays(countryCode, year)

		for _, holiday := range holidays {
			if (holiday.Date.Equal(startDate) || holiday.Date.After(startDate)) &&
				(holiday.Date.Equal(endDate) || holiday.Date.Before(endDate)) {
				result = append(result, holiday)
			}
		}
	}

	return result
}
