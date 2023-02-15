package day10

import (
	utils "aoc/utils"
	"strconv"
)

func lookAndSay(seq []int) []int {
	var newSeq []int = make([]int, 0, len(seq))

	var count int = 1
	var prevVal int = seq[0]
	for _, val := range seq[1:] {
		if val == prevVal {
			count++
		} else {
			newSeq = append(newSeq, count, prevVal)
			count = 1
			prevVal = val
		}
	}
	newSeq = append(newSeq, count, prevVal)

	return newSeq
}

func iterateN(input string, n int) int {
	var seq []int = make([]int, 0, len(input))
	for _, r := range input {
		var val int = int(r - '0')
		seq = append(seq, val)
	}

	for i := 0; i < n; i++ {
		seq = lookAndSay(seq)
	}

	return len(seq)
}

func partOne(input string) string {
	return strconv.Itoa(iterateN(input, 40))
}

func partTwo(input string) string {
	return strconv.Itoa(iterateN(input, 50))
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  10,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 10", &day)
}
