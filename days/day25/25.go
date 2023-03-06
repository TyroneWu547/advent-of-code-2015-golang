package day25

import (
	utils "aoc/utils"
	"strconv"
	"strings"
)

func partOne(input string) string {
	var tokens []string = strings.Split(input, " ")
	row, _ := strconv.Atoi(tokens[16][:(len(tokens[16]) - 1)])
	col, _ := strconv.Atoi(tokens[18][:(len(tokens[18]) - 2)])

	var tri int = row + col - 2
	var order int = (((tri + 1) * tri) / 2) + col

	const mul uint64 = 252533
	const div uint64 = 33554393

	var code uint64 = 20151125
	for i := 0; i < (order - 1); i++ {
		code = (code * mul) % div
	}

	return strconv.Itoa(int(code))
}

func partTwo(input string) string {
	return "none"
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  25,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 25", &day)
}
