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

	tp := ParseLines(lines)
	fmt.Println("p1=", tp.MaxHapiness())

	for _, p := range tp.People {
		tp.AddPair("me", p, 0)
		tp.AddPair(p, "me", 0)
	}
	tp.People = append(tp.People, "me")

	fmt.Println("p2=", tp.MaxHapiness())
}

func ParseLines(lines []string) *TablePlanner {
	uniquePeople := make(map[string]struct{})
	tp := NewTablePlanner()

	for _, l := range lines {
		var p1, p2, gl string
		var h int
		fmt.Sscanf(l, "%s would %s %d happiness units by sitting next to %s.", &p1, &gl, &h, &p2)
		if gl == "lose" {
			h *= -1
		}

		tp.AddPair(p1, p2[:len(p2)-1], h)
		uniquePeople[p1] = struct{}{}
	}

	tp.People = make([]string, 0, len(uniquePeople))
	for p := range uniquePeople {
		tp.People = append(tp.People, p)
	}

	return tp
}

type TablePlanner struct {
	People   []string
	pairings map[string]int
}

func NewTablePlanner() *TablePlanner {
	return &TablePlanner{pairings: make(map[string]int)}
}

func (tp *TablePlanner) AddPair(p1, p2 string, h int) {
	tp.pairings[p1+p2] = h
}

func (tp *TablePlanner) GetHappiness(p1, p2 string) int {
	return tp.pairings[p1+p2] + tp.pairings[p2+p1]
}

func (tp *TablePlanner) HapinessForTable(table []string) int {
	h := 0
	for i := 0; i < len(table); i++ {
		h += tp.GetHappiness(table[i], table[(i+1)%len(table)])
	}
	return h
}

func (tp *TablePlanner) MaxHapiness() int {
	sp := util.NewPermuter(tp.People)

	maxH := 0
	for t := range sp.Permutations() {
		h := tp.HapinessForTable(t)

		if h > maxH {
			maxH = h
		}
	}
	return maxH
}
