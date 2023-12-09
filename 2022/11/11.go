package main

import (
	"advent-of-code-go/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	fmt.Println("p1=", p1(lines))
	fmt.Println("p2=", p2(lines))
}

func p1(lines []string) int {
	var monkeys []*monkey
	for i := 0; i < len(lines); i += 7 {
		monkeys = append(monkeys, newMonkey(lines[i:i+6]))
	}

	for round := 0; round < 20; round++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				w := m.inspect(item)
				w /= 3
				monkeys[m.test(w)].addItem(w)
			}
			m.items = []int{}
		}
	}

	inspected := make([]int, len(monkeys))
	for i, m := range monkeys {
		inspected[i] = m.inspected
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspected)))

	return inspected[0] * inspected[1]
}

func p2(lines []string) int {
	var monkeys []*monkey
	for i := 0; i < len(lines); i += 7 {
		monkeys = append(monkeys, newMonkey(lines[i:i+6]))
	}

	divisors := make([]int, len(monkeys))
	for i, m := range monkeys {
		divisors[i] = m.divisor
	}
	lcm := util.LCM(divisors...)

	for round := 0; round < 10_000; round++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				w := m.inspect(item)
				w %= lcm
				monkeys[m.test(w)].addItem(w)
			}
			m.items = []int{}
		}
	}

	inspected := make([]int, len(monkeys))
	for i, m := range monkeys {
		inspected[i] = m.inspected
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspected)))

	return inspected[0] * inspected[1]
}

type monkey struct {
	id        int
	items     []int
	divisor   int
	inspect   func(int) int
	test      func(int) int
	inspected int
}

func newMonkey(lines []string) *monkey {
	m := &monkey{}
	fmt.Sscanf(lines[0], "Monkey %d", &m.id)

	m.items = util.ParseInts(lines[1][17:], ",")

	opParts := strings.SplitN(lines[2][23:], " ", 2)
	usesOld := opParts[1] == "old"
	n, _ := strconv.Atoi(opParts[1])
	m.inspect = func(x int) int {
		m.inspected++

		if usesOld {
			n = x
		}
		switch opParts[0] {
		case "+":
			return x + n
		case "*":
			return x * n
		default:
			panic(fmt.Sprintf("unknown op %s", opParts[0]))
		}
	}

	m.divisor = util.Atoi(lines[3][21:])
	targetTrue := util.Atoi(lines[4][29:])
	targetFalse := util.Atoi(lines[5][30:])
	m.test = func(x int) int {
		if x%m.divisor == 0 {
			return targetTrue
		}
		return targetFalse
	}

	return m
}

func (m *monkey) addItem(i int) {
	m.items = append(m.items, i)
}

func (m *monkey) String() string {
	return fmt.Sprintf("Monkey %d: (inspected %5d) %v", m.id, m.inspected, m.items)
}
