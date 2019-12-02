package main

import (
	"advent-of-code-go/util"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// example := &state{
	// 	floors: 4,
	// 	lift:   0,
	// 	pairs: []pair{
	// 		pair{'H', 0, 1},
	// 		pair{'L', 0, 2},
	// 	},
	// }

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")
	s := stateFromLines(lines)
	fmt.Println("p1=", solve(s))

	// part 2
	s.pairs = append(s.pairs,
		pair{'E', 0, 0},
		pair{'D', 0, 0},
	)
	s.sort()
	fmt.Println("p2=", solve(s))
}

func solve(initial *state) (moves int) {
	seen := make(map[string]struct{})
	states := []*state{initial}
	for {
		if len(states) == 0 {
			panic("no solution")
		}

		nextStates := make([]*state, 0, len(states))
		for _, s := range states {
			if s.finished() {
				// // print out each move
				// for t := s; t != nil; t = t.prev {
				// 	t.dump()
				// 	fmt.Println("---")
				// }
				return
			}

			h := s.hash()
			if _, ok := seen[h]; ok {
				continue
			}
			seen[h] = struct{}{}

			nextStates = append(nextStates, possibleNextStates(s)...)
		}
		states = nextStates
		moves++
	}
}

type move struct {
	chips []int
	gens  []int
}

func possibleMoves(s *state) []move {
	var chips, gens []int
	var moves []move
	for i, p := range s.pairs {
		// if chip and gen are on the floor then the pair is a valid move
		if p.c == s.lift && p.g == s.lift {
			moves = append(moves, move{
				[]int{i}, []int{i},
			})
		}
		if p.c == s.lift {
			chips = append(chips, i)
		}
		if p.g == s.lift {
			gens = append(gens, i)
		}
	}

	for i, c := range chips {
		moves = append(moves, move{chips: []int{c}})
		for _, c2 := range chips[i+1:] {
			moves = append(moves, move{chips: []int{c, c2}})
		}
	}

	for i, g := range gens {
		moves = append(moves, move{gens: []int{g}})
		for _, g2 := range gens[i+1:] {
			moves = append(moves, move{gens: []int{g, g2}})
		}
	}

	return moves
}

func possibleNextStates(s *state) (next []*state) {
	for _, d := range []int{-1, 1} {
		if s.lift+d < 0 || s.lift+d >= s.floors {
			continue
		}

		for _, m := range possibleMoves(s) {
			s2 := s.clone()
			s2.prev = s
			s2.move(d, m.chips, m.gens)
			if s2.valid() {
				next = append(next, s2)
			}
		}
	}

	return
}

type pair struct {
	el   rune
	c, g int
}

type state struct {
	floors int
	lift   int
	pairs  []pair
	prev   *state
}

func stateFromLines(lines []string) *state {
	rgxChip := regexp.MustCompile(`(\w)\w+-compatible microchip`)
	rgxGen := regexp.MustCompile(`(\w)\w+ generator`)

	chips := make(map[rune]int)
	gens := make(map[rune]int)

	for i, l := range lines {
		for _, m := range rgxChip.FindAllStringSubmatch(l, -1) {
			chips[rune(strings.ToUpper(m[1])[0])] = i
		}
		for _, m := range rgxGen.FindAllStringSubmatch(l, -1) {
			gens[rune(strings.ToUpper(m[1])[0])] = i
		}
	}

	s := &state{
		floors: len(lines),
		pairs:  make([]pair, 0, len(chips)),
	}
	for el := range chips {
		s.pairs = append(s.pairs, pair{el, chips[el], gens[el]})
	}

	s.sort()
	return s
}

func (s *state) hash() string {
	buf := make([]byte, 2+2*len(s.pairs))
	buf[0] = byte(s.floors)
	buf[1] = byte(s.lift)
	for i, p := range s.pairs {
		idx := (i + 1) * 2
		buf[idx] = byte(p.c)
		buf[idx+1] = byte(p.g)
	}
	return fmt.Sprintf("%v", buf)
}

func (s *state) valid() bool {
	floorGenCount := make([]int, s.floors)
	for _, p := range s.pairs {
		floorGenCount[p.g]++
	}

	for _, p := range s.pairs {
		if floorGenCount[p.c] > 0 && p.c != p.g {
			return false
		}
	}
	return true
}

func (s *state) finished() bool {
	for _, p := range s.pairs {
		if p.c != s.floors-1 || p.g != s.floors-1 {
			return false
		}
	}
	return true
}

func (s *state) move(dir int, chips, gens []int) {
	s.lift += dir
	for _, c := range chips {
		s.pairs[c].c += dir
	}
	for _, g := range gens {
		s.pairs[g].g += dir
	}
	s.sort()
}

func (s *state) clone() *state {
	s2 := &state{
		floors: s.floors,
		lift:   s.lift,
		pairs:  make([]pair, len(s.pairs)),
	}
	copy(s2.pairs, s.pairs)
	return s2
}

// sort orders the pairs by chip floor then gen floor.
// massively improves run-time as discounts many many ultimately equivalent states
func (s *state) sort() {
	sort.Slice(s.pairs, func(i, j int) bool {
		if s.pairs[i].c == s.pairs[j].c {
			return s.pairs[i].g < s.pairs[j].g
		}
		return s.pairs[i].c < s.pairs[j].c
	})
}

func (s *state) dump() {
	floors := make([][]byte, s.floors)
	for i := range floors {
		floors[i] = append(floors[i], 'F', byte('0'+i+1), ' ')

		if s.lift == i {
			floors[i] = append(floors[i], 'E', ' ')
		} else {
			floors[i] = append(floors[i], '.', ' ')
		}

		for _, p := range s.pairs {
			if p.c == i {
				floors[i] = append(floors[i], byte(p.el), 'M', ' ')
			} else {
				floors[i] = append(floors[i], '.', ' ', ' ')
			}
			if p.g == i {
				floors[p.g] = append(floors[p.g], byte(p.el), 'G', ' ')
			} else {
				floors[i] = append(floors[i], '.', ' ', ' ')
			}
		}
	}

	for i := len(floors) - 1; i >= 0; i-- {
		fmt.Println(string(floors[i]))
	}
}
