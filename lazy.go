package holidays

import "github.com/bobadilla-tech/holidays-per-country/providers"

func LazyLoad(countryCode string) {
	if _, exists := registry[countryCode]; exists {
		return
	}

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
	default:
	}
}
