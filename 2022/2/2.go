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

	p1 := 0
	p2 := 0
	for _, l := range lines {
		switch l {
		case "A X":
			p1 += 1 + 3
			p2 += 3 + 0
		case "A Y":
			p1 += 2 + 6
			p2 += 1 + 3
		case "A Z":
			p1 += 3 + 0
			p2 += 2 + 6
		case "B X":
			p1 += 1 + 0
			p2 += 1 + 0
		case "B Y":
			p1 += 2 + 3
			p2 += 2 + 3
		case "B Z":
			p1 += 3 + 6
			p2 += 3 + 6
		case "C X":
			p1 += 1 + 6
			p2 += 2 + 0
		case "C Y":
			p1 += 2 + 0
			p2 += 3 + 3
		case "C Z":
			p1 += 3 + 3
			p2 += 1 + 6
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}
