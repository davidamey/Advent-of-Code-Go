package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
)

func main() {
	// examples
	// for _, eg := range []string{
	// 	"ne,ne,ne",
	// 	"ne,ne,sw,sw",
	// 	"ne,ne,s,s",
	// 	"se,sw,se,sw,sw",
	// } {
	// 	fmt.Println(eg, p1(eg))
	// }

	input := string(util.MustReadFile("input"))

	p1, p2 := solve(input)
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func solve(input string) (finish, max int) {
	moves := strings.Split(input, ",")
	o := point{0, 0}
	p := point{0, 0}
	for _, m := range moves {
		p.move(m)
		if d := p.dist(o); d > max {
			max = d
		}
	}
	finish = p.dist(o)
	return
}

type point struct{ x, y int }

func (p *point) move(dir string) {
	switch dir {
	case "n":
		p.y++
	case "ne":
		p.x++
	case "se":
		p.x++
		p.y--
	case "s":
		p.y--
	case "sw":
		p.x--
	case "nw":
		p.x--
		p.y++
	default:
		panic(fmt.Errorf("unknown direction %s", dir))
	}
}

func (p *point) dist(q point) int {
	dx := q.x - p.x
	dy := q.y - p.y

	if (dx > 0) == (dy > 0) {
		return util.AbsInt(dx + dy)
	}
	return util.MaxInt(util.AbsInt(dx), util.AbsInt(dy))
}

//      \ n  /
//    nw +--+ ne
//      /    \
//    -+      +-
//      \    /
//    sw +--+ se
//      / s  \
