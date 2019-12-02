package main

import (
	"fmt"
	"local/advent/util"
)

func main() {
	prog := util.MustReadCSInts("input")

	fmt.Println("p1=", p1(prog))
	fmt.Println("p2=", p2(prog))
}

func p1(prog []int) int {
	prog[1] = 12
	prog[2] = 2
	return run(append([]int(nil), prog...))
}

func p2(prog []int) int {
	p := make([]int, len(prog))
	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			copy(p, prog)
			p[1] = n
			p[2] = v
			if run(p) == 19690720 {
				return 100*n + v
			}
		}
	}
	return -1
}

func run(prog []int) int {
	for i := 0; i < len(prog); i += 4 {
		switch prog[i] {
		case 1:
			prog[prog[i+3]] = prog[prog[i+1]] + prog[prog[i+2]]
		case 2:
			prog[prog[i+3]] = prog[prog[i+1]] * prog[prog[i+2]]
		case 99:
			return prog[0]
		default:
			return -1
		}
	}
	return -1
}
