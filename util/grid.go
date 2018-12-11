package util

import "fmt"

type Grid map[string]int

func (g Grid) Set(x, y, i int) {
	loc := fmt.Sprintf("%d,%d", x, y)
	g[loc] = i
}

func (g Grid) Get(x, y int) int {
	loc := fmt.Sprintf("%d,%d", x, y)
	return g[loc]
}

func (g Grid) Print(x1, y1, x2, y2 int) {
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			fmt.Printf(" %2d ", g.Get(x, y))
		}
		fmt.Println()
	}
	fmt.Println()
}
