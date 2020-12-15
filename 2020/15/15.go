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
	for i := 1; i <= 30000000; i++ {
		if i <= len(input) {
			turnSpoken[lastSpoken] = i - 1
			lastSpoken = input[i-1]
			continue
		}

		turnBefore, heardBefore := turnSpoken[lastSpoken]
		turnSpoken[lastSpoken] = i - 1

		if heardBefore {
			lastSpoken = i - 1 - turnBefore
		} else {
			lastSpoken = 0
		}

		if i == 2020 {
			p1 = lastSpoken
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p1=", lastSpoken)
}
