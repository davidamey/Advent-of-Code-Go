package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	s3 := newState(lines, 3)
	s4 := newState(lines, 4)

	for i := 0; i < 6; i++ {
		s3.evolve()
		s4.evolve()
	}

	p1 := 0
	for _, v := range s3.values {
		if v {
			p1++
		}
	}

	p2 := 0
	for _, v := range s4.values {
		if v {
			p2++
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

type state struct {
	min, max   vec
	dimensions int
	values     map[string]bool
}

func newState(lines []string, dimensions int) (s *state) {
	s = &state{
		make(vec, dimensions),
		make(vec, dimensions),
		dimensions,
		make(map[string]bool),
	}
	for y, l := range lines {
		for x, c := range l {
			v := make(vec, dimensions)
			v[0], v[1] = x, y
			val := c == '#'
			s.values[v.String()] = val
			if val {
				s.resizeFor(v)
			}
		}
	}
	return
}

func (s *state) resizeFor(v vec) {
	for i := range v {
		if v[i] < s.min[i] {
			s.min[i] = v[i]
		}
		if v[i] > s.max[i] {
			s.max[i] = v[i]
		}
	}
}

func (s *state) evolve() {
	vs := []vec{make(vec, s.dimensions)}
	for d := 0; d < s.dimensions; d++ {
		for _, v := range vs {
			for x := s.min[d] - 1; x <= s.max[d]+1; x++ {
				if x == 0 {
					continue
				}

				v2 := make(vec, s.dimensions)
				copy(v2, v)
				v2[d] = x
				vs = append(vs, v2)
			}
		}
	}

	newValues := make(map[string]bool) //, len(vs))
	for _, v := range vs {
		active := 0
		for _, n := range v.neighbours() {
			if s.values[n.String()] {
				active++
			}
		}

		vStr := v.String()
		if s.values[vStr] {
			newValues[vStr] = active == 2 || active == 3
		} else {
			newValues[vStr] = active == 3
		}

		if newValues[vStr] {
			s.resizeFor(v)
		}
	}
	s.values = newValues
}

func (s *state) print() {
	if s.dimensions != 3 {
		fmt.Println("too many dimensions to print")
		return
	}

	v := make(vec, s.dimensions)
	for z := s.min[2]; z <= s.max[2]; z++ {
		fmt.Println("z=", z)
		for y := s.min[1]; y <= s.max[1]; y++ {
			for x := s.min[0]; x <= s.max[0]; x++ {
				v[0] = x
				v[1] = y
				if s.values[v.String()] {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

type vec []int

func (v vec) String() string {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(v[0]))
	for _, x := range v[1:] {
		sb.WriteString(",")
		sb.WriteString(strconv.Itoa(x))
	}
	return sb.String()
}

func (v vec) neighbours() []vec {
	ns := []vec{v}
	for i := range v {
		for _, n := range ns {
			n1 := make(vec, len(n))
			copy(n1, n)
			n1[i]++
			n2 := make(vec, len(n))
			copy(n2, n)
			n2[i]--
			ns = append(ns, n1, n2)
		}
	}
	return ns[1:]
}
