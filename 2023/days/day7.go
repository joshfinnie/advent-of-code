package days

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/joshfinnie/advent-of-code/2023/utils"
)

type HandType int

const (
	oneOAK = iota
	onePair
	twoPair
	threeOAK
	fullHouse
	fourOAK
	fiveOAK
)

type Hands struct {
	hands []Hand
}

type Hand struct {
	cards []rune
	bet   int
	hand  HandType
}

func calculateHandType(hand Hand) Hand {
	cards := map[rune]int{}
	for _, r := range hand.cards {
		cards[r]++
	}
	counts := []int{}
	for _, i := range cards {
		counts = append(counts, i)
	}
	slices.Sort(counts)
	slices.Reverse(counts)

	switch len(counts) {
	case 1:
		hand.hand = fiveOAK
	case 2:
		if counts[0] == 4 {
			hand.hand = fourOAK
		}
		if counts[0] == 3 {
			hand.hand = fullHouse
		}
	case 3:
		if counts[0] == 3 {
			hand.hand = threeOAK
		}
		if counts[0] == 2 {
			hand.hand = twoPair
		}
	case 4:
		if counts[0] == 2 {
			hand.hand = onePair
		}
	case 5:
		hand.hand = oneOAK
	default:
		panic("AHH")
	}

	return hand
}

var cardVal = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func comp(a Hand, b Hand) int {
	if a.hand < b.hand {
		return -1
	}
	if a.hand > b.hand {
		return 1
	}
	for i := range a.cards {
		if cardVal[a.cards[i]] < cardVal[b.cards[i]] {
			return -1
		}
		if cardVal[a.cards[i]] > cardVal[b.cards[i]] {
			return 1
		}
	}
	return 0
}

func Day07A(file io.Reader) string {
	answer := 0
	scanner := bufio.NewScanner(file)
	hands := Hands{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineArr := strings.Split(line, " ")
		hand := Hand{
			cards: []rune(lineArr[0]),
			bet:   utils.ConvertStringToInt(lineArr[1]),
		}

		hand = calculateHandType(hand)

		hands.hands = append(hands.hands, hand)
	}

	slices.SortFunc(hands.hands, comp)
	for i, h := range hands.hands {
		answer += h.bet * (i + 1)
	}

	return strconv.Itoa(answer)
}

func calculateHandType2(hand Hand) Hand {
	cards := map[rune]int{}
	jokers := 0
	for _, r := range hand.cards {
		if r == 'J' {
			jokers++
		} else {
			cards[r]++
		}
	}
	counts := []int{}
	for _, i := range cards {
		counts = append(counts, i)
	}
	slices.Sort(counts)
	slices.Reverse(counts)

	if len(counts) > 0 {
		counts[0] += jokers
	} else {
		counts = []int{5} // all jokers
	}

	switch len(counts) {
	case 1:
		hand.hand = fiveOAK
	case 2:
		if counts[0] == 4 {
			hand.hand = fourOAK
		}
		if counts[0] == 3 {
			hand.hand = fullHouse
		}
	case 3:
		if counts[0] == 3 {
			hand.hand = threeOAK
		}
		if counts[0] == 2 {
			hand.hand = twoPair
		}
	case 4:
		if counts[0] == 2 {
			hand.hand = onePair
		}
	case 5:
		hand.hand = oneOAK
	default:
		panic("AHH")
	}

	return hand
}

var cardVal2 = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 1,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func comp2(a Hand, b Hand) int {
	if a.hand < b.hand {
		return -1
	}
	if a.hand > b.hand {
		return 1
	}
	for i := range a.cards {
		if cardVal2[a.cards[i]] < cardVal2[b.cards[i]] {
			return -1
		}
		if cardVal2[a.cards[i]] > cardVal2[b.cards[i]] {
			return 1
		}
	}
	return 0
}

func Day07B(file io.Reader) string {
	answer := 0
	scanner := bufio.NewScanner(file)
	hands := Hands{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineArr := strings.Split(line, " ")
		hand := Hand{
			cards: []rune(lineArr[0]),
			bet:   utils.ConvertStringToInt(lineArr[1]),
		}

		hand = calculateHandType2(hand)

		hands.hands = append(hands.hands, hand)
	}

	slices.SortFunc(hands.hands, comp2)
	for i, h := range hands.hands {
		answer += h.bet * (i + 1)
	}

	return strconv.Itoa(answer)
}
