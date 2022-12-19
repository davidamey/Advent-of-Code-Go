package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	visibleTrees := 0
	scenicScores := []int{}
	for y, row := range lines {
		for x := range row {
			visible, scenicScore := inspect(x, y, lines)
			if visible {
				visibleTrees++
			}
			scenicScores = append(scenicScores, scenicScore)
		}
	}

	fmt.Println("p1=", visibleTrees)
	fmt.Println("p2=", util.MaxInt(scenicScores...))
}

func inspect(tx, ty int, trees []string) (visible bool, scenicScore int) {
	v1, s1 := up(tx, ty, trees)
	v2, s2 := right(tx, ty, trees)
	v3, s3 := down(tx, ty, trees)
	v4, s4 := left(tx, ty, trees)
	return v1 || v2 || v3 || v4, s1 * s2 * s3 * s4
}

func up(tx, ty int, trees []string) (visible bool, ss int) {
	var hs []rune
	for y := ty; y >= 0; y-- {
		hs = append(hs, rune(trees[y][tx]))
	}
	return search(hs)
}

func down(tx, ty int, trees []string) (visible bool, ss int) {
	var hs []rune
	for y := ty; y < len(trees); y++ {
		hs = append(hs, rune(trees[y][tx]))
	}
	return search(hs)
}

func left(tx, ty int, trees []string) (visible bool, ss int) {
	var hs []rune
	for x := tx; x >= 0; x-- {
		hs = append(hs, rune(trees[ty][x]))
	}
	return search(hs)
}

func right(tx, ty int, trees []string) (visible bool, ss int) {
	return search([]rune(trees[ty][tx:]))
}

func search(haystack []rune) (visible bool, ss int) {
	t := haystack[0]
	for _, c := range haystack[1:] {
		ss++
		if c >= t {
			return
		}
	}
	return true, ss
}
