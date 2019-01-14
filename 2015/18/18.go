package main

import (
	"advent/util"
	"fmt"
)

const steps = 100

func main() {
	// file, _ := util.OpenExample()
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	initial := util.NewGrid()
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
	g1.ForEach(func(p util.Point, v interface{}) {
		if v.(rune) == '#' {
			p1++
		}
	})
	fmt.Println("p1=", p1)

	p2 := 0
	g2.ForEach(func(p util.Point, v interface{}) {
		if v.(rune) == '#' {
			p2++
		}
	})
	fmt.Println("p2=", p2)
}

func EvolveP1(g *util.Grid) *util.Grid {
	ng := g.Clone()
	g.ForEach(func(p util.Point, v interface{}) {
		ng.Set(p, EvolvePoint(g, p))
	})
	return ng
}

func EvolveP2(g *util.Grid) *util.Grid {
	ng := g.Clone()
	g.ForEach(func(p util.Point, v interface{}) {
		switch p {
		case
			util.Point{X: g.Min.X, Y: g.Min.Y},
			util.Point{X: g.Min.X, Y: g.Max.Y},
			util.Point{X: g.Max.X, Y: g.Min.Y},
			util.Point{X: g.Max.X, Y: g.Max.Y}:
			return
		default:
			ng.Set(p, EvolvePoint(g, p))
		}
	})
	return ng
}

func EvolvePoint(g *util.Grid, p util.Point) rune {
	sum := RuneToScore(g.GetRune(p))
	for _, a := range p.Adjacent(true) {
		if g.InBounds(a) {
			sum += RuneToScore(g.GetRune(a))
		}
	}

	switch sum {
	case 3:
		return '#'
	case 4:
		return g.GetRune(p)
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
