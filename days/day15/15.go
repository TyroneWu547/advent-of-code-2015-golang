package day15

import (
	utils "aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func parseInput(input string) []ingredient {
	var lines []string = strings.Split(input, "\n")
	var ingredients []ingredient = make([]ingredient, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		var tokens []string = strings.Split(line, " ")
		var name string = tokens[0]
		capacity, _ := strconv.Atoi(tokens[2][:(len(tokens[2]) - 1)])
		durability, _ := strconv.Atoi(tokens[4][:(len(tokens[4]) - 1)])
		flavor, _ := strconv.Atoi(tokens[6][:(len(tokens[6]) - 1)])
		texture, _ := strconv.Atoi(tokens[8][:(len(tokens[8]) - 1)])
		calories, _ := strconv.Atoi(tokens[10])

		var ingr ingredient = ingredient{name, capacity, durability, flavor, texture, calories}
		ingredients = append(ingredients, ingr)
	}

	return ingredients
}

func calcScore(ingredients []ingredient, counts []int, partTwo bool) int {
	if partTwo {
		var cal int = 0
		for i, c := range counts {
			cal += ingredients[i].calories * c
		}
		if cal != 500 {
			return 0
		}
	}

	var cap int = 0
	for i, c := range counts {
		cap += ingredients[i].capacity * c
	}
	if cap <= 0 {
		return 0
	}

	var dur int = 0
	for i, c := range counts {
		dur += ingredients[i].durability * c
	}
	if dur <= 0 {
		return 0
	}

	var fla int = 0
	for i, c := range counts {
		fla += ingredients[i].flavor * c
	}
	if fla <= 0 {
		return 0
	}

	var tex int = 0
	for i, c := range counts {
		tex += ingredients[i].texture * c
	}
	if tex <= 0 {
		return 0
	}

	return cap * dur * fla * tex
}

func partOne(input string) string {
	const teaspoons int = 100
	var ingredients []ingredient = parseInput(input)

	var maxScore int = 0
	for i := 0; i < teaspoons; i++ {
		for j := 0; j < (teaspoons - i); j++ {
			for k := 0; k < (teaspoons - i - j); k++ {
				var l int = teaspoons - i - j - k
				var score int = calcScore(ingredients, []int{i, j, k, l}, false)

				if maxScore < score {
					maxScore = score
				}
			}
		}
	}

	return strconv.Itoa(maxScore)
}

func partTwo(input string) string {
	const teaspoons int = 100
	var ingredients []ingredient = parseInput(input)

	fmt.Println(ingredients)

	var maxScore int = 0
	for i := 0; i < teaspoons; i++ {
		for j := 0; j < (teaspoons - i); j++ {
			for k := 0; k < (teaspoons - i - j); k++ {
				var l int = teaspoons - i - j - k
				var score int = calcScore(ingredients, []int{i, j, k, l}, true)

				if maxScore < score {
					maxScore = score
				}
			}
		}
	}

	return strconv.Itoa(maxScore)
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  15,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 15", &day)
}
