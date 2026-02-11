package providers

import (
	"time"

	"github.com/bobadilla-tech/holidays-per-country/providers/internal"
)

type SpainProvider struct{}

func (_ SpainProvider) RegisterHolidays(year int) []internal.Holiday {
	holiday := []internal.Holiday{
		internal.NewHoliday(year, time.January, 6, "Epiphany", nil),
		internal.NewHoliday(year, time.May, 1, "Labour Day", nil),
		internal.NewHoliday(year, time.October, 12, "National Day of Spain", nil),
		internal.NewHoliday(year, time.November, 1, "All Saints Day", nil),
		internal.NewHoliday(year, time.December, 6, "Constitution Day", nil),
		internal.NewHoliday(year, time.December, 8, "Immaculate Conception", nil),
		internal.NewHolidayFromTime(internal.CatholicGoodFriday(year), "Good Friday", nil),
	}

	holiday = append(holiday, newYearsDayES(year)...)
	if h := christmasDayES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := corpusChristiES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := easterMondayES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := maundyThursdayES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := assumptionES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := carnivalTuesdayES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := castileAndLeonDayES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := dayOfAndaluciaES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := dayOfAragonES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := dayOfAsturiasES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := dayOfCastillaLaManchaES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := dayOfExtremaduraES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := dayOfLaRiojaES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := dayOfMadridES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := dayOfMurciaES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := dayOfBalearicIslandsES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := dayOfCanaryIslandsES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := dayOfCantabrianInstitutionsES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := dayOfValencianCommunityES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := galicianLiteratureDayES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := laBienAparecidaES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := nationalDayOfCataloniaES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := santiagoApostolES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := stJohnsDayES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := stJosephsDayES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := stStephensDayES(year); h != nil {
		holiday = append(holiday, *h)
	}
	if h := whitMondayES(year); h != nil {
		holiday = append(holiday, *h)
	}

	return holiday
}

func newYearsDayES(year int) []internal.Holiday {
	base := internal.NewHoliday(year, time.January, 1, "New Year's Day", nil)
	if year == 2023 {
		return []internal.Holiday{
			base,
			internal.NewHoliday(year, time.January, 1, "New Year's Day", []string{"ES-AN", "ES-AR", "ES-AS", "ES-CL", "ES-MC"}),
		}
	}

	return []internal.Holiday{base}
}

func easterMondayES(year int) *internal.Holiday {
	holiday := internal.NewHolidayFromTime(internal.CatholicEasterMonday(year), "Easter Monday", []string{"ES-CT", "ES-IB", "ES-RI", "ES-NC", "ES-PV", "ES-VC"})
	return &holiday
}

func corpusChristiES(year int) *internal.Holiday {
	holiday := internal.NewHolidayFromTime(internal.CatholicCorpusChristi(year), "Corpus Christi", []string{"ES-CM"})
	return &holiday
}

func maundyThursdayES(year int) *internal.Holiday {
	holiday := internal.NewHolidayFromTime(internal.CatholicMaundyThursday(year), "Maundy Thursday", []string{"ES-AN", "ES-AR", "ES-CL", "ES-CM", "ES-CN", "ES-EX", "ES-GA", "ES-IB", "ES-RI", "ES-MD", "ES-MC", "ES-NC", "ES-AS", "ES-PV", "ES-CB"})
	return &holiday
}

func christmasDayES(year int) *internal.Holiday {
	if year == 2022 {
		holiday := internal.NewHoliday(year, time.December, 26, "Christmas Day", []string{"ES-AN", "ES-AR", "ES-AS", "ES-CN", "ES-CB", "ES-CL", "ES-CM", "ES-EX", "ES-GA", "ES-IB", "ES-RI", "ES-MD", "ES-MC", "ES-NC"})
		return &holiday
	}

	holiday := internal.NewHoliday(year, time.December, 25, "Christmas Day", nil)
	return &holiday
}

func whitMondayES(year int) *internal.Holiday {
	if year != 2022 {
		return nil
	}
	holiday := internal.NewHolidayFromTime(internal.CatholicWhitMonday(year), "Whit Monday", []string{"ES-CT"})
	return &holiday
}

func assumptionES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.August, 15, "Assumption", nil)
	return &holiday
}

func dayOfMadridES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.May, 2, "Day of Madrid", []string{"ES-MD"})
	return &holiday
}

