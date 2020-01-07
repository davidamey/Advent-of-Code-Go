package grid

import (
	vector "advent-of-code-go/util/vector"
	"fmt"
)

type Grid struct {
	Min, Max vector.Vec
	entries  map[vector.Vec]interface{}
}

func New() *Grid {
	g := &Grid{
		Min:     vector.NewMax(),
		Max:     vector.NewMin(),
		entries: make(map[vector.Vec]interface{}),
	}
	return g
}

func FromLines(lines []string) *Grid {
	g := New()
	for y, l := range lines {
		for x, r := range l {
			g.SetAt(x, y, r)
		}
	}
	return g
}

func (g *Grid) Clone() *Grid {
	ng := &Grid{
		Min:     g.Min,
		Max:     g.Max,
		entries: make(map[vector.Vec]interface{}, len(g.entries)),
	}
	for v, x := range g.entries {
		ng.entries[v] = x
	}
	return ng
}

func (g *Grid) resizeFor(v vector.Vec) {
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

func (g *Grid) Entry(v vector.Vec) interface{} {
	return g.entries[v]
}
func (g *Grid) Int(v vector.Vec) int {
	return g.entries[v].(int)
}
func (g *Grid) Rune(v vector.Vec) rune {
	return g.entries[v].(rune)
}
func (g *Grid) EntryAt(x, y int) interface{} {
	return g.entries[vector.New(x, y)]
}
func (g *Grid) IntAt(x, y int) int {
	return g.entries[vector.New(x, y)].(int)
}
func (g *Grid) RuneAt(x, y int) rune {
	return g.entries[vector.New(x, y)].(rune)
}

func (g *Grid) Col(x int) []interface{} {
	col := make([]interface{}, 1+g.Max.Y-g.Min.Y)
	for y := g.Min.Y; y <= g.Max.Y; y++ {
		col[y] = g.EntryAt(x, y)
	}
	return col
}

func (g *Grid) Row(y int) []interface{} {
	row := make([]interface{}, 1+g.Max.X-g.Min.X)
	for x := g.Min.X; x <= g.Max.X; x++ {
		row[x] = g.EntryAt(x, y)
	}
	return row
}

func (g *Grid) Set(v vector.Vec, i interface{}) {
	g.entries[v] = i
	g.resizeFor(v)
}
func (g *Grid) SetAt(x, y int, i interface{}) {
	g.Set(vector.New(x, y), i)
}

func (g *Grid) Fill(v vector.Vec, w, h int, i interface{}) {
	for y := v.Y; y < v.Y+h; y++ {
		for x := v.X; x < v.X+w; x++ {
			g.SetAt(x, y, i)
		}
	}
}
func (g *Grid) FillAt(x, y, w, h int, i interface{}) {
	g.Fill(vector.New(x, y), w, h, i)
}

func (g *Grid) InBounds(v vector.Vec) bool {
	return v.Within(g.Min, g.Max)
}
func (g *Grid) InBoundsAt(x, y int) bool {
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

func (g *Grid) SubGrid(x, y, w, h int) *Grid {
	sg := New()
	for sx := 0; sx < w; sx++ {
		for sy := 0; sy < h; sy++ {
			sg.SetAt(sx, sy, g.EntryAt(x+sx, y+sy))
		}
	}
	return sg
}

func (g *Grid) ForEach(fn func(v vector.Vec, i interface{})) {
	for y := g.Min.Y; y <= g.Max.Y; y++ {
		for x := g.Min.X; x <= g.Max.X; x++ {
			v := vector.New(x, y)
			fn(v, g.Entry(v))
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
			fmt.Printf(format, g.EntryAt(x, y))
		}
		fmt.Println()
	}
}

func (g *Grid) PrintRunes() {
	g.Print("%c", false)
}
