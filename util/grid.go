package util

import (
	"fmt"
)

type Grid struct {
	Min, Max Point
	entries  map[Point]interface{}
}

func NewGrid() *Grid {
	g := &Grid{
		Min:     NewMaxPoint(),
		Max:     NewMinPoint(),
		entries: make(map[Point]interface{}),
	}
	return g
}

func (g *Grid) Clone() *Grid {
	ng := &Grid{
		Min:     g.Min,
		Max:     g.Max,
		entries: make(map[Point]interface{}, len(g.entries)),
	}
	for p, v := range g.entries {
		ng.entries[p] = v
	}
	return ng
}

func (g *Grid) resizeFor(p Point) {
	switch {
	case p.X < g.Min.X:
		g.Min.X = p.X
	case p.Y < g.Min.Y:
		g.Min.Y = p.Y
	case p.X > g.Max.X:
		g.Max.X = p.X
	case p.Y > g.Max.Y:
		g.Max.Y = p.Y
	}
}

func (g *Grid) Get(p Point) interface{} {
	return g.entries[p]
}
func (g *Grid) GetInt(p Point) int {
	return g.entries[p].(int)
}
func (g *Grid) GetRune(p Point) rune {
	return g.entries[p].(rune)
}
func (g *Grid) GetAt(x, y int) interface{} {
	return g.entries[Point{X: x, Y: y}]
}
func (g *Grid) GetIntAt(x, y int) int {
	return g.entries[Point{X: x, Y: y}].(int)
}
func (g *Grid) GetRuneAt(x, y int) rune {
	return g.entries[Point{X: x, Y: y}].(rune)
}

func (g *Grid) Set(p Point, i interface{}) {
	g.entries[p] = i
	g.resizeFor(p)
}
func (g *Grid) SetAt(x, y int, i interface{}) {
	g.Set(Point{X: x, Y: y}, i)
}

func (g *Grid) InBounds(p Point) bool {
	return p.Within(g.Min, g.Max)
}
func (g *Grid) InBoundsAt(x, y int) bool {
	return g.InBounds(Point{X: x, Y: y})
}

// func (g *Grid) Corners() []Point {
// 	return []Point{
// 		g.Min,
// 		Point{g.Min.X, g.Max.Y},
// 		Point{g.Max.X, g.Min.Y},
// 		g.Max,
// 	}
// }

func (g *Grid) ForEach(fn func(p Point, v interface{})) {
	for y := g.Min.Y; y <= g.Max.Y; y++ {
		for x := g.Min.X; x <= g.Max.X; x++ {
			p := Point{X: x, Y: y}
			fn(p, g.Get(p))
		}
	}
}

func (g *Grid) Print(format string, clear bool) {
	if clear {
		fmt.Printf("\033[0;0H")
		fmt.Printf("\033[2J")
	}

	// spew.Dump(g)

	for y := g.Min.Y; y <= g.Max.Y; y++ {
		for x := g.Min.X; x <= g.Max.X; x++ {
			fmt.Printf(format, g.GetAt(x, y))
		}
		fmt.Println()
	}
}
