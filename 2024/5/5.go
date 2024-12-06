package main

import (
	"advent-of-code-go/util"
	"fmt"
	"slices"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// input := string(util.MustReadFile("example"))
	input := string(util.MustReadFile("input"))

	parts := strings.Split(input, "\n\n")

	rules := map[int]map[int]bool{}
	for _, r := range strings.Split(parts[0], "\n") {
		xs := util.ParseInts(r, "|")
		if _, exists := rules[xs[0]]; !exists {
			rules[xs[0]] = map[int]bool{}
		}
		rules[xs[0]][xs[1]] = true
	}

	sortFunc := func(a, b int) int {
		switch {
		case rules[a][b]:
			return -1
		case rules[b][a]:
			return 1
		default:
			return 0
		}
	}

	p1, p2 := 0, 0
	for _, update := range strings.Split(parts[1], "\n") {
		pages := util.ParseInts(update, ",")
		if slices.IsSortedFunc(pages, sortFunc) {
			p1 += pages[len(pages)/2]
		} else {
			slices.SortStableFunc(pages, sortFunc)
			p2 += pages[len(pages)/2]
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}
