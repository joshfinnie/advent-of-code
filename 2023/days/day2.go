package days

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func strCompare(str string, max int) bool {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return i > max
}

func isGamePossible(line string) int {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	substrings := strings.Split(line, ";")
	game := strings.TrimSpace(strings.Split(substrings[0], ":")[0])
	gameId := strings.Split(game, " ")[1]

	var outcome string
	possible := 0

	for i := 0; i < len(substrings); i++ {
		outcome = strings.TrimSpace(substrings[i])
		outcome = strings.TrimPrefix(outcome, game+": ")
		outcomeArray := strings.Split(outcome, ", ")
		for _, value := range outcomeArray {
			gameArray := strings.Split(value, " ")
			switch strings.TrimSpace(gameArray[1]) {
			case "blue":
				if strCompare(gameArray[0], maxBlue) {
					possible -= 1
				}
			case "red":
				if strCompare(gameArray[0], maxRed) {
					possible -= 1
				}
			case "green":
				if strCompare(gameArray[0], maxGreen) {
					possible -= 1
				}
			}

		}
	}

	if possible < 0 {
		return 0
	} else {
		i, err := strconv.Atoi(gameId)
		if err != nil {
			panic(err)
		}

		return i
	}

}

func Day02A(file io.Reader) string {
	scanner := bufio.NewScanner(file)

	games := 0

	for scanner.Scan() {
		line := scanner.Text()
		games += isGamePossible(line)
	}

	return strconv.Itoa(games)
}

func getGamePower(line string) int {
	minBlue := 0
	minRed := 0
	minGreen := 0

	substrings := strings.Split(line, ";")
	game := strings.TrimSpace(strings.Split(substrings[0], ":")[0])

	var outcome string
	for i := 0; i < len(substrings); i++ {
		outcome = strings.TrimSpace(substrings[i])
		outcome = strings.TrimPrefix(outcome, game+": ")
		outcomeArray := strings.Split(outcome, ", ")
		for _, value := range outcomeArray {
			gameArray := strings.Split(value, " ")
			count, err := strconv.Atoi(gameArray[0])
			if err != nil {
				panic(err)
			}
			switch strings.TrimSpace(gameArray[1]) {
			case "blue":
				if count > minBlue {
					minBlue = count
				}
			case "red":
				if count > minRed {
					minRed = count
				}
			case "green":
				if count > minGreen {
					minGreen = count
				}
			}

		}
	}

	return minBlue * minRed * minGreen
}

func Day02B(file io.Reader) string {
	scanner := bufio.NewScanner(file)

	games := 0

	for scanner.Scan() {
		line := scanner.Text()
		games += getGamePower(line)
	}

	return strconv.Itoa(games)
}
