package days

import (
	"strconv"
	"strings"
)

/*
 * A -> Rock    -> X
 * B -> Paper   -> Y
 * C -> Scissor -> Z
 */

const WIN = 6
const LOSE = 0
const TIE = 3

func calcScore(opp string, me string) int {
	if opp == "A" {
		switch me {
		case "X":
			return TIE + 1
		case "Y":
			return WIN + 2
		case "Z":
			return LOSE + 3
		default:
			return 0
		}
	} else if opp == "B" {
		switch me {
		case "X":
			return LOSE + 1
		case "Y":
			return TIE + 2
		case "Z":
			return WIN + 3
		default:
			return 0
		}
	} else {
		switch me {
		case "X":
			return WIN + 1
		case "Y":
			return LOSE + 2
		case "Z":
			return TIE + 3
		default:
			return 0
		}
	}
}

/*
 * X -> LOSE
 * Y -> TIE
 * Z -> WIN
 */
/*
 * A -> Rock    -> X
 * B -> Paper   -> Y
 * C -> Scissor -> Z
 */
func calcScore2(opp string, me string) int {
	if me == "X" {
		/*
		* LOSE
		* A -> Rock    3
		* B -> Paper   1
		* C -> Scissor 2
		 */
		switch opp {
		case "A":
			return LOSE + 3
		case "B":
			return LOSE + 1
		case "C":
			return LOSE + 2
		default:
			return 0
		}
	} else if me == "Y" {
		switch opp {
		case "A":
			return TIE + 1
		case "B":
			return TIE + 2
		case "C":
			return TIE + 3
		default:
			return 0
		}
	} else {
		/*
		* WIN
		* A -> Rock    2
		* B -> Paper   3
		* C -> Scissor 1
		 */
		switch opp {
		case "A":
			return WIN + 2
		case "B":
			return WIN + 3
		case "C":
			return WIN + 1
		default:
			return 0
		}
	}
}

func Day02Part1(input string) string {
	input = strings.TrimSuffix(input, "\n")
	scores := 0
	parts := strings.Split(input, "\n")
	for _, part := range parts {
		plays := strings.Fields(part)
		opponent := plays[0]
		me := plays[1]
		scores = scores + calcScore(opponent, me)
	}

	return strconv.Itoa(scores)
}

func Day02Part2(input string) string {
	input = strings.TrimSuffix(input, "\n")
	scores := 0
	parts := strings.Split(input, "\n")
	for _, part := range parts {
		plays := strings.Fields(part)
		opponent := plays[0]
		me := plays[1]
		scores = scores + calcScore2(opponent, me)
	}

	return strconv.Itoa(scores)
}
