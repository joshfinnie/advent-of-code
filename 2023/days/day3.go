package days

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type pos struct {
	x int
	y int
}

type num struct {
	val int
	size int
}

type grid struct {
	nums map[pos]num
	symb map[pos]rune
}

func parse(file io.Reader) grid {
	scanner := bufio.NewScanner(file)
	g := grid {
		nums: make(map[pos]num),
		symb: make(map[pos]rune),
	}

	x := 0
	n := 0
	numSize := 0
	for scanner.Scan() {
		line := scanner.Text()
		for y, r := range line {
			switch r {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				n = n * 10 + int(r - '0')
				numSize++
			default:
				g.symb[pos{x: x, y: y-1}] = r
				if numSize > 0 {
					g.nums[pos{x: x, y: y-1}] = num{val: n, size: numSize}
					n = 0
					numSize = 0
				}
			}
		}
		if numSize > 0 {
			g.nums[pos{x: x, y: len(line) - 1}] = num{val: n, size: numSize}
			n = 0
			numSize = 0
		}
		x++
	}

	return g
}

var adjPoints [][]int = [][]int{
	{-1, 0},  // Going up
	{1, 0},   // Going down
	{0, -1},  // Going left
	{0, 1},   // Going right
	{-1, -1}, // Top Left
	{-1, 1},  // Top Right
	{1, -1},  // Bottom Left
	{1, 1},   // Bottom Right
}

func hasNeigbor(g grid, p pos) bool {
	fmt.Println(p)
	for _, pt := range adjPoints {
		if r, ok := g.symb[pos{x: p.x + pt[0], y: p.y + pt[1]}]; ok && r != '.' {
			return true
		}
	}
	// Check up
	return false
}

func Day03A(file io.Reader) string {
	sum := 0
	g := parse(file)
	for p, n := range g.nums {
		if hasNeigbor(g, p) {
			sum += n.val
		}
	}

	return strconv.Itoa(sum)
}

func Day03B(file io.Reader) string {
	return "TODO"
}
