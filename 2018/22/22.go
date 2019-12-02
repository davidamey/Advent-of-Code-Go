package main

import (
	"advent-of-code-go/util"
	"fmt"
	"math"
)

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	var depth, tx, ty int
	fmt.Sscanf(lines[0], "depth: %d", &depth)
	fmt.Sscanf(lines[1], "target: %d,%d", &tx, &ty)

	g := NewGrid(depth, util.Vec{X: tx, Y: ty})
	sum := 0
	// Need the +10k to allow the A* algorithm for p2 to work
	for y := 0; y <= ty+10000; y++ {
		for x := 0; x <= tx+10000; x++ {
			var val int
			switch {
			case x == 0 && y == 0 || x == tx && y == ty:
				val = g.ToErrosion(0)

			case y == 0:
				val = g.ToErrosion(x * 16807)
			case x == 0:
				val = g.ToErrosion(y * 48271)
			default:
				a := g.GetAt(x-1, y)
				b := g.GetAt(x, y-1)
				val = g.ToErrosion(a * b)
			}
			g.SetAt(x, y, val)

			// part 1
			if x <= tx && y <= ty {
				sum += val % 3
			}
		}
	}

	fmt.Println("part1=", sum)

	g.FindPath(NewNode(0, 0, 'T'), NewNode(tx, ty, 'T'))
}

func (n *Node) ToLoc() string {
	return fmt.Sprintf("%d,%d,%c", n.Pos.X, n.Pos.Y, n.Tool)
}

func ToolSwitch(from, to rune) int {
	if from == to {
		return 0
	}
	return 7
}

type Node struct {
	Pos  util.Vec
	Tool rune
}

func NewNode(x, y int, t rune) Node {
	return Node{
		util.Vec{X: x, Y: y},
		t,
	}
}

type Grid struct {
	MinX, MinY int
	MaxX, MaxY int
	Depth      int
	Target     util.Vec
	entries    map[util.Vec]int
}

func NewGrid(depth int, target util.Vec) *Grid {
	g := &Grid{
		MinX:    math.MaxInt32,
		MinY:    math.MaxInt32,
		MaxX:    math.MinInt32,
		MaxY:    math.MinInt32,
		Depth:   depth,
		Target:  target,
		entries: make(map[util.Vec]int),
	}
	return g
}

func Cost(from, to Node) int {
	cost := from.Pos.ManhattanTo(to.Pos)
	if from.Tool != to.Tool {
		cost += 7
	}
	return cost
}

func GetScore(scores map[Node]int, key Node) int {
	if score, ok := scores[key]; ok {
		return score
	}
	return math.MaxInt32
}

const (
	rocky  = 0
	wet    = 1
	narrow = 2
)

// A* algorithm
func (g *Grid) FindPath(from, to Node) {
	closed := make(map[Node]bool)
	open := make(map[Node]bool)
	open[from] = true

	cameFrom := make(map[Node]Node)

	gScores := make(map[Node]int)
	gScores[from] = 0

	fScores := make(map[Node]int)
	fScores[from] = Cost(from, to)

	// for ; len(open) > 0; fmt.Println(len(open)) {
	for len(open) > 0 {
		var current Node
		minFS := math.MaxInt32
		for n := range open {
			fs := GetScore(fScores, n)
			if fs < minFS {
				current = n
				minFS = fs
			}
		}

		// fmt.Println("looking at", current)

		if current == to {
			// TracePath(cameFrom, current)
			break
		}

		delete(open, current)
		closed[current] = true

		currentRegion := g.Get(current.Pos) % 3

		var possibleNodes []Node
		for _, p := range current.Pos.Adjacent(false) {
			if p.X < 0 || p.Y < 0 {
				continue
			}

			targetRegion := g.Get(p) % 3

			if ICanUse(currentRegion, targetRegion, 'C') {
				possibleNodes = append(possibleNodes, Node{p, 'C'})
			}
			if ICanUse(currentRegion, targetRegion, 'T') {
				possibleNodes = append(possibleNodes, Node{p, 'T'})
			}
			if ICanUse(currentRegion, targetRegion, 'N') {
				possibleNodes = append(possibleNodes, Node{p, 'N'})
			}
		}

		for _, pn := range possibleNodes {
			if _, ok := closed[pn]; ok {
				continue
			}

			tentativeGScore := GetScore(gScores, current) + Cost(current, pn)

			if _, ok := open[pn]; !ok {
				open[pn] = true
			} else if tentativeGScore >= GetScore(gScores, pn) {
				continue
			}

			// This is the best path so far.
			cameFrom[pn] = current
			gScores[pn] = tentativeGScore
			fScores[pn] = tentativeGScore + Cost(pn, to)
		}
	}

	fmt.Println("part2=", fScores[to])
}

func ICanUse(regA, regB int, tool rune) bool {
	switch tool {
	case 'T':
		return regA != wet && regB != wet
	case 'C':
		return regA != narrow && regB != narrow
	case 'N':
		return regA != rocky && regB != rocky
	}
	panic("unknown tool")
}

func TracePath(cameFrom map[Node]Node, current Node) {
	// fmt.Println("tracing path")
	path := []Node{current}
	for {
		// fmt.Printf("=> (%d,%d,%c)\n", current.Pos.X, current.Pos.Y, current.Tool)
		var ok bool
		if current, ok = cameFrom[current]; ok {
			path = append(path, current)
		} else {
			break
		}
	}

	for i := len(path) - 1; i >= 0; i-- {
		fmt.Printf("(%d,%d,%c)\n", path[i].Pos.X, path[i].Pos.Y, path[i].Tool)
	}
}

func (g *Grid) ToErrosion(gi int) int {
	return (gi + g.Depth) % 20183
}

func (g *Grid) ResizeFor(p util.Vec) {
	if p.X < g.MinX {
		g.MinX = p.X
	}
	if p.Y < g.MinY {
		g.MinY = p.Y
	}
	if p.X > g.MaxX {
		g.MaxX = p.X
	}
	if p.Y > g.MaxY {
		g.MaxY = p.Y
	}
}

func (g *Grid) Set(p util.Vec, i int) {
	g.entries[p] = i
	g.ResizeFor(p)
}

func (g *Grid) Get(p util.Vec) int {
	return g.entries[p]
}

func (g *Grid) SetAt(x, y, i int) {
	g.Set(util.Vec{X: x, Y: y}, i)
}

func (g *Grid) GetAt(x, y int) int {
	return g.entries[util.Vec{X: x, Y: y}]
}

func (g *Grid) ForEach(fn func(v, x, y int)) {
	for y := g.MinY; y <= g.MaxY; y++ {
		for x := g.MinX; x <= g.MaxX; x++ {
			fn(g.GetAt(x, y), x, y)
		}
	}
}

func (g *Grid) Print(clear bool) {
	if clear {
		fmt.Printf("\033[0;0H")
		fmt.Printf("\033[2J")
	}

	// spew.Dump(g)

	for y := g.MinY; y <= g.MaxY; y++ {
		for x := g.MinX; x <= g.MaxX; x++ {
			r := '?'
			switch g.GetAt(x, y) % 3 {
			case 0: // rocky
				r = '.'
			case 1: // wet
				r = '='
			case 2: // narrow
				r = '|'
			}

			if x == 0 && y == 0 {
				r = 'M'
			} else if x == g.Target.X && y == g.Target.Y {
				r = 'T'
			}
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}
