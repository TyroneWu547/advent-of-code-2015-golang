package days

type task func(string) (string, error)

type Day struct {
	DayNum  int
	PartOne task
	PartTwo task
}

var AOCDays map[string]Day = map[string]Day{}

func RegisterDay(dayFmt string, day Day) {
	AOCDays[dayFmt] = day
}
