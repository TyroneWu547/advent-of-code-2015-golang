package day06

import (
	utils "aoc/utils"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

func setCoord(tokens []string) coord {
	x, _ := strconv.Atoi(tokens[0])
	y, _ := strconv.Atoi(tokens[1])
	return coord{x, y}
}

func parseCoord(tokens []string) (coord, coord) {
	var cStart coord
	var cEnd coord
	if tokens[0] == "toggle" {
		cStart = setCoord(strings.Split(tokens[1], ","))
		cEnd = setCoord(strings.Split(tokens[3], ","))
	} else {
		cStart = setCoord(strings.Split(tokens[2], ","))
		cEnd = setCoord(strings.Split(tokens[4], ","))
	}

	return cStart, cEnd
}

func partOne(input string) string {
	var grid [1000][1000]bool

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var tokens []string = strings.Split(line, " ")
		cStart, cEnd := parseCoord(tokens[:])

		for i := cStart.y; i <= cEnd.y; i++ {
			var slice []bool = grid[i][:]
			for j := cStart.x; j <= cEnd.x; j++ {
				if tokens[0] == "toggle" {
					slice[j] = !slice[j]
				} else if tokens[1] == "on" {
					slice[j] = true
				} else {
					slice[j] = false
				}
			}
		}
	}

	var lit uint64 = 0
	for _, row := range grid {
		for _, col := range row {
			if col {
				lit++
			}
		}
	}
	return strconv.FormatUint(lit, 10)
}

func partTwo(input string) string {
	var grid [1000][1000]int

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var tokens []string = strings.Split(line, " ")
		cStart, cEnd := parseCoord(tokens[:])

		for i := cStart.y; i <= cEnd.y; i++ {
			var slice []int = grid[i][:]
			for j := cStart.x; j <= cEnd.x; j++ {
				if tokens[0] == "toggle" {
					slice[j] += 2
				} else if tokens[1] == "on" {
					slice[j]++
				} else {
					if slice[j] == 0 {
						continue
					}
					slice[j]--
				}
			}
		}
	}

	var lit int = 0
	for _, row := range grid {
		for _, col := range row {
			lit += col
		}
	}
	return strconv.Itoa(lit)
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  6,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 06", &day)
}
