package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	// data := util.MustReadFileToInts("example")
	data := util.MustReadFileToInts("input")

	p1, p2 := 0, 0
	for i := range data {
		if i+1 < len(data) && data[i] < data[i+1] {
			p1++
		}

		if i+3 < len(data) && data[i] < data[i+3] {
			p2++
		}
	}
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}
