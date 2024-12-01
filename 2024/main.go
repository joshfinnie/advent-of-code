package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/joshfinnie/advent-of-code/2024/days"
)

type dayFunc func(io.Reader) string

var dayFunctions = map[string]struct {
	part1 dayFunc
	part2 dayFunc
}{
	"01": {days.Day01A, days.Day01B},
}

func getDay(day string, part2 bool) dayFunc {
	if functions, ok := dayFunctions[day]; ok {
		if part2 {
			return functions.part2
		}
		return functions.part1
	}
	return nil
}

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
		return
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
	if dayFunction == nil {
		fmt.Println("No function found for day", dayFlag)
		os.Exit(1)
	}

	result := dayFunction(reader)
	fmt.Println(result)
}
