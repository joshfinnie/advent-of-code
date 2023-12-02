package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/joshfinnie/advent-of-code/2023/days"
)

// notAvailable returns a message indicating that the day is not available.
func notAvailable(_ io.Reader) string {
	return "This day is not done yet..."
}

type dayFunc func(io.Reader) string

var dayFunctions = map[string]struct {
	part1 dayFunc
	part2 dayFunc
}{
	"01": {days.Day01A, days.Day01B},
	"02": {days.Day02A, days.Day02B},
}

// getDay returns the function corresponding to the specified day and part2 flag.
func getDay(day string, part2 bool) dayFunc {
	if functions, ok := dayFunctions[day]; ok {
		if part2 {
			return functions.part2
		}
		return functions.part1
	}
	return notAvailable
}

// getInput returns an IO Reader file or an error.
func getInput(day string, test bool) (io.Reader, error) {
	var filename string
	if test {
		filename = fmt.Sprintf("inputs/day%s_test.txt", day)
	} else {
		filename = fmt.Sprintf("inputs/day%s.txt", day)
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading input for day %s: %v", day, err)
	}
	return file, nil
}

func main() {
	var (
		helpFlag  bool
		dayFlag   string
		part2Flag bool
		testFlag  bool
	)

	flag.BoolVar(&helpFlag, "help", false, "Show help")
	flag.StringVar(&dayFlag, "day", "00", "Which day to run.")
	flag.BoolVar(&part2Flag, "part2", false, "Run part 2.")
	flag.BoolVar(&testFlag, "test", false, "Use test input")

	flag.Parse()

	if helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	reader, err := getInput(dayFlag, testFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		if closer, ok := reader.(io.Closer); ok {
			closer.Close()
		}
	}()

	dayFunction := getDay(dayFlag, part2Flag)
	result := dayFunction(reader)

	part := "Part 1:"
	if part2Flag {
		part = "Part 2:"
	}

	fmt.Printf("%s %s\n", part, result)
}
