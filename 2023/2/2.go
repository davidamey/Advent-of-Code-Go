package main

import (
	"advent-of-code-go/util"
	"fmt"
	"regexp"
	"time"
)

const LIMIT_RED = 12
const LIMIT_GREEN = 13
const LIMIT_BLUE = 14

var rgx = regexp.MustCompile(`(\d+) (red|green|blue)`)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	p1, p2 := 0, 0
	for _, l := range lines {
		var gameID int
		fmt.Sscanf(l, "Game %d:", &gameID)

		max := make(map[string]int)
		for _, m := range rgx.FindAllStringSubmatch(l, -1) {
			max[m[2]] = util.MaxInt(max[m[2]], util.Atoi(m[1]))
		}

		if max["red"] <= LIMIT_RED && max["green"] <= LIMIT_GREEN && max["blue"] <= LIMIT_BLUE {
			p1 += gameID
		}

		power := max["red"] * max["green"] * max["blue"]
		p2 += power
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}
