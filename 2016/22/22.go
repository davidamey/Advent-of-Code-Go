package main

import (
	"advent/util"
	"advent/util/grid"
	"advent/util/vector"
	"fmt"
)

func main() {
	lines := util.MustReadFileToLines("input")
	nodes := make([]*node, len(lines)-2)
	for i, l := range lines[2:] {
		nodes[i] = newNode(l)
	}

	p1(nodes)
	p2(nodes)
}

func p1(nodes []*node) {
	pairs := 0
	for _, a := range nodes {
		for _, b := range nodes {
			if viablePair(a, b) {
				pairs++
			}
		}
	}
	fmt.Println("p1=", pairs)
}

func p2(nodes []*node) {
	minSize := -1
	maxX := 0
	var empty *node
	for _, n := range nodes {
		if minSize == -1 || n.size < minSize {
			minSize = n.size
		}
		if n.x > maxX {
			maxX = n.x
		}
		if n.used == 0 {
			empty = n
		}
	}

	g := grid.New()
	for _, n := range nodes {
		c := '.'
		switch {
		case n.used == 0:
			c = '_'
		case n.used > minSize:
			c = '#'
			if n.y == 0 {
				panic("can't solve when blocks on top row")
			}
		case n.x == maxX && n.y == 0:
			c = 'G'
		}

		g.SetAt(n.x, n.y, c)
	}

	path := g.ShortestPath(vector.New(empty.x, empty.y), vector.New(maxX-1, 0), func(v interface{}, depth int) bool {
		return v.(rune) == '.'
	})

	// path.length puts the empty drive next to goal data
	// +1 to switch empty and goal
	// 5 moves for empty to go around to left of goal and switch with it again
	fmt.Println("p2=", 1+path.Length+(maxX-1)*5)
}

func viablePair(a, b *node) bool {
	if a.used == 0 {
		return false
	}

	if a == b {
		return false
	}

	return a.used <= b.avail
}

type node struct {
	x, y              int
	size, used, avail int
	usePct            int
}

func newNode(raw string) *node {
	var x, y, size, used, avail, usePct int
	fmt.Sscanf(raw, "/dev/grid/node-x%d-y%d %dT %dT %dT %d%%", &x, &y, &size, &used, &avail, &usePct)
	return &node{x, y, size, used, avail, usePct}
}
