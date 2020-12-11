package main

import (
	"advent-of-code-go/util"
	"fmt"
	"sort"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// ints := util.MustReadFileToInts("example")
	// ints := util.MustReadFileToInts("example2")
	ints := util.MustReadFileToInts("input")

	sort.Ints(ints)
	ints = append(ints, ints[len(ints)-1]+3)

	p1 := part1(ints)
	p2 := part2(0, ints, map[int]int{ints[len(ints)-1]: 1})

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func part1(ints []int) int {
	x, a, b := 0, 0, 0
	for _, y := range ints {
		switch y - x {
		case 1:
			a++
		case 3:
			b++
		}
		x = y
	}
	return a * b
}

func part2(joltage int, pool []int, cache map[int]int) int {
	if v, ok := cache[joltage]; ok {
		return v
	}

	score := 0
	for i, y := range pool[:3] {
		if d := y - joltage; d >= 1 && d <= 3 {
			score += part2(y, pool[i+1:], cache)
		}
	}

	cache[joltage] = score
	return score
}
