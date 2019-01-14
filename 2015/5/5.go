package main

import (
	"advent/util"
	"fmt"
)

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	///

	p1NiceCount := 0
	p2NiceCount := 0
	for _, l := range lines {
		if nice, _ := isNiceP1(l); nice {
			fmt.Println(l, nice)
			p1NiceCount++
		} else {
			fmt.Println(l, nice)
		}
		if isNiceP2(l) {
			p2NiceCount++
		}
	}

	fmt.Printf("p1 gives %d nice lines\n", p1NiceCount)
	fmt.Printf("p2 gives %d nice lines\n", p2NiceCount)

	///

	isNiceP2("xilodxfuxphuiiii")

	/** Part 2 examples **/
	// fmt.Println(isNiceP2("qjhvhtzxzqqjkmpb"))
	// fmt.Println(isNiceP2("xxyxx"))
	// fmt.Println(isNiceP2("uurcxstgmygtbstg"))
	// fmt.Println(isNiceP2("ieodomkazucvgmuy"))
}

func isNiceP1(str string) (bool, string) {
	vowelCount := 0
	haveDouble := false

	for i, c := range str {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowelCount++
		}

		if i > 0 {
			pre := rune(str[i-1])

			if pre == c {
				haveDouble = true
				continue
			}

			if c == 'b' && pre == 'a' ||
				c == 'd' && pre == 'c' ||
				c == 'q' && pre == 'p' ||
				c == 'y' && pre == 'x' {
				return false, "bad pair"
			}
		}
	}

	if vowelCount < 3 {
		return false, "too few vowels"
	}

	if !haveDouble {
		return false, "no double"
	}

	return true, ""
}

func isNiceP2(str string) bool {
	pairs := make(map[string]int)
	hasXYX := false

	for i := 0; i < len(str)-1; i++ {
		if i+2 < len(str) && str[i] == str[i+2] {
			hasXYX = true
		}

		pair := []byte{str[i], str[i+1]}
		if str[i] == str[i+1] {
			for ; i+2 < len(str) && str[i+2] == pair[0]; i++ {
				pair = append(pair, str[i+2])
			}
		}

		pairs[string(pair[:2])] += len(pair) / 2
	}

	hasRepeatedPair := false
	for _, count := range pairs {
		if count > 1 {
			hasRepeatedPair = true
		}
	}

	// fmt.Println(str, hasRepeatedPair, hasXYX)
	// for p, c := range pairs {
	// 	fmt.Printf("%s: %d\n", p, c)
	// }
	return hasRepeatedPair && hasXYX
}
