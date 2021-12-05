package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/vector"
	"fmt"
)

func main() {
	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	fmt.Println("p1=", solve(lines, false))
	fmt.Println("p2=", solve(lines, true))
}

func solve(lines []string, diagonals bool) (overlaps int) {
	g := make(map[vector.Vec]int)
	for _, l := range lines {
		var x1, y1, x2, y2 int
		fmt.Sscanf(l, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		dx, dy := 0, 0
		if x1 < x2 {
			dx = 1
		}
		if x2 < x1 {
			dx = -1
		}
		if y1 < y2 {
			dy = 1
		}
		if y2 < y1 {
			dy = -1
		}

		if !diagonals && dx != 0 && dy != 0 {
			continue
		}

		for x, y := x1, y1; true; x, y = x+dx, y+dy {
			g[vector.New(x, y)]++
			if x == x2 && y == y2 {
				break
			}
		}
	}
	for _, v := range g {
		if v >= 2 {
			overlaps++
		}
	}
	return
}
