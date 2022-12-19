package main

import (
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
)

const (
	buffer      = 10
	pathOpen    = '.'
	pathBlocked = '#'

	// example
	// designersNumber = 10
	// targetX         = 7
	// targetY         = 4

	// input
	designersNumber = 1352
	targetX         = 31
	targetY         = 39
)

func main() {
	g := grid.New()
	for y := 0; y <= targetY+buffer; y++ {
		for x := 0; x <= targetX+buffer; x++ {
			n := x*x + 3*x + 2*x*y + y + y*y
			n += designersNumber
			if countOneBits(n)%2 == 0 {
				g.SetAt(x, y, pathOpen)
			} else {
				g.SetAt(x, y, pathBlocked)
			}
		}
	}

	countBelow50 := 0
	path := g.ShortestPath(
		vector.New(1, 1),
		vector.New(targetX, targetY),
		func(v, _ interface{}, depth int) bool {
			valid := v.(rune) == pathOpen
			if valid && depth+1 <= 50 {
				countBelow50++
			}
			return valid
		},
	)

	fmt.Println("p1=", path.Length)
	fmt.Println("p2=", countBelow50)
}

func countOneBits(i int) (count int) {
	for i > 0 {
		if i&1 == 1 {
			count++
		}
		i >>= 1
	}
	return
}
