package main

import (
	"advent-of-code-go/util"
	"fmt"
	"sort"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// data := util.MustReadFileToLines("example")
	data := util.MustReadFileToLines("input")

	var elves []int
	elf := 0
	for _, d := range data {
		if d == "" {
			elves = append(elves, elf)
			elf = 0
			continue
		}
		elf += util.Atoi(d)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	fmt.Println("p1=", elves[0])
	fmt.Println("p2=", elves[0]+elves[1]+elves[2])
}
