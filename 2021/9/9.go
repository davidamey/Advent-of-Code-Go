package main

import (
	"advent-of-code-go/2021/intgrid"
	"advent-of-code-go/util"
	"fmt"
	"sort"
	"time"
)

var floor []int
var w, h int

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	floor, w, h = intgrid.Parse(lines)

	p1 := 0
	var lows []int
	for i, x := range floor {
		lowest := true
		for _, a := range intgrid.Adjacent(i, w, h, false) {
			if floor[a] <= x {
				lowest = false
				break
			}
		}
		if lowest {
			lows = append(lows, i)
			p1 += x + 1
		}
	}

	var sizes []int
	for _, i := range lows {
		sizes = append(sizes, search(i, make([]bool, len(floor))))
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[j] < sizes[i]
	})

	fmt.Println("p1=", p1)
	fmt.Println("p2=", sizes[0]*sizes[1]*sizes[2])

}

func search(i int, seen []bool) int {
	seen[i] = true
	size := 1
	for _, a := range intgrid.Adjacent(i, w, h, false) {
		if !seen[a] && floor[a] < 9 {
			size += search(a, seen)
		}
	}
	return size
}
