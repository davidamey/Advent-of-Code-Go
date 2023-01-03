package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
)

const steps = 100

func main() {
	lines := util.MustReadFileToLines("input")

	initial := grid.New[rune]()
	for y, l := range lines {
		for x, r := range l {
			initial.SetAt(x, y, r)
		}
	}

	g1 := initial.Clone()

	// g2 has each corner stuck on
	g2 := initial.Clone()
	g2.SetAt(g2.Min.X, g2.Min.Y, '#')
	g2.SetAt(g2.Min.X, g2.Max.Y, '#')
	g2.SetAt(g2.Max.X, g2.Min.Y, '#')
	g2.SetAt(g2.Max.X, g2.Max.Y, '#')

	// Evolve
	for i := 0; i < steps; i++ {
		g1 = EvolveP1(g1)
		g2 = EvolveP2(g2)
	}

	p1 := 0
	g1.ForEach(func(v vector.Vec, r rune) {
		if r == '#' {
			p1++
		}
	})
	fmt.Println("p1=", p1)

	p2 := 0
	g2.ForEach(func(v vector.Vec, r rune) {
		if r == '#' {
			p2++
		}
	})
	fmt.Println("p2=", p2)
}

func EvolveP1(g *grid.Grid[rune]) *grid.Grid[rune] {
	ng := g.Clone()
	g.ForEach(func(v vector.Vec, r rune) {
		ng.Set(v, EvolvePoint(g, v))
	})
	return ng
}

func EvolveP2(g *grid.Grid[rune]) *grid.Grid[rune] {
	ng := g.Clone()
	g.ForEach(func(v vector.Vec, r rune) {
		switch v {
		case
			vector.Vec{X: g.Min.X, Y: g.Min.Y},
			vector.Vec{X: g.Min.X, Y: g.Max.Y},
			vector.Vec{X: g.Max.X, Y: g.Min.Y},
			vector.Vec{X: g.Max.X, Y: g.Max.Y}:
			return
		default:
			ng.Set(v, EvolvePoint(g, v))
		}
	})
	return ng
}

func EvolvePoint(g *grid.Grid[rune], v vector.Vec) rune {
	sum := RuneToScore(g.Get(v))
	for _, a := range v.Adjacent(true) {
		if g.InBounds(a) {
			sum += RuneToScore(g.Get(a))
		}
	}

	switch sum {
	case 3:
		return '#'
	case 4:
		return g.Get(v)
	default:
		return '.'
	}
}

func RuneToScore(r rune) int {
	if r == '#' {
		return 1
	}
	return 0
}
