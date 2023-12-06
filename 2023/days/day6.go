package days

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/joshfinnie/advent-of-code/2023/utils"
)

func Day06A(file io.Reader) string {
	scanner := bufio.NewScanner(file)

	var times []int
	var distances []int

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		lineArray := strings.Fields(line)

		numSlice := utils.ParseIntArray(lineArray[1:])

		switch lineArray[0] {
		case "Time:":
			times = numSlice
		case "Distance:":
			distances = numSlice
		}
	}

	answer := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		recordDistance := distances[i]

		winners := 0
		for speed := 1; speed < time; speed++ {
			distance := (time - speed) * speed
			if distance > recordDistance {
				winners++
			}
		}
		answer *= winners
	}

	return strconv.Itoa(answer)
}

func Day06B(file io.Reader) string {
	scanner := bufio.NewScanner(file)

	var time int
	var distance int

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		lineArray := strings.Split(line, ":")

		num := strings.ReplaceAll(lineArray[1], " ", "")

		switch lineArray[0] {
		case "Time":
			time = utils.ConvertStringToInt(num)
		case "Distance":
			distance = utils.ConvertStringToInt(num)
		}
	}

	answer := 0
	for speed := 1; speed < time; speed++ {
		d := (time - speed) * speed
		if d > distance {
			answer++
		}
	}

	return strconv.Itoa(answer)
}
