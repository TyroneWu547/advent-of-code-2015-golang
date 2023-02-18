package day17

import (
	utils "aoc/utils"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	var lines []string = strings.Split(input, "\n")
	var containers []int = make([]int, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		val, _ := strconv.Atoi(line)
		containers = append(containers, val)
	}

	return containers
}

func partOne(input string) string {
	const capacity int = 150
	var containers []int = parseInput(input)

	var max int = (1 << len(containers)) - 1
	var combos int = 0
	for i := 0; i < max; i++ {
		var sum int = 0
		var idx int = 0
		for j := i; j > 0; j /= 2 {
			if (j % 2) == 1 {
				sum += containers[idx]

				if sum > capacity {
					break
				}
			}
			idx++
		}

		if sum == capacity {
			combos++
		}
	}

	return strconv.Itoa(combos)
}

func partTwo(input string) string {
	const capacity int = 150
	var containers []int = parseInput(input)

	var max int = (1 << len(containers)) - 1
	var combos []int = []int{}
	var minContainers int = len(containers)
	for i := 0; i < max; i++ {
		var sum int = 0
		var idx int = 0
		var containerCount int = 0
		for j := i; j > 0; j /= 2 {
			if (j % 2) == 1 {
				sum += containers[idx]
				containerCount++

				if sum > capacity {
					break
				}
			}
			idx++
		}

		if sum == capacity {
			combos = append(combos, i)

			if minContainers > containerCount {
				minContainers = containerCount
			}
		}
	}

	var minCount int = 0
	for _, c := range combos {
		var onBits int = 0
		for i := c; i > 0; i /= 2 {
			onBits += i & 1
		}

		if onBits == minContainers {
			minCount++
		}
	}

	return strconv.Itoa(minCount)
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  17,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 17", &day)
}
