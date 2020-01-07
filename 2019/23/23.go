package main

import (
	"advent-of-code-go/2019/intcode"
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	prog := intcode.Program(util.MustReadCSInts("input"))

	queue := make(chan packet, 3000)
	var computers [50]*computer
	for i := range computers {
		computers[i] = &computer{make(chan int), make(chan int, 30), 0}
		go prog.RunBuf(computers[i].in, computers[i].out)
		computers[i].in <- i
	}

	var nat, lastNat packet
	p1 := true
	for {
		for _, c := range computers {
			select {
			case p := <-queue:
				if p.dest == 255 {
					nat = packet{0, p.x, p.y}
					if p1 {
						fmt.Println("p1=", p.y)
						p1 = false
					}
				} else {
					computers[p.dest].idle = 0
					computers[p.dest].in <- p.x
					computers[p.dest].in <- p.y
				}
			case dest := <-c.out:
				x := <-c.out
				y := <-c.out
				queue <- packet{dest, x, y}
			default:
				c.in <- -1
				c.idle++
			}
		}

		idle := true
		for _, c := range computers {
			// would like this more scientific, but 10 seems
			// to consistently get the right answer
			if c.idle < 10 {
				idle = false
				break
			}
		}

		if !idle {
			continue
		}

		if len(queue) == 0 {
			if nat == lastNat {
				fmt.Println("p2=", nat.y)
				return
			}
			computers[0].in <- nat.x
			computers[0].in <- nat.y
			lastNat = nat
		}
	}
}

type computer struct {
	in, out chan int
	idle    int
}

type packet struct {
	dest, x, y int
}
