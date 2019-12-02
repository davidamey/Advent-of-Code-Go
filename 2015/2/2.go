package main

import (
	"advent-of-code-go/util"
	"fmt"
	"math"
)

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	paperReq := 0
	for _, l := range lines {
		paperReq += paperForBox(l)
	}

	fmt.Println("== part1 ==")
	fmt.Printf("paper required = %d\n", paperReq)

}

func part2(lines []string) {
	ribbonReq := 0
	for _, l := range lines {
		ribbonReq += ribbonForBox(l)
	}

	fmt.Println("== part2 ==")
	fmt.Printf("ribbon required = %d\n", ribbonReq)

}

func paperForBox(raw string) int {
	var l, w, h int
	fmt.Sscanf(raw, "%dx%dx%d", &l, &w, &h)

	d1 := l * w
	d2 := w * h
	d3 := h * l

	return 2*d1 + 2*d2 + 2*d3 + minInt(d1, d2, d3)
}

func ribbonForBox(raw string) int {
	var l, w, h int
	fmt.Sscanf(raw, "%dx%dx%d", &l, &w, &h)

	p1 := 2*l + 2*w
	p2 := 2*w + 2*h
	p3 := 2*h + 2*l
	v := l * w * h

	return minInt(p1, p2, p3) + v
}

func minInt(inputs ...int) int {
	min := math.MaxInt32
	for _, i := range inputs {
		if i < min {
			min = i
		}
	}
	return min
}
