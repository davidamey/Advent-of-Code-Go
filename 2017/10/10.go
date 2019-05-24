package main

import (
	"advent/2017/10/knothash"
	"advent/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// examples
	// fmt.Println("ex p1=", p1(5, "3, 4, 1, 5"))
	// fmt.Println("ex KnotHash=", KnotHash([]byte("")))
	// fmt.Println("ex KnotHash=", KnotHash([]byte("AoC 2017")))
	// fmt.Println("ex KnotHash=", KnotHash([]byte("1,2,3")))
	// fmt.Println("ex KnotHash=", KnotHash([]byte("1,2,4")))

	// actual
	input := util.MustReadFile("input")
	fmt.Println("p1=", p1(256, string(input)))
	fmt.Println("p2=", knothash.Compute(input))
}

func p1(listSize int, input string) int {
	var lengths []int
	for _, s := range strings.Split(input, ", ") {
		l, _ := strconv.Atoi(s)
		lengths = append(lengths, l)
	}

	l := make(knothash.List, listSize)
	for i := range l {
		l[i] = i
	}

	pos := 0
	skip := 0

	for _, ln := range lengths {
		l.Reverse(pos, pos+ln-1)
		pos += ln + skip
		skip++
	}

	return l[0] * l[1]
}
