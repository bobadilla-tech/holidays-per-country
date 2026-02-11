// Package internal provides shared utilities for calculating holidays across different countries.
// This file contains functions for finding specific weekdays and dates within months.
package internal

import "time"

// FindDay finds the nth occurrence of a specific weekday in a month
// occurrence is 1-indexed (1 = first, 2 = second, etc.)
func FindDay(year int, month time.Month, weekday time.Weekday, occurrence int) time.Time {
	if occurrence < 1 || occurrence > 5 {
		return time.Time{}
	}

	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	daysNeeded := int(weekday) - int(firstDay.Weekday())

	if daysNeeded < 0 {
		daysNeeded += 7
	}

	day := daysNeeded + 1 + (7 * (occurrence - 1))

	// Check if the day exists in this month
	if day > daysInMonth(year, month) {
		return time.Time{}
	}

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// FindLastDay finds the last occurrence of a specific weekday in a month
func FindLastDay(year int, month time.Month, weekday time.Weekday) time.Time {
	// Try 5th occurrence first, fallback to 4th
	result := FindDay(year, month, weekday, 5)
	if result.IsZero() {
		result = FindDay(year, month, weekday, 4)
	}
	return result
}

// FindNextDay finds the next occurrence of a specific weekday from a given date
func FindNextDay(date time.Time, weekday time.Weekday) time.Time {
	daysNeeded := int(weekday) - int(date.Weekday())

	if daysNeeded >= 0 {
		return date.AddDate(0, 0, daysNeeded)
	}

	return date.AddDate(0, 0, daysNeeded+7)
}

// FindDayBefore finds the most recent occurrence of a specific weekday before a given date
func FindDayBefore(date time.Time, weekday time.Weekday) time.Time {
	daysBack := int(date.Weekday()) - int(weekday)

	if daysBack < 0 {
		daysBack += 7
	}

	return date.AddDate(0, 0, -daysBack)
}

// FindDayBetween finds the first occurrence of a specific weekday between two dates (inclusive)
func FindDayBetween(startDate, endDate time.Time, weekday time.Weekday) time.Time {
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		if d.Weekday() == weekday {
			return d
		}
	}
	return time.Time{}
}

func daysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
