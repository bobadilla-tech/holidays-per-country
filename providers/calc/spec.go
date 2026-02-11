// Package calc provides shared utilities for calculating holiday dates across different countries.
package calc

import "time"

// Holiday represents a public holiday with its date, name, and applicable subdivisions.
type Holiday struct {
	Date         time.Time // The date of the holiday
	Name         string    // The holiday name
	Subdivisions []string  // Subdivision codes where the holiday applies (empty = nationwide)
	Fixed        bool      // Whether the holiday falls on a fixed date each year
}

// NewHoliday creates a new Holiday instance with a fixed date specified by year, month, and day.
func NewHoliday(year int, month time.Month, day int, name string, subdivisions []string) Holiday {
	return Holiday{
		Date:         time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
		Name:         name,
		Subdivisions: subdivisions,
		Fixed:        true,
	}
}

// NewHolidayFromTime creates a new Holiday instance from an existing time.Time value.
// This is typically used for holidays with variable dates that are calculated.
func NewHolidayFromTime(date time.Time, name string, subdivisions []string) Holiday {
	return Holiday{
		Date:         date,
		Name:         name,
		Subdivisions: subdivisions,
		Fixed:        false,
	}
}
