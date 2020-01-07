package main

import "advent-of-code-go/2019/intcode"

import "advent-of-code-go/util"

import "fmt"

func main() {
	prog := intcode.Program(util.MustReadCSInts("input"))

	p1(prog)
	p2(prog)
}

func p1(prog intcode.Program) {
	out := prog.RunScript([]string{
		"NOT C T",
		"NOT A J",
		"OR T J",
		"AND D J",
		"WALK",
	})
	if v := out[len(out)-1]; v > 127 {
		fmt.Println("p1=", v)
	} else {
		print(out)
	}
}

func p2(prog intcode.Program) {
	out := prog.RunScript([]string{
		"NOT A J",
		"NOT C T",
		"AND H T",
		"OR T J",
		"NOT B T",
		"AND A T",
		"AND C T",
		"OR T J",
		"AND D J",
		"RUN",
	})
	if v := out[len(out)-1]; v > 127 {
		fmt.Println("p2=", v)
	} else {
		print(out)
	}
}

func print(ascii []int) {
	for _, i := range ascii {
		fmt.Print(string(i))
	}
	fmt.Println()
}
