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

	// tm := &turingMachine{make(map[int]int), 0, make(map[rune]state), 'A'}
	// tm.states['A'] = func(tm *turingMachine) {
	// 	if tm.tape[tm.cursor] == 0 {
	// 		tm.tape[tm.cursor] = 1
	// 		tm.cursor++
	// 	} else {
	// 		tm.tape[tm.cursor] = 0
	// 		tm.cursor--
	// 	}
	// 	tm.currentState = 'B'
	// }

	// tm.states['B'] = func(tm *turingMachine) {
	// 	if tm.tape[tm.cursor] == 0 {
	// 		tm.tape[tm.cursor] = 1
	// 		tm.cursor--
	// 	} else {
	// 		tm.tape[tm.cursor] = 1
	// 		tm.cursor++
	// 	}
	// 	tm.currentState = 'A'
	// }

	tm := newTM(lines)
	tm.run(false)

	fmt.Println("p1=", tm.checksum())
}

type turingMachine struct {
	tape            map[int]int
	cursor          int
	states          map[rune]state
	currentState    rune
	diagnosticAfter int
}

type state func(tm *turingMachine)

func newState(lines []string) state {
	var (
		offV, onV int
		offC, onC int = 1, 1
		offS, onS rune
	)

	fmt.Sscanf(lines[1], "    - Write the value %d.", &offV)
	if lines[2] == "    - Move one slot to the left." {
		offC = -1
	}
	fmt.Sscanf(lines[3], "    - Continue with state %c.", &offS)

	fmt.Sscanf(lines[5], "    - Write the value %d.", &onV)
	if lines[6] == "    - Move one slot to the left." {
		onC = -1
	}
	fmt.Sscanf(lines[7], "    - Continue with state %c.", &onS)

	return func(tm *turingMachine) {
		if tm.tape[tm.cursor] == 0 {
			tm.tape[tm.cursor] = offV
			tm.cursor += offC
			tm.currentState = offS
		} else {
			tm.tape[tm.cursor] = onV
			tm.cursor += onC
			tm.currentState = onS
		}
	}
}

func newTM(lines []string) *turingMachine {
	tm := &turingMachine{
		tape:   make(map[int]int),
		cursor: 0,
		states: make(map[rune]state),
	}

	fmt.Sscanf(lines[0], "Begin in state %c.", &tm.currentState)
	fmt.Sscanf(lines[1], "Perform a diagnostic checksum after %d steps.", &tm.diagnosticAfter)

	for i := 3; i < len(lines); i += 10 {
		var s rune
		fmt.Sscanf(lines[i], "In state %c:", &s)
		tm.states[s] = newState(lines[i+1 : i+10])
	}

	return tm
}

func (tm *turingMachine) String() string {
	var sb strings.Builder
	sb.WriteString("...")

	l := len(tm.tape)
	if l < 6 {
		l = 6
	}
	l /= 2

	for i := -l; i <= l; i++ {
		if i == tm.cursor {
			sb.WriteString(fmt.Sprintf("[%d]", tm.tape[i]))
		} else {
			sb.WriteString(fmt.Sprintf(" %d ", tm.tape[i]))
		}
	}

	sb.WriteString("...")
	sb.WriteString(fmt.Sprintf(" | state %c", tm.currentState))
	return sb.String()
}

func (tm *turingMachine) run(debug bool) {
	if debug {
		fmt.Println(tm)
	}
	for i := 0; i < tm.diagnosticAfter; i++ {
		tm.states[tm.currentState](tm)
		if debug {
			fmt.Println(tm)
		}
	}
	if debug {
		fmt.Println()
	}
}

func (tm *turingMachine) checksum() (csum int) {
	for _, v := range tm.tape {
		if v == 1 {
			csum++
		}
	}
	return
}
