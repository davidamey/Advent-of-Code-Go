package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	p1, p2 := 0, 0
	for _, l := range lines {
		data, counts, _ := strings.Cut(l, " ")
		p1 += newProcessor(data, counts).process(0, 0, 0)

		expandedData, expandedCounts := make([]string, 5), make([]string, 5)
		for i := range expandedData {
			expandedData[i] = data
			expandedCounts[i] = counts
		}
		p2 += newProcessor(
			strings.Join(expandedData, "?"),
			strings.Join(expandedCounts, ","),
		).process(0, 0, 0)
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

type processor struct {
	cache  *map[string]int
	data   string
	counts []int
}

func newProcessor(data, counts string) *processor {
	return &processor{
		cache:  &map[string]int{},
		data:   data,
		counts: util.ParseInts(counts, ","),
	}
}

func (p *processor) process(dIdx, cIdx, count int) int {
	if dIdx >= len(p.data) {
		if cIdx >= len(p.counts) {
			return 1
		}

		if cIdx == len(p.counts)-1 && p.counts[cIdx] == count {
			return 1
		}

		return 0
	}

	switch p.data[dIdx] {
	case '.':
		if count == 0 {
			return p.process(dIdx+1, cIdx, count)
		}
		if cIdx >= len(p.counts) || p.counts[cIdx] != count {
			return 0
		}
		return p.process(dIdx+1, cIdx+1, 0)
	case '#':
		if cIdx >= len(p.counts) || p.counts[cIdx] < count+1 {
			return 0
		}
		return p.process(dIdx+1, cIdx, count+1)
	case '?':
		cacheKey := fmt.Sprintf("%d-%d-%d", dIdx, cIdx, count)
		if x, inCache := (*p.cache)[cacheKey]; inCache {
			return x
		}

		opts := 0

		if count == 0 {
			opts += p.process(dIdx+1, cIdx, count)
		}

		if cIdx < len(p.counts) {
			if count < p.counts[cIdx] {
				opts += p.process(dIdx+1, cIdx, count+1)
			}

			if count == p.counts[cIdx] {
				opts += p.process(dIdx+1, cIdx+1, 0)
			}
		}

		(*p.cache)[cacheKey] = opts
		return opts
	}

	return -1
}
