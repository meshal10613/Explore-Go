package main

import "fmt"

type Weekday int

const (
	SUNDAY Weekday = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

func getWorkDayStatus(day Weekday) string {
	switch day {
	case SUNDAY, MONDAY, TUESDAY, WEDNESDAY:
		return "Office is open"
	case THURSDAY:
		return "Office is half open"
	case FRIDAY, SATURDAY:
		return "Office is closed"
	default:
		return "Office is closed"
	}
}

func main() {
	//? Enum
	fmt.Println(getWorkDayStatus(SUNDAY))
}
