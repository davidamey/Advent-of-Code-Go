package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	var duplicateItems []int
	for _, l := range lines {
		duplicateItems = append(duplicateItems, findDuplicate(l))
	}
	fmt.Println("p1=", util.IntSum(duplicateItems...))

	var badges []int
	for i := 0; i < len(lines); i += 3 {
		badges = append(badges, findBadges(lines[i:i+3]))
	}
	fmt.Println("p2=", util.IntSum(badges...))
}

func findDuplicate(rucksack string) int {
	mid := len(rucksack) / 2
	cmp1 := rucksack[:mid]
	cmp2 := rucksack[mid:]

	for _, r := range cmp1 {
		if strings.ContainsRune(cmp2, r) {
			return toPriority(r)
		}
	}
	panic("No duplicate item found")
}

func findBadges(rucksacks []string) int {
	for _, r := range rucksacks[0] {
		if strings.ContainsRune(rucksacks[1], r) && strings.ContainsRune(rucksacks[2], r) {
			return toPriority(r)
		}
	}
	panic("No badges found")
}

func toPriority(r rune) int {
	if r > 'a' {
		return 1 + int(r-'a')
	}
	return 27 + int(r-'A')
}
