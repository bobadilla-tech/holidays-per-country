// Package common provides shared types used across the holidays library.
package common

import "time"

// Holiday represents a public holiday with its date, name, and applicable subdivisions.
type Holiday struct {
	Date         time.Time // The date of the holiday
	Name         string    // The holiday name
	Subdivisions []string  // Subdivision codes where the holiday applies (empty = nationwide)
	Fixed        bool      // Whether the holiday falls on a fixed date each year
}
