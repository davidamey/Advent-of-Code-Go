package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

type vec [2]int

func main() {
	defer util.Duration(time.Now())

	// grid := util.MustReadFileToLines("example")
	grid := util.MustReadFileToLines("input")

	size := len(grid)
	if size != len(grid[0]) {
		panic("Grid not square")
	}

	var start vec
findStart:
	for y, row := range grid {
		for x, r := range row {
			if r == 'S' {
				start = vec{x, y}
				grid[y] = row[:x] + "." + row[x+1:]
				break findStart
			}
		}
	}

	fmt.Println("p1=", walk(grid, start, 64))

	// P2
	// By inspection, the solution is quadratic: ax^2 + bx + c
	// 3 point rule will give us a, b and c
	//
	// Brilliant write-up: https://github.com/derailed-dash/Advent-of-Code/blob/master/src/AoC_2023/Dazbo's_Advent_of_Code_2023.ipynb

	steps := 26_501_365

	solns := []int{
		walk(grid, start, steps%size),
		walk(grid, start, steps%size+size),
		walk(grid, start, steps%size+size+size),
	}

	c := solns[0]
	b := (4*solns[1] - 3*c - solns[2]) / 2
	a := solns[1] - c - b
	x := (steps - size/2) / size

	fmt.Println("p2=", a*x*x+b*x+c)
}

func walk(grid []string, start vec, steps int) int {
	size := len(grid)

	queue := [][3]int{{start[0], start[1], steps}}
	reachable := map[vec]struct{}{}
	seen := map[vec]bool{}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		steps := p[2]

		if steps < 0 {
			continue
		}

		if steps%2 == 0 {
			reachable[vec{p[0], p[1]}] = struct{}{}
		}

		if steps > 0 {
			steps--

			for _, d := range []vec{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				x, y := p[0]+d[0], p[1]+d[1]

				gx, gy := util.Mod(x, size), util.Mod(y, size)

				if grid[gy][gx] != '.' {
					continue
				}

				if v := (vec{x, y}); !seen[v] {
					seen[v] = true
					queue = append(queue, [3]int{x, y, steps})
				}
			}

		}
	}

	return len(reachable)
}
