package main

import (
	"advent-of-code-go/util"
	"fmt"
	"sort"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	p1 := 0
	var fixScores []int
	for _, l := range lines {
		if corrupt, score, fixScore := check(l); corrupt {
			p1 += score
		} else {
			fixScores = append(fixScores, fixScore)
		}
	}

	sort.Ints(fixScores)

	fmt.Println("p1=", p1)
	fmt.Println("p2=", fixScores[len(fixScores)/2])
}

func check(chunk string) (bool, int, int) {
	pairs := map[rune]rune{
		')': '(', ']': '[',
		'}': '{', '>': '<',
	}
	scores := map[rune]int{
		')': 3, ']': 57, '}': 1197, '>': 25137,
		'(': 1, '[': 2, '{': 3, '<': 4,
	}

	chars := make([]rune, 0, len(chunk))
	for _, c := range chunk {
		i := len(chars) - 1

		switch c {
		case '(', '[', '{', '<':
			chars = append(chars, c)
		default:
			if chars[i] != pairs[c] { // corrupt
				return true, scores[c], 0
			}
			chars = chars[:i]
		}
	}

	fixScore := 0
	for len(chars) > 0 {
		i := len(chars) - 1
		fixScore *= 5
		fixScore += scores[chars[i]]
		chars = chars[:i]
	}
	return false, 0, fixScore
}
