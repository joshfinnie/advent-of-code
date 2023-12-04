package days

import (
	"bufio"
	"io"
	"strconv"
)

type pos struct {
	x int
	y int
}

type num struct {
	val  int
	size int
}

type grid struct {
	nums map[pos]num
	symb map[pos]rune
}

func parse(file io.Reader) grid {
	scanner := bufio.NewScanner(file)
	g := grid{
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
				n = n*10 + int(r-'0')
				numSize++
			default:
				g.symb[pos{x: x, y: y}] = r
				if numSize > 0 {
					g.nums[pos{x: x, y: y - 1}] = num{val: n, size: numSize}
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

func hasNeigbor(g grid, p pos, n num) bool {
	// test above and below num
	for i := 0; i < n.size; i++ {
		if r, ok := g.symb[pos{x: p.x - 1, y: p.y - i}]; ok && r != '.' {
			return true
		}
		if r, ok := g.symb[pos{x: p.x + 1, y: p.y - i}]; ok && r != '.' {
			return true
		}
	}

	// test ends
	if r, ok := g.symb[pos{x: p.x, y: p.y + 1}]; ok && r != '.' {
		return true
	}
	if r, ok := g.symb[pos{x: p.x, y: p.y - n.size}]; ok && r != '.' {
		return true
	}

	// check corners
	if r, ok := g.symb[pos{x: p.x + 1, y: p.y + 1}]; ok && r != '.' {
		return true
	}
	if r, ok := g.symb[pos{x: p.x - 1, y: p.y + 1}]; ok && r != '.' {
		return true
	}
	if r, ok := g.symb[pos{x: p.x + 1, y: p.y - n.size}]; ok && r != '.' {
		return true
	}
	if r, ok := g.symb[pos{x: p.x - 1, y: p.y - n.size}]; ok && r != '.' {
		return true
	}

	return false
}

func Day03A(file io.Reader) string {
	sum := 0
	g := parse(file)
	for p, n := range g.nums {
		if hasNeigbor(g, p, n) {
			sum += n.val
		}
	}

	return strconv.Itoa(sum)
}

func getNumNeigbors(g grid, p pos) []num {
	numNeigbors := []num{}
	for pCur, nCur := range g.nums {
		if p.x >= pCur.x-1 && p.x <= pCur.x+1 && p.y >= pCur.y-nCur.size && p.y <= pCur.y+1 {
			numNeigbors = append(numNeigbors, nCur)
		}
	}
	return numNeigbors

}

func Day03B(file io.Reader) string {
	sum := 0
	g := parse(file)
	for p, s := range g.symb {
		if s == '*' {
			numNeigbors := getNumNeigbors(g, p)
			if len(numNeigbors) == 2 {
				ratio := numNeigbors[0].val * numNeigbors[1].val
				sum += ratio
			}
		}
	}

	return strconv.Itoa(sum)
}
