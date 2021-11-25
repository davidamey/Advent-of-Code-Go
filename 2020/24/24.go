package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	f := newFloor()
	for _, l := range lines {
		p := parsePath(l)
		f.tiles[p] = !f.tiles[p]
		if f.tiles[p] {
			f.resizeFor(p)
		}
	}

	fmt.Println("p1=", f.countBlack())

	for i := 0; i < 100; i++ {
		f.evolve()
	}
	fmt.Println("p2=", f.countBlack())
}

type point struct{ x, y int }

func (p *point) neighbours() []point {
	return []point{
		{p.x + 1, p.y},     // e
		{p.x, p.y - 1},     // se
		{p.x - 1, p.y - 1}, // sw
		{p.x - 1, p.y},     // w
		{p.x, p.y + 1},     // nw
		{p.x + 1, p.y + 1}, // ne
	}
}

type floor struct {
	tiles    map[point]bool
	min, max point
}

func newFloor() *floor {
	return &floor{tiles: make(map[point]bool)}
}

func (f *floor) copy() *floor {
	g := &floor{make(map[point]bool), f.min, f.max}
	for p, t := range f.tiles {
		g.tiles[p] = t
	}
	return g
}

func (f *floor) countBlack() (cb int) {
	for _, t := range f.tiles {
		if t {
			cb++
		}
	}
	return
}

func (f *floor) resizeFor(p point) {
	if p.x < f.min.x {
		f.min.x = p.x
	}
	if p.y < f.min.y {
		f.min.y = p.y
	}
	if p.x > f.max.x {
		f.max.x = p.x
	}
	if p.y > f.max.y {
		f.max.y = p.y
	}
}

func (f *floor) evolve() {
	g := f.copy()

	for y := g.min.y - 1; y <= g.max.y+1; y++ {
		for x := g.min.x - 1; x <= g.max.x+1; x++ {
			p := point{x, y}
			cb := 0
			for _, n := range p.neighbours() {
				if g.tiles[n] {
					cb++
				}
			}

			switch {
			case g.tiles[p] && (cb == 0 || cb > 2):
				f.tiles[p] = false
			case !g.tiles[p] && cb == 2:
				f.tiles[p] = true
				f.resizeFor(p)
			}
		}
	}
}

func parsePath(path string) (p point) {
	i := 0
loop:
	for j := 1; j <= len(path); j++ {
		mv := path[i:j]
		switch mv {
		case "e":
			p.x++
		case "se":
			p.y--
		case "sw":
			p.x--
			p.y--
		case "w":
			p.x--
		case "nw":
			p.y++
		case "ne":
			p.x++
			p.y++
		default:
			continue loop
		}
		i = j
	}
	return p
}
