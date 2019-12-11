package main

import (
	"advent-of-code-go/2019/intcode"
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	prog := intcode.Program(util.MustReadCSInts("input"))
	fmt.Println("p1=", p1(prog))
	fmt.Println("p2=", p2(prog))
}

func p1(prog intcode.Program) int {
	out := prog.Run(1)
	return out[len(out)-1]
}

func p2(prog intcode.Program) int {
	return prog.Run(5)[0]
}
