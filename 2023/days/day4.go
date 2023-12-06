package days

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/joshfinnie/advent-of-code/2023/utils"
)

func parseFile(file io.Reader) ([]string, [][]int) {
	var cards []string
	var values [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "|")
		valuesStr := strings.Fields(parts[1])
		var v []int
		for _, val := range valuesStr {
			val, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			v = append(v, val)
		}
		values = append(values, v)

		cards = append(cards, parts[0])
	}

	return cards, values
}

func parseCard(card string, vals []int) int {
	game := strings.TrimSpace(strings.Split(card, ":")[0])
	outcome := strings.TrimPrefix(card, game+": ")
	outcomeArray := strings.Split(outcome, " ")
	results := 0
	for _, o := range outcomeArray {
		if o != "" {
			i := utils.ConvertStringToInt(o)
			if utils.IntInSlice(i, vals) {
				results++
			}
		}
	}

	return results
}

func Day04A(file io.Reader) string {
	answer := 0.0
	cards, values := parseFile(file)
	for i, card := range cards {
		p := parseCard(card, values[i])
		answer += math.Exp2(float64(p))
	}
	return strconv.Itoa(int(answer))
}

func Day04B(file io.Reader) string {
	answer := 0
	copies := map[int]int{}
	cards, values := parseFile(file)
	for i, card := range cards {
		matches := parseCard(card, values[i])
		for j := 1; j <= matches; j++ {
			copies[i+j] += 1 + copies[i]
		}
		answer += 1 + copies[i]
	}

	return strconv.Itoa(answer)
}
