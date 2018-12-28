package main

import (
	"advent/util"
	"fmt"
)

func main() {
	// file, _ := util.OpenExample()
	// file, _ := util.OpenFile("example2")
	// file, _ := util.OpenFile("example3")
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	points := make([]*Point, len(lines))
	for i, l := range lines {
		points[i] = NewPoint(l)
	}

	for _, p1 := range points {
		// fmt.Printf("(%2d,%2d,%2d,%2d)\n", p1.X, p1.Y, p1.Z, p1.T)
		joined := false
		for _, p2 := range points {
			if p1 == p2 {
				continue
			}

			d := p1.ManhattanTo(p2)
			// fmt.Printf("  (%2d,%2d,%2d,%2d) = %2d\n",
			// 	p2.X, p2.Y, p2.Z, p2.T, d)
			if d <= 3 {
				joined = true
				p1.JoinWith(p2)
			}
		}

		// Constellation on its own
		if !joined {
			NewConstellation(p1)
		}
	}

	galaxy := make(map[*Constellation]struct{})
	for _, p := range points {
		galaxy[p.C] = struct{}{}
	}

	fmt.Println("Constellations =", len(galaxy))
	// i := 0
	// for c := range galaxy {
	// 	i++
	// 	fmt.Printf("Constellation %d\n", i)
	// 	for p := range c.Points {
	// 		fmt.Printf("  (%d,%d,%d,%d)\n", p.X, p.Y, p.Z, p.T)
	// 	}
	// }

}

type Point struct {
	X, Y, Z, T int
	C          *Constellation
}

func NewPoint(raw string) *Point {
	var x, y, z, t int
	fmt.Sscanf(raw, "%d,%d,%d,%d", &x, &y, &z, &t)
	return &Point{x, y, z, t, nil}
}

func (p *Point) ManhattanTo(p2 *Point) int {
	return util.AbsInt(p2.X-p.X) +
		util.AbsInt(p2.Y-p.Y) +
		util.AbsInt(p2.Z-p.Z) +
		util.AbsInt(p2.T-p.T)
}

func (p1 *Point) JoinWith(p2 *Point) *Constellation {
	if p1.C == nil && p2.C == nil {
		return NewConstellation(p1, p2)
	}

	if p1.C != nil && p2.C != nil {
		return p1.C.MergeWith(p2.C)
	}

	if p1.C != nil {
		return p1.C.AddPoint(p2)
	} else {
		return p2.C.AddPoint(p1)
	}
}

type Constellation struct {
	Points map[*Point]struct{}
}

func NewConstellation(points ...*Point) *Constellation {
	c := &Constellation{
		make(map[*Point]struct{}),
	}
	for _, p := range points {
		c.AddPoint(p)
	}
	return c
}

// func NewConstellation(points ...*Point) *Constellation {
// 	c := &Constellation{points}
// 	for _, p := range points {
// 		p.C = c
// 	}
// 	return c
// }

func (c *Constellation) AddPoint(p *Point) *Constellation {
	c.Points[p] = struct{}{}
	// c.Points = append(c.Points, p)
	p.C = c
	return c
}

func (c1 *Constellation) MergeWith(c2 *Constellation) *Constellation {
	for p := range c2.Points {
		c1.AddPoint(p)
	}
	return c1
}
