package day05

import (
	utils "aoc/utils"
	"strconv"
	"strings"
)

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}

func isNiceP1(line []rune) bool {
	var vowels uint8 = 0
	var dupl bool = false

	var lim int = len(line) - 1
	for i := 0; i < lim; i++ {
		if (line[i] == 'a' || line[i] == 'c' || line[i] == 'p' || line[i] == 'x') && (line[i+1] == (line[i] + 1)) {
			return false
		}

		if line[i] == line[i+1] {
			dupl = true
		}
		if isVowel(line[i]) {
			vowels++
		}
	}

	if isVowel(line[lim]) {
		vowels++
	}

	return vowels >= 3 && dupl
}

func partOne(input string) string {
	var niceNum uint64 = 0
	for _, line := range strings.Split(input, "\n") {
		if line != "" && isNiceP1([]rune(line)) {
			niceNum++
		}
	}
	return strconv.FormatUint(niceNum, 10)
}

func isNiceP2(line []rune) bool {
	var dupl bool = false
	var pairsMap map[string]bool = make(map[string]bool)
	var tmpPair string = ""
	for i := 2; i <= len(line); i++ {
		var pair string = string(line[(i - 2):i])
		_, ok := pairsMap[pair]
		if ok {
			dupl = true
			break
		}
		pairsMap[tmpPair] = true
		tmpPair = pair
	}

	var lim int = len(line) - 1
	for i := 1; i < lim; i++ {
		if line[i-1] == line[i+1] {
			return dupl
		}
	}

	return false
}

func partTwo(input string) string {
	var niceNum uint64 = 0
	for _, line := range strings.Split(input, "\n") {
		if line != "" && isNiceP2([]rune(line)) {
			niceNum++
		}
	}
	return strconv.FormatUint(niceNum, 10)
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  5,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 05", &day)
}
