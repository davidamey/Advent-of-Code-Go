package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/grid"
	"advent-of-code-go/util/vector"
	"fmt"
)

func main() {
	// file, _ := util.OpenExample()
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	code := make([]rune, len(lines))
	pad := makePad()
	p := vector.New(1, 1)
	for i, l := range lines {
		for _, c := range l {
			var newp vector.Vec
			switch c {
			case 'U':
				newp = p.Up()
			case 'R':
				newp = p.Right()
			case 'D':
				newp = p.Down()
			case 'L':
				newp = p.Left()
			}

			if pad.InBounds(newp) {
				p = newp
			}
		}
		code[i] = pad.Get(p)
	}

	fmt.Println("p1=", string(code))
}

func part2(lines []string) {
	code := make([]rune, len(lines))
	pad := makePadP2()
	p := pad['5']
	for i, l := range lines {
		for _, c := range l {
			switch {
			case c == 'U' && p.U != nil:
				p = p.U
			case c == 'R' && p.R != nil:
				p = p.R
			case c == 'D' && p.D != nil:
				p = p.D
			case c == 'L' && p.L != nil:
				p = p.L
			}
		}
		code[i] = p.Val
	}

	fmt.Println("p2=", string(code))
}

func makePad() *grid.Grid[rune] {
	pad := grid.New[rune]()
	pad.SetAt(0, 0, '1')
	pad.SetAt(1, 0, '2')
	pad.SetAt(2, 0, '3')
	pad.SetAt(0, 1, '4')
	pad.SetAt(1, 1, '5')
	pad.SetAt(2, 1, '6')
	pad.SetAt(0, 2, '7')
	pad.SetAt(1, 2, '8')
	pad.SetAt(2, 2, '9')
	return pad
}

type p2pad map[rune]*key
type key struct {
	Val        rune
	U, R, D, L *key
}

func makePadP2() p2pad {
	pad := make(p2pad)
	for _, r := range []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D'} {
		pad[r] = &key{Val: r}
	}

	pad.link('1', '-', '-', '3', '-')
	pad.link('2', '-', '3', '6', '-')
	pad.link('3', '1', '4', '7', '2')
	pad.link('4', '-', '-', '8', '3')
	pad.link('5', '-', '6', '-', '-')
	pad.link('6', '2', '7', 'A', '5')
	pad.link('7', '3', '8', 'B', '6')
	pad.link('8', '4', '9', 'C', '7')
	pad.link('9', '-', '-', '-', '8')
	pad.link('A', '6', 'B', '-', '-')
	pad.link('B', '7', 'C', 'D', 'A')
	pad.link('C', '8', '-', '-', 'B')
	pad.link('D', 'B', '-', '-', '-')

	return pad
}

func (p p2pad) link(key, u, r, d, l rune) {
	p[key].U = p[u]
	p[key].R = p[r]
	p[key].D = p[d]
	p[key].L = p[l]
}
