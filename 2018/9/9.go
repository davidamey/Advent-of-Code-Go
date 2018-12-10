package main

import (
	"advent/util"
	"fmt"
)

func main() {
	// file, _ := util.OpenExample()
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	for _, l := range lines {
		var playerCount, lastMarble int
		fmt.Sscanf(l, "%d players; last marble is worth %d points", &playerCount, &lastMarble)

		fmt.Println("== part1 ==")
		play(playerCount, lastMarble)

		fmt.Println("== part2 ==")
		play(playerCount, lastMarble*100)
	}
}

func play(playerCount, lastMarble int) {
	fmt.Printf("%d players; last marble is worth %d points\n", playerCount, lastMarble)

	players := make([]int, playerCount)
	p := 0
	c := NewCircle()

	for m := 1; m <= lastMarble; m++ {
		if m%23 == 0 {
			players[p] += m
			players[p] += c.remove()
		} else {
			c.insert(m)
		}
		// fmt.Printf("[%d] ", p+1)
		// c.print()

		p = (p + 1) % playerCount
	}

	maxScore := 0
	for _, score := range players {
		if score > maxScore {
			maxScore = score
		}
	}
	fmt.Printf("max score: %d\n", maxScore)
}

type Circle struct {
	First   *CirclePoint
	Current *CirclePoint
}

type CirclePoint struct {
	Value int
	Next  *CirclePoint
	Prev  *CirclePoint
}

func NewCircle() *Circle {
	point := &CirclePoint{Value: 0}
	point.Next = point
	point.Prev = point

	return &Circle{
		First:   point,
		Current: point,
	}
}

func (c *Circle) print() {
	fmt.Printf("%d ", c.First.Value)

	for p := c.First.Next; p != c.First; p = p.Next {
		if p == c.Current {
			fmt.Printf("(%d) ", p.Value)
		} else {
			fmt.Printf("%d ", p.Value)
		}
	}
	fmt.Println()
}

func (c *Circle) insert(x int) {
	newPrev := c.Current.Next
	newNext := c.Current.Next.Next

	point := &CirclePoint{
		Value: x,
		Prev:  newPrev,
		Next:  newNext,
	}

	newPrev.Next = point
	newNext.Prev = point

	c.Current = point
}

func (c *Circle) remove() int {
	for i := 0; i < 6; i++ {
		c.Current = c.Current.Prev
	}

	toRemove := c.Current.Prev
	toRemove.Prev.Next = c.Current
	c.Current.Prev = toRemove.Prev

	return toRemove.Value
}

/** Circle from original part 1 below **/

// type Circle struct {
// 	vals    []int
// 	current int
// }

// func NewCircle() *Circle {
// 	return &Circle{
// 		vals:    []int{0},
// 		current: 0,
// 	}
// }

// func (c *Circle) print() {
// 	for i, v := range c.vals {
// 		if i == c.current {
// 			fmt.Printf("(%d) ", v)
// 		} else {
// 			fmt.Printf("%d ", v)
// 		}
// 	}
// 	fmt.Println()
// }

// func (c *Circle) insert(x int) {
// 	i := (c.current + 2) % len(c.vals)
// 	if i <= 0 {
// 		i += len(c.vals)
// 	}
// 	c.vals = append(c.vals, 0)
// 	copy(c.vals[i+1:], c.vals[i:])
// 	c.vals[i] = x
// 	c.current = i
// }

// func (c *Circle) remove() (val int) {
// 	i := (c.current - 7) % len(c.vals)
// 	if i < 0 {
// 		i += len(c.vals)
// 	}
// 	val = c.vals[i]

// 	copy(c.vals[i:], c.vals[i+1:])
// 	c.vals[len(c.vals)-1] = -1
// 	c.vals = c.vals[:len(c.vals)-1]

// 	c.current = i % len(c.vals)

// 	return
// }
