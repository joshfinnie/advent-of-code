package days

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func isValidSequence(nums []int) bool {
	// Determine if sequence is increasing or decreasing based on first two numbers
	increasing := nums[1] > nums[0]

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]

		// Check if difference is too large
		if diff < -3 || diff > 3 {
			return false
		}

		// Check if sequence maintains direction (increasing or decreasing)
		if increasing && diff <= 0 {
			return false
		}
		if !increasing && diff >= 0 {
			return false
		}
	}

	return true
}

func Day02A(file io.Reader) string {
	scanner := bufio.NewScanner(file)
	safe := 0

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)

		// Convert strings to integers
		var nums []int
		for _, numStr := range numbers {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}
			nums = append(nums, num)
		}

		if isValidSequence(nums) {
			safe += 1
		}
	}

	return strconv.Itoa(safe)
}

func isValidSequenceWithOneDeviation(nums []int) bool {
	// If the sequence is already valid, it doesn't need a deviation
	if isValidSequence(nums) {
		return true
	}

	// Try removing each number once and check if remaining sequence is valid
	for skipIdx := 0; skipIdx < len(nums); skipIdx++ {
		// Create new slice without the current number
		var testNums []int
		for i := 0; i < len(nums); i++ {
			if i != skipIdx {
				testNums = append(testNums, nums[i])
			}
		}

		// If removing this number makes the sequence valid, we found our solution
		if isValidSequence(testNums) {
			return true
		}
	}

	return false
}

func Day02B(file io.Reader) string {
	scanner := bufio.NewScanner(file)
	safe := 0

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)

		// Convert strings to integers
		var nums []int
		for _, numStr := range numbers {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}
			nums = append(nums, num)
		}

		if isValidSequenceWithOneDeviation(nums) {
			safe += 1
		}
	}

	return strconv.Itoa(safe)

}
