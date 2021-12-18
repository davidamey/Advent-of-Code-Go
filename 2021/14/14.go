package main

import (
	"advent-of-code-go/util"
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	template := lines[0]
	rules := make(map[string]string)
	for _, l := range lines[2:] {
		parts := strings.Split(l, " -> ")
		rules[parts[0]] = parts[1]
	}

	fmt.Println("p1=", process(template, rules, 10))
	fmt.Println("p2=", process(template, rules, 40))
}

func process(template string, rules map[string]string, iterations int) int {
	pairs := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]]++
	}

	for i := 0; i < iterations; i++ {
		newPairs := make(map[string]int)
		for p, v := range pairs {
			newPairs[string(p[0])+rules[p]] += v
			newPairs[rules[p]+string(p[1])] += v
		}
		pairs = newPairs
	}

	counts := make(map[byte]int)
	for p, v := range pairs {
		counts[p[0]] += v
	}
	counts[template[len(template)-1]]++ // Pair loop misses last element

	min, max := math.MaxInt, math.MinInt
	for _, v := range counts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return max - min
}
