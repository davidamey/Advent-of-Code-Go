package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
	"strings"
	"time"
)

const totalCycles = 1_000_000_000

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	g := grid.FromLines(lines)

	p1 := 0
	seenGrids := make(map[string]int)
	for i := 0; i < totalCycles; i++ {
		// This could be optimised, but overall it's a <1s runtime - good enough for AoC!
		h := hash(g)
		if prev, seen := seenGrids[h]; seen {
			cycleLength := i - prev
			remainingCycles := (totalCycles - i) % cycleLength
			i = totalCycles - remainingCycles
		}
		seenGrids[h] = i

		// North
		tilt(g)

		if i == 0 {
			p1 = load(g)
		}

		// West
		g.RotateCW()
		tilt(g)

		// South
		g.RotateCW()
		tilt(g)

		// East
		g.RotateCW()
		tilt(g)

		g.RotateCW()
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", load(g))

}

func tilt(g *grid.Grid[rune]) {
	for x := g.Min.X; x <= g.Max.X; x++ {
		canMoveBy := 0
		for y := g.Min.Y; y <= g.Max.Y; y++ {
			switch g.GetAt(x, y) {
			case '#':
				canMoveBy = 0
			case '.':
				canMoveBy++
			case 'O':
				if canMoveBy > 0 {
					g.SetAt(x, y-canMoveBy, 'O')
					g.SetAt(x, y, '.')
				}
			}
		}
	}
}

func load(g *grid.Grid[rune]) int {
	l := 0
	g.ForEach(func(v vector.Vec, r rune) {
		if r == 'O' {
			l += g.Max.Y - v.Y + 1
		}
	})
	return l
}

func hash(g *grid.Grid[rune]) string {
	var sb strings.Builder
	for y := g.Min.Y; y <= g.Max.Y; y++ {
		for x := g.Min.X; x <= g.Max.X; x++ {
			sb.WriteRune(g.GetAt(x, y))
		}
	}
	return sb.String()
}
