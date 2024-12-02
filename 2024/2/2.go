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
		levels := util.ParseInts(l, " ")
		if checkReport(levels) {
			p1++
		}
		if checkWithDampener(levels) {
			p2++
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func checkReport(levels []int) bool {
	ld := 0
	for i := range levels[1:] {
		d := levels[i+1] - levels[i]

		if ad := util.AbsInt(d); ad < 1 || ad > 3 {
			return false
		}

		if (ld < 0 && d > 0) || (ld > 0 && d < 0) {
			return false
		}

		ld = d
	}

	return true
}

func checkWithDampener(levels []int) bool {
	if checkReport(levels) {
		return true
	}

	dampened := make([]int, len(levels)-1)
	for i := range levels {
		k := 0
		for j := range levels {
			if j == i {
				continue
			}
			dampened[k] = levels[j]
			k++
		}

		if checkReport(dampened) {
			return true
		}
	}

	return false
}
