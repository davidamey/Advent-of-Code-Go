package main

import (
	"fmt"
	"time"
)

func main() {
	// example()
	solve(18, 300)
	solve(42, 300)
	solve(3463, 300)
	// part1()
	// part2()
}

func example() {
	fmt.Println("3,5 @ 8: ", powerFuncForSerial(8)(3, 5))           // 4
	fmt.Println("122,79 @ 57: ", powerFuncForSerial(57)(122, 79))   // -5
	fmt.Println("217,196 @ 39: ", powerFuncForSerial(39)(217, 196)) // 0
	fmt.Println("101,153 @ 71: ", powerFuncForSerial(71)(101, 153)) // 4

	// g1 := NewGrid(powerFuncForSerial(18), 300, 300)
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

func solve(serial, gridSize int) {
	start := time.Now()
	defer func() { fmt.Printf(" in %fs\n", time.Since(start).Seconds()) }()

	g := NewGrid(powerFuncForSerial(serial), gridSize, gridSize)

	maxSum := 0
	maxX := 0
	maxY := 0
	maxSize := 0
	sums := NewGrid(func(x, y int) int {
		if x+2 >= gridSize || y+2 >= gridSize {
			return 0
		}

		sum := 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				sum += g[y+j][x+i]
			}
		}

		if sum > maxSum {
			maxSum = sum
			maxX = x
			maxY = y
			maxSize = 3
		}

		return sum
	}, gridSize, gridSize)

	// fmt.Println("maxSum|maxX|maxY|maxSize", maxSum, maxX+1, maxY+1, maxSize)

	for size := 4; size < 300; size++ {
		for y := 0; y+size < gridSize; y++ {
			for x := 0; x+size < gridSize; x++ {
				for i := 0; i < size; i++ {
					sums[y][x] += g[y+size-1][x+i] + g[y+i][x+size-1]
				}
				sums[y][x] += g[y+size-1][x+size-1]
				if sums[y][x] > maxSum {
					maxSum = sums[y][x]
					maxX = x
					maxY = y
					maxSize = size
				}
			}
		}
	}

	fmt.Printf("%d: maxSum of %d found at %d,%d,%d", serial, maxSum, maxX+1, maxY+1, maxSize)
}

func powerFuncForSerial(serial int) func(x, y int) int {
	return func(x, y int) int {
		// problem uses 1-based
		x += 1
		y += 1

		rackID := x + 10
		power := rackID * y
		power += serial
		power *= rackID
		power = power / 100 % 10
		power -= 5

		return power
	}
}

type Grid [][]int

func NewGrid(f func(x, y int) int, dx, dy int) Grid {
	g := make(Grid, dy)
	for y := 0; y < dy; y++ {
		g[y] = make([]int, dx)
		for x := 0; x < dx; x++ {
			g[y][x] = f(x, y)
		}
	}
	return g
}

func (g Grid) Print(x1, y1, x2, y2 int) {
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			fmt.Printf(" %2d ", g[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}

/** original inefficient code below **/

// func part1() {
// 	fmt.Println("== part1 == ")
// 	maxPower, x, y := maxPowerForGridP1(makeGrid(3463))
// 	fmt.Printf("serial 3463 gives maxPower of %d at %d,%d\n", maxPower, x, y)
// }

// func part2() {
// 	start := time.Now()
// 	var powerMax, xMax, yMax, sizeMax int

// 	g := makeGrid(3463)
// 	summer := NewSummer(&g)

// 	for size := 1; size < 300; size++ {
// 		if size%10 == 0 {
// 			fmt.Printf("* checking size %d\n", size)
// 		}
// 		for x := 1; x <= 300-size+1; x++ {
// 			for y := 1; y <= 300-size+1; y++ {
// 				power := summer.Sum(size, x, y)
// 				if power > powerMax {
// 					powerMax = power
// 					sizeMax = size
// 					xMax = x
// 					yMax = y
// 				}
// 			}
// 		}
// 	}

// 	fmt.Println("== part2 ==")
// 	fmt.Printf("max power of %d achieved at %d,%d with square size %d\n", powerMax, xMax, yMax, sizeMax)
// 	fmt.Printf("took %fs", time.Since(start).Seconds())
// }

// type Summer struct {
// 	grid *util.Grid
// 	sums map[string]int
// }

// func NewSummer(grid *util.Grid) *Summer {
// 	return &Summer{
// 		grid,
// 		make(map[string]int),
// 	}
// }

// func (s *Summer) Loc(size, x, y int) string {
// 	return fmt.Sprintf("%d,%d,%d", size, x, y)
// }

// func (s *Summer) Sum(size, x, y int) int {
// 	loc := s.Loc(size, x, y)

// 	if size == 1 {
// 		return s.grid.Get(x, y)
// 	}

// 	if sum, ok := s.sums[loc]; ok {
// 		return sum
// 	}

// 	sum := s.Sum(size-1, x, y)
// 	for i := 0; i < size-1; i++ {
// 		sum += s.grid.Get(x+i, y+size-1)
// 		sum += s.grid.Get(x+size-1, y+i)
// 	}
// 	sum += s.grid.Get(x+size, y+size)
// 	s.sums[loc] = sum
// 	return sum
// }
