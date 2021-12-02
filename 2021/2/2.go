package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	// data := util.MustReadFileToLines("example")
	data := util.MustReadFileToLines("input")

	h, d1, d2 := 0, 0, 0 // d1 also equivalent to 'aim'
	for _, d := range data {
		dir, v := "", 0
		fmt.Sscanf(d, "%s %d", &dir, &v)

		switch dir {
		case "forward":
			h += v
			d2 += d1 * v
		case "down":
			d1 += v
		case "up":
			d1 -= v
		}
	}

	fmt.Println("p1=", h*d1)
	fmt.Println("p2=", h*d2)
}
