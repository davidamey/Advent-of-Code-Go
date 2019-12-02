package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// input := util.MustReadFileToLines("example")
	input := util.MustReadFileToLines("input")

	all := make(group)
	for i := range input {
		all[i] = newProgram(i)
	}

	for _, l := range input {
		parts := strings.Split(l, " <-> ")
		id, _ := strconv.Atoi(parts[0])

		p := all[id]
		for _, q := range strings.Split(parts[1], ", ") {
			qid, _ := strconv.Atoi(q)
			p.addPipe(all[qid])
		}
	}

	uniqueGroups := make(map[string]int)
	for _, p := range all {
		gid := fmt.Sprintf("%p", p.group)
		uniqueGroups[gid]++
	}

	fmt.Println("p1=", len(all[0].group))
	fmt.Println("p2=", len(uniqueGroups))
}

// func altP1(input []string) {
// 	g := make(map[int]bool)
// 	toSearch := []int{0}

// 	for len(toSearch) > 0 {
// 		var id int
// 		id, toSearch = toSearch[len(toSearch)-1], toSearch[:len(toSearch)-1]

// 		if g[id] {
// 			continue
// 		}
// 		fmt.Println()
// 		g[id] = true

// 		parts := strings.Split(input[id], " <-> ")
// 		for _, p := range strings.Split(parts[1], ", ") {
// 			pid, _ := strconv.Atoi(p)
// 			toSearch = append(toSearch, pid)
// 		}
// 	}

// 	fmt.Println(len(g))
// }

type program struct {
	id    int
	group group
}

func newProgram(id int) *program {
	p := &program{id: id, group: make(group)}
	p.group[id] = p
	return p
}

func (p *program) addPipe(q *program) {
	if &p.group == &q.group {
		// pipe already exists
		return
	}

	// merge the groups
	for k, v := range q.group {
		p.group[k] = v
		v.group = p.group
	}
}

type group map[int]*program
