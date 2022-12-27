package main

import (
	"advent-of-code-go/util"
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	var packets []*packet
	for _, l := range lines {
		if l == "" {
			continue
		}
		packets = append(packets, newPacket(l))
	}

	p1 := 0
	for i := 0; i < len(packets)-1; i += 2 {
		if compare(packets[i].parsed, packets[i+1].parsed) == -1 {
			p1 += i/2 + 1
		}
	}

	fmt.Println("p1=", p1)

	divPacket1 := newPacket("[[2]]")
	divPacket2 := newPacket("[[6]]")
	packets = append(packets, divPacket1, divPacket2)
	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i].parsed, packets[j].parsed) == -1
	})

	p2 := 1
	for i, p := range packets {
		if p == divPacket1 || p == divPacket2 {
			p2 *= (i + 1)
		}
	}
	fmt.Println("p2=", p2)
}

func compare(l, r interface{}) int {
	x, lIsNum := l.(float64)
	y, rIsNum := r.(float64)

	if lIsNum && rIsNum {
		return diffInt(x, y)
	}

	xs, _ := l.([]interface{})
	ys, _ := r.([]interface{})

	if lIsNum {
		return compare([]interface{}{x}, ys)
	}

	if rIsNum {
		return compare(xs, []interface{}{y})
	}

	for i := range xs {
		if i >= len(ys) {
			return 1
		}
		if r := compare(xs[i], ys[i]); r != 0 {
			return r
		}
	}

	return diffInt(len(xs), len(ys))
}

type packet struct {
	parsed interface{}
	str    string
}

func newPacket(s string) *packet {
	p := &packet{str: s}
	json.Unmarshal([]byte(s), &p.parsed)
	return p
}

func diffInt[T int | float64](x, y T) int {
	switch {
	case x < y:
		return -1
	case x > y:
		return 1
	default:
		return 0
	}
}
