package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	// lines, minutes := util.MustReadFileToLines("example"), 10
	lines, minutes := util.MustReadFileToLines("input"), 200

	fmt.Println("p1=", p1(lines))
	fmt.Println("p2=", p2(lines, minutes))
}

func p1(lines []string) int {
	a := newArea(lines)
	seen := make(map[area]bool)
	for {
		if seen[a] {
			break
		}
		seen[a] = true

		// tick
		var na area
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				bc := a.bugCount(x, y)
				i := y*5 + x
				switch {
				case a[i] && bc != 1:
					na[i] = false
				case !a[i] && bc == 1 || bc == 2:
					na[i] = true
				default:
					na[i] = a[i]
				}
			}
		}
		a = na
	}
	return a.biodiversity()
}

func p2(lines []string, minutes int) (bugs int) {
	areas := map[int]area{
		-1: *new(area),
		0:  newArea(lines),
		1:  *new(area),
	}
	limit := 1

	for m := 0; m < minutes; m++ {
		nareas := make(map[int]area)
		for l, a := range areas {
			var na area
			for y := 0; y < 5; y++ {
				for x := 0; x < 5; x++ {
					if x == 2 && y == 2 {
						continue
					}

					bc := a.bugCountRecusive(x, y, areas[l-1], areas[l+1])
					i := y*5 + x
					switch {
					case a[i] && bc != 1:
						na[i] = false
					case !a[i] && bc == 1 || bc == 2:
						na[i] = true
					default:
						na[i] = a[i]
					}
				}
			}
			nareas[l] = na
		}
		areas = nareas

		limit++
		areas[-limit] = *new(area)
		areas[limit] = *new(area)
	}

	// count the bugs!
	for _, a := range areas {
		for _, b := range a {
			if b {
				bugs++
			}
		}
	}
	return
}

type vec struct{ x, y int }
type area [25]bool

func newArea(lines []string) (a area) {
	for i := range a {
		x, y := i%5, i/5
		a[i] = lines[y][x] == '#'
	}
	return
}

func (a area) at(x, y int) bool {
	return a[y*5+x]
}

func (a area) print(recursive bool) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if recursive && x == 2 && y == 2 {
				fmt.Print("?")
				continue
			}

			if a.at(x, y) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (a area) bugCount(x, y int) (bc int) {
	for _, v := range []vec{{0, -1}, {-1, 0}, {1, 0}, {0, 1}} {
		xx, yy := x+v.x, y+v.y
		if xx >= 0 && xx < 5 && yy >= 0 && yy < 5 {
			if a.at(xx, yy) {
				bc++
			}
		}
	}
	return
}

func (a area) bugCountRecusive(x, y int, outer, inner area) (bc int) {
	for _, v := range []vec{{0, -1}, {-1, 0}, {1, 0}, {0, 1}} {
		xx, yy := x+v.x, y+v.y
		switch {
		case xx == 2 && yy == 2 && x == 2:
			for innerX := 0; innerX < 5; innerX++ {
				if (y == 1 && inner.at(innerX, 0)) || (y == 3 && inner.at(innerX, 4)) {
					bc++
				}
			}
		case xx == 2 && yy == 2 && y == 2:
			for innerY := 0; innerY < 5; innerY++ {
				if (x == 1 && inner.at(0, innerY)) || (x == 3 && inner.at(4, innerY)) {
					bc++
				}
			}
		case
			xx < 0 && outer.at(1, 2),
			yy < 0 && outer.at(2, 1),
			xx >= 5 && outer.at(3, 2),
			yy >= 5 && outer.at(2, 3):
			bc++
		case xx < 0, xx >= 5, yy < 0, yy >= 5:
			continue
		case a.at(xx, yy):
			bc++
		}
	}
	return
}

func (a area) biodiversity() (bd int) {
	for i, b := range a {
		if b {
			bd += 1 << i
		}
	}
	return
}
