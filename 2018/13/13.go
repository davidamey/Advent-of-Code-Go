package main

import (
	"advent/util"
	"fmt"
	"sort"
	"time"
)

func main() {
	file, _ := util.OpenExample()
	// file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	// part1(lines)
	// part2(lines)
	simulate(lines)
}

func simulate(lines []string) {
	g, carts := parse(lines)

	print(g, carts, true)
	t := 0
	for {
		t++
		tick(g, carts)
		print(g, carts, true)
		time.Sleep(500 * time.Millisecond)
	}
}

func part1(lines []string) {
	g, carts := parse(lines)

	for {
		crash := tick(g, carts)
		if crash {
			break
		}
	}

	var cx, cy int
	for _, c := range carts {
		if c.crashed {
			cx = c.x
			cy = c.y
			break
		}
	}
	fmt.Println("== part1 == ")
	fmt.Printf("crash at %d,%d\n", cx, cy)
}

func part2(lines []string) {
	g, carts := parse(lines)

	var lastCart cart
	for {
		crash := tick(g, carts)
		if crash {
			activeCarts := 0
			for i, c := range carts {
				if !c.crashed {
					activeCarts++
					lastCart = carts[i]
				}
			}

			if activeCarts <= 1 {
				break
			}
		}
	}

	fmt.Println("== part2 == ")
	fmt.Printf("only one cart left at %d,%d\n", lastCart.x, lastCart.y)
}

func tick(g grid, carts []cart) bool {
	sort.Slice(carts, func(i, j int) bool {
		if carts[i].y == carts[j].y {
			return carts[i].x < carts[j].x
		}
		return carts[i].y < carts[j].y
	})

	crash := false
	for i, c := range carts {
		if c.crashed {
			continue
		}

		nx := c.x + c.vx
		ny := c.y + c.vy

		if c2 := activeCartAtXY(carts, nx, ny); c2 != nil {
			// crash!
			crash = true
			carts[i].crashed = true
			c2.crashed = true
		}

		carts[i].x = nx
		carts[i].y = ny

		switch g[carts[i].y][carts[i].x] {
		case '\\':
			carts[i].vx, carts[i].vy = c.vy, c.vx
		case '/':
			carts[i].vx, carts[i].vy = -c.vy, -c.vx
		case '+':
			switch c.choice {
			case 0:
				carts[i].vx, carts[i].vy = c.vy, -c.vx
			case 1: // Keep straight
			case 2:
				carts[i].vx, carts[i].vy = -c.vy, c.vx
			}
			carts[i].choice = (c.choice + 1) % 3
		}
	}
	return crash
}

type grid [][]rune
type cart struct {
	x, y    int
	vx, vy  int
	choice  int
	crashed bool
}

func newCart(x, y, vx, vy int) cart {
	return cart{x, y, vx, vy, 0, false}
}

func (c *cart) rune() rune {
	switch {
	case c.crashed:
		return 'X'
	case c.vy == -1:
		return '^'
	case c.vy == 1:
		return 'v'
	case c.vx == 1:
		return '>'
	case c.vx == -1:
		return '<'
	}
	return '.' // shouldn't get here ^_^
}

func activeCartAtXY(carts []cart, x, y int) *cart {
	for i, c := range carts {
		if c.crashed {
			continue
		}
		if c.x == x && c.y == y {
			return &carts[i]
		}
	}
	return nil
}

func parse(lines []string) (grid, []cart) {
	var carts []cart
	g := make(grid, len(lines))
	for y, l := range lines {
		g[y] = make([]rune, len(l))
		for x, t := range l {
			switch t {
			case '^':
				carts = append(carts, newCart(x, y, 0, -1))
				t = '|'
			case '>':
				carts = append(carts, newCart(x, y, 1, 0))
				t = '-'
			case 'v':
				carts = append(carts, newCart(x, y, 0, 1))
				t = '|'
			case '<':
				carts = append(carts, newCart(x, y, -1, 0))
				t = '-'
			}

			g[y][x] = t
		}
	}
	return g, carts
}

func print(g grid, carts []cart, clear bool) {
	if clear {
		fmt.Printf("\033[0;0H")
		fmt.Printf("\033[2J")
	}
	for y := range g {
		for x := range g[y] {
			format := "%c"
			r := g[y][x]
			for _, c := range carts {
				if c.x == x && c.y == y {
					format = "\033[92m%c\033[0m"
					if c.crashed {
						format = "\033[91m%c\033[0m"
					}
					r = c.rune()
				}
			}
			fmt.Printf(format, r)
		}
		fmt.Println()
	}
	fmt.Println()
}
