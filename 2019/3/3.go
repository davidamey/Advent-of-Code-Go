package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
	"strings"
)

func main() {
	// lines := util.MustReadFileToLines("example")
	// lines := util.MustReadFileToLines("example2")
	// lines := util.MustReadFileToLines("example3")
	// lines := util.MustReadFileToLines("example4")
	lines := util.MustReadFileToLines("input")

	g1, g2 := grid.New(), grid.New()
	addWire(g1, lines[0])
	addWire(g2, lines[1])

	p1, p2 := p1p2(g1, g2)
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func p1p2(g1, g2 *grid.Grid) (p1, p2 int) {
	var vs []vector.Vec
	g1.ForEach(func(v vector.Vec, x interface{}) {
		if !v.IsOrigin() && x != nil && g2.Entry(v) != nil {
			vs = append(vs, v)
		}
	})

	p1, p2 = 1<<16, 1<<16
	for _, v := range vs {
		//p1
		if d := v.Manhattan(); d < p1 {
			p1 = d
		}

		//p2
		d1 := g1.Entry(v).(int)
		d2 := g2.Entry(v).(int)
		if d := d1 + d2; d < p2 {
			p2 = d
		}
	}
	return
}

func addWire(g *grid.Grid, path string) {
	v := vector.New(0, 0)
	l := 1
	for _, p := range strings.Split(path, ",") {
		var d rune
		var x int
		fmt.Sscanf(p, "%c%d", &d, &x)

		for i := 0; i < x; i++ {
			switch d {
			case 'U':
				v.Y--
			case 'R':
				v.X++
			case 'D':
				v.Y++
			case 'L':
				v.X--
			}

			if g.Entry(v) == nil {
				g.Set(v, l)
			}
			l++
		}
	}
}

func print(g *grid.Grid) {
	for y := g.Min.Y; y <= g.Max.Y; y++ {
		for x := g.Min.X; x <= g.Max.X; x++ {
			if x == 0 && y == 0 {
				fmt.Print(" o,")
				continue
			}
			e := g.EntryAt(x, y)
			if e != nil {
				fmt.Printf("%2d,", e)
			} else {
				fmt.Print("  ,")
			}
		}
		fmt.Println()
	}
}
