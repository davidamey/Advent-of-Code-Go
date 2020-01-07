package main

import (
	"advent-of-code-go/util"
	"advent/util/grid"
	"advent/util/vector"
	"fmt"
)

func main() {
	// input := util.MustReadFileToLines("example")
	input := util.MustReadFileToLines("input")

	// initial grid
	g := grid.New()
	for y, l := range input {
		for x, ch := range l {
			g.SetAt(x, y, ch)
		}
	}

	// Assume odd-sized square grid
	if (1+g.Max.X)%2 == 0 || (1+g.Max.Y)%2 == 0 {
		panic("even grid")
	}

	fmt.Println("p1=", p1(g.Clone()))
	fmt.Println("p2=", p2(g.Clone()))
}

func p1(g *grid.Grid) (p1 int) {
	c := &carrier{
		p: vector.New((1+g.Max.X)/2, (1+g.Max.Y)/2),
		v: vector.New(0, -1),
	}

	// bursts
	for b := 0; b < 10000; b++ {
		e := g.Entry(c.p)
		if e != nil && e.(rune) == '#' {
			c.v.X, c.v.Y = -c.v.Y, c.v.X // turn right
			g.Set(c.p, '.')
		} else {
			c.v.X, c.v.Y = c.v.Y, -c.v.X // turn left
			g.Set(c.p, '#')
			p1++
		}
		c.p = c.p.Add(c.v)
	}
	return
}

func p2(g *grid.Grid) (p2 int) {
	c := &carrier{
		p: vector.New((1+g.Max.X)/2, (1+g.Max.Y)/2),
		v: vector.New(0, -1),
	}

	// bursts
	for b := 0; b < 10000000; b++ {
		r := '.'
		if e := g.Entry(c.p); e != nil {
			r = e.(rune)
		}
		switch r {
		case '.':
			c.v.X, c.v.Y = c.v.Y, -c.v.X // turn left
			g.Set(c.p, 'W')
		case 'W':
			g.Set(c.p, '#')
			p2++
		case '#':
			c.v.X, c.v.Y = -c.v.Y, c.v.X // turn right
			g.Set(c.p, 'F')
		case 'F':
			c.v.X, c.v.Y = -c.v.X, -c.v.Y // reverse
			g.Set(c.p, '.')
		}
		c.p = c.p.Add(c.v)
	}
	return
}

type carrier struct {
	p, v vector.Vec
}

func printGrid(g *grid.Grid, c *carrier) {
	g = g.Clone()

	g.Min = g.Min.Add(vector.New(-3, -3))
	g.Max = g.Max.Add(vector.New(3, 3))

	printCloseBracket := false
	for y := g.Min.Y; y <= g.Max.Y; y++ {
		for x := g.Min.X; x <= g.Max.X; x++ {
			r := '.'
			if e := g.EntryAt(x, y); e != nil {
				r = e.(rune)
			}

			switch {
			case x == c.p.X && y == c.p.Y:
				fmt.Printf("[%c", r)
				printCloseBracket = true
			case printCloseBracket:
				fmt.Printf("]%c", r)
				printCloseBracket = false
			case x == g.Min.X:
				fmt.Printf("%c", r)
			default:
				fmt.Printf(" %c", r)
			}
		}
		fmt.Println()
	}
}
