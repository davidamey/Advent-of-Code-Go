package grid

import vector "advent-of-code-go/util/vector"

type PathNode struct {
	Pos    vector.Vec
	Length int
	Parent *PathNode
}

type PathValidator func(v, parent interface{}, depth int) bool

func (g *Grid) ShortestPath(start vector.Vec, end vector.Vec, valid PathValidator) (path *PathNode) {
	queue := []*PathNode{{start, 0, nil}}
	seen := make(map[vector.Vec]struct{})

	depth := 0
	shortest := -1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Length > depth {
			// Reached a new depth of search
			depth = node.Length
			if shortest >= 0 && depth > shortest {
				// We've found an answer and will forever be deeper than it now so return
				return
			}
		}

		// Have we reached the end?
		if node.Pos == end {
			shortest = depth
			path = node
		}

		// Add all unseen children to the queue
		parent := g.Entry(node.Pos)
		for _, c := range node.Pos.Adjacent(false) {
			if _, ok := seen[c]; ok {
				continue
			}

			if !g.InBounds(c) {
				continue
			}

			if valid(g.Entry(c), parent, depth) {
				n := &PathNode{c, depth + 1, node}
				queue = append(queue, n)
				seen[c] = struct{}{}
			}
		}
	}

	return
}
