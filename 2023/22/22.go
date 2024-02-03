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

	bricks := make([]*brick, len(lines))
	for i, l := range lines {
		b := brick{id: i + 1}
		fmt.Sscanf(l, "%d,%d,%d~%d,%d,%d",
			&b.start[0], &b.start[1], &b.start[2],
			&b.end[0], &b.end[1], &b.end[2],
		)
		bricks[i] = &b
	}

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].start[2] < bricks[j].start[2]
	})

	layers := make([][]*brick, 1)

	for _, b := range bricks {
		landed := false
		z := 1
		for l := len(layers) - 1; !landed && l >= 0; l-- {
			for _, b2 := range layers[l] {
				if b.intersects(b2) {
					landed = true
					z = util.MaxInt(z, b2.end[2]+1)
					b.supportedBy = append(b.supportedBy, b2)
					b2.supports = append(b2.supports, b)
				}
			}
		}

		h := b.end[2] - b.start[2]
		b.start[2], b.end[2] = z, z+h
		for b.end[2] >= len(layers) {
			layers = append(layers, []*brick{})
		}
		layers[b.end[2]] = append(layers[b.end[2]], b)
	}

	p1 := 0
	for _, b := range bricks {
		canDisintegrate := true
		for _, s := range b.supports {
			if len(s.supportedBy) == 1 {
				canDisintegrate = false
			}
		}
		if canDisintegrate {
			p1++
		}
	}
	fmt.Println("p1=", p1)

	p2 := 0
	for _, b := range bricks {
		p2 += b.topple()
	}
	fmt.Println("p2=", p2)
}

type vec3 [3]int
type brick struct {
	id          int
	start, end  vec3
	supports    []*brick
	supportedBy []*brick
}

func (b1 *brick) intersects(b2 *brick) bool {
	if b1.end[0] < b2.start[0] || b2.end[0] < b1.start[0] {
		return false
	}
	if b1.end[1] < b2.start[1] || b2.end[1] < b1.start[1] {
		return false
	}
	return true
}

func (b *brick) topple() int {
	falling := map[*brick]bool{b: true}
	queue := append([]*brick{}, b.supports...)

	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]

		if falling[x] {
			continue
		}

		supported := false
		for _, base := range x.supportedBy {
			if !falling[base] {
				supported = true
				break
			}
		}

		if !supported {
			falling[x] = true
			queue = append(queue, x.supports...)
		}
	}

	return len(falling) - 1 // Don't count original "disintegrated" brick
}

func (b *brick) String() string {
	// return strconv.Itoa(b.id)
	return fmt.Sprintf("%d,%d,%d~%d,%d,%d",
		b.start[0], b.start[1], b.start[2],
		b.end[0], b.end[1], b.end[2],
	)
}
