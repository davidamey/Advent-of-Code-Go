package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := 509671

	part1(input)

	fmt.Println("== part2 ==")
	part2("51589")
	part2("01245")
	part2("92510")
	part2("59414")
	part2("509671")
}

func part1(input int) {
	recipies := []int{3, 7}
	elf1 := 0
	elf2 := 1

	for len(recipies) < input+10 {
		// Make new recipes
		sum := recipies[elf1] + recipies[elf2]
		new := strings.Split(strconv.Itoa(sum), "")
		for _, c := range new {
			r, _ := strconv.Atoi(c)
			recipies = append(recipies, r)
		}

		// Update elves
		elf1 = (elf1 + 1 + recipies[elf1]) % len(recipies)
		elf2 = (elf2 + 1 + recipies[elf2]) % len(recipies)
	}

	fmt.Println("== part1 ==")
	for _, i := range recipies[input:] {
		fmt.Printf("%d", i)
	}
	fmt.Println()
}

func part2(input string) {
	recipies := []int{3, 7}
	elf1 := 0
	elf2 := 1

	found := false
	at := -1
	for !found {
		// Make new recipes
		sum := recipies[elf1] + recipies[elf2]
		new := strings.Split(strconv.Itoa(sum), "")
		for _, c := range new {
			r, _ := strconv.Atoi(c)
			recipies = append(recipies, r)
		}

		// Update elves
		elf1 = (elf1 + 1 + recipies[elf1]) % len(recipies)
		elf2 = (elf2 + 1 + recipies[elf2]) % len(recipies)

		found, at = contains(recipies, input)
	}

	// fmt.Println("== part2 ==")
	fmt.Println(input, at)
}

func contains(recipies []int, input string) (bool, int) {
	needle := strings.Split(input, "")

	if len(recipies) < len(needle) {
		return false, -1
	}

	offset := len(recipies) - len(needle)
	if len(recipies) > 2*len(needle) {
		offset = len(recipies) - 2*len(needle)
	}
	haystack := recipies[offset:]

	for i := range haystack {
		if i+len(needle) >= len(haystack) {
			break
		}

		for j := range needle {
			if strconv.Itoa(haystack[i+j]) != needle[j] {
				break
			}
			if j == len(needle)-1 {
				return true, offset + i
			}
		}
	}

	return false, -1
}
