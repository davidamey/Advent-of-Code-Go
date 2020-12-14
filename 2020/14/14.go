package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	memP1, memP2 := make(map[int]int), make(map[int]int)
	var mask string
	for _, l := range lines {
		if strings.HasPrefix(l, "mask") {
			mask = l[7:]
			continue
		}

		var i, v int
		fmt.Sscanf(l, "mem[%d] = %d", &i, &v)

		memP1[i] = applyMask(mask, v)
		for _, a := range computeAddresses(mask, i) {
			memP2[a] = v
		}
	}

	p1 := 0
	for _, v := range memP1 {
		p1 += v
	}

	p2 := 0
	for _, v := range memP2 {
		p2 += v
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func applyMask(mask string, in int) (out int) {
	out = in
	for i, m := range mask {
		j := len(mask) - 1 - i
		switch m {
		case '0':
			out &^= 1 << j
		case '1':
			out |= 1 << j
		}
	}
	return
}

func computeAddresses(mask string, base int) (addresses []int) {
	// handle 1s first
	for i, m := range mask {
		if m == '1' {
			base |= 1 << (len(mask) - 1 - i)
		}
	}

	addresses = append(addresses, base)
	for i, m := range mask {
		if m != 'X' {
			continue
		}

		j := len(mask) - 1 - i
		// for each address we already have, add its 'X-mirror'
		for _, a := range addresses {
			switch (a >> j) & 1 {
			case 0:
				addresses = append(addresses, a|(1<<j))
			case 1:
				addresses = append(addresses, a&^(1<<j))
			}
		}
	}

	return
}
