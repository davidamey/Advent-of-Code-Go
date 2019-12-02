package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	msgLen := len(lines[0])
	counters := make([]*counter, msgLen)
	for i := range lines[0] {
		counters[i] = newCounter()
	}

	for _, l := range lines {
		for i, r := range l {
			counters[i].Add(r)
		}
	}

	p1 := make([]rune, msgLen)
	p2 := make([]rune, msgLen)
	for i, c := range counters {
		p1[i] = c.MaxVal()
		p2[i] = c.MinVal()
	}

	fmt.Println("p1=", string(p1))
	fmt.Println("p2=", string(p2))
}

type counter struct {
	counts   map[rune]int
	maxCount int
	maxVal   rune
}

func newCounter() *counter {
	return &counter{counts: make(map[rune]int)}
}

func (c *counter) Add(r rune) {
	c.counts[r]++
	if c.counts[r] > c.maxCount {
		c.maxCount = c.counts[r]
		c.maxVal = r
	}
}

func (c *counter) MaxVal() rune {
	return c.maxVal
}

func (c *counter) MinVal() (minVal rune) {
	minCount := 1 << 32
	for r, v := range c.counts {
		if v < minCount {
			minCount = v
			minVal = r
		}
	}
	return
}
