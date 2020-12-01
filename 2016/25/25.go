package main

import (
	"advent-of-code-go/2016/12/assembunny"
	"advent-of-code-go/util"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	lines := util.MustReadFileToLines("input")
	p := assembunny.Compile(lines)

	found := false
	i := 0
search:
	for ; !found; i++ {
		var r assembunny.Registers
		r[0] = i

		p.RunWithOut(&r, false, 10)
		for j, o := range p.Out {
			if j%2 != o {
				continue search
			}
		}

		found = true
	}

	fmt.Println("p1=", i-1)
}
