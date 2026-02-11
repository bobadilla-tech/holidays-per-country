package holidays

import (
	"fmt"
	"sync"

	"github.com/bobadilla-tech/holidays-per-country/providers"
	"github.com/bobadilla-tech/holidays-per-country/providers/calc"
)

// Holiday represents a public holiday with its date, name, and applicable subdivisions.
type Holiday = calc.Holiday

// provider is the interface that all country holiday providers must implement.
type provider interface {
	RegisterHolidays(year int) []Holiday
}

// registry holds all registered country holiday providers.
var registry = map[string]provider{}

// registryMutex protects concurrent access to the registry map.
var registryMutex sync.RWMutex

// holidaysCache stores computed holidays by country and year for performance.
var holidaysCache = map[string][]Holiday{}

// cacheMutex protects concurrent access to the holidaysCache map.
var cacheMutex sync.RWMutex

// cacheKey generates a unique cache key for a country-year combination.
func cacheKey(countryCode string, year int) string {
	return fmt.Sprintf("%s:%d", countryCode, year)
}

// ensureProviderLoaded automatically loads a country provider if not already loaded.
// This function is thread-safe and uses double-checked locking for optimal performance.
func ensureProviderLoaded(countryCode string) {
	// Fast path: check with read lock
	registryMutex.RLock()
	_, exists := registry[countryCode]
	registryMutex.RUnlock()

	if exists {
		return
	}

	// Slow path: load with write lock
	registryMutex.Lock()
	defer registryMutex.Unlock()

	// Double-check after acquiring write lock to prevent duplicate loading
	if _, exists := registry[countryCode]; exists {
		return
	}

	// Load the provider
	switch countryCode {
	case "AU":
		registry[countryCode] = providers.AustraliaProvider{}
	case "BR":
		registry[countryCode] = providers.BrazilProvider{}
	case "CA":
		registry[countryCode] = providers.CanadaProvider{}
	case "CN":
		registry[countryCode] = providers.ChinaProvider{}
	case "DE":
		registry[countryCode] = providers.GermanyProvider{}
	case "ES":
		registry[countryCode] = providers.SpainProvider{}
	case "FR":
		registry[countryCode] = providers.FranceProvider{}
	case "GB":
		registry[countryCode] = providers.UnitedKingdomProvider{}
	case "IN":
		registry[countryCode] = providers.IndiaProvider{}
	case "JP":
		registry[countryCode] = providers.JapanProvider{}
	case "MX":
		registry[countryCode] = providers.MexicoProvider{}
	case "US":
		registry[countryCode] = providers.UnitedStatesProvider{}
	}
}
