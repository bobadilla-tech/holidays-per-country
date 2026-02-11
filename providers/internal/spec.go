// Package internal provides shared utilities for calculating holidays across different countries.
// This file contains constructors for creating Holiday instances.
package internal

import (
	"time"
)

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
