package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

type vec [2]int

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	w := len(lines[0])
	h := len(lines)

	antennas := map[rune][]vec{}
	for y, l := range lines {
		for x, r := range l {
			if r != '.' {
				antennas[r] = append(antennas[r], vec{x, y})
			}
		}
	}

	antinodesP1 := map[vec]struct{}{}
	antinodesP2 := map[vec]struct{}{}
	for _, n := range antennas {
		for _, an := range findAllAntinodes(n, w, h, false) {
			antinodesP1[an] = struct{}{}
		}
		for _, an := range findAllAntinodes(n, w, h, true) {
			antinodesP2[an] = struct{}{}
		}
	}

	fmt.Println("p1=", len(antinodesP1))
	fmt.Println("p2=", len(antinodesP2))
}

func findAllAntinodes(nodes []vec, w, h int, withResonance bool) (antinodes []vec) {
	for _, u := range nodes {
		for _, v := range nodes {
			antinodes = append(antinodes, findAntinodes(u, v, w, h, withResonance)...)
		}
	}
	return
}

func findAntinodes(u, v vec, w, h int, withResonance bool) (antinodes []vec) {
	if u == v {
		if withResonance {
			return []vec{u}
		}
		return []vec{}
	}

	dx := v[0] - u[0]
	dy := v[1] - u[1]

	x, y := u[0], u[1]
	for {
		x, y = x-dx, y-dy
		if x < 0 || x >= w || y < 0 || y >= h {
			break
		}
		antinodes = append(antinodes, vec{x, y})

		if !withResonance {
			break
		}
	}
	return
}
