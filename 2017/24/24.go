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

	p1, p2 := solve(lines)
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func solve(lines []string) (p1, p2 int) {
	cmps := make(map[int][]*cmp)

	for _, l := range lines {
		var v1, v2 int
		fmt.Sscanf(l, "%d/%d", &v1, &v2)
		c := &cmp{v1, v2}
		addCmp(cmps, v1, c)
		addCmp(cmps, v2, c)
	}

	p := newPath(0)
	r := p.findPath(cmps)
	return r.strength, r.longestStrength
}

func addCmp(pool map[int][]*cmp, i int, c *cmp) {
	if pool[i] == nil {
		pool[i] = []*cmp{c}
	} else {
		pool[i] = append(pool[i], c)
	}
}

type cmp struct {
	v1, v2 int
}

func (c *cmp) String() string {
	return fmt.Sprintf("%d/%d", c.v1, c.v2)
}

type path struct {
	val, strength int
	seen          map[*cmp]struct{}
	path          []string
}

func newPath(val int) *path {
	return &path{val: val, strength: 0, seen: make(map[*cmp]struct{})}
}

func (p *path) String() string {
	return strings.Join(p.path, "--")
}

func (p *path) move(c *cmp) *path {
	p2 := newPath(c.v1)
	if p2.val == p.val {
		p2.val = c.v2
	}

	p2.strength = p.strength + c.v1 + c.v2

	p2.seen[c] = struct{}{}
	for k, v := range p.seen {
		p2.seen[k] = v
	}

	p2.path = append(p2.path, p.path...)
	p2.path = append(p2.path, c.String())

	return p2
}

func (p *path) findPath(pool map[int][]*cmp) (r result) {
	r.strength = p.strength
	r.longest = len(p.seen)
	r.longestStrength = p.strength

	for _, c := range pool[p.val] {
		if _, ok := p.seen[c]; ok {
			continue
		}

		p2 := p.move(c)
		r2 := p2.findPath(pool)

		if r2.strength > r.strength {
			r.strength = r2.strength
		}

		if r2.longest > r.longest || r2.longest == r.longest && r2.longestStrength > r.longestStrength {
			r.longest = r2.longest
			r.longestStrength = r2.longestStrength
		}
	}
	return
}

type result struct {
	strength, longest, longestStrength int
}
