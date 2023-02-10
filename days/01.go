package days

import (
	"strconv"
)

func init() {
	var day Day = Day{
		DayNum: 1,

		PartOne: func(input string) string {
			var floor int = 0
			for _, r := range input {
				if r == '(' {
					floor++
				} else {
					floor--
				}
			}
			return strconv.Itoa(floor)
		},

		PartTwo: func(input string) string {
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
		},
	}
	RegisterDay("Day 01", &day)
}
