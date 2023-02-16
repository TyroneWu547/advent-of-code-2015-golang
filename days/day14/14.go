package day14

import (
	utils "aoc/utils"
	"strconv"
	"strings"
)

type reindeer struct {
	name     string
	speed    int
	duration int
	rest     int
}

func parseInput(input string) []reindeer {
	var lines []string = strings.Split(input, "\n")
	var reindeers []reindeer = make([]reindeer, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}

		var tokens []string = strings.Split(line, " ")
		var name string = tokens[0]
		speed, _ := strconv.Atoi(tokens[3])
		duration, _ := strconv.Atoi(tokens[6])
		rest, _ := strconv.Atoi(tokens[13])

		var r reindeer = reindeer{name, speed, duration, rest}
		reindeers = append(reindeers, r)
	}

	return reindeers
}

func distanceTraveled(r reindeer, seconds int) int {
	var period int = r.duration + r.rest
	var completeCycles int = seconds / period
	var amp int = r.speed * r.duration

	var traveled int = completeCycles * amp

	var timeleft int = seconds - (period * completeCycles)
	if timeleft < r.duration {
		traveled += timeleft * r.speed
	} else {
		traveled += amp
	}

	return traveled
}

func partOne(input string) string {
	var reindeer []reindeer = parseInput(input)

	const seconds int = 2503
	var maxDist int = 0
	for _, r := range reindeer {
		var traveled int = distanceTraveled(r, seconds)
		if traveled > maxDist {
			maxDist = traveled
		}
	}

	return strconv.Itoa(maxDist)
}

func partTwo(input string) string {
	var reindeer []reindeer = parseInput(input)
	var distances []int = make([]int, len(reindeer))
	var scores []int = make([]int, len(reindeer))

	const seconds int = 2503
	for i := 0; i < seconds; i++ {
		for j, r := range reindeer {
			var period int = r.duration + r.rest
			var frame int = i % period
			if frame < r.duration {
				distances[j] += r.speed
			}
		}

		var leadIdx int = 0
		for j, d := range distances {
			if distances[leadIdx] < d {
				leadIdx = j
			}
		}

		for j, d := range distances {
			if distances[leadIdx] == d {
				scores[j]++
			}
		}
	}

	var maxScore int = 0
	for _, s := range scores {
		if maxScore < s {
			maxScore = s
		}
	}
	return strconv.Itoa(maxScore)
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  14,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 14", &day)
}
