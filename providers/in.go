// Package providers implements holiday providers for various countries.
// This file contains the provider for India (IN).
package providers

import "github.com/bobadilla-tech/holidays-per-country/providers/calc"

// IndiaProvider provides public holidays for India (stub implementation)
type IndiaProvider struct{}

func (IndiaProvider) RegisterHolidays(year int) []calc.Holiday {
	return []calc.Holiday{}
}
