package days

import (
	"math"
	"sort"
	"strings"
	"strconv"
)

func parse(input string) []int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	elves := make([]int, 0)
	for _, part := range parts {
		sum := 0
		for _, line := range strings.Split(part, "\n") {
			i, _ := strconv.Atoi(line)
			sum += i
		}
		elves = append(elves, sum)
	}
	return elves
}

func Day01Part1(input string) string {
	elves := parse(input)
	max := math.MinInt
	for _, elf := range elves {
		if elf > max {
			max = elf
		}
	}
	return strconv.Itoa(max)
}

func Day01Part2(input string) string {
	elves := parse(input)
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	return strconv.Itoa(elves[0] + elves[1] + elves[2])
}
