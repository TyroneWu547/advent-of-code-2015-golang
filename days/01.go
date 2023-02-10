package days

import (
	"strconv"
)

func partOne(input string) string {
	var floor int = 0
	for _, r := range input {
		if r == '(' {
			floor++
		} else {
			floor--
		}
	}
	return strconv.Itoa(floor)
}

func partTwo(input string) string {
	var floor int = 0
	for i, r := range input {
		if r == '(' {
			floor++
		} else {
			floor--
		}

		if floor == -1 {
			return strconv.Itoa(i + 1)
		}
	}
	return ""
}

var day01 Day = Day{
	DayNum:  1,
	PartOne: partOne,
	PartTwo: partTwo,
}

func init() {
	RegisterDay("Day 01", day01)
}
