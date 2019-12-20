package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())
	fmt.Println("p1=", p1())
	fmt.Println("p2=", p2())
}

func p1() int {
	lines := util.MustReadFileToLines("inputP1")
	var starts []vector.Vec
	var keys []rune
	g := grid.New()
	for y, l := range lines {
		for x, r := range l {
			g.SetAt(x, y, rune(r))

			switch {
			case r == '@':
				starts = append(starts, vector.New(x, y))
			case r >= 'a' && r <= 'z':
				keys = append(keys, r)
			}
		}
	}

	p := &pather{g, make(map[string]int)}
	return p.findPath(starts, make(keymap))
}

func p2() int {
	lines := util.MustReadFileToLines("inputP2")
	var starts []vector.Vec
	var keys []rune
	g := grid.New()
	for y, l := range lines {
		for x, r := range l {
			g.SetAt(x, y, rune(r))

			switch {
			case r == '@':
				starts = append(starts, vector.New(x, y))
			case r >= 'a' && r <= 'z':
				keys = append(keys, r)
			}
		}
	}

	p := &pather{g, make(map[string]int)}
	return p.findPath(starts, make(keymap))
}

type pather struct {
	grid *grid.Grid
	seen map[string]int
}

func (p *pather) findPath(from []vector.Vec, ownedKeys keymap) (pathLength int) {
	h := hash(from, ownedKeys)
	if l, ok := p.seen[h]; ok {
		return l
	}

	foundKeys := p.reachableKeys(from, ownedKeys)
	if len(foundKeys) == 0 {
		return 0
	}

	min := 1 << 16
	for k, n := range foundKeys {
		var next []vector.Vec
		for i, f := range from {
			if i == n.robot {
				next = append(next, n.pos)
			} else {
				next = append(next, f)
			}
		}

		ks := ownedKeys.clone()
		ks[k] = true
		if d := n.depth + p.findPath(next, ks); d < min {
			min = d
		}
	}
	p.seen[h] = min
	return min
}

func hash(from []vector.Vec, ownedKeys keymap) string {
	var keys []byte
	for k := range ownedKeys {
		keys = append(keys, byte(k))
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	var sb strings.Builder
	for _, f := range from {
		sb.WriteString(f.String())
	}
	sb.WriteString(string(keys))
	return sb.String()
}

type keymap map[rune]bool

func (k keymap) clone() keymap {
	k2 := make(keymap)
	for k, v := range k {
		k2[k] = v
	}
	return k2
}

type node struct {
	pos   vector.Vec
	depth int
	robot int
}

func (p *pather) reachableKeys(from []vector.Vec, ownedKeys keymap) (foundKeys map[rune]node) {
	foundKeys = make(map[rune]node)
	for robot, f := range from {
		toSearch := append([]vector.Vec{}, f)
		distance := map[vector.Vec]int{f: 0}
		for len(toSearch) > 0 {
			var v vector.Vec
			v, toSearch = toSearch[0], toSearch[1:]

			for _, w := range v.Adjacent(false) {
				// off the grid
				if !p.grid.InBounds(w) {
					continue
				}

				// already processed
				if _, ok := distance[w]; ok {
					continue
				}

				r := p.grid.Rune(w)

				// wall
				if r == '#' {
					continue
				}

				// locked door
				if r >= 'A' && r <= 'Z' && !ownedKeys[r-'A'+'a'] {
					continue
				}

				// found a valid place to go to so add the distance
				distance[w] = distance[v] + 1

				// found a key
				if r >= 'a' && r <= 'z' && !ownedKeys[r] {
					foundKeys[r] = node{w, distance[w], robot}
				} else { // todo: think this can go
					// a new place to stand
					toSearch = append(toSearch, w)
				}
			}
		}
	}
	return
}
