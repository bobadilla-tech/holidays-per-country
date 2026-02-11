package holidays

type provider interface {
	RegisterHolidays(year int) []Holiday
}
