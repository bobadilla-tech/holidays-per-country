// Package internal provides shared utilities for calculating holidays across different countries.
// This file contains functions for calculating Catholic/Christian holiday dates.
package internal

import (
	"time"
)

var cache = map[int]time.Time{}

// CatholicEasterSunday calculates the date of Easter Sunday for a given year using the Meeus/Jones/Butcher algorithm.
func CatholicEasterSunday(year int) time.Time {
	if cached, ok := cache[year]; ok {
		return cached
	}

	// Meeus/Jones/Butcher algorithm
	g := year % 19
	c := year / 100
	h := (c - c/4 - (8*c+13)/25 + 19*g + 15) % 30
	i := h - h/28*(1-h/28*(29/(h+1))*((21-g)/11))

	day := i - (year+year/4+i+2-c+c/4)%7 + 28
	month := 3

	if day > 31 {
		month++
		day -= 31
	}

	result := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	cache[year] = result
	return result
}

// CatholicGoodFriday returns Good Friday (Easter - 2 days).
func CatholicGoodFriday(year int) time.Time {
	return CatholicEasterSunday(year).AddDate(0, 0, -2)
}

// CatholicEasterSaturday returns Holy Saturday (Easter - 1 day).
func CatholicEasterSaturday(year int) time.Time {
	return CatholicEasterSunday(year).AddDate(0, 0, -1)
}

// CatholicEasterMonday returns Easter Monday (Easter + 1 day).
func CatholicEasterMonday(year int) time.Time {
	return CatholicEasterSunday(year).AddDate(0, 0, 1)
}

// CatholicMaundyThursday returns Maundy Thursday (Easter - 3 days).
func CatholicMaundyThursday(year int) time.Time {
	return CatholicEasterSunday(year).AddDate(0, 0, -3)
}

// CatholicAscensionDay returns Ascension Day (Easter + 39 days).
func CatholicAscensionDay(year int) time.Time {
	return CatholicEasterSunday(year).AddDate(0, 0, 39)
}

// CatholicPentecost returns Pentecost (Easter + 49 days).
func CatholicPentecost(year int) time.Time {
	return CatholicEasterSunday(year).AddDate(0, 0, 49)
}

// CatholicWhitMonday returns Whit Monday (Easter + 50 days).
func CatholicWhitMonday(year int) time.Time {
	return CatholicEasterSunday(year).AddDate(0, 0, 50)
}

// CatholicCorpusChristi returns Corpus Christi (Easter + 60 days).
func CatholicCorpusChristi(year int) time.Time {
	return CatholicEasterSunday(year).AddDate(0, 0, 60)
}
