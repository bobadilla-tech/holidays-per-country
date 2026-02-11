package holidays

import "github.com/bobadilla-tech/holidays-per-country/common"

type provider interface {
	RegisterHolidays(year int) []common.Holiday
}
