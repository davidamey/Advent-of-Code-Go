package main

import (
	"advent-of-code-go/util"
	"fmt"
	"math"
)

func main() {
	// input := util.MustReadFileToLines("example")
	input := util.MustReadFileToLines("input")

	fmt.Println("p1=", p1(input))
	fmt.Println("p2=", p2(input))
}

func p1(lines []string) int {
	particles := make([]*particle, len(lines))
	for i, l := range lines {
		particles[i] = newParticle(l)
	}

	for t := 0; t < 1000; t++ {
		for _, p := range particles {
			p.tick()
		}
	}

	minDist := math.MaxInt32
	minI := -1
	for i, p := range particles {
		if d := p.dist(); d < minDist {
			minDist = d
			minI = i
		}
	}
	return minI
}

func p2(lines []string) int {
	particles := make([]*particle, len(lines))
	for i, l := range lines {
		particles[i] = newParticle(l)
	}

	for t := 0; t < 1000; t++ {
		positions := make(map[v3]int)
		for i, p := range particles {
			if p.collided {
				continue
			}

			p.tick()
			if j, ok := positions[p.p]; ok {
				particles[j].collided = true
				p.collided = true
			} else {
				positions[p.p] = i
			}
		}
	}

	survivingCount := 0
	for _, p := range particles {
		if !p.collided {
			survivingCount++
		}
	}

	return survivingCount
}

type v3 struct{ x, y, z int }

func (v *v3) add(w v3) {
	v.x += w.x
	v.y += w.y
	v.z += w.z
}

type particle struct {
	p, v, a  v3
	collided bool
}

func newParticle(raw string) *particle {
	var p, v, a v3
	fmt.Sscanf(raw, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
		&p.x, &p.y, &p.z,
		&v.x, &v.y, &v.z,
		&a.x, &a.y, &a.z,
	)
	return &particle{p, v, a, false}
}

func (p *particle) tick() {
	p.v.add(p.a)
	p.p.add(p.v)
}

func (p *particle) dist() int {
	return util.AbsInt(p.p.x) +
		util.AbsInt(p.p.y) +
		util.AbsInt(p.p.z)
}
