package main

import (
	"advent/util"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	// Examples
	// part1(">")
	// part1("^>v<")
	// part1("^v^v^v^v^v")
	// part2(">")
	// part2("^>v<")
	// part2("^v^v^v^v^v")

	file, _ := util.OpenInput()
	defer file.Close()
	input, _ := ioutil.ReadAll(file)
	part1(string(input))
	part2(string(input))
}

func part1(input string) {
	grid := make(map[string]int)
	x := 0
	y := 0

	deliver(grid, x, y)
	for _, c := range input {
		switch c {
		case '^':
			y++
		case '>':
			x++
		case 'v':
			y--
		case '<':
			x--
		}
		deliver(grid, x, y)
	}

	fmt.Println("== part1 ==")
	fmt.Printf("%d houses get at least 1 present (pattern %s)\n", len(grid), input[:10])
}

func part2(input string) {
	grid := make(map[string]int)
	x1 := 0
	y1 := 0
	x2 := 0
	y2 := 0

	deliver(grid, x1, y1)
	deliver(grid, x2, y2)
	for i, c := range input {
		var x, y *int
		if i%2 == 0 {
			x = &x1
			y = &y1
		} else {
			x = &x2
			y = &y2
		}

		switch c {
		case '^':
			*y++
		case '>':
			*x++
		case 'v':
			*y--
		case '<':
			*x--
		}
		deliver(grid, *x, *y)
	}

	fmt.Println("== part2 ==")
	fmt.Printf("%d houses get at least 1 present (pattern %s)\n", len(grid), input[:10])
}

func deliver(grid map[string]int, x, y int) {
	loc := strconv.Itoa(x) + strconv.Itoa(y)
	grid[loc]++
}
