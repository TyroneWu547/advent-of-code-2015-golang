package day24

import (
	utils "aoc/utils"
	"strconv"
	"strings"
)

type partition struct {
	comb []int
	rest []int
}

func parseInput(input *string) []int {
	var lines []string = strings.Split(*input, "\n")
	var pkgs []int = make([]int, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}

		val, _ := strconv.Atoi(line)
		pkgs = append(pkgs, val)
	}

	return pkgs
}

func remove(i int, a1 []int) (int, []int) {
	var a2 []int = make([]int, len(a1))
	copy(a2, a1)

	var element int = a2[i]
	a2[i] = a2[len(a2)-1]
	return element, a2[:(len(a2) - 1)]
}

func sum(list []int) int {
	var sum int = 0
	for _, e := range list {
		sum += e
	}
	return sum
}

func product(list []int) int {
	var product int = 1
	for _, e := range list {
		product *= e
	}
	return product
}

// func assignRest(start int, r int, weight int, size int, pkgs []int, target int) bool {
// 	if size == r {
// 		return target == sum(pkgs) && target == weight
// 	}

// 	for i := start; i < len(pkgs); i++ {
// 		val, result := remove(i, pkgs)
// 		if assignRest(i, r, (weight + val), (size + 1), result, target) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func assignFirstGroup(start int, r int, weight int, qe int, size int, pkgs []int, target int) int {
// 	if size == r {
// 		if weight == target {
// 			for j := 1; j < len(pkgs); j++ {
// 				if assignRest(0, j, 0, 0, pkgs, weight) {
// 					return qe
// 				}
// 			}
// 		}
// 		return 0
// 	}

// 	var minQE int = int(^uint32(0))
// 	for i := start; i < len(pkgs); i++ {
// 		val, result := remove(i, pkgs)
// 		var qe int = assignFirstGroup(i, r, (weight + val), (qe * val), (size + 1), result, target)
// 		if qe != 0 && qe < minQE {
// 			minQE = qe
// 		}
// 	}
// 	return minQE
// }

func factorial(start int, end int) int {
	for i := start + 1; i <= end; i++ {
		start *= i
	}
	return start
}

func nCr(n int, r int) int {
	if (n / 2) < r {
		return factorial(r+1, n) / factorial(1, n-r)
	}
	return factorial(n-r+1, n) / factorial(1, r)
}

func getCombinations(start int, r int, part []int, rest []int, combos *[]partition) {
	if len(part) == r {
		var partCopy []int = make([]int, len(part))
		copy(partCopy, part)
		*combos = append(*combos, partition{partCopy, rest})
	}

	for i := start; i < len(rest); i++ {
		val, result := remove(i, rest)
		getCombinations(i, r, append(part, val), result, combos)
	}
}

func partOne(input string) string {
	var pkgs []int = parseInput(&input)
	var total int = sum(pkgs)

	for i := 1; i <= len(pkgs); i++ {
		var g1Part []partition = make([]partition, 0, nCr(len(pkgs), i))
		getCombinations(0, i, make([]int, 0, i), pkgs, &g1Part)

		for _, g1 := range g1Part {
			if (sum(g1.comb) * 3) == total {
				for j := 1; j < len(g1.rest); j++ {
					var g2Part []partition = make([]partition, 0, nCr(len(g1.rest), j))
					getCombinations(0, j, make([]int, 0, j), g1.rest, &g2Part)

					for _, g2 := range g2Part {
						if sum(g2.comb) == sum(g2.rest) {
							return strconv.Itoa(product(g1.comb))
						}
					}
				}
			}
		}
	}

	return "none"
}

func partTwo(input string) string {
	var pkgs []int = parseInput(&input)
	var total int = sum(pkgs)

	for i := 1; i <= len(pkgs); i++ {
		var g1Part []partition = make([]partition, 0, nCr(len(pkgs), i))
		getCombinations(0, i, make([]int, 0, i), pkgs, &g1Part)

		for _, g1 := range g1Part {
			if (sum(g1.comb) * 4) == total {

				for j := 1; j < len(g1.rest); j++ {
					var g2Part []partition = make([]partition, 0, nCr(len(g1.rest), j))
					getCombinations(0, j, make([]int, 0, j), g1.rest, &g2Part)

					for _, g2 := range g2Part {
						if (sum(g2.comb) * 4) == total {

							for k := 1; k < len(g2.rest); k++ {
								var g3Part []partition = make([]partition, 0, nCr(len(g2.rest), k))
								getCombinations(0, k, make([]int, 0, k), g2.rest, &g3Part)

								for _, g3 := range g3Part {
									if sum(g3.comb) == sum(g3.rest) {
										return strconv.Itoa(product(g1.comb))
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return "none"
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  24,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 24", &day)
}
