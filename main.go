package main

import (
	days "aoc/days"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var args []string = os.Args
	if len(args) != 4 {
		invalidArgs()
	}
	fmt.Println(args)

	var dayNum int
	var err error
	dayNum, err = strconv.Atoi(os.Args[1])
	if err != nil {
		invalidArgs()
	}

	var part string = os.Args[2]
	if part != "p1" && part != "p2" && part != "both" {
		invalidArgs()
	}

	var task string = os.Args[3]
	if task != "solve" && task != "test" {
		invalidArgs()
	}

	var dayFmt string = fmt.Sprintf("day%02d", dayNum)
	fmt.Println(dayFmt)
	var day days.Day = days.AOCDays[dayFmt]

	fmt.Println(day.PartOne("asdfasdf"))
}

func invalidArgs() {
	fmt.Println("Usage: go run main.go <day-number> <p1|p2|both> <solve|test>")
	os.Exit(1)
}
