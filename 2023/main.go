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

// getDay returns the function corresponding to the specified day and part2 flag.
func getDay(day string, part2 bool) func(io.Reader) string {
	switch day {
	// This is where you add your next day's code
	case "01":
		if part2 {
			return days.Day01B
		}
		return days.Day01A
	default:
		return notAvailable
	}
}

// getInput returns an IO Reader file or an error.
func getInput(day string) (io.Reader, error) {
	filename := fmt.Sprintf("inputs/day%s.txt", day)
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
	)

	flag.BoolVar(&helpFlag, "help", false, "Show help")
	flag.StringVar(&dayFlag, "day", "00", "Which day to run.")
	flag.BoolVar(&part2Flag, "part2", false, "Run part 2.")

	flag.Parse()

	if helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	reader, err := getInput(dayFlag)
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
