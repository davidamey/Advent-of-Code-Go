package main

import (
	"advent/util"
	"fmt"
	"sort"
	"time"
)

const (
	// maxIP = 9 // example
	maxIP = 1<<32 - 1 // input
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")
	blocks := make([]*block, len(lines))
	for i, l := range lines {
		var min, max int
		fmt.Sscanf(l, "%d-%d", &min, &max)
		blocks[i] = newBlock(min, max)
	}

	sort.Slice(blocks, func(i, j int) bool {
		if blocks[i].min == blocks[j].min {
			return blocks[i].max < blocks[j].max
		}
		return blocks[i].min < blocks[j].min
	})

	for i := range blocks {
		if i > 0 {
			blocks[i].prev = blocks[i-1]
		}
		if i < len(blocks)-1 {
			blocks[i].next = blocks[i+1]
		}
	}

	p1 := -1
	p2 := 0
	b := blocks[0]
	for b.next != nil {
		if b.next.min <= b.max {
			merge(b, b.next)
			continue
		}

		unblocked := b.next.min - b.max - 1
		if p1 == -1 && unblocked > 0 {
			p1 = b.max + 1
		}
		p2 += unblocked
		b = b.next
	}

	p2 += util.MaxInt(0, blocks[0].min-1)
	p2 += maxIP - b.max

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

// combine ranges and remove the second
func merge(b1, b2 *block) {
	b1.min = util.MinInt(b1.min, b2.min)
	b1.max = util.MaxInt(b1.max, b2.max)
	b1.next = b2.next
	if b1.next != nil {
		b1.next.prev = b1
	}
}

type block struct {
	min, max   int
	prev, next *block
}

func newBlock(min, max int) *block {
	return &block{min: min, max: max}
}

func (b block) String() string {
	return fmt.Sprintf("[%d, %d]", b.min, b.max)
}
