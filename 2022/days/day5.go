package days

import (
	"fmt"
	"strings"
)

type Stack []string

func (s *Stack) pushFront(val string) {
	*s = append(Stack{val}, *s...)
}

func (s *Stack) push(val string) {
	*s = append(*s, val)
}

func (s *Stack) pop() string {
	length := len(*s)
	if length == 0 {
		return ""
	} else {
		element := (*s)[length - 1]
		*s = (*s)[:length - 1]
		return element
	}
}

func (s *Stack) peek() string {
	length := len(*s)
	if length == 0 {
		return ""
	} else {
		element := (*s)[length - 1]
		return element
	}
}

func parseStacks(part string) []Stack {
	stacks := make([]Stack, 10)

	crates := strings.Split(part, "\n")
	for _, crate := range crates {
		if crate[1] == '1' {
			return stacks
		}

		s := 1
		i := 1
		for i < len(crate) {
			if crate[i] != ' ' {
				stack := stacks[s]
				stack.pushFront(string(crate[i]))
				stacks[s] = stack
			}
			s++
			i += 4
		}
	}

	return stacks
}

func calcInstructions(part string, stacks []Stack) string {
	insts := strings.Split(part, "\n")
	for _, inst := range insts {
		crates := 0
		loc_a := 0
		loc_b := 0
		fmt.Sscanf(inst, "move %d from %d to %d", &crates, &loc_a, &loc_b)
		for i := 0; i < crates; i++ {
			stackA := stacks[loc_a]
			stackB := stacks[loc_b]
			v := stackA.pop()
			stackB.push(v)

			stacks[loc_a] = stackA
			stacks[loc_b] = stackB
		}
	}

	var answer strings.Builder
	for i := 0; i < len(stacks); i++ {
		answer.WriteString(stacks[i].peek())
	}
	return answer.String()
}

func calcInstructions2(part string, stacks []Stack) string {
	insts := strings.Split(part, "\n")
	for _, inst := range insts {
		crates := 0
		loc_a := 0
		loc_b := 0
		fmt.Sscanf(inst, "move %d from %d to %d", &crates, &loc_a, &loc_b)
		stackA := stacks[loc_a]
		stackB := stacks[loc_b]
		x := len(stackA) - crates
		stackB = append(stackB, stacks[loc_a][x:]...)

		stacks[loc_a] = stackA[:x]
		stacks[loc_b] = stackB
	}

	var answer strings.Builder
	for i := 0; i < len(stacks); i++ {
		answer.WriteString(stacks[i].peek())
	}
	return answer.String()
}

func Day05Part1(input string) string {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")

	stacks := parseStacks(parts[0])
	return calcInstructions(parts[1], stacks)
}

func Day05Part2(input string) string {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")

	stacks := parseStacks(parts[0])
	return calcInstructions2(parts[1], stacks)
}
