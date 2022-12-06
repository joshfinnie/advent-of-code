package days

import (
	"strconv"
	"strings"
)

func unique(input string) bool {
	chars := ""
	for _, char := range input {
		if strings.Contains(chars, string(char)) {
			return false
		}
		chars = chars + string(char)
	}
	return true
}

func calc(input string, val int) int {
	i := val - 1
	for i < (len(input) - val) {
		marker := input[i : i+val]
		if unique(marker) {
			return i + val
		}
		i++
	}
	return 0
}

func Day06Part1(input string) string {
	input = strings.TrimSuffix(input, "\n")
	return strconv.Itoa(calc(input, 4))
}

func Day06Part2(input string) string {
	input = strings.TrimSuffix(input, "\n")
	return strconv.Itoa(calc(input, 14))
}
