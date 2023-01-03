package main

import (
	"advent-of-code-go/2019/intcode"
	"advent-of-code-go/util"
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
	"strings"
)

func main() {
	prog := intcode.Program(util.MustReadCSInts("input"))
	prog[0] = 2

	// 100% by hand ^_^
	var in strings.Builder
	in.WriteString("A,B,A,B,A,C,B,C,A,C\n") // main
	in.WriteString("L,10,L,12,R,6\n")       // A
	in.WriteString("R,10,L,4,L,4,L,12\n")   // B
	in.WriteString("L,10,R,10,R,6,L,4\n")   // C
	in.WriteString("n\n")                   // video feed y/n

	bytes := []byte(in.String())
	ints := make([]int, len(bytes))
	for i, b := range bytes {
		ints[i] = int(b)
	}

	out := prog.Run(ints...)
	g := grid.New[rune]()
	p := vector.New(0, 0)
	for _, r := range out[:len(out)-1] {
		if r == 10 {
			p.Y++
			p.X = 0
			continue
		}
		g.Set(p, rune(r))
		p.X++
	}
	g.Print("%c", true)
	fmt.Println(out[len(out)-1])
}

func p1() {
	prog := intcode.Program(util.MustReadCSInts("input"))

	out := prog.Run()

	g := grid.New[rune]()
	p := vector.New(0, 0)
	for _, r := range out {
		if r == 10 {
			p.Y++
			p.X = 0
			continue
		}
		g.Set(p, rune(r))
		p.X++
	}

	g.PrintRunes()

	sum := 0
	g.ForEach(func(v vector.Vec, r rune) {
		if r != '#' {
			return
		}

		for _, w := range v.Adjacent(false) {
			if !g.InBounds(w) || g.Get(w) != '#' {
				return
			}
		}

		sum += v.X * v.Y
	})

	fmt.Println(sum)
}
