package main

import (
	"fmt"
)

const (
	safe = '.'
	trap = '^'
)

func main() {
	// m := makeMap(".^^.^.^^^^", 10)  // example
	m := makeMap(".^..^....^....^^.^^.^.^^.^.....^.^..^...^^^^^^.^^^^.^.^^^^^^^.^^^^^..^.^^^.^^..^.^^.^....^.^...^^.^.", 400000) // input

	p1, p2 := 0, 0
	for i, r := range m {
		for _, t := range r {
			if t == safe {
				if i < 40 {
					p1++
				}
				p2++
			}
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func makeMap(firstRow string, rows int) (m [][]rune) {
	m = make([][]rune, rows)
	m[0] = []rune(firstRow)
	for i := 1; i < rows; i++ {
		m[i] = nextRow(m[i-1])
	}
	return
}

func nextRow(prev []rune) (out []rune) {
	out = make([]rune, len(prev))
	for i, c := range prev {
		l, r := safe, safe
		if i > 0 {
			l = rune(prev[i-1])
		}
		if i < len(prev)-1 {
			r = rune(prev[i+1])
		}

		out[i] = tile(l, c, r)
	}
	return
}

func tile(l, c, r rune) rune {
	switch {
	case l == trap && c == trap && r == safe:
		return trap
	case l == safe && c == trap && r == trap:
		return trap
	case l == trap && c == safe && r == safe:
		return trap
	case l == safe && c == safe && r == trap:
		return trap
	default:
		return safe
	}
}
