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

	// g := grid.FromLines(util.MustReadFileToLines("example"))
	g := grid.FromLines(util.MustReadFileToLines("input"))

	p1 := evolve(g, func(g *grid.Grid[rune], v vector.Vec, r rune) (evolution rune, changed bool) {
		occupied := 0
		for _, w := range v.Adjacent(true) {
			if g.InBounds(w) && g.Get(w) == '#' {
				occupied++
			}
		}

		switch {
		case r == 'L' && occupied == 0:
			return '#', true
		case r == '#' && occupied >= 4:
			return 'L', true
		default:
			return r, false
		}
	})

	p2 := evolve(g, func(g *grid.Grid[rune], v vector.Vec, r rune) (evolution rune, changed bool) {
		occupied := 0
		for _, dir := range []vector.Vec{{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 1, Y: -1}, {X: -1, Y: 0}, {X: 1, Y: 0}, {X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1}} {
			for w := v.Add(dir); g.InBounds(w); w = w.Add(dir) {
				if s := g.Get(w); s == '#' {
					occupied++
					break
				} else if s == 'L' {
					break
				}
			}
		}

		switch {
		case r == 'L' && occupied == 0:
			return '#', true
		case r == '#' && occupied >= 5:
			return 'L', true
		default:
			return r, false
		}
	})

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

type evolveFn func(g *grid.Grid[rune], v vector.Vec, r rune) (evolution rune, changed bool)

func evolve(start *grid.Grid[rune], ef evolveFn) (occupied int) {
	g := start.Clone()
	h := grid.New[rune]()

	for {
		changed := false
		g.ForEach(func(v vector.Vec, r rune) {
			r, c := ef(g, v, r)
			h.Set(v, r)
			if c {
				changed = true
			}
		})

		g, h = h, g
		if !changed {
			break
		}
	}

	g.ForEach(func(v vector.Vec, r rune) {
		if r == '#' {
			occupied++
		}
	})

	return occupied
}
