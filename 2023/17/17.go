package main

import (
	"advent-of-code-go/util"
	"container/heap"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	w, h := len(lines[0]), len(lines)
	grid := make([][]int, h)
	for y, l := range lines {
		grid[y] = make([]int, w)
		for x, r := range l {
			grid[y][x] = int(r - '0')
		}
	}

	fmt.Println("p1=", search(grid, func(_, curr int) bool { return curr < 4 }))
	fmt.Println("p2=", search(grid, func(prev, curr int) bool {
		return (curr > prev || prev >= 4) && curr < 11
	}))
}

type (
	vec  [2]int
	node [6]int // 0:heatloss, 1:x, 2:y, 3:dir, 4:steps, 5:heuristic
)

var dirs []vec = []vec{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} // U R D L

func search(city [][]int, condition func(prev, curr int) bool) int {
	w, h := len(city[0]), len(city)
	start, end := vec{0, 0}, vec{w - 1, h - 1}

	open := priorityQueue{
		node{0, start[0], start[1], -1, 0},
	}
	heap.Init(&open)
	visited := make(map[[4]int]int)

	for open.Len() > 0 {
		n := heap.Pop(&open).(node)

		if n[1] == end[0] && n[2] == end[1] {
			return n[0]
		}

		for di, d := range dirs {
			// Can't reverse
			if (di+2)%4 == n[3] {
				continue
			}

			// If same dir, inc steps, otherwise 1
			steps := 1
			if di == n[3] {
				steps = n[4] + 1
			}

			// Step condition
			if !condition(n[4], steps) {
				continue
			}

			// Move
			m := node{n[0], n[1] + d[0], n[2] + d[1], di, steps}

			// Out of bounds
			if m[1] < 0 || m[1] >= w || m[2] < 0 || m[2] >= h {
				continue
			}

			// Update heatloss and heuristic
			m[0] += city[m[2]][m[1]]
			m[5] += m[0] + (end[0] - m[1]) + (end[1] - m[2])

			// If we've seen with a better heatloss, ignore
			s := [4]int{m[1], m[2], m[3], m[4]} // x, y, dir, steps
			if hl, seen := visited[s]; seen && hl <= m[0] {
				continue
			}
			visited[s] = m[0]

			heap.Push(&open, m)
		}
	}

	panic("no path found")
}

type priorityQueue []node

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i][5] < pq[j][5]
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x any) {
	*pq = append(*pq, x.(node))
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
