package main

import (
	"advent/util"
	"fmt"
)

func main() {
	// instructs := util.MustReadFileToInts("example")
	instructs := util.MustReadFileToInts("input")

	fmt.Println("p1=", p1(instructs))
	fmt.Println("p1=", p2(instructs))
}

func p1(initial []int) (steps int) {
	instructs := make([]int, len(initial))
	copy(instructs, initial)
	for i := 0; i >= 0 && i < len(instructs); {
		jump := instructs[i]
		instructs[i]++
		i += jump
		steps++
	}
	return
}

func p2(initial []int) (steps int) {
	instructs := make([]int, len(initial))
	copy(instructs, initial)
	for i := 0; i >= 0 && i < len(instructs); {
		jump := instructs[i]
		if jump >= 3 {
			instructs[i]--
		} else {
			instructs[i]++
		}
		i += jump
		steps++
	}
	return
}
