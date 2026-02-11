package holidays_test

import (
	"testing"
	"time"

	"github.com/bobadilla-tech/holidays-per-country"
)

func TestAutoLoad_IsHoliday(t *testing.T) {
	date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	// This should automatically load the US provider
	result := holidays.IsHoliday(date, "US")

	if !result {
		t.Error("Expected New Year's Day to be a holiday in US")
		t.Error("Auto-loading did not work!")
	}
}

func TestAutoLoad_GetHolidays(t *testing.T) {
	result := holidays.GetHolidays("CA", 2024)

	if len(result) == 0 {
		t.Error("Expected holidays for Canada 2024")
		t.Error("Auto-loading did not work!")
	}

	// Verify we got valid holidays
	foundNewYear := false
	for _, h := range result {
		if h.Date.Month() == time.January && h.Date.Day() == 1 {
			foundNewYear = true
			break
		}
	}

	if !foundNewYear {
		t.Error("Expected to find New Year's Day in Canadian holidays")
	}
}

func TestAutoLoad_GetHolidaysInRange(t *testing.T) {
	start := time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)

	result := holidays.GetHolidaysInRange("GB", start, end)

	// Should find at least Christmas
	if len(result) == 0 {
		t.Error("Expected holidays in December for GB")
		t.Error("Auto-loading did not work!")
	}
}

// TestAutoLoad_MultipleCountries verifies auto-loading works for multiple countries
func TestAutoLoad_MultipleCountries(t *testing.T) {
	// Note: IN (India) is excluded as it's a stub provider with no holidays
	countries := []string{"AU", "BR", "CN", "DE", "ES", "FR", "JP", "MX"}

	for _, country := range countries {
		t.Run(country, func(t *testing.T) {
			result := holidays.GetHolidays(country, 2024)

			if len(result) == 0 {
				t.Errorf("Expected holidays for %s, got empty result", country)
			}
		})
	}
}
