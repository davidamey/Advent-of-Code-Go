package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
	"time"
)

const maxLevel = 30

type maze struct {
	grid       *grid.Grid[rune]
	portals    map[vector.Vec]vector.Vec
	start, end vector.Vec
}

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	// lines := util.MustReadFileToLines("example2")
	lines := util.MustReadFileToLines("input")

	m := newMaze(lines)
	fmt.Println("p1=", m.findPath(false).depth)
	fmt.Println("p2=", m.findPath(true).depth)
}

type node struct {
	pos    vector.Vec
	depth  int
	level  int
	parent *node
}

type levelVec struct {
	v vector.Vec
	l int
}

func (m *maze) findPath(recursive bool) (path *node) {
	queue := []*node{&node{m.start, 0, 0, nil}}
	seen := map[levelVec]struct{}{
		levelVec{m.start, 0}: struct{}{},
	}

	depth := 0
	shortest := -1
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if n.depth > depth {
			depth = n.depth
			if shortest >= 0 && depth > shortest {
				return
			}
		}

		if n.pos == m.end && n.level == 0 {
			shortest = depth
			path = n
		}

		for _, v := range n.pos.Adjacent(false) {
			if _, ok := seen[levelVec{v, n.level}]; ok {
				continue
			}

			if !m.grid.InBounds(v) {
				continue
			}

			r := m.grid.Get(v)
			if r == '.' {
				nn := &node{v, depth + 1, n.level, n}
				queue = append(queue, nn)
				seen[levelVec{v, n.level}] = struct{}{}
				continue
			}

			if r < 'A' || r > 'Z' {
				continue
			}

			if w, exists := m.portals[n.pos]; exists {
				nn := &node{w, depth + 1, n.level, n}

				if recursive {
					if m.isOuter(n.pos) {
						nn.level--
					} else {
						nn.level++
					}
				}

				if nn.level >= 0 && nn.level <= maxLevel {
					queue = append(queue, nn)
					seen[levelVec{w, nn.level}] = struct{}{}
				}
			}
		}
	}

	return
}

func newMaze(lines []string) *maze {
	m := &maze{
		grid:    grid.New[rune](),
		portals: make(map[vector.Vec]vector.Vec),
	}

	entites := make(map[vector.Vec]rune)
	for y, l := range lines {
		for x, r := range l {
			m.grid.SetAt(x, y, r)
			if r >= 'A' && r <= 'Z' {
				entites[vector.New(x, y)] = r
			}
		}
	}

	//

	pNames := make(map[vector.Vec]string)
	for v, r := range entites {
		for _, w := range v.Adjacent(false) {
			if s, ok := entites[w]; ok {
				switch {
				case r == 'A' && s == 'A':
					m.start = m.placeEntity(v, w)
				case r == 'Z' && s == 'Z':
					m.end = m.placeEntity(v, w)
				case v.X < w.X || v.Y < w.Y:
					pNames[m.placeEntity(v, w)] = string([]rune{r, s})
				case v.X > w.X || v.Y > w.Y:
					pNames[m.placeEntity(w, v)] = string([]rune{s, r})
				}
				delete(entites, v)
				delete(entites, w)
				break
			}
		}
	}

	m.portals = make(map[vector.Vec]vector.Vec)
	for v, vN := range pNames {
		for w, wN := range pNames {
			if vN == wN && v != w {
				m.portals[v] = w
			}
		}
	}

	return m
}

func (m *maze) placeEntity(v1, v2 vector.Vec) vector.Vec {
	diff := v2.Sub(v1)
	x := v1.Sub(diff)
	if m.grid.Get(x) == '.' {
		return x
	}
	return v2.Add(diff)
}

func (m *maze) isOuter(v vector.Vec) bool {
	return v.X <= 3 || v.Y <= 3 || m.grid.Max.X-v.X <= 3 || m.grid.Max.Y-v.Y <= 3
}
