package util

import "math"

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}
func NewMaxPoint() Point {
	return Point{math.MaxInt32, math.MaxInt32}
}
func NewMinPoint() Point {
	return Point{math.MinInt32, math.MinInt32}
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
	return p.X >= b1.X && p.X <= b2.X &&
		p.Y >= b1.Y && p.Y <= b2.Y
}
