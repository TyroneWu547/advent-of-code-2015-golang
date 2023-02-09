package days

func partOne(input string) (string, error) {

	return input, nil
}

func partTwo(input string) (string, error) {

	return input, nil
}

var day01 Day = Day{
	DayNum:  1,
	PartOne: partOne,
	PartTwo: partTwo,
}

func init() {
	RegisterDay("day01", day01)
}
