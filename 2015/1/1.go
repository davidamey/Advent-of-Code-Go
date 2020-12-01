package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	input := util.MustReadFile("input")

	p1, p2 := 0, -1
	for i, c := range input {
		if c == '(' {
			p1++
		} else {
			p1--
		}

		if p1 == -1 && p2 == -1 {
			p2 = i + 1
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}
