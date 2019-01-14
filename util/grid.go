package util

import (
	"fmt"
)

type Grid struct {
	Min, Max Vec
	entries  map[Vec]interface{}
}

func NewGrid() *Grid {
	g := &Grid{
		Min:     NewMaxVec(),
		Max:     NewMinVec(),
		entries: make(map[Vec]interface{}),
	}
	return g
}

func (g *Grid) Clone() *Grid {
	ng := &Grid{
		Min:     g.Min,
		Max:     g.Max,
		entries: make(map[Vec]interface{}, len(g.entries)),
	}
	for v, x := range g.entries {
		ng.entries[v] = x
	}
	return ng
}

func (g *Grid) resizeFor(v Vec) {
	switch {
	case v.X < g.Min.X:
		g.Min.X = v.X
	case v.Y < g.Min.Y:
		g.Min.Y = v.Y
	case v.X > g.Max.X:
		g.Max.X = v.X
	case v.Y > g.Max.Y:
		g.Max.Y = v.Y
	}
}

func (g *Grid) Get(v Vec) interface{} {
	return g.entries[v]
}
func (g *Grid) GetInt(v Vec) int {
	return g.entries[v].(int)
}
func (g *Grid) GetRune(v Vec) rune {
	return g.entries[v].(rune)
}
func (g *Grid) GetAt(x, y int) interface{} {
	return g.entries[Vec{X: x, Y: y}]
}
func (g *Grid) GetIntAt(x, y int) int {
	return g.entries[Vec{X: x, Y: y}].(int)
}
func (g *Grid) GetRuneAt(x, y int) rune {
	return g.entries[Vec{X: x, Y: y}].(rune)
}

func (g *Grid) Set(v Vec, i interface{}) {
	g.entries[v] = i
	g.resizeFor(v)
}
func (g *Grid) SetAt(x, y int, i interface{}) {
	g.Set(Vec{X: x, Y: y}, i)
}

func (g *Grid) InBounds(v Vec) bool {
	return v.Within(g.Min, g.Max)
}
func (g *Grid) InBoundsAt(x, y int) bool {
	return g.InBounds(Vec{X: x, Y: y})
}

// func (g *Grid) Corners() []Vec {
// 	return []Vec{
// 		g.Min,
// 		Vec{g.Min.X, g.Max.Y},
// 		Vec{g.Max.X, g.Min.Y},
// 		g.Max,
// 	}
// }

func (g *Grid) ForEach(fn func(v Vec, i interface{})) {
	for y := g.Min.Y; y <= g.Max.Y; y++ {
		for x := g.Min.X; x <= g.Max.X; x++ {
			v := Vec{X: x, Y: y}
			fn(v, g.Get(v))
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
