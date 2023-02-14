package day09

import (
	utils "aoc/utils"
	"strconv"
	"strings"
)

type locState struct {
	currentlocation int
	visited         int
}

func addLoc(locIdx map[string]int, idxLoc map[int]string, loc string, idx *int) {
	_, ok := locIdx[loc]
	if !ok {
		locIdx[loc] = *idx
		idxLoc[*idx] = loc
		*idx++
	}
}

func dfs(current int, visited int, adjMat [][]int, cache map[locState]int, partTwo bool) int {
	var state locState = locState{currentlocation: current, visited: visited}
	distance, ok := cache[state]
	if ok {
		return distance
	}

	if partTwo {
		distance = 0
	} else {
		distance = int(^uint32(0))
	}

	for i := 0; i < len(adjMat); i++ {
		travelDist := adjMat[current][i]
		if travelDist == 0 {
			continue
		}

		var locationBitRep int = 1 << i
		if (visited & locationBitRep) != 0 {
			continue
		}

		var newVisited int = visited | locationBitRep
		var newDist int = travelDist
		if newVisited != ((1 << len(adjMat)) - 1) {
			newDist += dfs(i, newVisited, adjMat, cache, partTwo)
		}

		if partTwo && distance < newDist {
			distance = newDist
		} else if !partTwo && distance > newDist {
			distance = newDist
		}
	}

	cache[state] = distance
	return distance
}

func travelingSalesman(input string, partTwo bool) int {
	var lines []string = strings.Split(input, "\n")

	var locIdx map[string]int = make(map[string]int, len(lines))
	var idxLoc map[int]string = make(map[int]string, len(lines))
	var i int = 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		var tokens []string = strings.Split(line, " ")
		addLoc(locIdx, idxLoc, tokens[0], &i)
		addLoc(locIdx, idxLoc, tokens[2], &i)
	}

	var adjMat [][]int = make([][]int, len(locIdx))
	for i := range adjMat {
		adjMat[i] = make([]int, len(locIdx))
	}

	for _, line := range lines {
		if line == "" {
			continue
		}

		var tokens []string = strings.Split(line, " ")
		var s int = locIdx[tokens[0]]
		var e int = locIdx[tokens[2]]

		dist, _ := strconv.Atoi(tokens[4])
		adjMat[s][e] = dist
		adjMat[e][s] = dist
	}

	var distance int
	if partTwo {
		distance = 0
	} else {
		distance = int(^uint32(0))
	}

	var cache map[locState]int = map[locState]int{}
	for i := 0; i < len(adjMat); i++ {
		var locBitRep int = 1 << i
		var travelDist int = dfs(i, locBitRep, adjMat, cache, partTwo)

		if partTwo && distance < travelDist {
			distance = travelDist
		} else if !partTwo && distance > travelDist {
			distance = travelDist
		}
	}

	return distance
}

func partOne(input string) string {
	return strconv.Itoa(travelingSalesman(input, false))
}

func partTwo(input string) string {
	return strconv.Itoa(travelingSalesman(input, true))
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  9,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 09", &day)
}
