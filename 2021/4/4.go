package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	calls := util.ParseInts(lines[0], ",")

	var boards []*bingoBoard
	for i := 2; i < len(lines)-4; i += 6 {
		boards = append(boards, newBingoBoard(lines[i:i+5]))
	}

	p1, p2 := 0, 0
	finishedBoards := 0
	for _, c := range calls {
		for _, b := range boards {
			if b.finished {
				continue
			}

			if b.isWinnerWith(c) {
				if p1 == 0 {
					p1 = b.sumUnmarked() * c
				}

				b.finished = true
				finishedBoards++
				if finishedBoards == len(boards) {
					p2 = b.sumUnmarked() * c
				}
			}
		}
	}
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

type bingoBoard struct {
	values   [5][5]int
	matched  [5][5]bool
	finished bool
}

func newBingoBoard(lines []string) *bingoBoard {
	b := &bingoBoard{}
	for y, l := range lines {
		for x, a := range strings.Fields(l) {
			b.values[y][x], _ = strconv.Atoi(a)
		}
	}
	return b
}

func (b *bingoBoard) isWinnerWith(c int) bool {
	for y, r := range b.values {
		for x, v := range r {
			if v == c {
				b.matched[y][x] = true
				return b.isWinner(x, y)
			}
		}
	}
	return false
}

func (b *bingoBoard) isWinner(x, y int) bool {
	row, col := true, true
	for i := 0; i < 5; i++ {
		row = row && b.matched[y][i]
		col = col && b.matched[i][x]
	}
	return row || col
}

func (b *bingoBoard) sumUnmarked() (sum int) {
	for y, r := range b.matched {
		for x, m := range r {
			if !m {
				sum += b.values[y][x]
			}
		}
	}
	return
}
