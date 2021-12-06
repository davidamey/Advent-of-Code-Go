package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// fish := util.MustReadCSInts("example")
	fish := util.MustReadCSInts("input")

	fmt.Println("p1=", simulate(fish, 80))
	fmt.Println("p2=", simulate(fish, 256))
}

func simulate(startFish []int, days int) (endFish int) {
	var counts [9]int
	for _, f := range startFish {
		counts[f]++
	}

	for d := 0; d < days; d++ {
		spawn := counts[0]
		for i := 0; i < 8; i++ {
			counts[i] = counts[i+1]
		}
		counts[6] += spawn
		counts[8] = spawn
	}

	for _, v := range counts {
		endFish += v
	}

	return
}
