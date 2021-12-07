package main

import (
	"advent-of-code-go/util"
	"fmt"
	"math"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// crabs := util.MustReadCSInts("example")
	crabs := util.MustReadCSInts("input")

	p1, p2 := math.MaxInt, math.MaxInt
	for i := 0; i <= util.MaxInt(crabs...); i++ {
		f1, f2 := 0, 0
		for _, c := range crabs {
			n := util.AbsInt(c - i)
			f1 += n
			f2 += n * (n + 1) / 2 // triangular
		}
		if f1 < p1 {
			p1 = f1
		}
		if f2 < p2 {
			p2 = f2
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}
