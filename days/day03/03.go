package day03

import (
	utils "aoc/utils"
	"strconv"
)

type coord struct {
	x int
	y int
}

func step(santa *coord, move rune) {
	switch move {
	case '^':
		santa.y++
	case 'v':
		santa.y--
	case '>':
		santa.x++
	case '<':
		santa.x--
	}
}

func partOne(input string) string {
	const t = true
	var santa coord = coord{x: 0, y: 0}
	var houseSet map[coord]bool = map[coord]bool{santa: t}

	for _, r := range input {
		step(&santa, r)
		houseSet[santa] = t
	}

	return strconv.Itoa(len(houseSet))
}

func partTwo(input string) string {
	const t = true
	var turn bool = true
	var santa coord = coord{x: 0, y: 0}
	var roboSanta coord = coord{x: 0, y: 0}
	var houseSet map[coord]bool = map[coord]bool{santa: t, roboSanta: t}

	for _, r := range input {
		if turn {
			step(&santa, r)
			houseSet[santa] = t
		} else {
			step(&roboSanta, r)
			houseSet[roboSanta] = t
		}

		turn = !turn
	}

	return strconv.Itoa(len(houseSet))
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  3,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 03", &day)
}
