package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// grid := util.MustReadFileToLines("example")
	grid := util.MustReadFileToLines("input")

	fmt.Println("p1=", p1(grid))
	fmt.Println("p2=", p2(grid))
}

func p1(haystack []string) (found int) {
	for y, l := range haystack {
		for x := range l {
			for _, dx := range []int{-1, 0, 1} {
				for _, dy := range []int{-1, 0, 1} {
					if checkWord(haystack, "XMAS", 0, x, y, dx, dy) {
						found++
					}
				}
			}
		}
	}
	return
}

var xMas map[string]bool = map[string]bool{
	"MSAMS": true,
	"MMASS": true,
	"SSAMM": true,
	"SMASM": true,
}

func p2(haystack []string) (found int) {
	for y, l := range haystack[:len(haystack)-2] {
		for x := range l[:len(l)-2] {
			if xMas[string([]byte{
				haystack[y][x],
				haystack[y][x+2],
				haystack[y+1][x+1],
				haystack[y+2][x],
				haystack[y+2][x+2],
			})] {
				found++
			}
		}
	}
	return
}

func checkWord(haystack []string, needle string, i, x, y, dx, dy int) bool {
	if i == len(needle) {
		return true
	}

	if x < 0 || y < 0 || x >= len(haystack[0]) || y >= len(haystack) {
		return false
	}

	if haystack[y][x] != needle[i] {
		return false
	}

	return checkWord(haystack, needle, i+1, x+dx, y+dy, dx, dy)
}
