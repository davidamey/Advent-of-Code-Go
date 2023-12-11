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

	var galaxies []*galaxy
	g.ForEach(func(v vector.Vec, r rune) {
		if r == '#' {
			galaxies = append(galaxies, &galaxy{pos: v})
		}
	})

	for y := g.Min.Y; y <= g.Max.Y; y++ {
		allSpace := true
		for _, r := range g.Row(y) {
			if r != '.' {
				allSpace = false
				break
			}
		}
		if allSpace {
			for i, g := range galaxies {
				if g.pos.Y > y {
					galaxies[i].scale.Y++
				}
			}
		}
	}

	for x := g.Min.X; x <= g.Max.X; x++ {
		allSpace := true
		for _, r := range g.Col(x) {
			if r != '.' {
				allSpace = false
				break
			}
		}
		if allSpace {
			for i, g := range galaxies {
				if g.pos.X > x {
					galaxies[i].scale.X++
				}
			}
		}
	}

	fmt.Println("p1=", score(galaxies, 2))
	fmt.Println("p2=", score(galaxies, 1_000_000))
}

func score(galaxies []*galaxy, scale int) (s int) {
	for i, g1 := range galaxies {
		scaledG1 := g1.actual(scale)
		for _, g2 := range galaxies[i+1:] {
			s += scaledG1.ManhattanTo(g2.actual(scale))
		}
	}
	return
}

type galaxy struct {
	pos, scale vector.Vec
}

func (g *galaxy) actual(scaleFactor int) vector.Vec {
	return vector.New(
		g.pos.X+(g.scale.X*scaleFactor-g.scale.X),
		g.pos.Y+(g.scale.Y*scaleFactor-g.scale.Y),
	)
}
