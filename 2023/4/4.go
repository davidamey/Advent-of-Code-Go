package main

import (
	"advent-of-code-go/util"
	"fmt"
	"io"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	p1 := 0
	cards := make([]*Card, len(lines))
	for i, l := range lines {
		cards[i] = NewCard(l)
		p1 += cards[i].score()
	}

	for i, c := range cards {
		for j := c.matches(); j > 0; j-- {
			cards[i+j].multiplier += c.multiplier
		}
	}

	p2 := 0
	for _, c := range cards {
		p2 += c.multiplier
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

type Card struct {
	winning    map[int]struct{}
	actual     []int
	multiplier int
}

func NewCard(s string) *Card {
	c := &Card{winning: make(map[int]struct{}), multiplier: 1}
	r := strings.NewReader(s)
	for {
		if x, _, _ := r.ReadRune(); x == ':' {
			break
		}
	}

	for {
		var x int
		if n, _ := fmt.Fscanf(r, "%d", &x); n == 0 {
			break
		}
		c.winning[x] = struct{}{}
	}

	r.Seek(2, io.SeekCurrent)
	for {
		var x int
		if n, _ := fmt.Fscanf(r, "%d", &x); n == 0 {
			break
		}
		c.actual = append(c.actual, x)
	}

	return c
}

func (c *Card) matches() (matches int) {
	for _, x := range c.actual {
		if _, winner := c.winning[x]; winner {
			matches++
		}
	}
	return
}

func (c *Card) score() int {
	if matches := c.matches(); matches == 0 {
		return 0
	} else {
		return 1 << (matches - 1)
	}
}
