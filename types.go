package holidays

import "github.com/bobadilla-tech/holidays-per-country/providers/calc"

// Holiday represents a public holiday with its date, name, and applicable subdivisions.
// This is an alias to the calc.Holiday type to avoid import cycles.
type Holiday = calc.Holiday
