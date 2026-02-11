package holidays

import "github.com/bobadilla-tech/holidays-per-country/providers"

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
