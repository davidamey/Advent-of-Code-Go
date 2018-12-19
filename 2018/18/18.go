package main

import (
	"advent/util"
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
	var current, next *util.Grid

	current = util.NewGrid(lines, nil)
	next = current.Clone()
	// current.Print(false)
	for i := 1; i <= 10; i++ {
		for y, row := range *current {
			for x := range row {
				(*next)[y][x] = Evolve(current, x, y)
				// fmt.Scanln()
			}
		}
		// fmt.Printf("After %d minute(s):\n", i)
		// next.Print(false)
		current, next = next, current

		// fmt.Scanln()
	}

	area := make(map[rune]int, 3)
	for _, row := range *current {
		for _, r := range row {
			area[r]++
		}
	}

	fmt.Println("== part1 == ")
	fmt.Printf(". (%d), | (%d), # (%d)\n", area['.'], area['|'], area['#'])
	fmt.Printf("answer = %d\n", area['|']*area['#'])
	fmt.Println()
}

func part2(lines []string) {
	var current, next *util.Grid

	current = util.NewGrid(lines, nil)
	next = current.Clone()

	var results []string
	for i := 0; i < 2000; i++ {
		area := make(map[rune]int, 3)
		for _, row := range *current {
			for _, r := range row {
				area[r]++
			}
		}

		results = append(results, fmt.Sprintf("  %d => . (%d), | (%d), # (%d)", i, area['.'], area['|'], area['#']))

		for y, row := range *current {
			for x := range row {
				(*next)[y][x] = Evolve(current, x, y)
			}
		}
		current, next = next, current
	}

	fmt.Println("== part2 ==")
	for _, r := range results[1000:1029] {
		fmt.Println(r)
	}

	// The above loop shows the following repeated cycle of length 28
	// 1000 => . (1674), | (510), # (316)
	// 1001 => . (1678), | (511), # (311)
	// 1002 => . (1684), | (513), # (303)
	// 1003 => . (1679), | (513), # (308)
	// 1004 => . (1674), | (520), # (306)
	// 1005 => . (1673), | (523), # (304)
	// 1006 => . (1665), | (533), # (302)
	// 1007 => . (1656), | (539), # (305)
	// 1008 => . (1649), | (551), # (300)
	// 1009 => . (1633), | (563), # (304)
	// 1010 => . (1618), | (578), # (304)
	// 1011 => . (1603), | (586), # (311)
	// 1012 => . (1591), | (598), # (311)
	// 1013 => . (1580), | (597), # (323)
	// 1014 => . (1579), | (599), # (322)
	// 1015 => . (1575), | (592), # (333)
	// 1016 => . (1584), | (591), # (325)
	// 1017 => . (1587), | (585), # (328)
	// 1018 => . (1589), | (584), # (327)
	// 1019 => . (1588), | (579), # (333)
	// 1020 => . (1593), | (576), # (331)
	// 1021 => . (1598), | (566), # (336)
	// 1022 => . (1606), | (559), # (335)
	// 1023 => . (1616), | (544), # (340)
	// 1024 => . (1631), | (533), # (336)
	// 1025 => . (1639), | (519), # (342)
	// 1026 => . (1657), | (512), # (331)
	// 1027 => . (1671), | (508), # (321)
	// 1028 => . (1674), | (510), # (316)

	// As (1000000000 - 1000) % 28 = 0
	// The answer for the 1000000000 iteration matches that of the 1000th
	// i.e. 510 * 316 = 161160
	fmt.Println("answer = ", 510*316)
}

func Evolve(g *util.Grid, x, y int) rune {
	// fmt.Printf("Looking at cell (%d, %d) = %c\n", x, y, (*g)[y][x])
	b1, b2 := g.Bounds()
	p := util.Point{X: x, Y: y}

	area := make(map[rune]int, 3)
	// fmt.Println("  adjacent", p.Adjacent(true))
	for _, q := range p.Adjacent(true) {
		if !q.Within(b1, b2) {
			// fmt.Printf("  oob (%d, %d)\n", q.X, q.Y)
			continue
		}

		// fmt.Printf("  nearby (%d, %d) = %c\n", q.X, q.Y, (*g)[q.Y][q.X])
		area[(*g)[q.Y][q.X]]++
	}

	r := (*g)[y][x]
	// fmt.Printf("  processing %d (.:%d, |:%d, #:%d)\n", r, area['.'], area['|'], area['#'])
	switch r {
	case '.': // Open acre becomes Tree if 3+ Trees adjacent
		if area['|'] >= 3 {
			r = '|'
		} else {
			// fmt.Printf("   not enough | (count = %d)\n", area['|'])
		}
	case '|': // Tree becomes LumberYard if 3+ LumberYards adjacent
		if area['#'] >= 3 {
			r = '#'
		} else {
			// fmt.Printf("   not enough # (count = %d)\n", area['#'])
		}
	case '#': // LumberYard stays if 1+ LumberYard and 1+ Trees adjacent
		if area['#'] == 0 || area['|'] == 0 {
			r = '.'
		} else {
			// fmt.Printf("   enough # and | to stay # (count = %d,%d)\n", area['#'], area['|'])
		}
	}

	// fmt.Printf(" returning %c\n", r)

	return r
}
