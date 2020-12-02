package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	lines := util.MustReadFileToLines("input")
	p1, p2 := 0, 0
	for _, l := range lines {
		var a, b int
		var c rune
		var pass string
		fmt.Sscanf(l, "%d-%d %c: %s", &a, &b, &c, &pass)

		if isValidP1(a, b, c, pass) {
			p1++
		}
		if isValidP2(a, b, c, pass) {
			p2++
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func isValidP1(min, max int, c rune, pass string) bool {
	cCount := 0
	for _, p := range pass {
		if p == c {
			cCount++
		}
	}
	return cCount >= min && cCount <= max
}

func isValidP2(i, j int, c rune, pass string) bool {
	b := byte(c)
	if (pass[i-1] == b || pass[j-1] == b) && pass[i-1] != pass[j-1] {
		return true
	}
	return false
}
