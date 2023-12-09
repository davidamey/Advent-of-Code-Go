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
	for _, line := range lines {
		l, r := nextNumber(util.ParseInts(line, " "))
		p1 += r
		p2 += l
	}
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func nextNumber(xs []int) (l, r int) {
	var diffs []int
	allZero := true
	for i := 1; i < len(xs); i++ {
		d := xs[i] - xs[i-1]
		diffs = append(diffs, d)
		if d != 0 {
			allZero = false
		}
	}
	if allZero {
		return xs[0], xs[len(xs)-1]
	} else {
		dl, dr := nextNumber(diffs)
		return xs[0] - dl, xs[len(xs)-1] + dr
	}
}
