package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/vector"
	"fmt"
	"strconv"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// instructions := util.MustReadFileToLines("example")
	instructions := util.MustReadFileToLines("input")

	fmt.Println("p1=", part1(instructions))
	fmt.Println("p2=", part2(instructions))
}

func part1(instructions []string) int {
	ship := vector.New(0, 0)
	dir := vector.New(1, 0)
	for _, move := range instructions {
		d := move[0]
		i, _ := strconv.Atoi(move[1:])
		switch d {
		case 'N':
			ship.Y -= i
		case 'S':
			ship.Y += i
		case 'E':
			ship.X += i
		case 'W':
			ship.X -= i
		case 'L':
			for a := 0; a < i; a += 90 {
				dir.X, dir.Y = dir.Y, -dir.X
			}
		case 'R':
			for a := 0; a < i; a += 90 {
				dir.X, dir.Y = -dir.Y, dir.X
			}
		case 'F':
			ship.X += dir.X * i
			ship.Y += dir.Y * i
		}
	}
	return ship.Manhattan()
}

func part2(instructions []string) int {
	ship := vector.New(0, 0)
	wp := vector.New(10, -1)
	for _, move := range instructions {
		d := move[0]
		i, _ := strconv.Atoi(move[1:])
		switch d {
		case 'N':
			wp.Y -= i
		case 'S':
			wp.Y += i
		case 'E':
			wp.X += i
		case 'W':
			wp.X -= i
		case 'L':
			for a := 0; a < i; a += 90 {
				wp.X, wp.Y = wp.Y, -wp.X
			}
		case 'R':
			for a := 0; a < i; a += 90 {
				wp.X, wp.Y = -wp.Y, wp.X
			}
		case 'F':
			ship.X += wp.X * i
			ship.Y += wp.Y * i
		}
	}
	return ship.Manhattan()
}
