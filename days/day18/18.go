package day18

import (
	utils "aoc/utils"
	"strconv"
	"strings"
)

func parseInput(input string) [][]bool {
	var lines []string = strings.Split(input, "\n")
	var grid [][]bool = make([][]bool, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}

		var row []bool = make([]bool, len(line))
		for j, r := range line {
			if r == '#' {
				row[j] = true
			}
		}

		grid = append(grid, row)
	}

	return grid
}

func neighborsLit(x int, y int, grid [][]bool) int {
	var count int = 0
	for i := x - 1; i <= (x + 1); i++ {
		if i < 0 || i >= len(grid) {
			continue
		}

		for j := y - 1; j <= (y + 1); j++ {
			if (j < 0 || j >= len(grid[i])) || (i == x && j == y) {
				continue
			}

			if grid[j][i] {
				count++
			}
		}
	}

	return count
}

func animate(grid [][]bool) [][]bool {
	var newGrid [][]bool = make([][]bool, len(grid))
	for i, _ := range newGrid {
		newGrid[i] = make([]bool, len(grid[i]))
	}

	for i, row := range grid {
		for j, light := range row {
			var neighborsLit int = neighborsLit(j, i, grid)
			if light {
				if neighborsLit == 2 || neighborsLit == 3 {
					newGrid[i][j] = true
				}
			} else {
				if neighborsLit == 3 {
					newGrid[i][j] = true
				}
			}
		}
	}

	return newGrid
}

func partOne(input string) string {
	var grid [][]bool = parseInput(input)
	for i := 0; i < 100; i++ {
		grid = animate(grid)
	}

	var lit int = 0
	for _, row := range grid {
		for _, light := range row {
			if light {
				lit++
			}
		}
	}

	return strconv.Itoa(lit)
}

func partTwo(input string) string {
	var grid [][]bool = parseInput(input)
	var ySize int = len(grid) - 1
	var xSize int = len(grid[0]) - 1
	for i := 0; i < 100; i++ {
		grid = animate(grid)

		grid[0][0] = true
		grid[ySize][0] = true
		grid[0][xSize] = true
		grid[ySize][xSize] = true
	}

	var lit int = 0
	for _, row := range grid {
		for _, light := range row {
			if light {
				lit++
			}
		}
	}

	return strconv.Itoa(lit)
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  18,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 18", &day)
}
