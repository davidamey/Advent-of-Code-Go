package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// raw := string(util.MustReadFile("example"))
	raw := string(util.MustReadFile("input"))

	var patterns [][]string
	for _, p := range strings.Split(raw, "\n\n") {
		patterns = append(patterns, strings.Fields(p))
	}

	p1, p2 := 0, 0
	for _, p := range patterns {
		p1 += findSymmetry(p, 0)
		p2 += findSymmetry(p, 1)
	}
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func findSymmetry(pattern []string, smudges int) int {
	// Horizontal
	for i := 0; i < len(pattern)-1; i++ {
		if isSymmetryAt(pattern, i, smudges) {
			return 100 * (i + 1)
		}
	}

	// Vertical
	rotated := rotate(pattern)
	for i := 0; i < len(rotated)-1; i++ {
		if isSymmetryAt(rotated, i, smudges) {
			return i + 1
		}
	}

	panic("no symmetry found")
}

func isSymmetryAt(lines []string, i int, smudges int) bool {
	diffCount := 0
	for x, y := i, i+1; x >= 0 && y < len(lines); x, y = x-1, y+1 {
		for j := range lines[x] {
			if lines[x][j] != lines[y][j] {
				diffCount++
			}
			if diffCount > smudges {
				return false
			}
		}
	}
	return diffCount == smudges

}

func rotate(pattern []string) []string {
	rotated := make([]string, len(pattern[0]))
	for i := range pattern[0] {
		line := make([]byte, len(pattern))
		for j := range line {
			line[j] = pattern[len(line)-1-j][i]
		}
		rotated[i] = string(line)
	}
	return rotated
}
