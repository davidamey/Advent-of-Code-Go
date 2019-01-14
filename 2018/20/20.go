package main

import (
	"advent/util"
	"fmt"
	"io/ioutil"
	"math"
)

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	steps, _ := ioutil.ReadAll(file)

	examples := [][]byte{
		[]byte("^E(N|S)E$"), // doesn't work
		[]byte("^WNE$"),
		[]byte("^ENWWW(NEEE|SSE(EE|N))$"),
		[]byte("^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$"),
		[]byte("^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$"),
	}

	fmt.Println("== examples ==")
	for _, e := range examples {
		fmt.Println(string(e))
		runSteps(e, true)
	}
	fmt.Println()

	fmt.Println("== part1&2 ==")
	runSteps(steps, false)
}

// After writing I realise this approach doesn't handle
// branches that finish in different places, like ^(N/S)E$ above
// Luckily, that doesn't matter for the given examples/input.
func runSteps(steps []byte, drawGrid bool) {
	g := NewGrid()
	stack := []Node{*g.current}
	for _, s := range steps[1 : len(steps)-1] {
		switch s {
		case 'N':
			g.MoveUp()
		case 'E':
			g.MoveRight()
		case 'S':
			g.MoveDown()
		case 'W':
			g.MoveLeft()
		case '(':
			stack = append(stack, *g.current)
		case ')':
			g.current = &stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case '|':
			g.current = &stack[len(stack)-1]
		}
	}

	if drawGrid {
		g.Print(false)
	}

	longestPath := 0
	atLeast1kDoors := 0
	for _, n := range g.nodes {
		if n.D > longestPath {
			longestPath = n.D
		}
		if n.D >= 1000 {
			atLeast1kDoors++
		}
	}

	fmt.Println("  longest path=", longestPath)
	fmt.Println("  rooms over 1k doors away=", atLeast1kDoors)
	fmt.Println()
}

type Node struct {
	X, Y, D int
}

type Grid struct {
	current    *Node
	minX, minY int
	maxX, maxY int
	nodes      map[string]*Node
	doors      map[string]rune
}

func NewGrid() *Grid {
	g := &Grid{
		minX:  math.MaxInt32,
		minY:  math.MaxInt32,
		maxX:  math.MinInt32,
		maxY:  math.MinInt32,
		nodes: make(map[string]*Node),
		doors: make(map[string]rune),
	}
	g.current = g.UpdateNode(0, 0, 0)
	return g
}

func (g *Grid) ResizeFor(x, y int) {
	if x < g.minX {
		g.minX = x
	}
	if y < g.minY {
		g.minY = y
	}
	if x > g.maxX {
		g.maxX = x
	}
	if y > g.maxY {
		g.maxY = y
	}
}

func (g *Grid) MoveUp() {
	// fmt.Println("moving up")
	g.SetDoor(g.current.X, g.current.Y-1, '-')
	g.current = g.UpdateNode(g.current.X, g.current.Y-2, g.current.D+1)
}

func (g *Grid) MoveDown() {
	// fmt.Println("moving down")
	g.SetDoor(g.current.X, g.current.Y+1, '-')
	g.current = g.UpdateNode(g.current.X, g.current.Y+2, g.current.D+1)
}

func (g *Grid) MoveLeft() {
	// fmt.Println("moving left")
	g.SetDoor(g.current.X-1, g.current.Y, '|')
	g.current = g.UpdateNode(g.current.X-2, g.current.Y, g.current.D+1)
}

func (g *Grid) MoveRight() {
	// fmt.Println("moving right")
	g.SetDoor(g.current.X+1, g.current.Y, '|')
	g.current = g.UpdateNode(g.current.X+2, g.current.Y, g.current.D+1)
}

func (g *Grid) GetDoorOrWall(x, y int) rune {
	if val, ok := g.doors[fmt.Sprintf("%d,%d", x, y)]; ok {
		return val
	}
	return '#'
}

func (g *Grid) SetDoor(x, y int, d rune) {
	g.doors[fmt.Sprintf("%d,%d", x, y)] = d
}

func (g *Grid) GetNode(x, y int) *Node {
	return g.nodes[fmt.Sprintf("%d,%d", x, y)]
}

func (g *Grid) UpdateNode(x, y, d int) *Node {
	loc := fmt.Sprintf("%d,%d", x, y)
	n := g.nodes[loc]
	if n == nil {
		n = &Node{x, y, d}
		g.nodes[loc] = n
		g.ResizeFor(x, y)
	} else if d < n.D {
		n.D = d
	}
	return n
}

func (g *Grid) Print(clear bool) {
	if clear {
		fmt.Printf("\033[0;0H")
		fmt.Printf("\033[2J")
	}

	// spew.Dump(g)

	for y := g.minY - 1; y <= g.maxY+1; y++ {
		for x := g.minX - 1; x <= g.maxX+1; x++ {
			if x == 0 && y == 0 {
				fmt.Print("X")
				continue
			}

			if n := g.GetNode(x, y); n != nil {
				fmt.Print(".")
				continue
			}

			fmt.Printf("%c", g.GetDoorOrWall(x, y))

			// p := util.Vec{X: x, Y: y}
			// if g.current != nil && x == g.current.Pos.X && y == g.current.Pos.Y {
			// 	fmt.Printf("\033[%d;%dm%c\033[0m", 92, 49, g.Get(p))
			// 	continue
			// }
			// fmt.Printf("%c", g.Get(p))
		}
		fmt.Println()
	}
}
