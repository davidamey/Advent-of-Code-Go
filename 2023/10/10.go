package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	g := grid.FromLines(lines)

	var start vector.Vec
	g.ForEach(func(v vector.Vec, r rune) {
		if g.Get(v) == 'S' {
			start = v
		}
	})

	var dir vector.Vec
	switch {
	case strings.ContainsRune("|7F", g.Get(start.Up())):
		dir.Y = -1
	case strings.ContainsRune("-J7", g.Get(start.Right())):
		dir.X = 1
	case strings.ContainsRune("|LJ", g.Get(start.Down())):
		dir.Y = 1
	case strings.ContainsRune("-LF", g.Get(start.Left())):
		dir.X = -1
	}

	path := []vector.Vec{start}
	for {
		v := path[len(path)-1].Add(dir)
		if v == start {
			break
		}

		path = append(path, v)
		r := g.Get(v)

		switch {
		case r == 'L' && dir.Y == 0: // approaching from the east
			dir.X, dir.Y = 0, -1
		case r == 'L' && dir.X == 0: // approaching from the north
			dir.X, dir.Y = 1, 0
		case r == 'J' && dir.Y == 0: // approaching from the west
			dir.X, dir.Y = 0, -1
		case r == 'J' && dir.X == 0: // approaching from the north
			dir.X, dir.Y = -1, 0
		case r == '7' && dir.Y == 0: // approaching from the west
			dir.X, dir.Y = 0, 1
		case r == '7' && dir.X == 0: // approaching from the south
			dir.X, dir.Y = -1, 0
		case r == 'F' && dir.Y == 0: // approaching from the east
			dir.X, dir.Y = 0, 1
		case r == 'F' && dir.X == 0: // approaching from the south
			dir.X, dir.Y = 1, 0
		}
	}

	fmt.Println("p1=", len(path)/2)

	// Shoelace algorithm
	area := 0
	for i, v := range path {
		w := path[(i+1)%len(path)]
		area += v.X*w.Y - v.Y*w.X
	}
	area = util.AbsInt(area) / 2

	// Pick's theorem
	fmt.Println("p2=", area-len(path)/2+1)
}
