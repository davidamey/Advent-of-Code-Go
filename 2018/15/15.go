package main

import (
	"advent/util"
	"fmt"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(`---`)
		fmt.Printf("Ran in %f seconds\n", time.Since(start).Seconds())
	}()
	// file, _ := util.OpenExample()
	// file, _ := util.OpenFile("example2")
	// file, _ := util.OpenFile("example3")

	// file, _ := util.OpenInput()
	file, _ := util.OpenFile("input_alt")
	defer file.Close()
	lines, _ := util.ReadLines(file)

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	g := Parse(lines, 3)
	g.Run(false, true)

	totalHP := 0
	for _, u := range g.Units {
		if u.Alive() {
			totalHP += u.HP
		}
	}

	fmt.Println("== Part1 ==")
	fmt.Println("Last complete round:", g.Round)
	fmt.Println("Remaining HP:", totalHP)
	fmt.Println("Outcome: ", g.Round*totalHP)
	fmt.Println()
}

func part2(lines []string) {
	elfAtk := 3

	var g *Game
	for {
		g = Parse(lines, elfAtk)
		if !g.Run(false, false) {
			break
		}
		elfAtk++
	}

	totalHP := 0
	for _, u := range g.Units {
		if u.Alive() {
			totalHP += u.HP
		}
	}

	fmt.Println("== Part2 ==")
	fmt.Println("Elf attack:", elfAtk)
	fmt.Println("Last complete round:", g.Round)
	fmt.Println("Remaining HP:", totalHP)
	fmt.Println("Outcome: ", g.Round*totalHP)
	fmt.Println()
}

type Game struct {
	Grid      [][]rune
	Units     []*Unit
	Round     int
	LastActed *Unit
}

func (g *Game) Run(visual, allowElfDeath bool) (elfDied bool) {
	for {
		if visual {
			Print(g, true)
		}

		sort.Slice(g.Units, func(i, j int) bool {
			if g.Units[i].Pos.Y == g.Units[j].Pos.Y {
				return g.Units[i].Pos.X < g.Units[j].Pos.X
			}
			return g.Units[i].Pos.Y < g.Units[j].Pos.Y
		})

		for j := range g.Units {
			success := g.Units[j].TakeTurn()
			if !success {
				return false
			}
		}

		if visual {
			time.Sleep(300 * time.Millisecond)
			// fmt.Scanln()
		}

		if allowElfDeath == false {
			for _, u := range g.Units {
				if u.Rune == 'E' && !u.Alive() {
					return true
				}
			}
		}

		g.Round++
	}
}

type PathNode struct {
	Pos    util.Vec
	Length int
	Parent *PathNode
}

func ShortestPath(start util.Vec, ends []util.Vec, grid [][]rune) []PathNode {
	var found []PathNode

	queue := make([]PathNode, 1, 4)
	queue[0] = PathNode{start, 0, nil}

	processed := make(map[util.Vec]bool)

	depth := 0
	shortest := -1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// grid[node.Pos.Y][node.Pos.X] = rune(48 + node.Length)

		if node.Length > depth {
			// Reached a new depth of search
			depth = node.Length
			if shortest >= 0 && depth > shortest {
				// We've found an answer and will forever be higher than it now so return
				return found
			}
		}

		// Check if we've reached an end
		for _, e := range ends {
			if e == node.Pos {
				// fmt.Println("adding", node, e)
				shortest = depth
				found = append(found, node)
			}
		}

		// Add all children to the queue if we haven't processed them already
		for _, c := range node.Pos.Adjacent(false) {
			if _, ok := processed[c]; ok {
				// fmt.Println("skipping", c)
				continue
			}

			if grid[c.Y][c.X] == '.' {
				queue = append(queue, PathNode{c, depth + 1, &node})
				processed[c] = true
			}
		}
	}

	return found
}

func Parse(lines []string, elfAtk int) *Game {
	g := Game{}
	g.Grid = make([][]rune, len(lines))
	for y, l := range lines {
		g.Grid[y] = make([]rune, len(l))
		for x, r := range l {
			switch r {
			case 'E':
				g.Units = append(g.Units, &Unit{&g, r, util.Vec{X: x, Y: y}, 200, elfAtk})
			case 'G':
				g.Units = append(g.Units, &Unit{&g, r, util.Vec{X: x, Y: y}, 200, 3})
			}
			g.Grid[y][x] = r
		}
	}
	return &g
}

func Print(g *Game, clear bool) {
	if clear {
		fmt.Printf("\033[0;0H")
		fmt.Printf("\033[2J")
	}

	if g.Round == 0 {
		fmt.Println("Initial:")
	} else {
		plural := ""
		if g.Round > 1 {
			plural = "s"
		}
		fmt.Printf("After %d round%s:\n", g.Round, plural)
	}

	for y := range g.Grid {
		rowUnits := Filter(g.Units, func(u *Unit) bool {
			return u.Pos.Y == y && u.Alive()
		})
		sort.Slice(rowUnits, func(i, j int) bool {
			return rowUnits[i].Pos.X < rowUnits[j].Pos.X
		})

		for x := range g.Grid[y] {
			r := g.Grid[y][x]
			fg := 39
			bg := 49

			for _, u := range rowUnits {
				if u.Pos.X == x {
					r = u.Rune

					switch u.Rune {
					case 'E':
						fg = 93
					case 'G':
						fg = 92
					}

					if !u.Alive() {
						r = '.'
						// r = 'x'
						// fg = 91
					}

					// if u.Game.LastActed != nil && *u.Game.LastActed == u {
					// 	bg = 105
					// }
				}
			}

			fmt.Printf("\033[%d;%dm%c\033[0m", fg, bg, r)
		}

		fmt.Print("    ")
		for i, u := range rowUnits {
			fmt.Printf("(%c %3d)", u.Rune, u.HP)
			if i+1 != len(rowUnits) {
				fmt.Print(", ")
			}
		}

		fmt.Println()
	}
	fmt.Println()
}
