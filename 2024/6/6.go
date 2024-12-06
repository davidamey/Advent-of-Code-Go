package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

type vec [2]int

func main() {
	defer util.Duration(time.Now())

	// input := util.MustReadFileToLines("example")
	input := util.MustReadFileToLines("input")

	w := len(input[0])
	h := len(input)

	grid := make([][]rune, h)
	for y, l := range input {
		grid[y] = []rune(l)
	}
	start := findGuard(grid)

	// p1
	visited := map[vec]bool{}
	guard := start
	dir := vec{0, -1}
	for {
		visited[guard] = true
		next := move(guard, dir)
		if !inBounds(next, w, h) {
			break
		}

		for grid[next[1]][next[0]] == '#' {
			dir.turn()
			next = move(guard, dir)
		}
		guard = next
	}

	// p2
	p2 := 0
	for v := range visited {
		if v == start {
			continue
		}

		grid[v[1]][v[0]] = '#'
		if detectLoop(grid, w, h, start) {
			p2++
		}
		grid[v[1]][v[0]] = '.'
	}

	fmt.Println("p1=", len(visited))
	fmt.Println("p2=", p2)
}

func findGuard(grid [][]rune) vec {
	for y, l := range grid {
		for x, r := range l {
			if r == '^' {
				grid[y][x] = '.'
				return vec{x, y}
			}
		}
	}
	panic("no guard found")
}

func inBounds(pos vec, w, h int) bool {
	if pos[0] < 0 || pos[0] >= w {
		return false
	}
	if pos[1] < 0 || pos[1] >= h {
		return false
	}
	return true
}

func detectLoop(grid [][]rune, w, h int, start vec) bool {
	visited := map[vec]vec{}
	guard := start
	dir := vec{0, -1}
	for {
		if visited[guard] == dir {
			return true
		}

		visited[guard] = dir
		next := move(guard, dir)
		if !inBounds(next, w, h) {
			break
		}

		for grid[next[1]][next[0]] == '#' {
			dir.turn()
			next = move(guard, dir)
		}
		guard = next
	}
	return false
}

func move(pos, dir vec) vec {
	return vec{pos[0] + dir[0], pos[1] + dir[1]}
}

func (dir *vec) turn() {
	dir[0], dir[1] = -dir[1], dir[0]
}
