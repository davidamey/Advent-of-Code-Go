package main

import (
	"advent-of-code-go/util"
	"fmt"
	"math"
	"sort"
)

type point struct {
	X, Y int
	Area int
	ID   int
}

type box struct {
	X1, X2, Y1, Y2 int
}

type grid map[string]int

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	points := make([]*point, len(lines))
	for i, l := range lines {
		p := &point{ID: i}
		fmt.Sscanf(l, "%d, %d", &p.X, &p.Y)
		points[i] = p
	}

	p1points := make([]*point, len(lines))
	copy(p1points, points)

	part1(p1points)
	part2(points)
}

func part1(points []*point) {
	b := getBoundingBox(points)
	g := make(grid)

	for x := b.X1; x <= b.X2; x++ {
		for y := b.Y1; y <= b.Y2; y++ {
			min := math.MaxInt16
			closest := -1
			for i, p := range points {
				d := dist(p.X, p.Y, x, y)

				if d == 0 {
					g.set(x, y, i)
					closest = i
					break
				}

				if d == min {
					closest = -1
					g.set(x, y, -1)
					continue
				}

				if d < min {
					min = d
					g.set(x, y, i)
					closest = i
					continue
				}
			}

			if closest >= 0 {
				points[closest].Area++
			}
		}
	}

	fmt.Println("== part1 ==")
	// g.print(b)

	sort.Slice(points, func(i, j int) bool {
		return points[i].Area > points[j].Area
	})
	for i, p := range points[:3] {
		if !isInfinite(g, b, p) {
			fmt.Printf("%d: %d\n", i, p.Area)
		}
	}
	fmt.Printf("\n")

	// for y := b.Y1; y <= b.Y2; y++ {
	// 	for x := b.X1; x <= b.X2; x++ {
	// 		if grid[y][x] == 0 {
	// 			grid[y][x] = '.'
	// 		}
	// 		fmt.Printf("%c", grid[y][x])
	// 	}
	// 	fmt.Println()
	// }
}

func part2(points []*point) {
	boundary := 10000
	region := 0
	for x := -boundary; x <= boundary; x++ {
		for y := -boundary; y <= boundary; y++ {
			d := 0
			for _, p := range points {
				d += dist(x, y, p.X, p.Y)
			}
			if d < boundary {
				region++
			}
		}
	}

	fmt.Println("== part2 ==")
	fmt.Printf("region size: %d\n", region)
}

func dist(x1, y1, x2, y2 int) int {
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	return dx + dy
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func getBoundingBox(points []*point) *box {
	b := newBox(points[0])
	for _, p := range points {
		if p.X < b.X1 {
			b.X1 = p.X
		}
		if p.X > b.X2 {
			b.X2 = p.X
		}
		if p.Y < b.Y1 {
			b.Y1 = p.Y
		}
		if p.Y > b.Y2 {
			b.Y2 = p.Y
		}
	}
	return b
}

func newBox(p *point) *box {
	return &box{
		X1: p.X,
		X2: p.X,
		Y1: p.Y,
		Y2: p.Y,
	}
}

func isInfinite(g grid, b *box, p *point) bool {
	for x := b.X1; x <= b.X2; x++ {
		for y := b.Y1; y <= b.Y2; y++ {
			if x != b.X1 && x != b.X2 {
				// Not on the x-boundaries so skip middle y's
				if y != b.Y1 && y != b.Y2 {
					continue
				}
			}

			if g.get(x, y) == p.ID {
				return true
			}
		}
	}
	return false
}

func (g grid) set(x, y, i int) {
	loc := fmt.Sprintf("%d,%d", x, y)
	g[loc] = i
}

func (g grid) get(x, y int) int {
	loc := fmt.Sprintf("%d,%d", x, y)
	return g[loc]
}

func (g grid) print(b *box) {
	for y := b.Y1; y <= b.Y2; y++ {
		for x := b.X1; x <= b.X2; x++ {
			fmt.Printf(" %2d ", g.get(x, y))
		}
		fmt.Println()
	}
	fmt.Println()
}
