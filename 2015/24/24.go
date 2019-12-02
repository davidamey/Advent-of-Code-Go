package main

import (
	"advent-of-code-go/util"
	"fmt"
	"math"
)

func main() {
	// file, _ := util.OpenExample()
	file, _ := util.OpenInput()
	defer file.Close()
	weights, _ := util.ReadLinesToInts(file)

	fmt.Println("p1=", solve(weights, 3))
	fmt.Println("p2=", solve(weights, 4))
}

func solve(weights []int, parts int) int {
	groupSum := sum(weights) / parts

	minQE := math.MaxInt64
	for i := range weights {
		for c := range util.Combinations(weights, i) {
			if sum(c) == groupSum {
				if qec := qe(c); qec < minQE {
					minQE = qec
				}
			}
		}
		if minQE < math.MaxInt64 {
			return minQE
		}
	}

	panic("no solution")
}

func dump(g1, g2, g3 []int) {
	fmt.Println(g1, qe(g1), g2, g3)
}

func sum(ints []int) (sum int) {
	for _, i := range ints {
		sum += i
	}
	return
}

func qe(ints []int) (qe int) {
	qe = 1
	for _, i := range ints {
		qe *= i
	}
	return
}
