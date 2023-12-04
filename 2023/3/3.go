package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
	"time"
	"unicode"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	g := grid.FromLines(lines)

	p1 := 0
	var parts []*Part
	var possibleGears []vector.Vec
	for y := g.Min.Y; y <= g.Max.Y; y++ {
		for x := g.Min.X; x <= g.Max.X; x++ {
			v := vector.New(x, y)
			if g.Get(v) == '*' {
				possibleGears = append(possibleGears, v)
				continue
			}

			if p, valid := NewPart(g, v); valid {
				parts = append(parts, p)
				if len(p.adjSym) > 0 {
					p1 += p.val
				}
				x += p.len()
			}
		}
	}

	p2 := 0
	for _, v := range possibleGears {
		var adjParts []int
		for _, p := range parts {
			if _, isAdj := p.adjSym[v]; isAdj {
				adjParts = append(adjParts, p.val)
			}
		}
		if len(adjParts) == 2 {
			p2 += adjParts[0] * adjParts[1]
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

type Part struct {
	g          *grid.Grid[rune]
	start, end vector.Vec
	val        int
	adjSym     map[vector.Vec]rune
}

func NewPart(g *grid.Grid[rune], from vector.Vec) (part *Part, validPart bool) {
	r := g.Get(from)

	if !unicode.IsDigit(r) {
		return nil, false
	}

	part = &Part{g: g, start: from, end: from}

	num := string(r)
	for w := from.Right(); g.InBounds(w); w = w.Right() {
		r := g.Get(w)
		if !unicode.IsDigit(r) {
			break
		}
		part.end = w
		num += string(r)
	}
	part.val = util.Atoi(num)

	part.adjSym = part.getAdjacentSym()
	return part, true
}

func (p *Part) len() int {
	return p.end.X - p.start.X
}

func (p *Part) getAdjacentSym() map[vector.Vec]rune {
	adj := make(map[vector.Vec]rune)
	for v := p.start; v != p.end.Right(); v = v.Right() {
		for _, w := range v.Adjacent(true) {
			if !p.g.InBounds(w) {
				continue
			}

			if r := p.g.Get(w); r != '.' && !unicode.IsDigit(r) {
				adj[w] = r
			}
		}
	}
	return adj
}
