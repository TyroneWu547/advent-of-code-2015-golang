package day04

import (
	utils "aoc/utils"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func solve(key string, prefix string) int {
	var i int
	var hashStr string = ""
	var minLen = len(hashStr)

	for i = 0; hashStr[:minLen] != prefix; i++ {
		var message string = fmt.Sprintf("%s%d", key, i)
		hash := md5.Sum([]byte(message))
		hashStr = hex.EncodeToString(hash[:])

		if len(hashStr) < len(prefix) {
			minLen = len(hashStr)
		} else {
			minLen = len(prefix)
		}
	}

	return i - 1
}

func partOne(input string) string {
	const prefix string = "00000"
	return strconv.Itoa(solve(input, prefix))
}

func partTwo(input string) string {
	const prefix string = "000000"
	return strconv.Itoa(solve(input, prefix))
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  4,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 04", &day)
}
