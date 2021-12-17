package main

import (
	"advent-of-code-go/2021/intgrid"
	"advent-of-code-go/util"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	c, w, h := intgrid.Parse(lines)

	p1, p2 := 0, 0
	for step := 0; ; step++ {
		for i := range c {
			c[i]++
		}

		for loop := 0; ; loop++ {
			flashed := false

			for i, v := range c {
				if v < 10 {
					continue
				}
				c[i] = 0
				flashed = true

				if step < 100 {
					p1++
				}

				for _, n := range intgrid.Adjacent(i, w, h, true) {
					if c[n] != 0 {
						c[n]++
					}
				}
			}

			if !flashed {
				break
			}
		}

		// // Visual
		// fmt.Printf("\033[0;0H")
		// fmt.Printf("\033[2J")
		// PrintCavern(c)
		// fmt.Println()
		// time.Sleep(150 * time.Millisecond)

		allFlashed := true
		for _, v := range c {
			allFlashed = allFlashed && v == 0
		}
		if allFlashed {
			p2 = step + 1
			break
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func PrintCavern(c []int) {
	var sb strings.Builder
	for i, v := range c {
		if i > 0 && i%10 == 0 {
			sb.WriteByte('\n')
		}
		switch {
		case v >= 10:
			sb.WriteByte('x')
		case v == 0:
			sb.WriteString("\033[92m0\033[0m")
		default:
			sb.WriteByte(byte(v + 48))
		}
	}
	fmt.Println(sb.String())
}
