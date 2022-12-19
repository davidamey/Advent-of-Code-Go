package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	split := 0
	for i, l := range lines {
		if l == "" {
			split = i
			break
		}
	}

	p1Stacks := parseStacks(lines[0:split])
	p2Stacks := parseStacks(lines[0:split])
	for _, instruct := range lines[split+1:] {
		var n, source, target int
		fmt.Sscanf(instruct, "move %d from %d to %d", &n, &source, &target)

		for i := 0; i < n; i++ {
			p1Stacks[target-1].push(p1Stacks[source-1].pop(1)...)
		}
		p2Stacks[target-1].push(p2Stacks[source-1].pop(n)...)
	}

	p1 := make([]rune, len(p1Stacks))
	p2 := make([]rune, len(p2Stacks))
	for i := range p1Stacks {
		p1[i] = p1Stacks[i].last()
		p2[i] = p2Stacks[i].last()
	}
	fmt.Println("p1=", string(p1))
	fmt.Println("p2=", string(p2))
}

type stack []rune

func (s *stack) push(rs ...rune) {
	(*s) = append(*s, rs...)
}

func (s *stack) pop(n int) (rs []rune) {
	rs, (*s) = (*s)[len((*s))-n:], (*s)[:len((*s))-n]
	return
}

func (s *stack) last() rune {
	return (*s)[len(*s)-1]
}

func parseStacks(raw []string) []stack {
	stacks := make([]stack, 0)

	for i := len(raw) - 2; i >= 0; i-- {
		row := raw[i]
		for s, j := 0, 0; j < len(row); s, j = s+1, j+4 {
			if row[j] == '[' {
				if s >= len(stacks) {
					stacks = append(stacks, stack{rune(row[j+1])})
				} else {
					stacks[s].push(rune(row[j+1]))
				}
			}
		}
	}

	return stacks
}
