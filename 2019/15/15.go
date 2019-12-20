package main

import (
	"advent-of-code-go/2019/intcode"
	"advent-of-code-go/util"
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
	"time"
)

func main() {
	prog := intcode.Program(util.MustReadCSInts("input"))

	in, out := make(chan int), make(chan int)
	go prog.RunBuf(in, out)

	d := &droid{
		grid: grid.New(),
		in:   in,
		out:  out,
		path: []node{node{}},
	}
	p1 := d.findOxy()
	fmt.Println("p1=", p1)
	fmt.Println("p2=", part2(d.oxy, d.grid))

}

func part2(oxy vector.Vec, g *grid.Grid) (minutes int) {
	minutes = -1
	queue := []vector.Vec{oxy}
	for len(queue) > 0 {
		var nextQueue []vector.Vec
		for _, v := range queue {
			g.Set(v, 2) // oxygenated
			for _, a := range v.Adjacent(false) {
				if g.Int(a) == 1 {
					nextQueue = append(nextQueue, a)
				}
			}
		}
		queue = nextQueue
		minutes++
		// print(g)
	}
	return
}

func print(g *grid.Grid) {
	fmt.Printf("\033[0;0H")
	fmt.Printf("\033[2J")

	for y := g.Min.Y; y <= g.Max.Y; y++ {
		for x := g.Min.X; x <= g.Max.X; x++ {
			e := g.EntryAt(x, y)
			r := '#'
			s := "%c"
			if e != nil {
				switch e.(int) {
				case 0:
					r = '#'
				case 1:
					r = ' '
				case 2:
					r = 'o'
					s = "\033[1;36m%c\033[0m"
				}
			}
			fmt.Printf(s, r)
		}
		fmt.Println()
	}

	time.Sleep(60 * time.Millisecond)
}

type node struct {
	dir int
	pos vector.Vec
}

type droid struct {
	grid    *grid.Grid
	pos     vector.Vec
	start   vector.Vec
	oxy     vector.Vec
	path    []node
	in, out chan int
}

func (d *droid) move(dir int) int {
	d.in <- dir
	return <-d.out
}

func (d *droid) moveBack() {
	i := len(d.path) - 1
	n := d.path[i]
	d.path = d.path[:i]
	d.in <- n.dir - 1 ^ 1 + 1
	<-d.out
	d.pos = d.path[i-1].pos
}

func (d *droid) findOxy() (min int) {
	min = 1 << 16
	for {
		moved := false
		for i, v := range []vector.Vec{d.pos.Up(), d.pos.Down(), d.pos.Left(), d.pos.Right()} {
			dir := i + 1

			if d.grid.Entry(v) != nil {
				// already visited
				continue
			}

			r := d.move(dir)
			d.grid.Set(v, r)
			if r == 0 {
				continue
			}

			if r == 2 {
				d.oxy = v
				min = util.MinInt(min, len(d.path))
			}

			d.pos = v
			d.path = append(d.path, node{dir, d.pos})
			moved = true
			break
		}

		if moved {
			continue
		} else if d.pos.IsOrigin() {
			break
		}

		d.moveBack()
	}

	return
}