func stJosephsDayES(year int) *internal.Holiday {
	if year < 2000 {
		return nil
	}

	var subdivisions []string
	switch year {
	case 2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010, 2011, 2012, 2013, 2014:
		subdivisions = []string{"ES-AR", "ES-CL", "ES-CM", "ES-EX", "ES-GA", "ES-MD", "ES-PV", "ES-VC"}
	case 2015:
		subdivisions = []string{"ES-CM", "ES-MD", "ES-PV", "ES-VC"}
	case 2016:
		subdivisions = []string{"ES-PV", "ES-VC"}
	case 2017:
		subdivisions = []string{"ES-EX", "ES-MD"}
	case 2018, 2019:
		subdivisions = []string{"ES-GA", "ES-PV", "ES-VC"}
	case 2020:
		subdivisions = []string{"ES-CM", "ES-GA", "ES-PV", "ES-VC"}
	case 2021:
		subdivisions = []string{"ES-EX", "ES-GA", "ES-MD", "ES-PV", "ES-VC"}
	case 2022:
		subdivisions = []string{"ES-VC"}
	case 2023:
		subdivisions = []string{"ES-MD"}
	default:
		return nil
	}

	holiday := internal.NewHoliday(year, time.March, 19, "Saint Joseph's Day", subdivisions)
	return &holiday
}

func santiagoApostolES(year int) *internal.Holiday {
	var subdivisions []string
	switch year {
	case 2017:
		subdivisions = []string{"ES-CL", "ES-CN", "ES-GA", "ES-MD", "ES-PV"}
	case 2018:
		subdivisions = []string{"ES-GA"}
	case 2019:
		subdivisions = []string{"ES-GA", "ES-PV"}
	case 2020:
		subdivisions = []string{"ES-GA", "ES-PV"}
	case 2022:
		subdivisions = []string{"ES-GA", "ES-MD", "ES-PV"}
	case 2023:
		subdivisions = []string{"ES-CL", "ES-GA", "ES-NC", "ES-PV"}
	case 2024, 2025, 2026, 2027:
		subdivisions = []string{"ES-GA", "ES-PV"}
	default:
		return nil
	}

	holiday := internal.NewHoliday(year, time.July, 25, "Santiago Apostol", subdivisions)
	return &holiday
}

func stStephensDayES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.December, 26, "St. Stephen's Day", []string{"ES-CT"})
	return &holiday
}

func dayOfValencianCommunityES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.October, 9, "Day of the Valencian Community", []string{"ES-VC"})
	return &holiday
}

func laBienAparecidaES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.September, 15, "Feast of Our Lady of Bien Aparecida", []string{"ES-CB"})
	return &holiday
}

func nationalDayOfCataloniaES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.September, 11, "National Day of Catalonia", []string{"ES-CT"})
	return &holiday
}

func dayOfExtremaduraES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.September, 8, "Day of Extremadura", []string{"ES-EX"})
	return &holiday
}

func dayOfAsturiasES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.September, 8, "Day of Asturias", []string{"ES-AS"})
	return &holiday
}

func dayOfCantabrianInstitutionsES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.July, 28, "Day of the Cantabrian Institutions", []string{"ES-CB"})
	return &holiday
}

func stJohnsDayES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.June, 24, "St. John's Day", []string{"ES-CT", "ES-VC"})
	return &holiday
}

func dayOfLaRiojaES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.June, 9, "Day of La Rioja", []string{"ES-RI"})
	return &holiday
}

func dayOfMurciaES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.June, 9, "Day of Murcia", []string{"ES-MC"})
	return &holiday
}

func dayOfCastillaLaManchaES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.May, 31, "Day of Castilla-La Mancha", []string{"ES-CM"})
	return &holiday
}

func castileAndLeonDayES(year int) *internal.Holiday {
	if year == 2023 {
		return nil
	}
	holiday := internal.NewHoliday(year, time.April, 23, "Castile and Leon Day", []string{"ES-CL"})
	return &holiday
}

func galicianLiteratureDayES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.May, 17, "Galician Literature Day", []string{"ES-GA"})
	return &holiday
}

func dayOfAragonES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.April, 23, "Day of Aragon", []string{"ES-AR"})
	return &holiday
}

func dayOfBalearicIslandsES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.March, 1, "Day of the Balearic Islands", []string{"ES-IB"})
	return &holiday
}

func dayOfAndaluciaES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.February, 28, "Day of Andalucia", []string{"ES-AN"})
	return &holiday
}

func dayOfCanaryIslandsES(year int) *internal.Holiday {
	holiday := internal.NewHoliday(year, time.May, 30, "Day of the Canary Islands", []string{"ES-CN"})
	return &holiday
}

func carnivalTuesdayES(year int) *internal.Holiday {
	if year != 2023 {
		return nil
	}
	holiday := internal.NewHoliday(year, time.February, 21, "Carnival Tuesday", []string{"ES-EX"})
	return &holiday
}
