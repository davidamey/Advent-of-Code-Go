package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	lines := util.MustReadFileToLines("input")
	p1, p2 := solve(lines)

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func solve(codes []string) (maxSeat, missingSeat int) {
	seats := make(map[int]bool)

	for _, c := range codes {
		row := decode(c[:len(c)-3])
		col := decode(c[len(c)-3:])
		seat := 8*row + col
		seats[seat] = true
		if seat > maxSeat {
			maxSeat = seat
		}
	}

	for i := 0; i < maxSeat; i++ {
		if !seats[i] && seats[i-1] && seats[i+1] {
			return maxSeat, i
		}
	}

	panic("failed to solve")
}

func decode(code string) (b int) {
	for _, c := range code {
		switch c {
		case 'F', 'L':
			b = b << 1
		case 'B', 'R':
			b = b<<1 + 1
		}
	}
	return
}
