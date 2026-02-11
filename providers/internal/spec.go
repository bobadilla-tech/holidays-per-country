package internal

import (
	"time"
)

func NewHoliday(year int, month time.Month, day int, name string) Holiday {
	return Holiday{
		Date:  time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
		Name:  name,
		Fixed: true,
	}
}

func NewHolidayFromTime(date time.Time, name string) Holiday {
	return Holiday{
		Date:  date,
		Name:  name,
		Fixed: false,
	}
}
