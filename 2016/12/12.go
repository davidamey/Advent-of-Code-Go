package main

import (
	"advent/2016/12/assembunny"
	"advent-of-code-go/util"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")
	p := assembunny.Compile(lines)

	var p1 assembunny.Registers
	p.Run(&p1, false)
	fmt.Println("p1=", p1[0])

	var p2 assembunny.Registers
	p2[2] = 1 // 'c' to 1
	p.Run(&p2, false)
	fmt.Println("p2=", p2[0])
}
