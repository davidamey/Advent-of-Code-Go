package main

import (
	"advent-of-code-go/util"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	lines := util.MustReadFileToLines("input")

	p1, p2 := 0, 0
	for _, l := range lines {
		v1, v2 := isValid(l)
		if v1 {
			p1++
		}
		if v2 {
			p2++
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func isValid(phrase string) (bool, bool) {
	words := make(map[string]struct{})
	scanner := bufio.NewScanner(strings.NewReader(phrase))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		w := scanner.Text()
		if _, ok := words[w]; ok {
			return false, false
		}
		words[w] = struct{}{}
	}

	for w1 := range words {
		for w2 := range words {
			if w1 == w2 || len(w1) != len(w2) {
				continue
			}

			runeCount := make(map[rune]int)
			for _, c := range w1 {
				runeCount[c]++
			}
			for _, c := range w2 {
				runeCount[c]--
				if runeCount[c] == 0 {
					delete(runeCount, c)
				}
			}

			if len(runeCount) == 0 {
				return true, false
			}
		}
	}

	return true, true
}
