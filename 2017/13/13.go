package main

import (
	"advent/util"
	"fmt"
)

func main() {
	// input := util.MustReadFileToLines("example")
	input := util.MustReadFileToLines("input")

	fmt.Println("p1=", p1(input))
	fmt.Println("p2=", p2(input))
}

func p1(input []string) int {
	tripSeverity := 0
	for _, l := range input {
		var d, r int
		fmt.Sscanf(l, "%d: %d", &d, &r)
		if caught, severity := severity(0, d, r); caught {
			tripSeverity += severity
		}
	}
	return tripSeverity
}

func p2(input []string) int {
outer:
	for delay := 0; ; delay++ {
		for _, l := range input {
			var d, r int
			fmt.Sscanf(l, "%d: %d", &d, &r)
			if caught, _ := severity(delay, d, r); caught {
				continue outer
			}
		}

		return delay
	}
}

func severity(delay, d, r int) (caught bool, severity int) {
	// r=1 is a special case, although this shouldn't be in the data
	// as p2 wouldn't be solvable
	if r == 1 {
		return true, d
	}

	period := 2 * (r - 1)
	if (d+delay)%period == 0 {
		return true, d * r
	}

	return false, 0
}
