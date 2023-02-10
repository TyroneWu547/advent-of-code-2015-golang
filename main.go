package main

import (
	days "aoc/days"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Format inspired by https://github.com/fspoettel/advent-of-code-rust
func main() {
	// Args error checking
	var args []string = os.Args
	if len(args) != 4 {
		invalidArgs()
	}

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

	var fileContents []byte
	fileContents, err = os.ReadFile(os.Args[3])
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	var dayFmt string = fmt.Sprintf("Day %02d", dayNum)
	var day *days.Day = days.AOCDays[dayFmt]

	fmt.Printf("----------\n"+
		"| %s |\n"+
		"----------\n", dayFmt)

	var input string = string(fileContents)
	if part == "p1" || part == "both" {
		fmt.Println("\n🎄 Part 1 🎄")
		benchmarkTask(day.PartOne, input)
	}
	if part == "p2" || part == "both" {
		fmt.Println("\n🎄 Part 2 🎄")
		benchmarkTask(day.PartTwo, input)
	}
}

// Benchmark function
func benchmarkTask(task days.Task, input string) {
	var start time.Time = time.Now()
	var output string = task(string(input))
	var duration time.Duration = time.Since(start)

	fmt.Printf("%s (elapsed: %s)\n", output, duration)
}

func invalidArgs() {
	fmt.Println("Usage: go run main.go <day-number> <p1|p2|both> <file-path-input>")
	os.Exit(1)
}
