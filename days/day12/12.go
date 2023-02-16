package day12

import (
	utils "aoc/utils"
	"encoding/json"
	"regexp"
	"strconv"
)

func partOne(input string) string {
	re := regexp.MustCompile(`[-]?[0-9]+`)
	var parsedNums []string = re.FindAllString(input, -1)
	var sum int = 0
	for _, n := range parsedNums {
		val, _ := strconv.Atoi(n)
		sum += val
	}
	return strconv.Itoa(sum)
}

func parseArray(jsonArray []interface{}) int {
	var sum int
	for _, j := range jsonArray {
		sum += traverseJson(j)
	}
	return sum
}

func parseJson(jsonObject map[string]interface{}) int {
	var sum int = 0
	for _, v := range jsonObject {
		if v == "red" {
			return 0
		}
		sum += traverseJson(v)
	}
	return sum
}

func traverseJson(jsonObjects interface{}) int {
	switch t := jsonObjects.(type) {
	case []interface{}:
		return parseArray(t)
	case map[string]interface{}:
		return parseJson(t)
	case float64:
		return int(t)
	default:
		return 0
	}
}

func partTwo(input string) string {
	var jsonInput interface{}
	json.Unmarshal([]byte(input), &jsonInput)
	return strconv.Itoa(traverseJson(jsonInput))
}

func init() {
	var day utils.Day = utils.Day{
		DayNum:  12,
		PartOne: partOne,
		PartTwo: partTwo,
	}
	utils.RegisterDay("Day 12", &day)
}
