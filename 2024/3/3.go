package main

import (
	"advent-of-code-go/util"
	"fmt"
	"regexp"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// memory := string(util.MustReadFile("example"))
	// memory := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	memory := string(util.MustReadFile("input"))

	fmt.Println("p1=", p1(memory))
	fmt.Println("p2=", p2(memory))
}

func p1(memory string) (p1 int) {
	rgx := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, m := range rgx.FindAllStringSubmatch(memory, -1) {
		p1 += util.Atoi(m[1]) * util.Atoi(m[2])
	}
	return
}

func p2(memory string) (p2 int) {
	rgx := regexp.MustCompile(`(do\(\))|(don't\(\))`)

	enabled := true
	previousId := 0
	for _, m := range rgx.FindAllStringIndex(memory, -1) {
		if enabled {
			p2 += p1(memory[previousId:m[0]])
		}
		previousId = m[1]
		enabled = m[1]-m[0] == 4 // do() = 4 chars
	}

	// Parse remaining memory
	if enabled {
		p2 += p1(memory[previousId:])
	}

	return
}
