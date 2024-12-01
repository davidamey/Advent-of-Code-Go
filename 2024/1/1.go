package main

import (
	"advent-of-code-go/util"
	"fmt"
	"slices"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	xs := make([]int, len(lines))
	ys := make([]int, len(lines))
	for i, l := range lines {
		fmt.Sscanf(l, "%d   %d", &xs[i], &ys[i])
	}

	slices.Sort(xs)
	slices.Sort(ys)

	counts := map[int]int{}
	for _, y := range ys {
		counts[y]++
	}

	p1, p2 := 0, 0
	for i, x := range xs {
		p1 += util.AbsInt(ys[i] - x)
		p2 += x * counts[x]
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}
