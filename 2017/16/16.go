package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
)

var byPos []*program = make([]*program, 16)
var byName map[rune]*program = make(map[rune]*program)

func main() {
	input := string(util.MustReadFile("input"))

	for i := range byPos {
		n := rune('a' + i)
		p := &program{n, i}
		byPos[i] = p
		byName[n] = p
	}

	dance(input)
	p1 := toString()

	fmt.Println("p1=", p1)

	cycleLength := 0
	for {
		cycleLength++
		dance(input)
		if toString() == p1 {
			break
		}
	}

	// if we're here, then we found a cycle
	// only need to run the remainder from 1bil / cycle-length
	for i := (1000000000 - 1) % cycleLength; i > 0; i-- {
		dance(input)
	}

	fmt.Println("p2=", toString())
}

func dance(input string) {
	for _, m := range strings.Split(input, ",") {
		switch m[0] {
		case 's':
			var x int
			fmt.Sscanf(m, "s%d", &x)
			spin(x)
		case 'x':
			var a, b int
			fmt.Sscanf(m, "x%d/%d", &a, &b)
			exchange(a, b)
		case 'p':
			var a, b rune
			fmt.Sscanf(m, "p%c/%c", &a, &b)
			partner(a, b)
		}
	}
}

func toString() string {
	var sb strings.Builder
	for _, p := range byPos {
		sb.WriteRune(p.name)
	}
	return sb.String()
}

type program struct {
	name rune
	pos  int
}

func spin(x int) {
	i := len(byPos) - x
	byPos = append(byPos[i:], byPos[:i]...)
	for i, p := range byPos {
		p.pos = i
	}
}

func exchange(a, b int) {
	pA, pB := byPos[a], byPos[b]
	pA.pos, pB.pos = b, a
	byPos[a], byPos[b] = pB, pA
}

func partner(a, b rune) {
	pA, pB := byName[a], byName[b]
	byPos[pA.pos], byPos[pB.pos] = pB, pA
	pA.pos, pB.pos = pB.pos, pA.pos
}
