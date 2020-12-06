package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
)

func main() {
	// raw := string(util.MustReadFile("example"))
	raw := string(util.MustReadFile("input"))

	p1, p2 := 0, 0
	for _, group := range strings.Split(raw, "\n\n") {
		numPeople := 1
		answers := make(map[rune]int)
		for _, ans := range group {
			if ans == '\n' {
				numPeople++
				continue
			}
			answers[ans]++
		}

		p1 += len(answers)
		for _, x := range answers {
			if x == numPeople {
				p2++
			}
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}
