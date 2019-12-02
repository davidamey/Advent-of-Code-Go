package main

import (
	"advent-of-code-go/util"
	"fmt"
	"io/ioutil"
)

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	input, _ := ioutil.ReadAll(file)

	part1(input)
	part2(input)
}

func part1(input []byte) {
	floor := 0
	for _, c := range input {
		if c == '(' {
			floor++
		} else {
			floor--
		}
	}

	fmt.Println("== part1 ==")
	fmt.Printf("resultant floor: %d\n", floor)
}

func part2(input []byte) {
	floor := 0
	position := -1
	for i, c := range input {
		if c == '(' {
			floor++
		} else {
			floor--
		}

		if floor == -1 {
			position = i + 1
			break
		}
	}

	fmt.Println("== part2 ==")
	fmt.Printf("position when enters basement: %d\n", position)
}
