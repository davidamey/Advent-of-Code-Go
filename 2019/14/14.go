package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
)

func main() {
	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	fmt.Println("p1=", p1(lines))
	fmt.Println("p2=", p2(lines))
}

func p1(lines []string) int {
	nf := newNanoFactory(lines)
	nf.ensureQuantity("FUEL", 1)
	return nf.usedOre
}

func p2(lines []string) int {
	nf := newNanoFactory(lines)

	ore := 1000000000000
	start, end := 0, ore
	for i := 0; i < 1000; i++ { // 1k should be enough
		guess := start + (end-start)/2
		nf.reset()
		nf.ensureQuantity("FUEL", guess)

		if nf.usedOre > ore {
			end = guess
		} else if nf.usedOre < ore {
			start = guess
		} else {
			break
		}
	}
	return nf.store["FUEL"]
}

type nanofactory struct {
	recipes map[chemical]recipe
	store   map[chemical]int
	usedOre int
}

func newNanoFactory(reactions []string) (nf *nanofactory) {
	nf = &nanofactory{
		store:   make(map[chemical]int),
		recipes: make(map[chemical]recipe),
	}
	for _, l := range reactions {
		c, r := newRecipe(l)
		nf.recipes[c] = r
	}
	return
}

func (nf *nanofactory) reset() {
	nf.store = make(map[chemical]int)
	nf.usedOre = 0
}

func (nf *nanofactory) ensureQuantity(c chemical, quantity int) {
	// ORE is free...
	if c == "ORE" {
		nf.store[c] += quantity
		nf.usedOre += quantity // ...but we want to keep track of how much
		return
	}

	// have enough chemical already
	if nf.store[c] >= quantity {
		return
	}

	// need to make some more
	r := nf.recipes[c]
	toMake := quantity - nf.store[c]
	m := toMake / r.quantity
	if toMake%r.quantity > 0 {
		m++
	}
	for in, q := range r.ingredients {
		nf.ensureQuantity(in, q*m)
		nf.store[in] -= q * m
	}

	// all ingredients used so we have made something
	nf.store[c] += m * r.quantity
}

type chemical string
type recipe struct {
	quantity    int
	ingredients map[chemical]int
}

func newRecipe(s string) (c chemical, r recipe) {
	r.ingredients = make(map[chemical]int)
	parts := strings.Split(s, " => ")
	c, r.quantity = parseItem(parts[1])
	for _, in := range strings.Split(parts[0], ", ") {
		n, q := parseItem(in)
		r.ingredients[n] = q
	}
	return
}

func parseItem(s string) (name chemical, quantity int) {
	fmt.Sscanf(s, "%d %s", &quantity, &name)
	return
}
