package days

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func extractValues(line string) int {
	var firstDigit, lastDigit rune

	// Find the first and last digits in the line
	for _, char := range line {
		if unicode.IsDigit(char) {
			if firstDigit == 0 {
				firstDigit = char
			}
			lastDigit = char
		}
	}

	// Convert the digits to integers and form the two-digit number
	firstDigitInt := int(firstDigit - '0')
	lastDigitInt := int(lastDigit - '0')
	calibrationValue := firstDigitInt*10 + lastDigitInt

	return calibrationValue
}

func Day01A(file io.Reader) string {
	scanner := bufio.NewScanner(file)

	totalCalibration := 0

	for scanner.Scan() {
		line := scanner.Text()
		calibrationValue := extractValues(line)
		totalCalibration += calibrationValue
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return strconv.Itoa(totalCalibration)
}

func toNumber(s string) int {
	switch s {
	case "0", "zero":
		return 0
	case "1", "one":
		return 1
	case "2", "two":
		return 2
	case "3", "three":
		return 3
	case "4", "four":
		return 4
	case "5", "five":
		return 5
	case "6", "six":
		return 6
	case "7", "seven":
		return 7
	case "8", "eight":
		return 8
	case "9", "nine":
		return 9
	default:
		panic("unknown value: " + s)
	}
}

func extractValues2(line string) int {
	numbers := `[0-9]|zero|one|two|three|four|five|six|seven|eight|nine`
	reFirst := regexp.MustCompile(`^.*?(` + numbers + `).*$`)
	reLast := regexp.MustCompile(`^.*(` + numbers + `).*?$`)
	firstDigit := reFirst.FindStringSubmatch(line)
	lastDigit := reLast.FindStringSubmatch(line)

	return (toNumber(firstDigit[1]) * 10) + toNumber(lastDigit[1])
}

func Day01B(file io.Reader) string {
	scanner := bufio.NewScanner(file)

	totalCalibration := 0

	for scanner.Scan() {
		line := scanner.Text()
		calibrationValue := extractValues2(line)
		totalCalibration += calibrationValue
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return strconv.Itoa(totalCalibration)
}
