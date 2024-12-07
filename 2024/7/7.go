package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	equations := make([]equation, len(lines))
	for i, l := range lines {
		equations[i] = newEquation(l)
	}

	p1, p2 := 0, 0
	for _, e := range equations {
		if e.isSolvable(false) {
			p1 += e.target
			p2 += e.target
		} else if e.isSolvable(true) {
			p2 += e.target
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

type equation struct {
	target int
	values []int
}

func newEquation(s string) equation {
	parts := strings.Split(s, ": ")
	return equation{
		target: util.Atoi(parts[0]),
		values: util.ParseInts(parts[1], " "),
	}
}

func (e *equation) isSolvable(allowConcat bool) bool {
	return e.solvable(e.values[0], 1, allowConcat)
}

func (e *equation) solvable(total, i int, allowConcat bool) bool {
	if i == len(e.values) {
		return total == e.target
	}

	if e.solvable(add(total, e.values[i]), i+1, allowConcat) {
		return true
	}

	if e.solvable(multiply(total, e.values[i]), i+1, allowConcat) {
		return true
	}

	return allowConcat && e.solvable(concat(total, e.values[i]), i+1, allowConcat)
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func concat(a, b int) int {
	for x := b; x > 0; x /= 10 {
		a *= 10
	}
	return a + b
}
