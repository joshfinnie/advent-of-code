package days

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

func sumOfDifferences(leftNums, rightNums []int) int {
	total := 0
	for i := 0; i < len(leftNums); i++ {
		diff := leftNums[i] - rightNums[i]
		if diff < 0 {
			diff = -diff // Get absolute value
		}
		total += diff
	}
	return total
}

func Day01A(file io.Reader) string {
	scanner := bufio.NewScanner(file)

	totalDistance := 0
	var leftNumbers []int
	var rightNumbers []int

	for scanner.Scan() {
		line := scanner.Text()
		var left, right int
		// Split the line and convert to integers
		fmt.Sscanf(line, "%d %d", &left, &right)

		leftNumbers = append(leftNumbers, left)
		rightNumbers = append(rightNumbers, right)
	}

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	totalDistance = sumOfDifferences(leftNumbers, rightNumbers)

	return strconv.Itoa(totalDistance)
}

func sumOfSimularity(leftNums, rightNums []int) int {
	total := 0
	for _, num := range leftNums {
		count := 0
		for _, target := range rightNums {
			if num == target {
				count++
			}
		}
		total += num * count
	}

	return total
}

func Day01B(file io.Reader) string {
	scanner := bufio.NewScanner(file)

	simularityScore := 0
	var leftNumbers []int
	var rightNumbers []int

	for scanner.Scan() {
		line := scanner.Text()
		var left, right int

		// Split the line and convert to integers
		fmt.Sscanf(line, "%d %d", &left, &right)

		leftNumbers = append(leftNumbers, left)
		rightNumbers = append(rightNumbers, right)
	}

	simularityScore = sumOfSimularity(leftNumbers, rightNumbers)

	return strconv.Itoa(simularityScore)
}
