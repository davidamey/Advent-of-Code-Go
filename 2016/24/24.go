package main

import (
	"advent-of-code-go/util"
	"advent/util/grid"
	"advent/util/vector"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	poiVecs := make(map[int]vector.Vec)
	g := grid.New()
	for y, l := range lines {
		for x, c := range l {
			g.SetAt(x, y, c)
			if c >= '0' && c <= '9' {
				poiVecs[int(c-'0')] = vector.New(x, y)
			}
		}
	}

	// pre-compute dists
	dists := make([][]int, len(poiVecs))
	for p1 := range poiVecs {
		dists[p1] = make([]int, len(poiVecs))

		for p2 := range poiVecs {
			d := 0
			if p1 != p2 {
				path := g.ShortestPath(poiVecs[p1], poiVecs[p2], pv)
				d = path.Length
			}
			dists[p1][p2] = d
		}
	}

	// calc dist for each permutation
	pois := make([]int, len(poiVecs)-1)
	for i := range pois {
		pois[i] = i + 1
	}

	p1 := 1 << 32
	p2 := 1 << 32
	for route := range util.NewIntPermuter(pois).Permutations() {
		cost := dists[0][route[0]]
		for i := 0; i < len(route)-1; i++ {
			cost += dists[route[i]][route[i+1]]
		}
		if cost < p1 {
			p1 = cost
		}
		cost += dists[route[len(route)-1]][0]
		if cost < p2 {
			p2 = cost
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func pv(v interface{}, depth int) bool {
	return v.(rune) != '#'
}
