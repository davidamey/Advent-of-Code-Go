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

	var crt [240]rune
	signal, toAdd := 0, 0
	x := 1
	for c, i := 1, 0; c <= 240; c++ {
		h := (c-1)%40 + 1
		if x == h || x+1 == h || x+2 == h {
			crt[c-1] = '#'
		} else {
			crt[c-1] = ' '
		}

		if c%40 == 20 {
			signal += c * x
		}

		if toAdd != 0 {
			x += toAdd
			toAdd = 0
			continue
		}

		if i < len(lines) {
			l := lines[i]
			i++

			switch l[:4] {
			case "addx":
				fmt.Sscanf(l, "addx %d", &toAdd)
			case "noop":
				// no op
			}
		}
	}

	fmt.Println("p1=", signal)
	fmt.Println("p2=")
	for i, r := range crt {
		fmt.Printf("%c", r)
		if i > 0 && (i+1)%40 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
