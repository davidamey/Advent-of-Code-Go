package main

import (
	"advent-of-code-go/2021/intgrid"
	"advent-of-code-go/util"
	"container/heap"
	"fmt"
	"math"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	cavern, w, h := intgrid.Parse(lines)
	fmt.Println("p1=", findPath(cavern, w, h, 0, len(cavern)-1))

	bw, bh := 5*w, 5*h
	bigCavern := make([]int, bw*bh)
	for i := range bigCavern {
		scaleX := (i % bw) / w
		scaleY := i / bw / h
		y := ((i / bw) % h) * w
		bigCavern[i] = (cavern[y+i%w]+scaleX+scaleY-1)%9 + 1
	}

	fmt.Println("p2=", findPath(bigCavern, bw, bh, 0, len(bigCavern)-1))
}

// A* algorithm
func findPath(cavern []int, w, h, start, end int) int {
	open := priorityQueue{&node{start, 0}}
	heap.Init(&open)

	from := make([]int, len(cavern))

	gScores := make([]int, len(cavern))
	fScores := make([]int, len(cavern))
	for i := range cavern {
		gScores[i] = math.MaxInt
		fScores[i] = math.MaxInt
	}

	for open.Len() > 0 {
		x := heap.Pop(&open).(*node).i

		if x == end {
			return calcRisk(cavern, from, start, end)
		}

		for _, n := range intgrid.Adjacent(x, w, h, false) {
			g := gScores[x] + cavern[n]
			if g >= gScores[n] {
				continue
			}

			from[n] = x
			gScores[n] = g
			fScores[n] = g + intgrid.Manhattan(n, end, w, h)

			if !open.contains(n) {
				heap.Push(&open, &node{n, fScores[n]})
			}
		}
	}

	panic("no path found")
}

func calcRisk(cavern []int, from []int, start, end int) (risk int) {
	risk += cavern[end]
	for x := from[end]; x != start; x = from[x] {
		risk += cavern[x]
	}
	return
}

type node struct {
	i, priority int
}

type priorityQueue []*node

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	node := x.(*node)
	*pq = append(*pq, node)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return node
}

func (pq priorityQueue) contains(i int) bool {
	for _, n := range pq {
		if n.i == i {
			return true
		}
	}
	return false
}
