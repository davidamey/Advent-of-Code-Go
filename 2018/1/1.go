package main

import (
	"advent/util"
	"fmt"
)

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	inputs, _ := util.ReadLinesToInts(file)

	fmt.Println("Part 1: Sum the inputs")
	part1(inputs)
	fmt.Println()
	fmt.Println("Part 2: Find repeated sum")
	part2(inputs)
	fmt.Println()
}

func part1(inputs []int) {
	sum := 0
	for _, i := range inputs {
		sum += i
	}

	fmt.Printf("Sum: %d\n", sum)
}

func part2(inputs []int) {
	sums := make(map[int]int)
	sum := 0
	i := 0

	for sums[sum] < 2 {
		sum += inputs[i]
		sums[sum]++
		i = (i + 1) % len(inputs)
	}

	fmt.Printf("Repeated sum %d after %d iterations", sum, i)
}
