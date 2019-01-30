package main

import (
	"advent/util"
	"fmt"
	"strings"
	"time"
)

func logln(a ...interface{}) {
	// fmt.Println(a...)
}

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")
	ops := compileOps(lines)

	// input := "abcde" // example
	input := "abcdefgh" // input

	fmt.Println("p1=", scramble(input, ops))

	permuter := util.NewBytePermuter([]byte("fbgdceah"))
	for p := range permuter.Permutations() {
		if scramble(string(p), ops) == "fbgdceah" {
			fmt.Println("p2=", string(p))
			break
		}
	}
}

func scramble(input string, ops []op) string {
	s := newScrambler(input)
	for _, o := range ops {
		o(s)
	}
	return s.String()
}

type op func(s *scrambler)

func compileOps(lines []string) (ops []op) {
	ops = make([]op, len(lines))
	for i, l := range lines {
		switch {
		case strings.HasPrefix(l, "swap position"):
			var x, y int
			fmt.Sscanf(l, "swap position %d with position %d", &x, &y)
			ops[i] = func(s *scrambler) {
				s.swapPos(x, y)
			}
		case strings.HasPrefix(l, "swap letter"):
			var x, y rune
			fmt.Sscanf(l, "swap letter %c with letter %c", &x, &y)
			ops[i] = func(s *scrambler) {
				s.swapLetter(x, y)
			}
		case strings.HasPrefix(l, "rotate based"):
			var x rune
			fmt.Sscanf(l, "rotate based on position of letter %c", &x)
			ops[i] = func(s *scrambler) {
				s.rotateAbout(x)
			}
		case strings.HasPrefix(l, "rotate"):
			var x int
			var dir string
			fmt.Sscanf(l, "rotate %s %d", &dir, &x)
			if dir == "left" {
				ops[i] = func(s *scrambler) {
					s.rotateLeft(x)
				}
			} else {
				ops[i] = func(s *scrambler) {
					s.rotateRight(x)
				}
			}
		case strings.HasPrefix(l, "reverse"):
			var x, y int
			fmt.Sscanf(l, "reverse positions %d through %d", &x, &y)
			ops[i] = func(s *scrambler) {
				s.reverse(x, y)
			}
		case strings.HasPrefix(l, "move"):
			var x, y int
			fmt.Sscanf(l, "move position %d to position %d", &x, &y)
			ops[i] = func(s *scrambler) {
				s.move(x, y)
			}
		}
	}
	return
}

type letter struct {
	val        rune
	prev, next *letter
}

type scrambler struct {
	orig   string
	length int
	first  *letter
}

func newScrambler(orig string) *scrambler {
	var first, l *letter
	for i, c := range orig {
		if i == 0 {
			first = &letter{val: c}
			l = first
			continue
		}

		l.next = &letter{val: c, prev: l}
		l = l.next
	}
	l.next = first
	first.prev = l

	return &scrambler{
		orig:   orig,
		length: len(orig),
		first:  first,
	}
}

func (s scrambler) String() string {
	out := make([]rune, s.length)
	for i, l := 0, s.first; i < s.length; i, l = i+1, l.next {
		out[i] = l.val
	}
	return string(out)
}

func (s *scrambler) swapPos(x, y int) {
	logln("swapPos", x, y)
	var lx, ly *letter
	for i, l := 0, s.first; lx == nil || ly == nil; i, l = i+1, l.next {
		if i == x {
			lx = l
		}
		if i == y {
			ly = l
		}
	}
	lx.val, ly.val = ly.val, lx.val
}

func (s *scrambler) swapLetter(x, y rune) {
	logln("swapLetter", string(x), string(y))
	var lx, ly *letter
	for l := s.first; lx == nil || ly == nil; l = l.next {
		if l.val == x {
			lx = l
		}
		if l.val == y {
			ly = l
		}
	}
	lx.val, ly.val = ly.val, lx.val
}

func (s *scrambler) rotateLeft(steps int) {
	logln("rotateLeft", steps)
	for i := 0; i < steps; i++ {
		s.first = s.first.next
	}
}

func (s *scrambler) rotateRight(steps int) {
	logln("rotateRight", steps)
	for i := 0; i < steps; i++ {
		s.first = s.first.prev
	}
}

func (s *scrambler) rotateAbout(x rune) {
	logln("rotateAbout", string(x))
	i := 0
	for l := s.first; i < s.length; i, l = i+1, l.next {
		if l.val == x {
			break
		}
	}
	steps := 1 + i
	if i >= 4 {
		steps++
	}
	s.rotateRight(steps)
}

func (s *scrambler) reverse(x, y int) {
	logln("reverse", x, y)
	var lx, ly *letter
	for i, l := 0, s.first; lx == nil || ly == nil; i, l = i+1, l.next {
		if i == x {
			lx = l
		}
		if i == y {
			ly = l
		}
	}

	for x < y {
		lx.val, ly.val = ly.val, lx.val
		lx = lx.next
		ly = ly.prev
		x++
		y--
	}
}

func (s *scrambler) move(x, y int) {
	logln("move", x, y)
	var lx, ly *letter
	for i, l := 0, s.first; lx == nil || ly == nil; i, l = i+1, l.next {
		if i == x {
			lx = l
		}
		if i == y {
			ly = l
		}
	}

	if x == 0 {
		s.first = lx.next
	}
	if y == 0 {
		s.first = lx
	}

	lx.prev.next = lx.next
	lx.next.prev = lx.prev

	if x < y {
		ly.next.prev = lx
		lx.next = ly.next
		ly.next = lx
		lx.prev = ly
	} else {
		ly.prev.next = lx
		lx.prev = ly.prev
		ly.prev = lx
		lx.next = ly
	}
}
