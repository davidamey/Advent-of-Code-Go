package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	bots := make([]NanoBot, len(lines))
	strongest := &bots[0]
	for i, l := range lines {
		var x, y, z, r int
		fmt.Sscanf(l, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)
		bots[i] = NanoBot{x, y, z, r}
		if bots[i].Radius > strongest.Radius {
			strongest = &bots[i]
		}
	}

	// Part1
	inRangeCount := 0
	for _, b := range bots {
		if strongest.DistanceTo(b) <= strongest.Radius {
			inRangeCount++
		}
	}
	fmt.Println("part1=", inRangeCount)

	// Part2
	// "Cheat" with Z3!
}

type NanoBot struct {
	X, Y, Z int
	Radius  int
}

func (n *NanoBot) DistanceTo(n2 NanoBot) int {
	return util.AbsInt(n2.X-n.X) +
		util.AbsInt(n2.Y-n.Y) +
		util.AbsInt(n2.Z-n.Z)
}

type Grid map[string]int

func Loc(x, y, z int) string {
	return fmt.Sprintf("%d,%d,%d", x, y, z)
}

func (g *Grid) AddRange(x, y, z, r int) {
	for i := x - r; i <= x+r; i++ {
		for j := y - r + util.AbsInt(i); j <= y+r-util.AbsInt(i); j++ {
			for k := z - r + util.AbsInt(i) + util.AbsInt(j); k <= z+r-util.AbsInt(i)-util.AbsInt(j); k++ {
				(*g)[Loc(i, j, k)]++
			}
		}
	}
}
