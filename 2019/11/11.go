package main

import (
	"advent-of-code-go/2019/intcode"
	"advent-of-code-go/util"
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
)

func main() {
	p := intcode.Program(util.MustReadCSInts("input"))
	p1(p)
	p2(p)
}

func p1(p intcode.Program) {
	pIn, pOut := make(chan int), make(chan int, 2)
	running := true
	go func() {
		p.RunBuf("prog", pIn, pOut)
		close(pIn)
		running = false
	}()

	g := grid.New()
	rP := vector.New(0, 0)
	rV := vector.New(0, -1)

	painted := make(map[vector.Vec]struct{})
	for running {
		painted[rP] = struct{}{}
		act(g, &rP, &rV, pIn, pOut)
	}
	fmt.Println("p1=", len(painted))
}

func p2(p intcode.Program) {
	pIn, pOut := make(chan int), make(chan int, 2)
	running := true
	go func() {
		p.RunBuf("prog", pIn, pOut)
		close(pIn)
		running = false
	}()

	g := grid.New()
	rP := vector.New(0, 0)
	rV := vector.New(0, -1)
	g.Set(rP, 1)

	for running {
		act(g, &rP, &rV, pIn, pOut)
	}

	fmt.Println("p2=")
	for y := g.Min.Y; y <= g.Max.Y; y++ {
		for x := g.Min.X; x <= g.Max.X; x++ {
			v := 0
			if e := g.EntryAt(x, y); e != nil {
				v = e.(int)
			}
			if v == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func act(g *grid.Grid, rP, rV *vector.Vec, pIn, pOut chan int) {
	current := 0
	if e := g.Entry(*rP); e != nil {
		current = e.(int)
	}

	pIn <- current
	colour, turn := <-pOut, <-pOut
	g.Set(*rP, colour)
	if turn == 0 { // left
		rV.X, rV.Y = rV.Y, -rV.X
	} else { // right
		rV.X, rV.Y = -rV.Y, rV.X
	}
	rP.Add(*rV)
}
