package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	discs := make([]disc, len(lines))
	for i, l := range lines {
		var n, pos, start int
		fmt.Sscanf(l, "Disc #%d has %d positions; at time=0, it is at position %d.", &n, &pos, &start)
		discs[i] = disc{pos, start}
	}

	fmt.Println("p1=", solve(discs))

	discs = append(discs, disc{11, 0})
	fmt.Println("p2=", solve(discs))
}

type disc struct {
	positions int
	start     int
}

func solve(discs []disc) (t int) {
tick:
	for ; ; t++ {
		for i, d := range discs {
			if (d.start+i+t+1)%d.positions > 0 {
				continue tick
			}
		}
		break
	}
	return
}
