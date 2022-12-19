package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// input := util.MustReadFile("example")
	input := util.MustReadFile("input")

	fmt.Println("p1=", findMarker(input, 4))
	fmt.Println("p2=", findMarker(input, 14))
}

func findMarker(input []byte, size int) int {
	for i := range input {
		if i < size {
			continue
		}
		if allUnique(input[i-size : i]) {
			return i
		}
	}
	panic("no marker found")
}

func allUnique(data []byte) bool {
	counts := make(map[byte]struct{}, len(data))
	for _, b := range data {
		counts[b] = struct{}{}
	}
	return len(counts) == len(data)
}
