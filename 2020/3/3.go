package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	p1 := traverse(lines, 3, 1)
	p2 := p1 *
		traverse(lines, 1, 1) *
		traverse(lines, 5, 1) *
		traverse(lines, 7, 1) *
		traverse(lines, 1, 2)

	fmt.Println("p1=", p1)
	fmt.Println("p1=", p2)
}

func traverse(lines []string, dx, dy int) (treeCount int) {
	x, y := 0, 0
	for y < len(lines) {
		if lines[y][x] == '#' {
			treeCount++
		}
		x = (x + dx) % len(lines[0])
		y += dy
	}
	return
}
