package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	entries := util.MustReadFileToInts("input")
	fmt.Println("p1=", p1(entries))
	fmt.Println("p2=", p2(entries))
}

func p1(entries []int) int {
	for i, e := range entries {
		for _, f := range entries[i+1:] {
			if e+f == 2020 {
				return e * f
			}
		}
	}
	panic("no answer found")
}

func p2(entries []int) int {
	for i, e := range entries {
		for j, f := range entries[i+1:] {
			for _, g := range entries[j+1:] {
				if e+f+g == 2020 {
					return e * f * g
				}
			}
		}
	}
	panic("no answer found")
}
