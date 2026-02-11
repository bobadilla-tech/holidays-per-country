package holidays_test

import (
	"sync"
	"testing"
	"time"

	"github.com/bobadilla-tech/holidays-per-country"
)

// TestAutoLoadRaceCondition tests for race conditions in automatic provider loading
// Run with: go test -race -run TestAutoLoadRaceCondition
func TestAutoLoadRaceCondition(t *testing.T) {
	const numGoroutines = 10

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Simulate multiple concurrent requests trying to auto-load the same country
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			// Auto-loading happens automatically
			holidays.GetHolidays("US", 2024)
		}()
	}

	wg.Wait()
}

// TestConcurrentDifferentCountries tests concurrent loading of different countries
func TestConcurrentDifferentCountries(t *testing.T) {
	countries := []string{"US", "CA", "GB", "AU", "DE", "FR", "JP", "MX"}

	var wg sync.WaitGroup
	wg.Add(len(countries))

	for _, country := range countries {
		go func(c string) {
			defer wg.Done()
			// Auto-loading happens automatically
			holidays.GetHolidays(c, 2024)
		}(country)
	}

	wg.Wait()
}

// TestConcurrentReads tests that concurrent reads are safe
func TestConcurrentReads(t *testing.T) {
	// Pre-warm cache
	holidays.GetHolidays("US", 2024)

	const numReaders = 100
	var wg sync.WaitGroup
	wg.Add(numReaders)

	date := time.Date(2024, 7, 4, 0, 0, 0, 0, time.UTC)

	for i := 0; i < numReaders; i++ {
		go func() {
			defer wg.Done()
			// These should all be safe concurrent reads
			holidays.IsHoliday(date, "US")
			holidays.GetHolidays("US", 2024)
			holidays.GetHolidaysInRange("US", date, date)
		}()
	}

	wg.Wait()
}

// TestConcurrentFirstLoad tests the scenario where multiple goroutines
// try to load and use a country for the first time simultaneously
func TestConcurrentFirstLoad(t *testing.T) {
	// This test specifically tests the FIRST load scenario with auto-loading

	const numGoroutines = 50
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Use a country that hasn't been loaded yet
	// We use BR since it's less commonly used in other tests
	country := "BR"
	year := 2024

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			// Each goroutine triggers auto-loading
			result := holidays.GetHolidays(country, year)
			if len(result) == 0 {
				t.Errorf("Expected holidays for %s %d, got empty result", country, year)
			}
		}()
	}

	wg.Wait()
}
