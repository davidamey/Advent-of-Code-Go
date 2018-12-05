package main

import (
	"advent/2018/util"
	"fmt"
	"log"
)

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	inputs, _ := util.ReadLines(file)

	part1(inputs)
	part2(inputs)
}

func part1(inputs []string) {
	twos := 0
	threes := 0
	for _, s := range inputs {
		hasExactTwo, hasExactThree := checkCounts(s)

		if hasExactTwo {
			twos++
		}

		if hasExactThree {
			threes++
		}
	}

	fmt.Printf("found %d with exactly two and %d with exactly three\n", twos, threes)
	fmt.Printf("checksum: %d\n", twos*threes)
}

func part2(inputs []string) {
	diff, ok := findDiffByOne(inputs)
	if !ok {
		log.Fatal("failed")
	}

	fmt.Printf("common letters: %s\n", diff)
}

func findDiffByOne(inputs []string) (string, bool) {
	for _, s1 := range inputs {
		for _, s2 := range inputs {
			if s1 == s2 {
				continue
			}

			byOne, sameChars := diffStrings(s1, s2)
			if byOne {
				return sameChars, true
			}
		}
	}
	return "", false
}

func diffStrings(s1, s2 string) (bool, string) {
	diffCount := 0
	var sameChars []byte
	for i := range s1 {
		if s1[i] == s2[i] {
			sameChars = append(sameChars, s1[i])
		} else {
			diffCount++
		}

		if diffCount > 1 {
			return false, string(sameChars)
		}
	}

	return diffCount == 1, string(sameChars)
}

func checkCounts(s string) (hasExactTwo bool, hasExactThree bool) {
	letterCounts := make(map[rune]int, 26)
	for _, c := range s {
		letterCounts[c]++
	}

	for _, count := range letterCounts {
		if count == 2 {
			hasExactTwo = true
		}

		if count == 3 {
			hasExactThree = true
		}
	}

	return
}
