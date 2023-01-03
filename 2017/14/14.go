package main

import (
	"advent-of-code-go/2017/10/knothash"
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
)

// const input = "flqrgnkx" // example
const input = "hwlqcszp"

func main() {
	g := grid.New[rune]()

	used := 0
	for y := 0; y < 128; y++ {
		bits := hexToBits(knothash.Compute([]byte(fmt.Sprintf("%s-%d", input, y))))
		for x, b := range bits {
			g.SetAt(x, y, rune(b))
			if b == '1' {
				used++
			}
		}
	}

	regionCount := 0
	g.ForEach(func(v vector.Vec, r rune) {
		if r == '0' || r == '2' {
			return
		}

		// found a new region
		regionCount++

		// spider this region setting all runes to '2'
		queue := []vector.Vec{v}
		for len(queue) > 0 {
			q := queue[0]
			queue = queue[1:]

			g.Set(q, '2')

			for _, a := range q.Adjacent(false) {
				if g.InBounds(a) && g.Get(a) == '1' {
					queue = append(queue, a)
				}
			}
		}
	})

	fmt.Println("p1=", used)
	fmt.Println("p2=", regionCount)
}

func hexToBits(hex string) []byte {
	result := make([]byte, 0, 128)
	for _, h := range hex {
		var r rune
		switch {
		case h >= '0' && h <= '9':
			r = h - '0'
		case h >= 'a' && h <= 'f':
			r = 10 + h - 'a'
		}

		result = append(result, fmt.Sprintf("%04b", r)...)
	}
	return result
}
