package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// notes := util.MustReadFileToLines("example")
	notes := util.MustReadFileToLines("input")

	fmt.Println("p1=", part1(notes))
	fmt.Println("p2=", part2(notes))
}

func part1(notes []string) int {
	earliest := util.Atoi(notes[0])

	bestBus, bestTime := 0, 1<<32
	for _, b := range strings.Split(notes[1], ",") {
		if b == "x" {
			continue
		}

		bus := util.Atoi(b)
		t := bus * ((earliest + bus - 1) / bus)
		if t < bestTime {
			bestBus, bestTime = bus, t
		}
	}
	return bestBus * (bestTime - earliest)
}

func part2(notes []string) int {
	var remainders, mods []int
	product := 1
	for i, b := range strings.Split(notes[1], ",") {
		if b == "x" {
			continue
		}

		bus := util.Atoi(b)
		remainders = append(remainders, util.Mod(-i, bus))
		mods = append(mods, bus)
		product *= bus
	}

	// Gauss's algorithm
	sum := 0
	for i, m := range mods {
		p := product / m
		sum += remainders[i] * p * util.ModInverse(p, m)
	}
	return sum % product
}
