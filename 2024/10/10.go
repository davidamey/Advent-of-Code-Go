package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	g := grid.FromLines(lines)

	var trailheads []vector.Vec
	g.ForEach(func(v vector.Vec, r rune) {
		if r == '0' {
			trailheads = append(trailheads, v)
		}
	})

	p1, p2 := 0, 0
	for _, th := range trailheads {
		s, r := scoreTrail(g, th)
		p1 += s
		p2 += r
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func scoreTrail(g *grid.Grid[rune], start vector.Vec) (score, rating int) {
	finishes := map[vector.Vec]struct{}{}

	queue := []vector.Vec{start}
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		x := g.Get(q)
		if x == '9' {
			rating++
			finishes[q] = struct{}{}
			continue
		}

		for _, a := range q.Adjacent(false) {
			if g.InBounds(a) && g.Get(a)-x == 1 {
				queue = append(queue, a)
			}
		}
	}
	return len(finishes), rating
}
