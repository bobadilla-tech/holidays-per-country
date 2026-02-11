# Holidays Per Country

[![CI](https://github.com/bobadilla-tech/holidays-per-country/actions/workflows/ci.yml/badge.svg)](https://github.com/bobadilla-tech/holidays-per-country/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/bobadilla-tech/holidays-per-country.svg)](https://pkg.go.dev/github.com/bobadilla-tech/holidays-per-country)
[![Go Report Card](https://goreportcard.com/badge/github.com/bobadilla-tech/holidays-per-country)](https://goreportcard.com/report/github.com/bobadilla-tech/holidays-per-country)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A pure Go library for determining public holidays across different countries and
their subdivisions (states, provinces, regions). Provides accurate holiday
calculations including movable dates (Easter-based, lunar calendar, astronomical
events) and country-specific rules.

## Features

- **Subdivision Handling**: Support for state, province, and region-specific
  holidays
- **Movable Holidays**: Accurate calculation of Easter-based Catholic holidays,
  Chinese lunar festivals, Japanese equinoxes
- **Historical Accuracy**: Year-specific rules for holidays that change dates or
  names over time
- **Zero Dependencies**: Pure Go implementation with no external dependencies
- **Calendar Utilities**: Helper functions for finding specific days (first
  Monday, last Friday, etc.)

## Installation

```bash
go get github.com/bobadilla-tech/holidays-per-country
```

## Usage

### Basic Usage

Check if a specific date is a holiday:

```go
package main

import (
    "fmt"
    "time"

    "github.com/bobadilla-tech/holidays-per-country"
)

func main() {
    date := time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)

    if holidays.IsHoliday(date, "US") {
        fmt.Println("December 25, 2024 is a holiday in the United States")
    }
}
```

### Get All Holidays for a Year

Retrieve all holidays for a specific country and year:

```go
package main

import (
    "fmt"

    "github.com/bobadilla-tech/holidays-per-country"
)

func main() {
    // Canada holidays for 2024
    holidays := holidays.GetHolidays("CA", 2024)

    for _, holiday := range holidays {
        fmt.Printf("%s: %s\n", holiday.Date.Format("2006-01-02"), holiday.Name)
    }
}
```

### Get Holidays in a Date Range

Find all holidays within a specific date range:

```go
package main

import (
    "fmt"
    "time"

    "github.com/bobadilla-tech/holidays-per-country"
)

func main() {
    start := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2024, 8, 31, 0, 0, 0, 0, time.UTC)

    holidays := holidays.GetHolidaysInRange("US", start, end)

    fmt.Printf("Summer holidays in the US (2024):\n")
    for _, holiday := range holidays {
        fmt.Printf("%s: %s\n", holiday.Date.Format("2006-01-02"), holiday.Name)
    }
}
```

### Working with Subdivisions

Filter holidays by subdivision (state, province, region):

```go
package main

import (
    "fmt"

    "github.com/bobadilla-tech/holidays-per-country"
)

func main() {
    holidays := holidays.GetHolidays("AU", 2024)

    // Filter holidays for Victoria
    for _, holiday := range holidays {
        // Check if holiday applies to Victoria
        if len(holiday.Subdivisions) == 0 || contains(holiday.Subdivisions, "AU-VIC") {
            fmt.Printf("%s: %s\n", holiday.Date.Format("2006-01-02"), holiday.Name)
        }
    }
}

func contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}
```

## API Reference

### Types

#### `Holiday`

Represents a public holiday:

```go
type Holiday struct {
    Date         time.Time // The date of the holiday
    Name         string    // The holiday name
    Subdivisions []string  // Subdivision codes where the holiday applies (empty = nationwide)
    Fixed        bool      // Whether the holiday falls on a fixed date each year
}
```

### Functions

#### `IsHoliday(date time.Time, countryCode string) bool`

Checks if a specific date is a public holiday in the given country.

- **Parameters:**
  - `date`: The date to check
  - `countryCode`: ISO 3166-1 alpha-2 country code (e.g., "US", "CA", "GB")
- **Returns:** `true` if the date is a holiday, `false` otherwise

#### `GetHolidays(countryCode string, year int) []Holiday`

Returns all public holidays for a specific country and year.

- **Parameters:**
  - `countryCode`: ISO 3166-1 alpha-2 country code
  - `year`: The year to query
- **Returns:** Slice of `Holiday` objects

#### `GetHolidaysInRange(countryCode string, startDate, endDate time.Time) []Holiday`

Returns all holidays within a date range for a specific country.

- **Parameters:**
  - `countryCode`: ISO 3166-1 alpha-2 country code
  - `startDate`: Start of the date range (inclusive)
  - `endDate`: End of the date range (inclusive)
- **Returns:** Slice of `Holiday` objects in chronological order

### Subdivision Codes

Many providers support subdivision-specific holidays using ISO 3166-2 codes:

- **Australia**: AU-ACT, AU-NSW, AU-NT, AU-QLD, AU-SA, AU-TAS, AU-VIC, AU-WA
- **Canada**: CA-AB, CA-BC, CA-MB, CA-NB, CA-NL, CA-NT, CA-NS, CA-NU, CA-ON,
  CA-PE, CA-QC, CA-SK, CA-YT
- **Germany**: DE-BB, DE-BE, DE-BW, DE-BY, DE-HB, DE-HE, DE-HH, DE-MV, DE-NI,
  DE-NW, DE-RP, DE-SH, DE-SL, DE-SN, DE-ST, DE-TH
- **Spain**: ES-AN, ES-AR, ES-AS, ES-CB, ES-CE, ES-CL, ES-CM, ES-CN, ES-CT,
  ES-EX, ES-GA, ES-IB, ES-MC, ES-MD, ES-ML, ES-NC, ES-PV, ES-RI, ES-VC
- **United Kingdom**: GB-ENG, GB-NIR, GB-SCT, GB-WLS
- **United States**: US-CT, US-DE, US-FL, US-HI, US-IL, US-IN, US-KY, US-LA,
  US-MO, US-MT, US-NC, US-NJ, US-TN, and more

## How It Works

### Provider Pattern

Each country is implemented as a provider that registers holidays for a given
year:

```go
type Provider interface {
    RegisterHolidays(year int) []Holiday
}
```

### Holiday Calculation

The library handles different types of holidays:

1. **Fixed Holidays**: Same date every year (e.g., New Year's Day on January 1)
2. **Movable Holidays**:
   - **Catholic Easter-based**: Good Friday, Easter Monday, Ascension Day,
     Pentecost, Corpus Christi, etc.
   - **Lunar Calendar**: Chinese Spring Festival, Dragon Boat Festival,
     Mid-Autumn Festival
   - **Astronomical**: Japanese vernal and autumnal equinoxes
3. **Rule-Based**: Holidays that shift based on:
   - Day of week (e.g., first Monday in September)
   - Conditional logic (e.g., presidential inauguration years)
   - Historical changes (e.g., holiday name changes, new holidays added)

### Date Utilities

The library provides utilities for finding specific days:

- `FindDay(year, month, dayOfWeek, occurrence)`: Find the Nth occurrence of a
  weekday (e.g., 3rd Monday of February)
- `FindLastDay(year, month, dayOfWeek)`: Find the last occurrence of a weekday
  in a month
- `FindNextDay(date, dayOfWeek)`: Find the next occurrence of a specific weekday
- `FindDayBefore(date, dayOfWeek)`: Find the previous occurrence of a specific
  weekday

### Special Calendar Support

- **Catholic Calendar**: Accurate Easter calculation using Gauss's algorithm and
  derived Catholic holidays
- **Chinese Lunisolar Calendar**: Conversion from lunar dates to Gregorian for
  years 1900-2100
- **Japanese Equinoxes**: Mathematical calculation of vernal and autumnal
  equinoxes

## Testing

Run the test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

Run tests for a specific provider:

```bash
go test ./providers -run TestUnitedStatesProvider
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for
detailed guidelines on:

- Adding new country providers
- Reporting bugs and requesting features
- Code style and testing requirements
- Pull request process

Quick start for adding a country:

1. Create a new file in `providers/` named `{countrycode}.go`
2. Implement the provider following the pattern in existing providers
3. Use helper functions from `providers/internal/` for date calculations
4. Add tests in `providers/{countrycode}_test.go`
5. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file
for details.

## Acknowledgments

Holiday data and rules derived from:

- [Nager.Date](https://date.nager.at/) - Public holiday API
