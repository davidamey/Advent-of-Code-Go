package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines, steps := util.MustReadFileToLines("example"), 10
	lines, steps := util.MustReadFileToLines("input"), 1000

	var moons []*moon
	for _, l := range lines {
		moons = append(moons, newMoon(l))
	}

	p1, p2 := solve(moons, steps)
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func solve(sys system, p1Steps int) (p1, p2 int) {
	step := 0
	var cycleLength [3]int
	pastStates := [3]map[string]struct{}{
		make(map[string]struct{}),
		make(map[string]struct{}),
		make(map[string]struct{}),
	}
	for step < p1Steps || cycleLength[0] == 0 || cycleLength[1] == 0 || cycleLength[2] == 0 {
		state := sys.tick()

		if step == p1Steps {
			for _, m := range sys {
				p1 += m.energy()
			}
		}

		for i := 0; i < 3; i++ {
			if cycleLength[i] == 0 {
				if _, exists := pastStates[i][state[i]]; exists {
					cycleLength[i] = step
				} else {
					pastStates[i][state[i]] = struct{}{}
				}
			}
		}

		step++
	}

	fmt.Println(cycleLength)
	p2 = util.LCM(cycleLength[:]...)
	return
}

type system []*moon

func (sys system) tick() (state [3]string) {
	// gravity
	for _, m1 := range sys {
		for _, m2 := range sys {
			m1.gravityFrom(m2)
		}
	}

	// movement
	for _, m := range sys {
		m.move()
		for i := 0; i < 3; i++ {
			state[i] += fmt.Sprintf("%2d%2d", m.p[i], m.v[i])
		}
	}

	return
}

type moon struct {
	p, v [3]int
}

func (m *moon) String() string {
	return fmt.Sprintf("pos=<x=%3d, y=%3d, z=%3d>, vel=<x=%3d, y=%3d, z=%3d>",
		m.p[0], m.p[1], m.p[2],
		m.v[0], m.v[1], m.v[2],
	)
}

func newMoon(raw string) (m *moon) {
	m = &moon{}
	fmt.Sscanf(raw, "<x=%d, y=%d, z=%d>", &m.p[0], &m.p[1], &m.p[2])
	return
}

func (m *moon) move() {
	for i := 0; i < 3; i++ {
		m.p[i] += m.v[i]
	}
}

func (m1 *moon) gravityFrom(m2 *moon) {
	for i := 0; i < 3; i++ {
		switch {
		case m1.p[i] < m2.p[i]:
			m1.v[i]++
		case m1.p[i] > m2.p[i]:
			m1.v[i]--
		}
	}
}

func (m *moon) energy() int {
	pe := 0
	ke := 0
	for i := 0; i < 3; i++ {
		pe += util.AbsInt(m.p[i])
		ke += util.AbsInt(m.v[i])
	}
	return pe * ke
}
