package main

import (
	"advent-of-code-go/2019/intcode"
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	prog := intcode.Program(util.MustReadCSInts("input"))

	count := 0
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if prog.Run(x, y)[0] == 1 {
				count++
			}
		}
	}
	fmt.Println("p1=", count)

	x, y := 0, 100
	for {
		for prog.Run(x, y)[0] == 0 {
			x++
		}
		if prog.Run(x+99, y-99)[0] == 1 {
			break
		}
		y++
	}
	fmt.Println("p2=", 10000*x+y-99)
}
