package days

import (
	"strconv"
	"strings"
)

func comp(a []string, b []string) bool {
	a0, _ := strconv.Atoi(a[0])
	a1, _ := strconv.Atoi(a[1])
	b0, _ := strconv.Atoi(b[0])
	b1, _ := strconv.Atoi(b[1])
	return a0 <= b0 && a1 >= b1
}

func comp2(a []string, b []string) bool {
	// 6-6,4-5 -> NO OVERLAP
	a0, _ := strconv.Atoi(a[0])
	a1, _ := strconv.Atoi(a[1])
	b0, _ := strconv.Atoi(b[0])
	b1, _ := strconv.Atoi(b[1])
	return a1 >= b0 && a0 <= b1
}

func Day04Part1(input string) string {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n")
	count := 0
	for _, part := range parts {
		ass := strings.Split(part, ",")
		a := strings.Split(ass[0], "-")
		b := strings.Split(ass[1], "-")
		if comp(a, b) || comp(b, a) {
			count = count + 1
		}
	}
	return strconv.Itoa(count)
}

func Day04Part2(input string) string {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n")
	count := 0
	for _, part := range parts {
		ass := strings.Split(part, ",")
		a := strings.Split(ass[0], "-")
		b := strings.Split(ass[1], "-")
		if comp2(a, b) || comp2(b, a) {
			count = count + 1
		}
	}
	return strconv.Itoa(count)
}
