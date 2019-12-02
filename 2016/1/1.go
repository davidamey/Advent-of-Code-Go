package main

import (
	"advent-of-code-go/util"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// for _, eg := range []string{
	// 	"R2, L3",
	// 	"R2, R2, R2",
	// 	"R8, R4, R4, R8",
	// } {
	// 	fmt.Printf("e.g. \"%s\":\n", eg)
	// 	fmt.Printf("  p1= %d\n", p1(eg))
	// 	fmt.Printf("  p2= %d\n", p2(eg))
	// }

	file, _ := util.OpenInput()
	defer file.Close()
	input, _ := ioutil.ReadAll(file)

	fmt.Println("p1=", p1(string(input)))
	fmt.Println("p2=", p2(string(input)))
}

func p1(input string) int {
	p := util.NewVec(0, 0)
	v := util.NewVec(0, -1)
	for _, m := range ParseMoves(input) {
		turn(m.Turn, &v)
		for i := 0; i < m.Walk; i++ {
			p.Add(v)
		}
	}

	return p.ManhattanTo(util.NewVec(0, 0))
}

func p2(input string) int {
	p := util.NewVec(0, 0)
	v := util.NewVec(0, -1)
	places := make(map[util.Vec]int)
	for _, m := range ParseMoves(input) {
		turn(m.Turn, &v)
		for i := 0; i < m.Walk; i++ {
			p.Add(v)
			places[p]++
			if places[p] == 2 {
				return p.ManhattanTo(util.NewVec(0, 0))
			}
		}
	}

	return -1
}

func turn(dir rune, v *util.Vec) {
	switch dir {
	case 'R':
		v.X, v.Y = -v.Y, v.X
	case 'L':
		v.X, v.Y = v.Y, -v.X
	}
}

type Move struct {
	Turn rune
	Walk int
}

func ParseMoves(raw string) []Move {
	parts := strings.Split(raw, ", ")
	moves := make([]Move, len(parts))
	for i, p := range parts {
		var t rune
		var w int
		fmt.Sscanf(p, "%c%d", &t, &w)
		moves[i] = Move{t, w}
	}
	return moves
}
