package main

import (
	"advent/util"
	"fmt"
	"strconv"
)

func main() {
	// file, _ := util.OpenExample()
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	p1 := 0
	p2 := 0
	for _, l := range lines {
		s1, _ := strconv.Unquote(l)
		s2 := strconv.Quote(l)

		p1 += len(l)
		p1 -= len(s1)

		p2 += len(s2)
		p2 -= len(l)
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}
