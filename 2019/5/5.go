package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	prog := util.MustReadCSInts("input")
	p1(prog)
	fmt.Println()
	p2(prog)
}

func p1(prog []int) int {
	fmt.Println("==p1==")
	return run(append([]int(nil), prog...), 1)
}

func p2(prog []int) int {
	fmt.Println("==p2==")
	return run(append([]int(nil), prog...), 5)
}

func run(prog []int, input int) int {
	for i := 0; prog[i] != 99; {
		oc, modes := parseInstruct(prog[i])

		switch oc {
		case 1:
			prog[prog[i+3]] = getVal(prog, i+1, modes[0]) + getVal(prog, i+2, modes[1])
			i += 4
		case 2:
			prog[prog[i+3]] = getVal(prog, i+1, modes[0]) * getVal(prog, i+2, modes[1])
			i += 4
		case 3:
			prog[prog[i+1]] = input
			i += 2
		case 4:
			fmt.Println("out:", getVal(prog, i+1, modes[0]))
			i += 2
		case 5:
			if getVal(prog, i+1, modes[0]) != 0 {
				i = getVal(prog, i+2, modes[1])
			} else {
				i += 3
			}
		case 6:
			if getVal(prog, i+1, modes[0]) == 0 {
				i = getVal(prog, i+2, modes[1])
			} else {
				i += 3
			}
		case 7:
			if getVal(prog, i+1, modes[0]) < getVal(prog, i+2, modes[1]) {
				prog[prog[i+3]] = 1
			} else {
				prog[prog[i+3]] = 0
			}
			i += 4
		case 8:
			if getVal(prog, i+1, modes[0]) == getVal(prog, i+2, modes[1]) {
				prog[prog[i+3]] = 1
			} else {
				prog[prog[i+3]] = 0
			}
			i += 4
		default:
			panic("unknown opcode")
		}
	}
	return prog[0]
}

func parseInstruct(inst int) (oc int, modes []int) {
	oc = inst % 100 // last 2 digits
	inst /= 100
	for i := 0; i < 4; i++ {
		modes = append(modes, inst%10)
		inst /= 10
	}
	return
}

func getVal(prog []int, ref, mode int) int {
	if mode == 1 {
		return prog[ref]
	}
	return prog[prog[ref]]
}
