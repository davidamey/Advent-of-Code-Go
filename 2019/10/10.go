package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/vector"
	"fmt"
	"math"
	"sort"
)

func main() {
	// lines := util.MustReadFileToLines("example")
	// lines := util.MustReadFileToLines("example2")
	lines := util.MustReadFileToLines("input")

	var asteroids []asteroid
	for y, l := range lines {
		for x, c := range l {
			if c == '#' {
				asteroids = append(asteroids, asteroid{pos: vector.New(x, y)})
			}
		}
	}

	p1, p2 := solve(asteroids)
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

type asteroid struct {
	pos        vector.Vec
	sightlines map[float64][]asteroid
}

func solve(asteroids []asteroid) (p1, p2 int) {
	var base asteroid

	for _, a := range asteroids {
		a.sightlines = make(map[float64][]asteroid)
		for _, b := range asteroids {
			if a.pos == b.pos {
				continue
			}
			dx := b.pos.X - a.pos.X
			dy := b.pos.Y - a.pos.Y
			θ := math.Atan2(-float64(dy), float64(dx))
			a.sightlines[θ] = append(a.sightlines[θ], b)
		}
		if canSee := len(a.sightlines); canSee > p1 {
			base = a
			p1 = canSee
		}
	}

	var angles []float64
	for θ, as := range base.sightlines {
		angles = append(angles, θ)
		sort.Slice(as, func(i, j int) bool {
			return as[i].pos.ManhattanTo(base.pos) < as[j].pos.ManhattanTo(base.pos)
		})
	}
	sort.Slice(angles, func(i, j int) bool {
		θi, θj := angles[i], angles[j]
		if (θi <= math.Pi/2) == (θj <= math.Pi/2) {
			return θi >= θj
		}
		return θi <= math.Pi/2
	})

	count := 0
	for {
		for _, θ := range angles {
			as := base.sightlines[θ]
			if len(as) > 0 {
				var a asteroid
				a, base.sightlines[θ] = as[0], as[1:]
				// fmt.Println("laz0ring", a.pos)
				count++
				if count == 200 {
					p2 = a.pos.X*100 + a.pos.Y
					return
				}
			}
		}
	}
}
