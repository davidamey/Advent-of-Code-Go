package main

import (
	"advent/util"
	"fmt"
	"strings"
)

func main() {
	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")
	t := makeTower(lines)

	fmt.Println("p1=", t.root().name)
	fmt.Println("p2=", p2(t))
}

func p2(t *tower) int {
	computeTotalWeight(t.root())
	p, correctWeight := findImbalance(t.root())

	for {
		if imb, w := findImbalance(p); imb != nil {
			p, correctWeight = imb, w
		} else {
			break
		}
	}

	// fmt.Printf("%s, %d -> %d\n", p.name, p.totalWeight, correctWeight)

	adjustment := correctWeight - p.totalWeight
	return p.weight + adjustment
}

func findImbalance(p *program) (*program, int) {
	weightToCount := make(map[int]int)
	weightToProg := make(map[int]*program)
	for _, c := range p.children {
		// fmt.Printf("  %s: %d\n", c.name, c.totalWeight)
		weightToCount[c.totalWeight]++
		weightToProg[c.totalWeight] = c
	}

	// as there can be only one imbalance we look for the weight with count 1
	// if count is bigger than 1 then that's the "correct" weight for programs on this disc
	var imbalanced *program
	var correctWeight int
	for w, c := range weightToCount {
		if c == 1 {
			imbalanced = weightToProg[w]
		} else {
			correctWeight = w
		}
	}

	return imbalanced, correctWeight
}

func computeTotalWeight(p *program) {
	p.totalWeight = p.weight
	for _, c := range p.children {
		computeTotalWeight(c)
		p.totalWeight += c.totalWeight
	}
}

type program struct {
	name        string
	weight      int
	totalWeight int
	parent      *program
	children    []*program
}

type tower map[string]*program

func makeTower(lines []string) *tower {
	t := make(tower, len(lines))
	for _, l := range lines {
		parts := strings.SplitN(l, " -> ", 2)

		var name string
		var weight int
		fmt.Sscanf(parts[0], "%s (%d)", &name, &weight)

		p := t.getOrMake(name)
		p.weight = weight

		// If no children, process the next node
		if len(parts) == 1 {
			continue
		}

		for _, c := range strings.Split(parts[1], ", ") {
			cp := t.getOrMake(c)
			cp.parent = p
			p.children = append(p.children, cp)
		}
	}
	return &t
}

func (t *tower) getOrMake(name string) *program {
	if p, ok := (*t)[name]; ok {
		return p
	}
	(*t)[name] = &program{name: name}
	return (*t)[name]
}

func (t *tower) root() *program {
	for _, p := range *t {
		if p.parent == nil {
			return p
		}
	}
	return nil
}
