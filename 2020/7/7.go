package main

import (
	"advent-of-code-go/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const theBag = "shiny gold"

var rgxBag = regexp.MustCompile(`(\d+) (.+?) bags?[,.]`)
var bags map[colour]bag

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	bags = make(map[colour]bag)
	for _, l := range lines {
		parts := strings.Split(l, " bags contain ")
		bags[colour(parts[0])] = newBag(parts[1])
	}

	p1 := 0
	for _, b := range bags {
		if b.canContainTheBag() {
			p1++
		}
	}

	p2 := bags[theBag].size()
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

type colour string
type bag map[colour]int

func newBag(contents string) bag {
	b := make(bag)
	for _, c := range rgxBag.FindAllStringSubmatch(contents, -1) {
		b[colour(c[2])], _ = strconv.Atoi(c[1])
	}
	return b
}

func (b bag) canContainTheBag() bool {
	// contains directly
	if _, ok := b[theBag]; ok {
		return true
	}

	// contains indirectly
	for c := range b {
		if bags[c].canContainTheBag() {
			return true
		}
	}

	return false
}

func (b bag) size() (count int) {
	for colour, amount := range b {
		count += amount * (1 + bags[colour].size())
	}
	return
}
