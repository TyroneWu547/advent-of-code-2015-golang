package day11

import (
	utils "aoc/utils"
)

func incrementPassword(password []rune, incIdx int) {
	if password[incIdx] == 'z' {
		password[incIdx] = 'a'
		incrementPassword(password, incIdx-1)
	} else {
		password[incIdx]++
	}
}

func checkPassword(password []rune) bool {
	for _, r := range password {
		if r == 'i' || r == 'o' || r == 'l' {
			return false
		}
	}

	var hasIncr bool = false
	var lim int = len(password) - 3
	for i := 0; i < lim; i++ {
		hasIncr = true
		for j := 0; j < 2; j++ {
			if (password[i+j] + 1) != password[i+j+1] {
				hasIncr = false
				break
			}
		}
		if hasIncr {
			break
		}
	}
	if !hasIncr {
		return hasIncr
	}

	var pairCount int = 0
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			pairCount++
			i++
		}

		if pairCount >= 2 {
			return true
		}
	}

	return false
}

func newPassword(password []rune) []rune {
	var incIdx int = len(password) - 1
	for !checkPassword(password) {
		incrementPassword(password, incIdx)
	}

	return password
}

func partOne(input string) string {
	var oldPass []rune = []rune(input)
	return string(newPassword(oldPass))
}

func partTwo(input string) string {
	var oldPass []rune = []rune(input)

	var prevPass []rune = newPassword(oldPass)
	incrementPassword(prevPass, len(prevPass)-1)

	return string(newPassword(prevPass))
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  11,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 11", &day)
}
