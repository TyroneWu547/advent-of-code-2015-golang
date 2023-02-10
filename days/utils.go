package days

type Task func(string) string

type Day struct {
	DayNum  int
	PartOne Task
	PartTwo Task
}

// Holds the functions for each aoc day
var AOCDays map[string]Day = map[string]Day{}

// Adds the day to the map
func RegisterDay(dayFmt string, day Day) {
	AOCDays[dayFmt] = day
}
