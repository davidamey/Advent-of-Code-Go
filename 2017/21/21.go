package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
	"math"
	"strings"
)

const startGrid = ".#./..#/###"

func main() {
	// input, iter := util.MustReadFileToLines("example"), 2
	input, iter := util.MustReadFileToLines("input"), 18

	twoRules, threeRules := parseRules(input)

	g := stringToGrid(startGrid)

	for i := 0; i < iter; i++ {
		parts, size := gridToParts(g)
		transformedParts := make([]*grid.Grid[rune], len(parts))

		rules := threeRules
		if size == 2 {
			rules = twoRules
		}

	parts:
		for i, part := range parts {
			str := gridToString(part)
			for _, r := range rules {
				for _, p := range r.patterns {
					if p == str {
						transformedParts[i] = r.output
						continue parts
					}
				}
			}
		}

		g = partsToGrid(transformedParts)

		if i == 4 {
			fmt.Println("p1=", onCount(g))
		}
	}

	fmt.Println("p2=", onCount(g))
}

func onCount(g *grid.Grid[rune]) (onCount int) {
	g.ForEach(func(v vector.Vec, r rune) {
		if r == '#' {
			onCount++
		}
	})
	return
}

func gridToString(g *grid.Grid[rune]) string {
	var sb strings.Builder
	var lastY int
	g.ForEach(func(v vector.Vec, r rune) {
		if v.Y != lastY {
			sb.WriteRune('/')
			lastY = v.Y
		}
		sb.WriteRune(r)
	})
	return sb.String()
}

func stringToGrid(str string) *grid.Grid[rune] {
	g := grid.New[rune]()
	for y, row := range strings.Split(str, "/") {
		for x, ch := range row {
			g.SetAt(x, y, ch)
		}
	}
	return g
}

func gridToParts(g *grid.Grid[rune]) (parts []*grid.Grid[rune], partSize int) {
	if (g.Max.X+1)%2 == 0 {
		partSize = 2
	} else {
		partSize = 3
	}

	for y := 0; y <= g.Max.Y; y += partSize {
		for x := 0; x <= g.Max.X; x += partSize {
			parts = append(parts, g.SubGrid(x, y, partSize, partSize))
		}
	}
	return
}

func partsToGrid(parts []*grid.Grid[rune]) *grid.Grid[rune] {
	partSize := parts[0].Max.X + 1
	x := 0
	y := 0
	g := grid.New[rune]()
	gw := int(math.Sqrt(float64(len(parts))))
	for i, p := range parts {
		if i > 0 && i%gw == 0 {
			x = 0
			y += partSize
		}

		p.ForEach(func(v vector.Vec, r rune) {
			g.SetAt(x+v.X, y+v.Y, r)
		})

		x += partSize
	}

	return g
}

type rule struct {
	patterns []string
	output   *grid.Grid[rune]
}

func newTwoRule(in, out string) *rule {
	return &rule{
		[]string{
			string([]byte{in[0], in[1], '/', in[3], in[4]}),
			string([]byte{in[1], in[4], '/', in[0], in[3]}),
			string([]byte{in[4], in[3], '/', in[1], in[0]}),
			string([]byte{in[3], in[0], '/', in[4], in[1]}),
			string([]byte{in[1], in[0], '/', in[4], in[3]}),
			string([]byte{in[0], in[3], '/', in[1], in[4]}),
			string([]byte{in[3], in[4], '/', in[0], in[1]}),
			string([]byte{in[4], in[1], '/', in[3], in[0]}),
		},
		stringToGrid(out),
	}
}

func newThreeRule(in, out string) *rule {
	return &rule{
		[]string{
			string([]byte{in[0], in[1], in[2], '/', in[4], in[5], in[6], '/', in[8], in[9], in[10]}),
			string([]byte{in[2], in[6], in[10], '/', in[1], in[5], in[9], '/', in[0], in[4], in[8]}),
			string([]byte{in[10], in[9], in[8], '/', in[6], in[5], in[4], '/', in[2], in[1], in[0]}),
			string([]byte{in[8], in[4], in[0], '/', in[9], in[5], in[1], '/', in[10], in[6], in[2]}),
			string([]byte{in[2], in[1], in[0], '/', in[6], in[5], in[4], '/', in[10], in[9], in[8]}),
			string([]byte{in[0], in[4], in[8], '/', in[1], in[5], in[9], '/', in[2], in[6], in[10]}),
			string([]byte{in[8], in[9], in[10], '/', in[4], in[5], in[6], '/', in[0], in[1], in[2]}),
			string([]byte{in[10], in[6], in[2], '/', in[9], in[5], in[1], '/', in[8], in[4], in[0]}),
		},
		stringToGrid(out),
	}
}

func parseRules(raw []string) (twoRules, threeRules []*rule) {
	for _, l := range raw {
		parts := strings.Split(l, " => ")
		if len(parts[0]) == 5 {
			twoRules = append(twoRules, newTwoRule(parts[0], parts[1]))
		} else {
			threeRules = append(threeRules, newThreeRule(parts[0], parts[1]))
		}
	}
	return
}
