package holidays_test

import (
	"testing"
	"time"

	"github.com/bobadilla-tech/holidays-per-country"
)

func TestIsHoliday(t *testing.T) {
	tests := []struct {
		name        string
		date        time.Time
		countryCode string
		expected    bool
	}{
		{
			name:        "New Year's Day in US",
			date:        time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			countryCode: "US",
			expected:    true,
		},
		{
			name:        "Christmas Day in US",
			date:        time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC),
			countryCode: "US",
			expected:    true,
		},
		{
			name:        "Independence Day in US",
			date:        time.Date(2024, 7, 4, 0, 0, 0, 0, time.UTC),
			countryCode: "US",
			expected:    true,
		},
		{
			name:        "Regular day in US",
			date:        time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC),
			countryCode: "US",
			expected:    false,
		},
		{
			name:        "Canada Day in Canada",
			date:        time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC),
			countryCode: "CA",
			expected:    true,
		},
		{
			name:        "Canada Day not a holiday in US",
			date:        time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC),
			countryCode: "US",
			expected:    false,
		},
		{
			name:        "Christmas Day in UK",
			date:        time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC),
			countryCode: "GB",
			expected:    true,
		},
		{
			name:        "New Year's Day in Japan",
			date:        time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			countryCode: "JP",
			expected:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Auto-loading enabled - LazyLoad no longer needed
			result := holidays.IsHoliday(tt.date, tt.countryCode)
			if result != tt.expected {
				t.Errorf("IsHoliday(%v, %s) = %v, want %v",
					tt.date.Format("2006-01-02"), tt.countryCode, result, tt.expected)
			}
		})
	}
}

func TestGetHolidays(t *testing.T) {
	tests := []struct {
		name        string
		countryCode string
		year        int
		minCount    int // Minimum number of holidays expected
	}{
		{
			name:        "US holidays for 2024",
			countryCode: "US",
			year:        2024,
			minCount:    10, // Federal holidays
		},
		{
			name:        "Canada holidays for 2024",
			countryCode: "CA",
			year:        2024,
			minCount:    5, // National holidays
		},
		{
			name:        "UK holidays for 2024",
			countryCode: "GB",
			year:        2024,
			minCount:    8,
		},
		{
			name:        "Japan holidays for 2024",
			countryCode: "JP",
			year:        2024,
			minCount:    15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := holidays.GetHolidays(tt.countryCode, tt.year)

			if len(result) < tt.minCount {
				t.Errorf("GetHolidays(%s, %d) returned %d holidays, want at least %d",
					tt.countryCode, tt.year, len(result), tt.minCount)
			}

			// Verify all holidays are in the correct year
			for _, holiday := range result {
				if holiday.Date.Year() != tt.year {
					t.Errorf("Holiday %s has year %d, want %d",
						holiday.Name, holiday.Date.Year(), tt.year)
				}
			}

			// Verify holidays are sorted by date
			for i := 1; i < len(result); i++ {
				if result[i].Date.Before(result[i-1].Date) {
					t.Errorf("Holidays not in chronological order: %s (%v) comes after %s (%v)",
						result[i].Name, result[i].Date, result[i-1].Name, result[i-1].Date)
				}
			}
		})
	}
}

func TestGetHolidays_SpecificDates(t *testing.T) {
	usHolidays := holidays.GetHolidays("US", 2024)

	expectedHolidays := map[string]string{
		"2024-01-01": "New Year's Day",
		"2024-07-04": "Independence Day",
		"2024-12-25": "Christmas Day",
	}

	for dateStr, expectedName := range expectedHolidays {
		found := false
		for _, holiday := range usHolidays {
			if holiday.Date.Format("2006-01-02") == dateStr {
				found = true
				if holiday.Name != expectedName {
					t.Errorf("Holiday on %s is named %q, want %q",
						dateStr, holiday.Name, expectedName)
				}
				break
			}
		}
		if !found {
			t.Errorf("Expected holiday %q on %s not found", expectedName, dateStr)
		}
	}
}

func TestGetHolidaysInRange(t *testing.T) {
	tests := []struct {
		name        string
		countryCode string
		start       time.Time
		end         time.Time
		expected    int // Expected number of holidays
	}{
		{
			name:        "US holidays in January 2024",
			countryCode: "US",
			start:       time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			end:         time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
			expected:    2, // New Year's Day + MLK Day
		},
		{
			name:        "US holidays in Q4 2024",
			countryCode: "US",
			start:       time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC),
			end:         time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
			expected:    5, // Veterans Day, Thanksgiving, Christmas, Columbus Day, Indigenous Peoples' Day
		},
		{
			name:        "Single day range with holiday",
			countryCode: "US",
			start:       time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC),
			end:         time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC),
			expected:    1, // Christmas
		},
		{
			name:        "Single day range without holiday",
			countryCode: "US",
			start:       time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC),
			end:         time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC),
			expected:    0,
		},
		{
			name:        "Multi-year range",
			countryCode: "US",
			start:       time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
			end:         time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
			expected:    3, // Christmas 2023, New Year 2024, MLK 2024
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := holidays.GetHolidaysInRange(tt.countryCode, tt.start, tt.end)

			if len(result) != tt.expected {
				t.Errorf("GetHolidaysInRange(%s, %v, %v) returned %d holidays, want %d",
					tt.countryCode, tt.start.Format("2006-01-02"), tt.end.Format("2006-01-02"),
					len(result), tt.expected)

				// Debug output
				t.Logf("Holidays found:")
				for _, h := range result {
					t.Logf("  - %s: %s", h.Date.Format("2006-01-02"), h.Name)
				}
			}

			// Verify all holidays are within range
			for _, holiday := range result {
				if holiday.Date.Before(tt.start) || holiday.Date.After(tt.end) {
					t.Errorf("Holiday %s on %v is outside range [%v, %v]",
						holiday.Name, holiday.Date, tt.start, tt.end)
				}
			}

			// Verify holidays are sorted
			for i := 1; i < len(result); i++ {
				if result[i].Date.Before(result[i-1].Date) {
					t.Errorf("Holidays not in chronological order")
				}
			}
		})
	}
}

