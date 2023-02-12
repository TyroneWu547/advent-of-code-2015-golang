package day08

import (
	utils "aoc/utils"
	"strconv"
	"strings"
)

func partOne(input string) string {
	var total int = 0
	var memTotal int = 0

	for _, line := range strings.Split(input, "\n") {
		var maxIdx int = len(line) - 1
		var memCount int = 0
		for i := 1; i < maxIdx; i++ {
			if line[i] == '\\' {
				var nextRune rune = rune(line[i+1])
				if nextRune == '\\' || nextRune == '"' {
					i++
				} else {
					i += 3
				}
			}

			memCount++
		}

		total += len(line)
		memTotal += memCount
	}

	return strconv.Itoa(total - memTotal)
}

func partTwo(input string) string {
	var encodedTotal int = 0
	var originTotal int = 0

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var encodedCount int = 0
		for _, r := range line {
			if r == '"' || r == '\\' {
				encodedCount += 2
			} else {
				encodedCount++
			}
		}

		encodedTotal += encodedCount + 2
		originTotal += len(line)
	}

	return strconv.Itoa(encodedTotal - originTotal)
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  8,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 08", &day)
}
