package day02

import (
	utils "aoc/utils"
	"strconv"
	"strings"
)

func partOne(input string) string {
	var sum uint64 = 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var dimStr []string = strings.Split(line, "x")
		length, _ := strconv.ParseUint(dimStr[0], 10, 8)
		width, _ := strconv.ParseUint(dimStr[1], 10, 8)
		height, _ := strconv.ParseUint(dimStr[2], 10, 8)

		s1 := length * width
		s2 := width * height
		s3 := height * length

		var min uint64 = s1
		if s2 < min {
			min = s2
		}
		if s3 < min {
			min = s3
		}

		sum += 2*(s1+s2+s3) + min
	}

	return strconv.FormatUint(uint64(sum), 10)
}

func partTwo(input string) string {
	var sum uint64 = 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var dimStr []string = strings.Split(line, "x")
		length, _ := strconv.ParseUint(dimStr[0], 10, 8)
		width, _ := strconv.ParseUint(dimStr[1], 10, 8)
		height, _ := strconv.ParseUint(dimStr[2], 10, 8)

		var total uint64 = length + width + height
		var max uint64 = length
		if width > max {
			max = width
		}
		if height > max {
			max = height
		}

		sum += length*width*height + 2*(total-max)
	}

	return strconv.FormatUint(uint64(sum), 10)
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  1,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 02", &day)
}
