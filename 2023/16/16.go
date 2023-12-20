package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// g := util.MustReadFileToLines("example")
	g := util.MustReadFileToLines("input")

	h, w := len(g), len(g[0])

	p1, p2 := 0, 0
	for y := 0; y < h; y++ {
		e1 := energized(g, []beam{{0, y, 1, 0}})
		e2 := energized(g, []beam{{w - 1, y, -1, 0}})
		if e1 > p2 {
			p2 = e1
		}
		if e2 > p2 {
			p2 = e2
		}
		if y == 0 {
			p1 = e1
		}
	}
	for x := 0; x < w; x++ {
		e1 := energized(g, []beam{{x, 0, 0, 1}})
		e2 := energized(g, []beam{{x, h - 1, 0, -1}})
		if e1 > p2 {
			p2 = e1
		}
		if e2 > p2 {
			p2 = e2
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func inBounds(b beam, w, h int) bool {
	if b[0] < 0 || b[0] >= w {
		return false
	}
	if b[1] < 0 || b[1] >= h {
		return false
	}
	return true
}

func energized(g []string, beams []beam) int {
	h, w := len(g), len(g[0])
	visited := make(map[beam]bool)
	energized := make(map[[2]int]struct{})

	for beamsActive := true; beamsActive; {
		beamsActive = false
		for i, b := range beams {
			if !inBounds(b, w, h) || visited[b] {
				continue
			}

			beamsActive = true
			visited[b] = true
			energized[[2]int{b[0], b[1]}] = struct{}{}

			switch g[b[1]][b[0]] {
			case '.':
				// no-op
			case '/':
				if b[3] == 0 {
					beams[i][2], beams[i][3] = b[3], -b[2]
				} else {
					beams[i][2], beams[i][3] = -b[3], b[2]
				}
			case '\\':
				beams[i][2], beams[i][3] = b[3], b[2]
			case '-':
				if b[2] == 0 {
					beams[i][2], beams[i][3] = 1, 0
					splitBeam := beam{b[0], b[1], -1, 0}
					beams = append(beams, splitBeam)
				}
			case '|':
				if b[3] == 0 {
					beams[i][2], beams[i][3] = 0, 1
					splitBeam := beam{b[0], b[1], 0, -1}
					beams = append(beams, splitBeam)
				}
			}

			beams[i][0] += beams[i][2]
			beams[i][1] += beams[i][3]
		}
	}
	return len(energized)
}

type beam [4]int // x, y, dx, dy
