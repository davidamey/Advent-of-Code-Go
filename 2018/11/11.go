package main

import (
	"advent/util"
	"fmt"
	"time"
)

func main() {
	// example()
	part1()
	part2()
}

func example() {
	fmt.Println("3,5 @ 8: ", power(3, 5, 8))           // 4
	fmt.Println("122,79 @ 57: ", power(122, 79, 57))   // -5
	fmt.Println("217,196 @ 39: ", power(217, 196, 39)) // 0
	fmt.Println("101,153 @ 71: ", power(101, 153, 71)) // 4

	// g1 := makeGrid(18)
	// g1.Print(32, 44, 36, 50)
	// fmt.Println("33,45", g1["33,45"])
	// fmt.Println("33,46", g1["33,46"])
	// fmt.Println("33,47", g1["33,47"])

	// g2 := makeGrid(42)
	// fmt.Println("21,61", g2["21,61"])
	// fmt.Println("21,62", g2["21,62"])
	// fmt.Println("21,63", g2["21,63"])

	// fmt.Printf("serial 18 gives maxPower of %d at %d,%d\n", maxPowerForGridP1(makeGrid(18)))
	// fmt.Printf("serial 42 gives maxPower of %d at %d,%d\n", maxPowerForGridP1(makeGrid(42)))
}

func part1() {
	fmt.Println("== part1 == ")
	maxPower, x, y := maxPowerForGridP1(makeGrid(3463))
	fmt.Printf("serial 3463 gives maxPower of %d at %d,%d\n", maxPower, x, y)
}

func part2() {
	start := time.Now()
	var powerMax, xMax, yMax, sizeMax int

	g := makeGrid(3463)
	summer := NewSummer(&g)

	for size := 1; size < 300; size++ {
		if size%10 == 0 {
			fmt.Printf("* checking size %d\n", size)
		}
		for x := 1; x <= 300-size+1; x++ {
			for y := 1; y <= 300-size+1; y++ {
				power := summer.Sum(size, x, y)
				if power > powerMax {
					powerMax = power
					sizeMax = size
					xMax = x
					yMax = y
				}
			}
		}
	}

	fmt.Println("== part2 ==")
	fmt.Printf("max power of %d achieved at %d,%d with square size %d\n", powerMax, xMax, yMax, sizeMax)
	fmt.Printf("took %fs", time.Since(start).Seconds())
}

type Summer struct {
	grid *util.Grid
	sums map[string]int
}

func NewSummer(grid *util.Grid) *Summer {
	return &Summer{
		grid,
		make(map[string]int),
	}
}

func (s *Summer) Loc(size, x, y int) string {
	return fmt.Sprintf("%d,%d,%d", size, x, y)
}

func (s *Summer) Sum(size, x, y int) int {
	loc := s.Loc(size, x, y)

	if size == 1 {
		return s.grid.Get(x, y)
	}

	if sum, ok := s.sums[loc]; ok {
		return sum
	}

	sum := s.Sum(size-1, x, y)
	for i := 0; i < size; i++ {
		sum += s.grid.Get(x+i, y+size-1)
	}
	for j := 0; j < size-1; j++ {
		sum += s.grid.Get(x+size-1, y+j)
	}
	s.sums[loc] = sum
	return sum
}

func maxPowerForGridP1(g util.Grid) (maxPower, xMax, yMax int) {
	for x := 1; x <= 300-2; x++ {
		for y := 1; y <= 300-2; y++ {
			power := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					loc := fmt.Sprintf("%d,%d", x+i, y+j)
					power += g[loc]
				}
			}
			if power > maxPower {
				maxPower = power
				xMax = x
				yMax = y
			}
		}
	}
	return
}

func makeGrid(serial int) util.Grid {
	grid := make(util.Grid)
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			grid.Set(x, y, power(x, y, serial))
		}
	}
	return grid
}

func power(x, y, serial int) int {
	rackID := x + 10
	power := rackID * y
	power += serial
	power *= rackID
	power = power / 100 % 10
	power -= 5

	return power
}
