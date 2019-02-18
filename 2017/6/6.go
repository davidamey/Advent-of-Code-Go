package main

import "fmt"

func main() {
	exampleBanks := []int{0, 2, 7, 0}
	inputBanks := []int{5, 1, 10, 0, 1, 7, 13, 14, 3, 12, 8, 10, 7, 12, 0, 6}

	exIterations, exCycleLn := solve(exampleBanks)
	fmt.Println("ex=", exIterations, exCycleLn)
	inputIterations, inputCycleLn := solve(inputBanks)
	fmt.Println("p1=", inputIterations, inputCycleLn)
}

func solve(initial []int) (iterations, cycleLn int) {
	banks := make([]int, len(initial))
	copy(banks, initial)

	seen := make(map[string]int)
	for seen[fmt.Sprintf("%v", banks)] == 0 { // If the cycle includes the first one, this won't work...^_^
		seen[fmt.Sprintf("%v", banks)] = iterations
		redistribute(banks)
		iterations++
	}
	cycleLn = iterations - seen[fmt.Sprintf("%v", banks)]
	return
}

func redistribute(banks []int) {
	maxIdx := len(banks) - 1
	for i := maxIdx - 1; i >= 0; i-- {
		if banks[i] >= banks[maxIdx] {
			maxIdx = i
		}
	}

	blocksToAllocate := banks[maxIdx]
	banks[maxIdx] = 0
	for i := maxIdx; blocksToAllocate > 0; blocksToAllocate-- {
		i = (i + 1) % len(banks)
		banks[i]++
	}
}
