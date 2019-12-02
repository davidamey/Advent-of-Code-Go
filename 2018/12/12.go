package main

import (
	"advent-of-code-go/util"
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
	state := make([]*state, 21)
	state[0] = stateFromInput(lines[0])

	transforms := make([]transform, len(lines)-2)
	for i, l := range lines[2:] {
		transforms[i] = transformFromInput(l)
	}

	// fmt.Print("0: ")
	// state[0].print()
	for s := 1; s <= 20; s++ {
		state[s] = transformState(state[s-1], transforms)
		// fmt.Printf("%d: ", s)
		// state[s].print()
	}

	sum := 0
	for i, p := range state[20].Entries {
		if p == '#' {
			sum += i - state[20].zero
		}
	}
	fmt.Println("== part1 ==")
	fmt.Println("answer", sum)
}

func part2(lines []string) {
	// s1 := stateFromInput(lines[0])
	// transforms := make([]transform, len(lines)-2)
	// for i, l := range lines[2:] {
	// 	transforms[i] = transformFromInput(l)
	// }

	// var s2 *state
	// for i := 1; i <= 180; i++ {
	// 	s2 = transformState(s1, transforms)
	// 	s1 = s2

	// // fmt.Printf("%d: ", i)
	// // s2.print()

	// 	if i > 150 {
	// 		sum := 0
	// 		for i, p := range s2.Entries {
	// 			if p == '#' {
	// 				sum += i - s2.zero
	// 			}
	// 		}
	// 		fmt.Printf("%d: %d\n", i, sum)
	// 	}
	// }

	// From observation of the above commented code
	sum := func(iter int64) int64 {
		return 62*iter + 655
	}
	fmt.Println("== part2 ==")
	fmt.Println("answer", sum(50000000000))
}

func transformState(s *state, transforms []transform) *state {
	ns := NewState()
	for i := range s.Entries {
		if i < 2 || i > len(s.Entries)-3 {
			ns.Entries[i] = s.Entries[i]
			continue
		}

		ns.Entries[i] = '.'
		for _, t := range transforms {
			if result, acted := t(s.Entries[i-2 : i+3]); acted {
				// fmt.Printf("transform %d matched\n", j)
				ns.Entries[i] = result
			}
		}
	}
	return ns
}

type state struct {
	Entries []rune
	zero    int
}

type transform func([]rune) (rune, bool)

func NewState() *state {
	return &state{
		Entries: make([]rune, 300),
		zero:    10,
	}
}

func (s *state) print() {
	for _, c := range s.Entries {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}

func stateFromInput(input string) *state {
	s := NewState()

	strippedInput := input[15:]
	for i := range s.Entries {
		si := i - s.zero
		if si >= 0 && si < len(strippedInput) {
			s.Entries[i] = rune(strippedInput[si])
		} else {
			s.Entries[i] = '.'
		}
	}

	return s
}

func transformFromInput(input string) transform {
	var match string
	var result rune
	fmt.Sscanf(input, "%s => %c", &match, &result)

	return func(in []rune) (rune, bool) {
		// fmt.Printf("comparing %s to %s: ", match, string(in))
		if match != string(in) {
			// fmt.Println("diff")
			return in[2], false
		}
		// fmt.Println("match", string(result))
		return result, true
	}
}
