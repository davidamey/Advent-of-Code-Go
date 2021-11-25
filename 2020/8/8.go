package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strconv"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	p1, _ := run(lines, -1)

	var p2 int
	var term bool
	for i := range lines {
		if lines[i][0:3] == "jmp" || lines[i][0:3] == "nop" {
			p2, term = run(lines, i)
			if term {
				break
			}
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func run(instructions []string, switchInstruct int) (accumulator int, terminated bool) {
	seen := make(map[int]bool)
	i := 0
	for !seen[i] && i < len(instructions) {
		seen[i] = true
		op := instructions[i][0:3]
		arg, _ := strconv.Atoi(instructions[i][4:])
		// fmt.Printf("%d: %s %2d | %d\n", i, op, arg, accumulator)
		switch op {
		case "nop":
			if i == switchInstruct {
				i += arg
			} else {
				i++
			}
		case "acc":
			accumulator += arg
			i++
		case "jmp":
			if i == switchInstruct {
				i++
			} else {
				i += arg
			}
		}
	}
	return accumulator, i >= len(instructions)
}
