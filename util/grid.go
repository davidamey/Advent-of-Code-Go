package util

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) Left() Point {
	return Point{p.X - 1, p.Y}
}
func (p *Point) Right() Point {
	return Point{p.X + 1, p.Y}
}
func (p *Point) Up() Point {
	return Point{p.X, p.Y - 1}
}
func (p *Point) Down() Point {
	return Point{p.X, p.Y + 1}
}

func (p *Point) ManhattanTo(to Point) int {
	return AbsInt(to.X-p.X) + AbsInt(to.Y-p.Y)
}

func (p *Point) Adjacent(withCorners bool) []Point {
	if withCorners {
		return []Point{
			Point{p.X - 1, p.Y - 1},
			Point{p.X, p.Y - 1},
			Point{p.X + 1, p.Y - 1},
			Point{p.X - 1, p.Y},
			Point{p.X + 1, p.Y},
			Point{p.X - 1, p.Y + 1},
			Point{p.X, p.Y + 1},
			Point{p.X + 1, p.Y + 1},
		}
	}
	return []Point{
		Point{p.X, p.Y - 1},
		Point{p.X - 1, p.Y},
		Point{p.X + 1, p.Y},
		Point{p.X, p.Y + 1},
	}
}

func (p *Point) Within(b1, b2 Point) bool {
	return p.X >= b1.X && p.X < b2.X &&
		p.Y >= b1.Y && p.Y < b2.Y
}

type Grid [][]rune

func NewGrid(lines []string, fn func(r rune) rune) *Grid {
	g := make(Grid, len(lines))
	for y, l := range lines {
		g[y] = make([]rune, len(l))
		for x, r := range l {
			if fn != nil {
				r = fn(r)
			}
			g[y][x] = r
		}
	}
	return &g
}

func (g *Grid) Clone() *Grid {
	clone := make(Grid, len(*g))
	for y, row := range *g {
		clone[y] = make([]rune, len(row))
		copy(clone[y], row)
	}
	return &clone
}

func (g *Grid) Bounds() (b1, b2 Point) {
	h := len(*g)
	w := 0
	if h > 0 {
		w = len((*g)[0])
	}
	return Point{0, 0}, Point{w, h}
}

func (g *Grid) Print(clear bool) {
	if clear {
		fmt.Printf("\033[0;0H")
		fmt.Printf("\033[2J")
	}

	for _, row := range *g {
		for _, r := range row {
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
	fmt.Println()
}
