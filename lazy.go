package holidays

import "github.com/bobadilla-tech/holidays-per-country/providers"

func LazyLoad(countryCode string) {
	if _, exists := registry[countryCode]; exists {
		return
	}

	switch countryCode {
	case "US":
		registry[countryCode] = providers.UnitedStatesProvider{}
		// Add more country providers here as they are implemented
	}
}
