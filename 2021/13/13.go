package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	var xs, ys []int
	var folds []string

	for _, l := range lines {
		switch {
		case l == "":
			continue
		case strings.HasPrefix(l, "fold"):
			folds = append(folds, l)
		default:
			var x, y int
			fmt.Sscanf(l, "%d,%d", &x, &y)
			xs = append(xs, x)
			ys = append(ys, y)
		}
	}

	w, h := util.MaxInt(xs...)+1, util.MaxInt(ys...)+1

	p := newPaper(w, h)
	for i, y := range ys {
		p[y][xs[i]] = '#'
	}

	p1 := 0
	for i, f := range folds {
		var dir rune
		var a int
		fmt.Sscanf(f, "fold along %c=%d", &dir, &a)

		switch dir {
		case 'x':
			p.foldLeft(a)
		case 'y':
			p.foldUp(a)
		}

		if i == 0 {
			p1 = p.dotCount()
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=")
	p.print()
}

type paper [][]rune

func newPaper(w, h int) paper {
	p := make(paper, h)
	for y := range p {
		p[y] = make([]rune, w)
		for x := range p[y] {
			p[y][x] = '.'
		}
	}
	return p
}

func (p *paper) foldUp(y int) {
	for d := 1; d <= y; d++ {
		for x, v := range (*p)[y+d] {
			if v == '#' {
				(*p)[y-d][x] = '#'
			}
		}
	}
	(*p) = (*p)[:y]
}

func (p *paper) foldLeft(x int) {
	for y := range *p {
		for d := 1; d <= x; d++ {
			if (*p)[y][x+d] == '#' {
				(*p)[y][x-d] = '#'
			}
		}
		(*p)[y] = (*p)[y][:x]
	}
}

func (p paper) dotCount() (dots int) {
	for _, r := range p {
		for _, v := range r {
			if v == '#' {
				dots++
			}
		}
	}
	return
}

func (p paper) print() {
	for _, r := range p {
		fmt.Println(string(r))
	}
	fmt.Println()
}
