package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

// var input = []int{0, 3, 6} // example
var input = []int{9, 6, 0, 10, 18, 2, 1} // actual

func main() {
	defer util.Duration(time.Now())

	turnSpoken := make(map[int]int) // number spoken to turn spoken
	lastSpoken := 0
	var p1 int
	for i := 0; i < 30000000; i++ {
		if i < len(input) {
			turnSpoken[lastSpoken] = i
			lastSpoken = input[i]
			continue
		}

		turnBefore, heardBefore := turnSpoken[lastSpoken]
		turnSpoken[lastSpoken] = i

		if heardBefore {
			lastSpoken = i - turnBefore
		} else {
			lastSpoken = 0
		}

		if i+1 == 2020 {
			p1 = lastSpoken
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", lastSpoken)
}
