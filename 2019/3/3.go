package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/vector"
	"fmt"
	"strconv"
	"strings"
)

type path map[vector.Vec]int

func main() {
	// lines := util.MustReadFileToLines("example")
	// lines := util.MustReadFileToLines("example2")
	// lines := util.MustReadFileToLines("example3")
	// lines := util.MustReadFileToLines("example4")
	lines := util.MustReadFileToLines("input")

	p1, p2 := makePath(lines[0]), makePath(lines[1])

	sol1, sol2 := solve(p1, p2)
	fmt.Println("p1=", sol1)
	fmt.Println("p2=", sol2)
}

func solve(p1, p2 path) (sol1, sol2 int) {
	var vs []vector.Vec
	for v := range p1 {
		if _, exists := p2[v]; exists {
			vs = append(vs, v)
		}
	}

	sol1, sol2 = 1<<16, 1<<16
	for _, v := range vs {
		// part 1
		if d := v.Manhattan(); d < sol1 {
			sol1 = d
		}

		// part 2
		d1, d2 := p1[v], p2[v]
		if d := d1 + d2; d < sol2 {
			sol2 = d
		}
	}
	return
}

func makePath(wire string) (p path) {
	dirs := map[byte]vector.Vec{
		'U': vector.New(0, -1),
		'R': vector.New(1, 0),
		'D': vector.New(0, 1),
		'L': vector.New(-1, 0),
	}

	p = make(path)
	v := vector.New(0, 0)
	l := 1
	for _, w := range strings.Split(wire, ",") {
		d := w[0]
		x, _ := strconv.Atoi(w[1:])
		for i := 0; i < x; i++ {
			v.Add(dirs[d])

			if _, exists := p[v]; !exists {
				p[v] = l
			}
			l++
		}
	}
	return
}
