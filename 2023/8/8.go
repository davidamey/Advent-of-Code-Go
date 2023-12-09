package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	// lines := util.MustReadFileToLines("example2") // Doesn't work with P1
	lines := util.MustReadFileToLines("input")

	instructions := lines[0]

	nodes := make(map[string]*node)
	for _, l := range lines[2:] {
		id, lID, rID := l[0:3], l[7:10], l[12:15]
		n := createIfNeeded(nodes, id)
		n.L = createIfNeeded(nodes, lID)
		n.R = createIfNeeded(nodes, rID)
	}

	fmt.Println("p1=", followPath(instructions, nodes["AAA"], func(id string) bool { return id == "ZZZ" }))

	var startNodes []*node
	for k, n := range nodes {
		if k[2] == 'A' {
			startNodes = append(startNodes, n)
		}
	}

	var pathLengths []int
	for _, n := range startNodes {
		pathLengths = append(pathLengths, followPath(instructions, n, func(id string) bool { return id[2] == 'Z' }))
	}
	fmt.Println("p2=", util.LCM(pathLengths...))
}

func followPath(instructions string, start *node, isEnd func(id string) bool) (pathLength int) {
	for n := start; !isEnd(n.id); pathLength++ {
		switch instructions[pathLength%len(instructions)] {
		case 'L':
			n = n.L
		case 'R':
			n = n.R
		}
	}
	return
}

type node struct {
	id   string
	L, R *node
}

func createIfNeeded(nodes map[string]*node, id string) *node {
	if _, exists := nodes[id]; !exists {
		nodes[id] = &node{id: id}
	}
	return nodes[id]
}
