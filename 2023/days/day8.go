package days

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type day8Map struct {
	instructions []rune
	nodes        map[string]node
}

type node struct {
	left  string
	right string
}

var REGEX = regexp.MustCompile(`^([A-Z1-3]+) = \(([A-Z1-3]+), ([A-Z1-3]+)\)$`)

func traverse(m day8Map, start string, end string) int {
	count := 0
	cur := start
	for cur != end {
		n, _ := m.nodes[cur]
		switch m.instructions[count%len(m.instructions)] {
		case 'L':
			cur = n.left
		case 'R':
			cur = n.right
		default:
			panic("panic")
		}
		count++
	}

	return count
}

func Day08A(file io.Reader) string {
	answer := 0
	m := day8Map{
		nodes: map[string]node{},
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := strings.TrimSpace(scanner.Text())
	m.instructions = []rune(instructions)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		match := REGEX.FindStringSubmatch(line)
		m.nodes[match[1]] = node{left: match[2], right: match[3]}
	}

	answer = traverse(m, "AAA", "ZZZ")
	return strconv.Itoa(answer)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a int, b int) int {
	result := a * b / GCD(a, b)
	return result
}

func traverse2(m day8Map, start string, end string) int {
	count := 0
	cur := start
	for cur[2:] != end {
		n, _ := m.nodes[cur]
		switch m.instructions[count%len(m.instructions)] {
		case 'L':
			cur = n.left
		case 'R':
			cur = n.right
		default:
			panic("panic")
		}
		count++
	}

	return count
}

func Day08B(file io.Reader) string {
	answer := 1
	m := day8Map{
		nodes: map[string]node{},
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := strings.TrimSpace(scanner.Text())
	m.instructions = []rune(instructions)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		match := REGEX.FindStringSubmatch(line)
		m.nodes[match[1]] = node{left: match[2], right: match[3]}
	}

	for node := range m.nodes {
		if node[2:] == "A" {
			answer = LCM(answer, traverse2(m, node, "Z"))
		}
	}

	return strconv.Itoa(answer)
}
