package day23

import (
	utils "aoc/utils"
	"math"
	"strconv"
	"strings"
)

type instruction int

const (
	hlf instruction = iota
	tpl
	inc
	jmp
	jie
	jio
)

type operation struct {
	instr    instruction
	register rune
	value    int8
}

func parseInput(input string) []operation {
	var lines []string = strings.Split(input, "\n")
	var ops []operation = make([]operation, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}

		var tokens []string = strings.Split(line, " ")
		switch tokens[0] {
		case "hlf":
			ops = append(ops, operation{hlf, rune(tokens[1][0]), 0})
		case "tpl":
			ops = append(ops, operation{tpl, rune(tokens[1][0]), 0})
		case "inc":
			ops = append(ops, operation{inc, rune(tokens[1][0]), 0})
		case "jmp":
			val, _ := strconv.Atoi(tokens[1][1:])
			if tokens[1][0] == '-' {
				val = -val
			}
			ops = append(ops, operation{jmp, '\000', int8(val)})
		case "jie":
			val, _ := strconv.Atoi(tokens[2][1:])
			if tokens[2][0] == '-' {
				val = -val
			}
			ops = append(ops, operation{jie, rune(tokens[1][0]), int8(val)})
		case "jio":
			val, _ := strconv.Atoi(tokens[2][1:])
			if tokens[2][0] == '-' {
				val = -val
			}
			ops = append(ops, operation{jio, rune(tokens[1][0]), int8(val)})
		default:
			panic("Invalid instruction parsed")
		}
	}

	return ops
}

func executeInstruction(op *operation, instrIdx *int, a *float64, b *float64) {
	switch op.instr {
	case hlf:
		if op.register == 'a' {
			*a /= 2
		} else {
			*b /= 2
		}
	case tpl:
		if op.register == 'a' {
			*a *= 3
		} else {
			*b *= 3
		}
	case inc:
		if op.register == 'a' {
			*a++
		} else {
			*b++
		}
	case jmp:
		*instrIdx += int(op.value) - 1
	case jie:
		if (op.register == 'a' && math.Mod(*a, 2) == 0) || (op.register == 'b' && math.Mod(*b, 2) == 0) {
			*instrIdx += int(op.value) - 1
		}
	case jio:
		if (op.register == 'a' && *a == 1) || (op.register == 'b' && *b == 1) {
			*instrIdx += int(op.value) - 1
		}
	}

	*instrIdx++
}

func partOne(input string) string {
	var ops []operation = parseInput(input)
	var a float64 = 0
	var b float64 = 0

	var instrIdx int = 0
	for 0 <= instrIdx && instrIdx < len(ops) {
		executeInstruction(&ops[instrIdx], &instrIdx, &a, &b)
	}
	return strconv.FormatFloat(b, 'f', -1, 64)
}

func partTwo(input string) string {
	var ops []operation = parseInput(input)
	var a float64 = 1
	var b float64 = 0

	var instrIdx int = 0
	for 0 <= instrIdx && instrIdx < len(ops) {
		executeInstruction(&ops[instrIdx], &instrIdx, &a, &b)
	}
	return strconv.FormatFloat(b, 'f', -1, 64)
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  23,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 23", &day)
}
