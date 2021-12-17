package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
	"time"
	"unicode"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	nodes := make(map[string]*node)
	for _, l := range lines {
		edge := strings.Split(l, "-")
		addNodeIfNeeded(nodes, edge[0])
		addNodeIfNeeded(nodes, edge[1])
		nodes[edge[0]].edges = append(nodes[edge[0]].edges, nodes[edge[1]])
		nodes[edge[1]].edges = append(nodes[edge[1]].edges, nodes[edge[0]])
	}

	fmt.Println("p1=", searchP1(nodes["start"], []*node{nodes["start"]}))
	fmt.Println("p2=", searchP2(nodes["start"], []*node{nodes["start"]}, false))
}

type node struct {
	name  string
	small bool
	edges []*node
}

func searchP1(n *node, path []*node) int {
	if n.name == "end" {
		return 1
	}

	size := 0
	for _, e := range n.edges {
		if e.small && contains(path, e) {
			continue
		}
		size += searchP1(e, append([]*node{n}, path...))
	}
	return size
}

func searchP2(n *node, path []*node, doubleVisited bool) int {
	if n.name == "end" {
		return 1
	}

	size := 0
	for _, e := range n.edges {
		dv := doubleVisited

		if e.name == "start" {
			continue
		}

		if e.small && contains(path, e) {
			if dv {
				continue
			}
			dv = true
		}

		size += searchP2(e, append([]*node{n}, path...), dv)
	}
	return size
}

func addNodeIfNeeded(nodes map[string]*node, name string) {
	if _, exists := nodes[name]; !exists {
		nodes[name] = &node{name: name, small: unicode.IsLower(rune(name[0]))}
	}
}

func contains(haystack []*node, needle *node) bool {
	for _, n := range haystack {
		if n == needle {
			return true
		}
	}
	return false
}
