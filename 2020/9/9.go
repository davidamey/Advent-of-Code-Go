package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// ints, preamble := util.MustReadFileToInts("example"), 5
	ints, preamble := util.MustReadFileToInts("input"), 25

	p1 := 0
	for i := preamble; i < len(ints); i++ {
		if !isSum(ints[i], ints[i-preamble:i]) {
			p1 = ints[i]
			break
		}
	}

	p2 := findSummingSet(p1, ints)

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func isSum(target int, pool []int) bool {
	for i, x := range pool {
		for _, y := range pool[i+1:] {
			if x+y == target {
				return true
			}
		}
	}
	return false
}

func findSummingSet(target int, pool []int) int {
	for i, x := range pool {
		min, max := x, x
		sum := x
		for _, y := range pool[i+1:] {
			switch {
			case y < min:
				min = y
			case y > max:
				max = y
			}

			sum += y
			if sum == target {
				return min + max
			} else if sum > target {
				break
			}
		}
	}

	panic("no summing set found)")
}
