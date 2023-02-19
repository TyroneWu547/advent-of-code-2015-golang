package day19

import (
	utils "aoc/utils"
	"strconv"
	"strings"
	"unicode"
)

func parseMolecule(input string) []string {
	var molecule []string = []string{}

	var e []rune
	for _, r := range input {
		if unicode.IsUpper(r) && len(e) != 0 {
			molecule = append(molecule, string(e))
			e = []rune{}
		}
		e = append(e, r)
	}
	molecule = append(molecule, string(e))

	return molecule
}

func parseInputOne(input string) (map[string][][]string, []string) {
	var lines []string = strings.Split(input, "\n")
	var replacements map[string][][]string = make(map[string][][]string, len(lines))

	for _, line := range lines {
		if line == "" {
			break
		}

		var tokens []string = strings.Split(line, " ")
		_, ok := replacements[string(tokens[0])]
		if !ok {
			replacements[string(tokens[0])] = [][]string{}
		}
		replacements[string(tokens[0])] = append(replacements[string(tokens[0])], parseMolecule(tokens[2]))
	}

	var molecule []string = parseMolecule(lines[len(lines)-2])

	return replacements, molecule
}

func transmuteElement(molecule []string, idx int, replacement []string) []string {
	var molCopy []string = make([]string, len(molecule))
	copy(molCopy, molecule)

	molCopy[idx] = replacement[0]
	if len(replacement) == 1 {
		return molCopy
	} else {
		// https://github.com/golang/go/wiki/SliceTricks#insert
		molCopy = append(molCopy[:(idx+1)], append(replacement[1:], molCopy[(idx+1):]...)...)
		return molCopy
	}
}

func partOne(input string) string {
	repl, mol := parseInputOne(input)
	var distMol map[string]bool = map[string]bool{}

	for i, e := range mol {
		for _, r := range repl[e] {
			var transMol []string = transmuteElement(mol, i, r)
			distMol[strings.Join(transMol, "")] = true
		}
	}

	return strconv.Itoa(len(distMol))
}

func parseInputTwo(input string) (map[string]string, string) {
	var lines []string = strings.Split(input, "\n")
	var replacements map[string]string = make(map[string]string, len(lines))

	for _, line := range lines {
		if line == "" {
			break
		}

		var tokens []string = strings.Split(line, " ")
		replacements[tokens[2]] = tokens[0]
	}

	var molecule string = lines[len(lines)-2]

	return replacements, molecule
}

func dfs(molecule string, replacements map[string]string, muts int) int {
	if molecule == "e" {
		return muts
	}

	for k, v := range replacements {
		var lim int = len(molecule) - len(k)
		for i := 0; i <= lim; i++ {
			if molecule[i:(i+len(k))] == k {
				var revMut string = strings.Join([]string{molecule[:i], v, molecule[(i + len(k)):]}, "")
				var counts int = dfs(revMut, replacements, muts+1)
				if counts > 0 {
					return counts
				}
			}
		}
	}

	return -1
}

func partTwo(input string) string {
	repl, mol := parseInputTwo(input)
	return strconv.Itoa(dfs(mol, repl, 0))
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  19,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 19", &day)
}
