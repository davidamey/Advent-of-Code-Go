package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// input := "target area: x=20..30, y=-10..-5"
	input := "target area: x=155..182, y=-117..-67"

	var target [4]int
	fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d",
		&target[0],
		&target[1],
		&target[2],
		&target[3],
	)

	// What goes up...
	p1 := target[2] * (target[2] + 1) / 2
	fmt.Println("p1=", p1)

	p2 := 0
	for vy := target[2]; vy <= p1; vy++ {
		for vx := -target[1]; vx <= target[1]; vx++ {
			if shoot(target, vx, vy) {
				p2++
			}
		}
	}
	fmt.Println("p2=", p2)
}

func shoot(target [4]int, vx, vy int) bool {
	x, y := 0, 0
	for {
		x, y, vy = x+vx, y+vy, vy-1
		if vx < 0 {
			vx++
		} else if vx > 0 {
			vx--
		}

		if contains(target, x, y) {
			return true
		}

		if x < target[0] && vx <= 0 {
			break
		}

		if x > target[1] && vx >= 0 {
			break
		}

		if y < target[2] && vy <= 0 {
			break
		}
	}
	return false
}

func contains(area [4]int, x, y int) bool {
	return x >= area[0] && x <= area[1] && y >= area[2] && y <= area[3]
}
