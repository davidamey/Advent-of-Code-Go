package main

import (
	"advent-of-code-go/util"
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	p1 := 0
	for _, l := range lines {
		parts := strings.Split(l, " | ")
		for _, d := range strings.Split(parts[1], " ") {
			switch len(d) {
			case 2, 4, 3, 7:
				p1++
			}
		}
	}
	fmt.Println("p1=", p1)

	p2 := 0
	for _, l := range lines {
		parts := strings.Split(l, " | ")
		input := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")

		digits := make(map[int]string)
		for _, d := range input {
			switch len(d) {
			case 2:
				digits[1] = sortWires(d)
			case 3:
				digits[7] = sortWires(d)
			case 4:
				digits[4] = sortWires(d)
			case 7:
				digits[8] = sortWires(d)
			}
		}
		for _, d := range input {
			switch {
			case len(d) == 5 && overlap(digits[7], d) == 3:
				digits[3] = sortWires(d)
			case len(d) == 5 && overlap(digits[4], d) == 3:
				digits[5] = sortWires(d)
			case len(d) == 5:
				digits[2] = sortWires(d)
			case len(d) == 6 && overlap(digits[4], d) == 4:
				digits[9] = sortWires(d)
			case len(d) == 6 && overlap(digits[7], d) == 3:
				digits[0] = sortWires(d)
			case len(d) == 6:
				digits[6] = sortWires(d)
			}
		}

		v := 0
		for _, d := range output {
			ds := sortWires(d)
			for i := range digits {
				if digits[i] == ds {
					v = 10*v + i
				}
			}
		}

		p2 += v
	}

	fmt.Println("p2=", p2)
}

func overlap(s1, s2 string) int {
	return len(util.Intersect(strings.Split(s1, ""), strings.Split(s2, "")))
}

func sortWires(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}
