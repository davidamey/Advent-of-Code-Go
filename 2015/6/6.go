package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	ops := make(map[string]op)
	ops["on"] = func(in int) int {
		return 1
	}
	ops["off"] = func(in int) int {
		return 0
	}
	ops["toggle"] = func(in int) int {
		if in == 1 {
			return 0
		}
		return 1
	}

	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	g := grid{}
	for _, l := range lines {
		var o string
		var x1, y1, x2, y2 int

		l = strings.Replace(l, "turn ", "", 1)
		fmt.Sscanf(l, "%s %d,%d through %d,%d", &o, &x1, &y1, &x2, &y2)
		// fmt.Println(o, x1, y1, x2, y2)
		g.runOp(ops[o], x1, y1, x2, y2)
	}
	// g.print()

	onCount := 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			onCount += g[y][x]
		}
	}

	fmt.Println("== part1 ==")
	fmt.Println("lights on =", onCount)
}

func part2() {
	ops := make(map[string]op)
	ops["on"] = func(in int) int {
		return in + 1
	}
	ops["off"] = func(in int) int {
		out := in - 1
		if out >= 0 {
			return out
		}
		return 0
	}
	ops["toggle"] = func(in int) int {
		return in + 2
	}

	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	g := grid{}
	for _, l := range lines {
		var o string
		var x1, y1, x2, y2 int

		l = strings.Replace(l, "turn ", "", 1)
		fmt.Sscanf(l, "%s %d,%d through %d,%d", &o, &x1, &y1, &x2, &y2)
		g.runOp(ops[o], x1, y1, x2, y2)
	}

	brightness := 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			brightness += g[y][x]
		}
	}

	fmt.Println("== part1 ==")
	fmt.Println("brightness =", brightness)
}

type grid [1000][1000]int
type op func(int) int

func (g *grid) runOp(op op, x1, y1, x2, y2 int) {
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			g[y][x] = op(g[y][x])
		}
	}
}

func (g *grid) print() {
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			fmt.Print(g[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}
