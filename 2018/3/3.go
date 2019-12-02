package main

import (
	"advent-of-code-go/util"
	"fmt"
)

const fabricSize = 2000

type fabric [fabricSize][fabricSize]int

type claim struct {
	ID, X, Y, W, H int
}

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	inputs, _ := util.ReadLines(file)

	claims := make([]*claim, len(inputs))
	for i, s := range inputs {
		claims[i] = strToClaim(s)
	}

	f := fabric{}
	for _, c := range claims {
		addClaimToFabric(&f, c)
	}

	part1(&f, claims)
	part2(&f, claims)
}

func part1(f *fabric, claims []*claim) {
	overlap := 0
	for y := range f {
		for x := range f[y] {
			if f[y][x] > 1 {
				overlap++
			}
		}
	}

	fmt.Printf("overlapping square inches = %d\n", overlap)
}

func part2(f *fabric, claims []*claim) {
	var c *claim
	for i := range claims {
		if !claimIsOverlapped(f, claims[i]) {
			c = claims[i]
			break
		}
	}

	if c == nil {
		fmt.Println("all claims overlapped :'(")
		return
	}

	fmt.Printf("claim %d not overlapped\n", c.ID)
}

func claimIsOverlapped(f *fabric, c *claim) bool {
	x1 := c.X
	x2 := c.X + c.W
	y1 := c.Y
	y2 := c.Y + c.H

	if x2 >= fabricSize || y2 >= fabricSize {
		fmt.Println("out of bounds", c.ID)
	}

	for y := y1; y < y2; y++ {
		for x := x1; x < x2; x++ {
			if f[y][x] > 1 {
				return true
			}
		}
	}
	return false
}

func addClaimToFabric(f *fabric, c *claim) {
	x1 := c.X
	x2 := c.X + c.W
	y1 := c.Y
	y2 := c.Y + c.H

	if x2 >= fabricSize || y2 >= fabricSize {
		fmt.Println("out of bounds", c.ID)
	}

	for y := y1; y < y2; y++ {
		for x := x1; x < x2; x++ {
			f[y][x]++
		}
	}
}

func strToClaim(s string) *claim {
	c := &claim{}
	fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &c.ID, &c.X, &c.Y, &c.W, &c.H)
	return c
}
