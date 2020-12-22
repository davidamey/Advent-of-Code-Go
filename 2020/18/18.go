package main

import (
	"advent-of-code-go/util"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	opAdd = 0
	opMul = 1
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	p1, p2 := 0, 0
	for _, l := range lines {
		p1 += eval(l, false)
		p2 += eval(l, true)
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func eval(eq string, addFirst bool) int {
	eq = strings.ReplaceAll(eq, "(", "( ")
	eq = strings.ReplaceAll(eq, ")", " )")
	s := bufio.NewScanner(strings.NewReader(eq))
	s.Split(bufio.ScanWords)
	return evalPart(s, addFirst)
}

func evalPart(s *bufio.Scanner, addFirst bool) (result int) {
	var ops, vals []int

	exec := func() {
		x := pop(&vals)
		if pop(&ops) == opAdd {
			vals[len(vals)-1] += x
		} else {
			vals[len(vals)-1] *= x
		}
	}

	defer func() {
		for len(ops) > 0 {
			exec()
		}
		result = vals[len(vals)-1]
	}()

	for s.Scan() {
		switch t := s.Text(); t {
		case "+":
			if len(ops) > 0 {
				if !addFirst || ops[len(ops)-1] == opAdd {
					exec()
				}
			}
			ops = append(ops, opAdd)
		case "*":
			if len(ops) > 0 {
				exec()
			}
			ops = append(ops, opMul)
		case "(":
			vals = append(vals, evalPart(s, addFirst))
		case ")":
			return
		default: // number
			x, _ := strconv.Atoi(t)
			vals = append(vals, x)
		}
	}

	return
}

func pop(xs *[]int) (x int) {
	x, (*xs) = (*xs)[len((*xs))-1], (*xs)[:len((*xs))-1]
	return
}
