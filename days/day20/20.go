package day20

import (
	utils "aoc/utils"
	"strconv"
)

func partOne(input string) string {
	target, _ := strconv.Atoi(input)
	target /= 10

	var houses []int = make([]int, target)

	for i := 1; i <= target; i++ {
		for j := i; j <= target; j += i {
			houses[j-1] += i
		}

		if houses[i-1] >= target {
			return strconv.Itoa(i)
		}
	}
	return "none"
}

func partTwo(input string) string {
	target, _ := strconv.Atoi(input)

	var houses []int = make([]int, target)

	for i := 1; i <= target; i++ {
		var h int = i
		for j := 0; (j < 50) && (h <= target); j++ {
			houses[h-1] += i * 11
			h += i
		}

		if houses[i-1] >= target {
			return strconv.Itoa(i)
		}
	}
	return "none"
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  20,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 20", &day)
}
