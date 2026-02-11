package holidays_test

import (
	"testing"
	"time"

	"github.com/bobadilla-tech/holidays-per-country"
)

func BenchmarkIsHoliday_Cached(b *testing.B) {
	holidays.LazyLoad("US")
	date := time.Date(2024, 7, 4, 0, 0, 0, 0, time.UTC)

	// Pre-warm cache
	holidays.GetHolidays("US", 2024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		holidays.IsHoliday(date, "US")
	}
}

func BenchmarkGetHolidays_Cached(b *testing.B) {
	holidays.LazyLoad("US")

	// Pre-warm cache
	holidays.GetHolidays("US", 2024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		holidays.GetHolidays("US", 2024)
	}
}

func BenchmarkGetHolidays_Uncached(b *testing.B) {
	holidays.LazyLoad("US")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Use different years to avoid cache hits
		year := 2000 + (i % 100)
		holidays.GetHolidays("US", year)
	}
}

func BenchmarkGetHolidaysInRange_OneMonth(b *testing.B) {
	holidays.LazyLoad("US")
	start := time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 7, 31, 0, 0, 0, 0, time.UTC)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		holidays.GetHolidaysInRange("US", start, end)
	}
}

func BenchmarkGetHolidaysInRange_OneYear(b *testing.B) {
	holidays.LazyLoad("US")
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		holidays.GetHolidaysInRange("US", start, end)
	}
}

func BenchmarkConcurrent_IsHoliday(b *testing.B) {
	holidays.LazyLoad("US")
	date := time.Date(2024, 7, 4, 0, 0, 0, 0, time.UTC)

	// Pre-warm cache
	holidays.GetHolidays("US", 2024)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			holidays.IsHoliday(date, "US")
		}
	})
}

func BenchmarkSubdivisionFiltering(b *testing.B) {
	holidays.LazyLoad("AU")

	// Pre-warm cache
	allHolidays := holidays.GetHolidays("AU", 2024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Simulate filtering for Victoria
		filtered := 0
		for _, h := range allHolidays {
			if len(h.Subdivisions) == 0 || contains(h.Subdivisions, "AU-VIC") {
				filtered++
			}
		}
		_ = filtered
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
