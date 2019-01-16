package main

import (
	"advent/util"
	"fmt"
)

const (
	// example
	// gW = 7
	// gH = 3

	// actual
	gW = 50
	gH = 6
)

func main() {
	// example
	// g := util.NewGrid()
	// rectWithRune(g, gW, gH, '.')
	// rect(g, 3, 2)
	// rotCol(g, 1, 1)
	// rotRow(g, 0, 4)
	// rotCol(g, 1, 1)
	// g.Print("%c", false)

	g := util.NewGrid()
	rectWithRune(g, gW, gH, ' ')

	lines := util.MustReadFileToLines("input")
	for _, l := range lines {
		switch {
		case l[:4] == "rect":
			var w, h int
			fmt.Sscanf(l[5:], "%dx%d", &w, &h)
			rect(g, w, h)
		case l[:10] == "rotate row":
			var y, n int
			fmt.Sscanf(l[11:], "y=%d by %d", &y, &n)
			rotRow(g, y, n)
		case l[:13] == "rotate column":
			var x, n int
			fmt.Sscanf(l[14:], "x=%d by %d", &x, &n)
			rotCol(g, x, n)
		}
	}

	p1 := 0
	g.ForEach(func(v util.Vec, i interface{}) {
		if i.(rune) == '#' {
			p1++
		}
	})

	fmt.Println("p1=", p1)
	fmt.Println("p2=")
	g.Print("%c", false)

}

func rect(g *util.Grid, w, h int) {
	rectWithRune(g, w, h, '#')
}

func rectWithRune(g *util.Grid, w, h int, r rune) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			g.SetAt(x, y, r)
		}
	}
}

func rotCol(g *util.Grid, x, n int) {
	col := g.Col(x)
	for y, r := range col {
		g.SetAt(x, (y+n)%gH, r)
	}
}

func rotRow(g *util.Grid, y, n int) {
	row := g.Row(y)
	for x, r := range row {
		g.SetAt((x+n)%gW, y, r)
	}
}
