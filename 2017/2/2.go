package main

import (
	"advent-of-code-go/util"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	checksumP1 := 0
	checksumP2 := 0
	for _, l := range lines {
		p1, p2 := minMax(l)
		checksumP1 += p1
		checksumP2 += p2
	}

	fmt.Println("p1=", checksumP1)
	fmt.Println("p2=", checksumP2)
}

func minMax(s string) (p1, p2 int) {
	min := 1 << 16
	max := -1 << 16
	strs := regexp.MustCompile(`\s+`).Split(s, -1)

	nums := make([]int, len(strs))
	for i, n := range strs {
		nums[i], _ = strconv.Atoi(n)
	}

	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}

		if p2 == 0 {
			for _, n2 := range nums {
				if n != n2 && n%n2 == 0 {
					p2 = n / n2
					break
				}
			}
		}
	}
	return max - min, p2
}
