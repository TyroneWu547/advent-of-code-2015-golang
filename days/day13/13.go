package day13

import (
	utils "aoc/utils"
	"strconv"
	"strings"
)

type happyState struct {
	person  int
	visited int
}

func addPerson(perIdx map[string]int, person string, idx *int) {
	_, ok := perIdx[person]
	if !ok {
		perIdx[person] = *idx
		*idx++
	}
}

func generatePersonMap(lines []string) map[string]int {
	var personIdx map[string]int = make(map[string]int, len(lines))
	var i int = 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		var tokens []string = strings.Split(line, " ")
		addPerson(personIdx, tokens[0], &i)
		addPerson(personIdx, tokens[10][:(len(tokens[10])-1)], &i)
	}

	return personIdx
}

func generateAdjMat(personIdx map[string]int, lines []string) ([][]int, int) {
	var adjMat [][]int = make([][]int, len(personIdx))
	for i := range adjMat {
		adjMat[i] = make([]int, len(personIdx))
	}

	for _, line := range lines {
		if line == "" {
			continue
		}

		var tokens []string = strings.Split(line, " ")
		hap, _ := strconv.Atoi(tokens[3])
		if tokens[2] == "lose" {
			hap *= -1
		}

		var p1 int = personIdx[tokens[0]]
		var p2 int = personIdx[tokens[10][:(len(tokens[10])-1)]]

		adjMat[p1][p2] += hap
		adjMat[p2][p1] += hap
	}

	var minHap int = 0
	for i := 0; i < (len(adjMat) - 1); i++ {
		for j := (i + 1); j < len(adjMat); j++ {
			if adjMat[i][j] < minHap {
				minHap = adjMat[i][j]
			}
		}
	}
	if minHap < 0 {
		minHap *= -1
		for i := 0; i < (len(adjMat) - 1); i++ {
			for j := (i + 1); j < len(adjMat); j++ {
				adjMat[i][j] += minHap
				adjMat[j][i] += minHap
			}
		}
	}

	return adjMat, minHap
}

func dfs(current int, visited int, adjMat [][]int, cache map[happyState]int) int {
	var state happyState = happyState{person: current, visited: visited}
	val, ok := cache[state]
	if ok {
		return val
	}

	if visited == ((1 << len(adjMat)) - 1) {
		return adjMat[current][0]
	}

	var maxHap int = 0
	for i := 0; i < len(adjMat); i++ {
		if i == current {
			continue
		}

		var bitRep int = 1 << i
		if (visited & bitRep) != 0 {
			continue
		}

		var newVisited int = visited | bitRep
		var cumHap int = adjMat[current][i] + dfs(i, newVisited, adjMat, cache)
		if maxHap < cumHap {
			maxHap = cumHap
		}
	}

	cache[state] = maxHap
	return maxHap
}

func partOne(input string) string {
	var lines []string = strings.Split(input, "\n")

	var personIdx map[string]int = generatePersonMap(lines)
	adjMat, minHap := generateAdjMat(personIdx, lines)

	var cache map[happyState]int = map[happyState]int{}
	var maxHap int = dfs(0, 1, adjMat, cache)
	return strconv.Itoa(maxHap - (minHap * len(adjMat)))
}

func partTwo(input string) string {
	var lines []string = strings.Split(input, "\n")

	var personIdx map[string]int = generatePersonMap(lines)
	personIdx["Me"] = len(personIdx)
	adjMat, minHap := generateAdjMat(personIdx, lines)

	var cache map[happyState]int = map[happyState]int{}
	var maxHap int = dfs(0, 1, adjMat, cache)
	return strconv.Itoa(maxHap - (minHap * len(adjMat)))
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  13,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 13", &day)
}
