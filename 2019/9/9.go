package main

import (
	"advent-of-code-go/2019/intcode"
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	// prog := intcode.New(util.MustReadCSInts("example"))
	prog := intcode.Program(util.MustReadCSInts("input"))
	fmt.Println("p1=", prog.Run(1)[0])
	fmt.Println("p2=", prog.Run(2)[0])
}
