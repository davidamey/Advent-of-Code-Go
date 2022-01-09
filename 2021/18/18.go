package main

import (
	"advent-of-code-go/util"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	fish := make([]*snailfish, len(lines))
	for i, l := range lines {
		fish[i] = newSnailFish(l)
	}

	sum := fish[0]
	for _, f := range fish[1:] {
		sum = add(sum, f)
	}
	fmt.Println("p1=", sum.magnitude())

	max := 0
	for _, f1 := range fish {
		for _, f2 := range fish {
			if f1 == f2 {
				continue
			}

			sum := add(f1, f2).magnitude()
			if sum > max {
				max = sum
			}
		}
	}
	fmt.Println("p2=", max)
}

func add(a, b *snailfish) *snailfish {
	sf := &snailfish{l: a.clone(), r: b.clone()}
	sf.l.parent = sf
	sf.r.parent = sf
	sf.reduce()
	return sf
}

type snailfish struct {
	l, r, parent *snailfish
	value        int
}

func newSnailFish(s string) *snailfish {
	var parse func(*bufio.Reader, byte) *snailfish
	parse = func(r *bufio.Reader, delim byte) *snailfish {
		if b, _ := r.Peek(1); b[0] == '[' {
			r.Discard(1)
			sf := &snailfish{
				l: parse(r, ','),
				r: parse(r, ']'),
			}
			sf.l.parent = sf
			sf.r.parent = sf
			r.Discard(1) // trailing , or ]
			return sf
		}

		data, _ := r.ReadBytes(delim)
		return &snailfish{value: util.Atoi(string(data[:len(data)-1]))}
	}

	return parse(bufio.NewReader(strings.NewReader(s)), '?')
}

func (sf *snailfish) String() string {
	if sf.l == nil && sf.r == nil {
		return strconv.Itoa(sf.value)
	}
	return fmt.Sprintf("[%s,%s]", sf.l, sf.r)
}

func (sf *snailfish) magnitude() int {
	if sf.l == nil && sf.r == nil {
		return sf.value
	}
	return 3*sf.l.magnitude() + 2*sf.r.magnitude()
}

func (sf *snailfish) clone() *snailfish {
	if sf.l == nil && sf.r == nil {
		return &snailfish{value: sf.value}
	}
	x := &snailfish{l: sf.l.clone(), r: sf.r.clone()}
	x.l.parent, x.r.parent = x, x
	return x
}

func (sf *snailfish) explode() bool {
	var search func(*snailfish, int) *snailfish
	search = func(sf *snailfish, depth int) *snailfish {
		if sf.l == nil && sf.r == nil {
			return nil
		}

		if depth == 4 {
			return sf
		}

		if found := search(sf.l, depth+1); found != nil {
			return found
		}

		return search(sf.r, depth+1)
	}

	toExplode := search(sf, 0)
	if toExplode == nil {
		return false
	}

	lt := toExplode
	for lt != nil {
		if p := lt.parent; p != nil && p.r == lt {
			lt = p.l
			break
		} else {
			lt = p
		}
	}
	for lt != nil && lt.r != nil {
		lt = lt.r
	}

	rt := toExplode
	for rt != nil {
		if p := rt.parent; p != nil && p.l == rt {
			rt = p.r
			break
		} else {
			rt = p
		}
	}
	for rt != nil && rt.l != nil {
		rt = rt.l
	}

	if lt != nil {
		lt.value += toExplode.l.value
	}
	if rt != nil {
		rt.value += toExplode.r.value
	}
	toExplode.l, toExplode.r, toExplode.value = nil, nil, 0

	return true
}

func (sf *snailfish) split() bool {
	var search func(*snailfish) *snailfish
	search = func(sf *snailfish) *snailfish {
		if sf.l == nil && sf.r == nil {
			if sf.value >= 10 {
				return sf
			}
			return nil
		}

		if found := search(sf.l); found != nil {
			return found
		}

		return search(sf.r)
	}

	toSplit := search(sf)
	if toSplit == nil {
		return false
	}

	toSplit.l = &snailfish{parent: toSplit, value: toSplit.value / 2}
	toSplit.r = &snailfish{parent: toSplit, value: (toSplit.value + 1) / 2}
	return true
}

func (sf *snailfish) reduce() {
	for {
		if sf.explode() {
			continue
		}
		if sf.split() {
			continue
		}
		break
	}
}
