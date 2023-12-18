package grid

import (
	vector "advent-of-code-go/util/vector"
	"fmt"
)

type Grid[T any] struct {
	Min, Max vector.Vec
	entries  map[vector.Vec]T
}

func New[T any]() *Grid[T] {
	g := &Grid[T]{
		Min:     vector.NewMax(),
		Max:     vector.NewMin(),
		entries: make(map[vector.Vec]T),
	}
	return g
}

func FromLines(lines []string) *Grid[rune] {
	g := New[rune]()
	for y, l := range lines {
		for x, r := range l {
			g.SetAt(x, y, r)
		}
	}
	return g
}

func (g *Grid[T]) Clone() *Grid[T] {
	ng := &Grid[T]{
		Min:     g.Min,
		Max:     g.Max,
		entries: make(map[vector.Vec]T, len(g.entries)),
	}
	for v, x := range g.entries {
		ng.entries[v] = x
	}
	return ng
}

func (g *Grid[T]) resizeFor(v vector.Vec) {
	if v.X < g.Min.X {
		g.Min.X = v.X
	}
	if v.Y < g.Min.Y {
		g.Min.Y = v.Y
	}
	if v.X > g.Max.X {
		g.Max.X = v.X
	}
	if v.Y > g.Max.Y {
		g.Max.Y = v.Y
	}
}

func (g *Grid[T]) Get(v vector.Vec) T {
	return g.entries[v]
}

func (g *Grid[T]) GetAt(x, y int) T {
	return g.entries[vector.New(x, y)]
}

func (g *Grid[T]) Set(v vector.Vec, i T) {
	g.entries[v] = i
	g.resizeFor(v)
}
func (g *Grid[T]) SetAt(x, y int, i T) {
	g.Set(vector.New(x, y), i)
}

func (g *Grid[T]) Col(x int) []T {
	col := make([]T, 1+g.Max.Y-g.Min.Y)
	for y := g.Min.Y; y <= g.Max.Y; y++ {
		col[y] = g.GetAt(x, y)
	}
	return col
}

func (g *Grid[T]) Row(y int) []T {
	row := make([]T, 1+g.Max.X-g.Min.X)
	for x := g.Min.X; x <= g.Max.X; x++ {
		row[x] = g.GetAt(x, y)
	}
	return row
}

func (g *Grid[T]) Fill(v vector.Vec, w, h int, i T) {
	for y := v.Y; y < v.Y+h; y++ {
		for x := v.X; x < v.X+w; x++ {
			g.SetAt(x, y, i)
		}
	}
}
func (g *Grid[T]) FillAt(x, y, w, h int, i T) {
	g.Fill(vector.New(x, y), w, h, i)
}

func (g *Grid[T]) InBounds(v vector.Vec) bool {
	return v.Within(g.Min, g.Max)
}
func (g *Grid[T]) InBoundsAt(x, y int) bool {
	return g.InBounds(vector.New(x, y))
}

// func (g *Grid) Corners() []Vec {
// 	return []Vec{
// 		g.Min,
// 		Vec{g.Min.X, g.Max.Y},
// 		Vec{g.Max.X, g.Min.Y},
// 		g.Max,
// 	}
// }

func (g *Grid[T]) SubGrid(x, y, w, h int) *Grid[T] {
	sg := New[T]()
	for sx := 0; sx < w; sx++ {
		for sy := 0; sy < h; sy++ {
			sg.SetAt(sx, sy, g.GetAt(x+sx, y+sy))
		}
	}
	return sg
}

func (g *Grid[T]) ForEach(fn func(v vector.Vec, i T)) {
	for y := g.Min.Y; y <= g.Max.Y; y++ {
		for x := g.Min.X; x <= g.Max.X; x++ {
			v := vector.New(x, y)
			fn(v, g.Get(v))
		}
	}
}

func (g *Grid[T]) RotateCW() {
	h := g.Clone()
	g.entries = make(map[vector.Vec]T, len(g.entries))
	g.Min.X, g.Min.Y = g.Min.Y, g.Min.X
	g.Max.X, g.Max.Y = g.Max.Y, g.Max.X
	for v, r := range h.entries {
		g.entries[vector.New(g.Max.X-v.Y, v.X)] = r
	}
}

func (g *Grid[T]) Print(format string, clear bool) {
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

func (g *Grid[T]) PrintRunes() {
	g.Print("%c", false)
}
