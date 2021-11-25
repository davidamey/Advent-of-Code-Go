package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	// input = []int{3, 2, 4, 1, 5} // example
	// input = []int{3, 8, 9, 1, 2, 5, 4, 6, 7} // example
	input = []int{7, 1, 2, 6, 4, 3, 5, 8, 9} // actual
)

func main() {
	defer util.Duration(time.Now())

	fmt.Println("p1=", p1())
	fmt.Println("p2=", p2())
}

func p1() string {
	g := newGame(input)
	for i := 0; i < 100; i++ {
		g.move()
	}
	return g.label()
}

func p2() int {
	g := newGame(input)
	g.cur = g.cur.p
	for v := len(g.cups) + 1; v <= 1000000; v++ {
		g.addCup(v)
		g.cur = g.cur.n
	}
	g.cur = g.cur.n

	for i := 0; i < 10000000; i++ {
		g.move()
	}

	// fmt.Println(g.cups[1].n.v, g.cups[1].n.n.v)
	return g.cups[1].n.v * g.cups[1].n.n.v
}

type cup struct {
	v    int
	p, n *cup
}

type game struct {
	cups map[int]*cup
	cur  *cup
}

func newGame(vs []int) *game {
	g := &game{cups: make(map[int]*cup, len(vs))}
	for _, v := range vs {
		g.addCup(v)
		g.cur = g.cur.n
	}
	g.cur = g.cur.n
	return g
}

func (g *game) addCup(v int) {
	c := &cup{v: v}
	g.cups[v] = c
	if g.cur == nil {
		c.p = c
		c.n = c
		g.cur = c
		return
	}

	c.p = g.cur
	c.n = g.cur.n
	c.n.p = c
	g.cur.n = c

	if g.cur.p == g.cur {
		g.cur.p = c
	}
}

func (g *game) move() {
	p := []*cup{g.cur.n, g.cur.n.n, g.cur.n.n.n}
	g.cur.n = p[2].n
	g.cur.n.p = g.cur

	dest := g.cur.v - 1
	for {
		if dest == 0 {
			dest = len(g.cups)
			continue
		}

		if dest == p[0].v || dest == p[1].v || dest == p[2].v {
			dest--
			continue
		}

		break
	}

	dc := g.cups[dest]
	p[0].p = dc
	p[2].n = dc.n
	dc.n = p[0]
	p[2].n.p = p[2]

	g.cur = g.cur.n
}

func (g *game) label() string {
	var sb strings.Builder
	for i, c := 0, g.cups[1].n; i < len(g.cups)-1; i, c = i+1, c.n {
		sb.WriteString(strconv.Itoa(c.v))
	}
	return sb.String()
}
