package main

import (
	"advent-of-code-go/2019/intcode"
	"advent-of-code-go/util"
	"advent-of-code-go/util/grid"
	"fmt"
	"time"
)

const (
	tEmpty = iota
	tWall
	tBlock
	tPaddle
	tBall
)

func main() {
	defer util.Duration(time.Now())

	prog := intcode.Program(util.MustReadCSInts("input"))

	fmt.Println("p1=", p1(prog))
	fmt.Println("p2=", p2(prog))
}

func p1(prog intcode.Program) (count int) {
	for i, out := 0, prog.Run(); i < len(out); i += 3 {
		if out[i+2] == tBlock {
			count++
		}
	}
	return
}

func p2(prog intcode.Program) (score int) {
	g := game{
		in:      make(chan int),
		out:     make(chan int),
		running: true,
		grid:    grid.New[rune](),
	}

	prog[0] = 2
	go func() {
		prog.RunBuf(g.in, g.out)
		g.running = false
	}()

	var e [3]int
	read := 0
	for g.running {
		select {
		case g.in <- g.getJoyDir():
		case v := <-g.out:
			e[read] = v
			read++
			if read == 3 {
				read = 0
				g.parseEntity(e)
			}
		}

		// // print:
		// g.grid.Print("%c", true)
		// fmt.Println("score=", g.score)
		// time.Sleep(10 * time.Millisecond)
	}

	return g.score
}

type game struct {
	in, out     chan int
	ballX, padX int
	score       int
	running     bool
	grid        *grid.Grid[rune]
}

func (g *game) getJoyDir() int {
	switch {
	case g.ballX > g.padX:
		return 1
	case g.ballX < g.padX:
		return -1
	default:
		return 0
	}
}

func (g *game) parseEntity(e [3]int) {
	if e[0] == -1 && e[1] == 0 {
		g.score = e[2]
	} else {
		switch e[2] {
		case tBall:
			g.ballX = e[0]
		case tPaddle:
			g.padX = e[0]
		}
		g.grid.SetAt(e[0], e[1], mapType(e[2]))
	}
}

func mapType(t int) rune {
	switch t {
	case tWall:
		return '█'
	case tBlock:
		return '▭'
	case tPaddle:
		return '_'
	case tBall:
		// return '◯'
		return '☃' // snowman!
	default:
		return ' '
	}
}
