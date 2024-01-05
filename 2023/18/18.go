package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strconv"
	"time"
)

var dirs map[int]vec = map[int]vec{
	'U': {0, -1},
	'R': {1, 0},
	'D': {0, 1},
	'L': {-1, 0},
	0:   {1, 0},
	1:   {0, 1},
	2:   {-1, 0},
	3:   {0, -1},
}

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	stepsP1 := make([]step, len(lines))
	stepsP2 := make([]step, len(lines))
	for i, l := range lines {
		var (
			d   rune
			n   int
			hex string
		)
		fmt.Sscanf(l, "%c %d %s", &d, &n, &hex)

		stepsP1[i] = step{dirs[int(d)], n}
		stepsP2[i] = hexToStep(hex)
	}

	fmt.Println("p1=", dig(stepsP1))
	fmt.Println("p2=", dig(stepsP2))
}

type vec [2]int

type step struct {
	dir vec
	len int
}

func hexToStep(hex string) step {
	// hex of the format (#abc123)
	i, _ := strconv.ParseInt(hex[2:7], 16, 0)
	d, _ := strconv.ParseInt(hex[7:8], 16, 0)
	return step{dirs[int(d)], int(i)}
}

func dig(steps []step) int {
	p := vec{0, 0}
	area, edge := 0, 0
	for _, s := range steps {
		area += s.len * (p[0]*s.dir[1] - p[1]*s.dir[0])
		edge += s.len
		p[0] += s.dir[0] * s.len
		p[1] += s.dir[1] * s.len
	}

	if p != (vec{0, 0}) {
		panic("not closed")
	}

	return util.AbsInt(area)/2 + edge/2 + 1
}
