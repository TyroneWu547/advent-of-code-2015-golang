package day16

import (
	utils "aoc/utils"
	"strconv"
	"strings"
)

var key map[string]int = map[string]int{
	"children:":    3,
	"cats:":        7,
	"samoyeds:":    2,
	"pomeranians:": 3,
	"akitas:":      0,
	"vizslas:":     0,
	"goldfish:":    5,
	"trees:":       3,
	"cars:":        2,
	"perfumes:":    1,
}

func partOne(input string) string {
	for i, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var split []string = strings.Split(line, ", ")
		var match bool = true
		for _, s := range split {
			var tokens []string = strings.Split(s, " ")

			val, _ := strconv.Atoi(tokens[len(tokens)-1])
			if key[tokens[len(tokens)-2]] != val {
				match = false
				break
			}
		}

		if match {
			return strconv.Itoa(i + 1)
		}
	}
	return "none"
}

func partTwo(input string) string {
	for i, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		var split []string = strings.Split(line, ", ")
		var match bool = true
		for _, s := range split {
			var tokens []string = strings.Split(s, " ")
			var k string = tokens[len(tokens)-2]
			val, _ := strconv.Atoi(tokens[len(tokens)-1])

			if (k == "cats:" || k == "trees:") && key[k] >= val {
				match = false
				break
			} else if (k == "pomeranians:" || k == "goldfish:") && key[k] <= val {
				match = false
				break
			} else if (k != "cats:" && k != "trees:" && k != "pomeranians:" && k != "goldfish:") && key[k] != val {
				match = false
				break
			}
		}

		if match {
			return strconv.Itoa(i + 1)
		}
	}
	return "none"
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  16,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 16", &day)
}
