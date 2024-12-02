package days

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Converts a space-separated line of numbers into a slice of integers
func parseNumbers(line string) ([]int, error) {
	fields := strings.Fields(line)
	numbers := make([]int, 0, len(fields))
	for _, numStr := range fields {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("invalid number %q: %v", numStr, err)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

// Checks if a sequence is strictly increasing or decreasing with small steps
func isValidSequence(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	increasing := nums[1] > nums[0]
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff < -3 || diff > 3 || (increasing && diff <= 0) || (!increasing && diff >= 0) {
			return false
		}
	}
	return true
}

// Processes the file and counts safe sequences for Day02A
func Day02A(file io.Reader) string {
	scanner := bufio.NewScanner(file)
	safeCount := 0
	for scanner.Scan() {
		numbers, err := parseNumbers(scanner.Text())
		if err != nil {
			continue // Ignore invalid lines
		}
		if isValidSequence(numbers) {
			safeCount++
		}
	}
	return strconv.Itoa(safeCount)
}

// Checks if a sequence is valid with up to one deviation
func isValidSequenceWithOneDeviation(nums []int) bool {
	if isValidSequence(nums) {
		return true
	}
	for skipIdx := 0; skipIdx < len(nums); skipIdx++ {
		previous := nums[0]
		valid := true
		for i := 1; i < len(nums); i++ {
			if i == skipIdx {
				continue
			}
			diff := nums[i] - previous
			if diff < -3 || diff > 3 {
				valid = false
				break
			}
			previous = nums[i]
		}
		if valid {
			return true
		}
	}
	return false
}

// Processes the file and counts safe sequences for Day02B
func Day02B(file io.Reader) string {
	scanner := bufio.NewScanner(file)
	safeCount := 0
	for scanner.Scan() {
		numbers, err := parseNumbers(scanner.Text())
		if err != nil {
			continue // Ignore invalid lines
		}
		if isValidSequenceWithOneDeviation(numbers) {
			safeCount++
		}
	}
	return strconv.Itoa(safeCount)
}
