package main

import (
	"flag"
	"fmt"
	"os"

	"aoc/days"
)

func notAvailable(_ string) string {
	return "This day is not available yet..."
}

func getDay(day string, part2 bool) func(string) string {
	switch day {
	// This is where you add your next day's code
	case "01":
		if part2 {
			return days.Day01Part2
		}
		return days.Day01Part1
	// This is the default so that something happens no matter what
	default:
		return notAvailable
	}
}

var help = flag.Bool("help", false, "Show help")
var dayFlag string
var part2Flag bool

func main() {
	flag.StringVar(&dayFlag, "day", "00", "Which day we want to run.")
	flag.BoolVar(&part2Flag, "part2", false, "A boolean flag to run part 2.")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// Read input for the day.
	filename := fmt.Sprintf("inputs/day%s.txt", dayFlag)
	dat, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("No input found for day:", dayFlag)
	}

	f := getDay(dayFlag, part2Flag)
	results := f(string(dat))

	var part string
	if part2Flag {
		part = "Part 2:"
	} else {
		part = "Part 1:"
	}

	fmt.Println(part, results)
}
