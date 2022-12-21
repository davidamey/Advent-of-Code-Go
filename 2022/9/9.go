package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/vector"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	var h vector.Vec
	var ts [9]vector.Vec
	var visited [9]map[vector.Vec]struct{}
	for i := range visited {
		visited[i] = map[vector.Vec]struct{}{}
	}

	for _, l := range lines {
		var d rune
		var n int
		fmt.Sscanf(l, "%c %d", &d, &n)

		for i := 0; i < n; i++ {
			switch d {
			case 'U':
				h.Y++
			case 'R':
				h.X++
			case 'D':
				h.Y--
			case 'L':
				h.X--
			}

			ts[0] = moveTail(ts[0], h)
			for i := range ts {
				if i > 0 {
					ts[i] = moveTail(ts[i], ts[i-1])
				}
				visited[i][ts[i]] = struct{}{}
			}
		}
	}

	fmt.Println("p1=", len(visited[0]))
	fmt.Println("p2=", len(visited[8]))
}

func moveTail(from, to vector.Vec) vector.Vec {
	switch {
	case from.Touches(to):
		return from
	case from.X == to.X:
		return vector.New(from.X, (to.Y+from.Y)/2)
	case from.Y == to.Y:
		return vector.New((to.X+from.X)/2, from.Y)
	}

	for _, v := range []vector.Vec{
		{X: from.X - 1, Y: from.Y - 1},
		{X: from.X + 1, Y: from.Y + 1},
		{X: from.X + 1, Y: from.Y - 1},
		{X: from.X - 1, Y: from.Y + 1},
	} {
		if v.Touches(to) {
			return v
		}
	}

	panic("Rope snapped?")
}
