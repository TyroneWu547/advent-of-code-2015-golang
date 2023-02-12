package day07

import (
	utils "aoc/utils"
	"strconv"
	"strings"
)

func ast(variables map[string]int, instructions map[string][]string, operand string) int {
	val, err := strconv.Atoi(operand)
	if err == nil {
		return val
	}

	val, ok := variables[operand]
	if ok {
		return val
	}

	var result int
	var instr []string = instructions[operand]
	if len(instr) == 2 {
		result = ^ast(variables, instructions, instr[1])
	} else if len(instr) == 3 {
		var left int = ast(variables, instructions, instr[0])
		var right int = ast(variables, instructions, instr[2])

		if instr[1] == "AND" {
			result = left & right
		} else if instr[1] == "OR" {
			result = left | right
		} else if instr[1] == "RSHIFT" {
			result = left >> right
		} else {
			result = left << right
		}
	} else {
		result = ast(variables, instructions, instr[0])
	}

	variables[operand] = result
	return result
}

func partOne(input string) string {
	var lines []string = strings.Split(input, "\n")

	var variables map[string]int = make(map[string]int, len(lines))
	var instructions map[string][]string = make(map[string][]string, len(lines))
	for _, l := range lines {
		if l == "" {
			continue
		}

		var tokens []string = strings.Split(l, " ")

		val, err := strconv.Atoi(tokens[0])
		if len(tokens) == 3 && err == nil {
			variables[tokens[2]] = val
		} else {
			var maxIdx int = len(tokens) - 1
			instructions[tokens[maxIdx]] = tokens[:(maxIdx - 1)]
		}
	}

	return strconv.Itoa(ast(variables, instructions, "a"))
}

func partTwo(input string) string {
	var lines []string = strings.Split(input, "\n")

	var variables map[string]int = make(map[string]int, len(lines))
	var instructions map[string][]string = make(map[string][]string, len(lines))
	for _, l := range lines {
		if l == "" {
			continue
		}

		var tokens []string = strings.Split(l, " ")

		val, err := strconv.Atoi(tokens[0])
		if len(tokens) == 3 && err == nil {
			variables[tokens[2]] = val
		} else {
			var maxIdx int = len(tokens) - 1
			instructions[tokens[maxIdx]] = tokens[:(maxIdx - 1)]
		}
	}

	var variablesOriginal map[string]int = make(map[string]int, len(variables))
	for k, v := range variables {
		variablesOriginal[k] = v
	}

	var a int = ast(variables, instructions, "a")
	variablesOriginal["b"] = a

	return strconv.Itoa(ast(variablesOriginal, instructions, "a"))
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  7,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 07", &day)
}
