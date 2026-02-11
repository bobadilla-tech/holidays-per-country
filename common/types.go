package common

import "time"

type Holiday struct {
	Date         time.Time
	Name         string
	Subdivisions []string
	Fixed        bool
}
