package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
	"math"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	g := grid.FromLines(lines)

	var starts []vector.Vec
	var end vector.Vec
	g.ForEach(func(v vector.Vec, r rune) {
		switch r {
		case 'S':
			starts = append([]vector.Vec{v}, starts...)
			g.Set(v, 'a')
		case 'a':
			starts = append(starts, v)
		case 'E':
			end = v
			g.Set(v, 'z')
		}
	})

	pathLengths := make([]int, len(starts))
	for i, s := range starts {
		pathLengths[i] = solve(g, s, end)
	}

	fmt.Println("p1=", pathLengths[0])
	fmt.Println("p2=", util.MinInt(pathLengths...))
}

func solve(g *grid.Grid[rune], start, end vector.Vec) int {
	n := g.ShortestPath(start, end, func(v, parent rune, depth int) bool {
		return v <= parent+1
	})
	if n == nil {
		return math.MaxInt
	}
	return n.Length
}
