package days

import (
	"strconv"
	"strings"
)

func calcVal(c rune) int {
	i := int(c)
	if i >= 65 && i <= 90 {
		return i - 38
	} else {
		return i - 96
	}
}

func Day03Part1(input string) string {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n")
	count := 0
	for _, part := range parts {
		var dups rune
		l := len(part)
		s1 := part[0:l/2]
		s2 := part[l/2:l]
		for _, char := range s1 {
			if strings.Contains(s2, string(char)) {
				dups = char
			}
		}
		
		count = count + calcVal(dups)
	}
	return strconv.Itoa(count)
}

func Day03Part2(input string) string {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n")
	max := len(parts)
	min := 0
	count := 0
	for min < max {
		s1 := parts[min]
		s2 := parts[min + 1]
		s3 := parts[min + 2]
		var dups rune
		for _, char := range s1 {
			if strings.Contains(s2, string(char)) {
				if strings.Contains(s3, string(char)) {
					dups = char
				}
			}
		}
		count = count + calcVal(dups)
		min = min + 3
	}
	return strconv.Itoa(count)
}