func TestGetHolidaysInRange_InvalidRange(t *testing.T) {
	start := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	result := holidays.GetHolidaysInRange("US", start, end)

	if len(result) != 0 {
		t.Errorf("GetHolidaysInRange with end before start should return empty slice, got %d holidays", len(result))
	}
}

func TestGetHolidaysInRange_BoundaryInclusive(t *testing.T) {
	// Test that start date is inclusive
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC) // New Year's Day
	end := time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)
	result := holidays.GetHolidaysInRange("US", start, end)

	found := false
	for _, h := range result {
		if h.Date.Equal(start) && h.Name == "New Year's Day" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Start date should be inclusive - New Year's Day not found")
	}

	// Test that end date is inclusive
	start = time.Date(2024, 12, 24, 0, 0, 0, 0, time.UTC)
	end = time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC) // Christmas
	result = holidays.GetHolidaysInRange("US", start, end)

	found = false
	for _, h := range result {
		if h.Date.Equal(end) && h.Name == "Christmas Day" {
			found = true
			break
		}
	}
	if !found {
		t.Error("End date should be inclusive - Christmas Day not found")
	}
}

func TestMultipleCountries(t *testing.T) {
	countries := []string{"US", "CA", "GB", "JP", "DE", "FR"}
	year := 2024

	for _, country := range countries {
		t.Run(country, func(t *testing.T) {
			result := holidays.GetHolidays(country, year)

			if len(result) == 0 {
				t.Errorf("GetHolidays(%s, %d) returned no holidays", country, year)
			}

			// Verify New Year's Day exists for all countries
			foundNewYear := false
			for _, h := range result {
				if h.Date.Month() == time.January && h.Date.Day() == 1 {
					foundNewYear = true
					break
				}
			}
			if !foundNewYear {
				t.Errorf("New Year's Day not found for country %s", country)
			}
		})
	}
}

func TestHolidayStructure(t *testing.T) {
	usHolidays := holidays.GetHolidays("US", 2024)

	if len(usHolidays) == 0 {
		t.Fatal("No holidays returned for US 2024")
	}

	for _, holiday := range usHolidays {
		// Verify required fields are not empty
		if holiday.Name == "" {
			t.Errorf("Holiday has empty name on date %v", holiday.Date)
		}

		if holiday.Date.IsZero() {
			t.Errorf("Holiday %s has zero date", holiday.Name)
		}

		// Fixed field should be set (either true or false is valid)
		// Just verify the struct is properly populated
		t.Logf("Holiday: %s on %v (Fixed: %v, Subdivisions: %v)",
			holiday.Name, holiday.Date.Format("2006-01-02"), holiday.Fixed, holiday.Subdivisions)
	}
}

func TestGetHolidays_CachePerYear(t *testing.T) {
	// Get holidays for 2024
	holidays2024 := holidays.GetHolidays("US", 2024)
	if len(holidays2024) == 0 {
		t.Fatal("No holidays returned for US 2024")
	}

	// Get holidays for 2025
	holidays2025 := holidays.GetHolidays("US", 2025)
	if len(holidays2025) == 0 {
		t.Fatal("No holidays returned for US 2025")
	}

	// Verify all 2024 holidays are in 2024
	for _, h := range holidays2024 {
		if h.Date.Year() != 2024 {
			t.Errorf("Holiday %s from GetHolidays(US, 2024) has year %d",
				h.Name, h.Date.Year())
		}
	}

	// Verify all 2025 holidays are in 2025
	for _, h := range holidays2025 {
		if h.Date.Year() != 2025 {
			t.Errorf("Holiday %s from GetHolidays(US, 2025) has year %d",
				h.Name, h.Date.Year())
		}
	}

	// Get 2024 again to verify cache returns correct year
	holidays2024Again := holidays.GetHolidays("US", 2024)
	for _, h := range holidays2024Again {
		if h.Date.Year() != 2024 {
			t.Errorf("Holiday %s from cached GetHolidays(US, 2024) has year %d",
				h.Name, h.Date.Year())
		}
	}
}

func TestGetHolidays_Sorted(t *testing.T) {
	countries := []string{"US", "CA", "GB", "JP", "DE", "FR"}

	for _, country := range countries {
		t.Run(country, func(t *testing.T) {
			result := holidays.GetHolidays(country, 2024)

			if len(result) == 0 {
				t.Fatalf("No holidays returned for %s 2024", country)
			}

			// Verify holidays are in chronological order
			for i := 1; i < len(result); i++ {
				if result[i].Date.Before(result[i-1].Date) {
					t.Errorf("%s: Holidays not sorted - %s (%s) comes before %s (%s)",
						country,
						result[i].Name, result[i].Date.Format("2006-01-02"),
						result[i-1].Name, result[i-1].Date.Format("2006-01-02"))
				}
			}
		})
	}
}
