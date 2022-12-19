package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	p1, p2 := 0, 0
	for _, l := range lines {
		var a, b, x, y int
		fmt.Sscanf(l, "%d-%d,%d-%d", &a, &b, &x, &y)
		if rangesFullyOverlap(a, b, x, y) {
			p1++
		}
		if rangesOverlap(a, b, x, y) {
			p2++
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func rangesFullyOverlap(a, b, x, y int) bool {
	return (a <= x && b >= y) || (x <= a && y >= b)
}

func rangesOverlap(a, b, x, y int) bool {
	return a <= y && x <= b
}
