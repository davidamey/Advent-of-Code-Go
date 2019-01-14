package main

import "fmt"

const input = 36000000

func main() {
	p1()
	p2()
}

func p1() {
	houses := make([]int, 1+input/10)
	for e := 1; e <= input/10; e++ { // for each elf
		for h := e; h <= input/10; h += e { // for each house for said elf
			houses[h] += e * 10
		}
	}

	for h, s := range houses {
		if s > input {
			fmt.Println("p1=", h, s)
			break
		}
	}
}

func p2() {
	houses := make([]int, 1+input/10)
	for e := 1; e <= input/10; e++ { // for each elf
		deliveredCount := 0
		for h := e; h <= input/10; h += e { // for each house for said elf
			houses[h] += e * 11
			deliveredCount++
			if deliveredCount == 50 {
				break
			}
		}
	}

	for h, s := range houses {
		if s > input {
			fmt.Println("p2=", h, s)
			break
		}
	}
}
