package main

import (
	"advent-of-code-go/util"
	"fmt"
	"io"
)

func main() {
	// file, _ := util.OpenExample()
	file, _ := util.OpenInput()
	defer file.Close()

	root := ReadNode(file)

	part1(root)
	part2(root)
}

func part1(root *Node) {
	metaSum := 0
	for _, n := range AllNodes {
		metaSum += n.MetaSum
	}

	fmt.Println("== part1 ==")
	fmt.Printf("sum of meta entries = %d\n", metaSum)

}

func part2(root *Node) {
	fmt.Println("== part2 ==")
	fmt.Printf("value of root node = %d\n", root.Value())
}

type Node struct {
	Children []*Node
	Meta     []int
	MetaSum  int
	value    int
}

var AllNodes []*Node

func ReadNode(r io.Reader) *Node {
	var lenChildren, lenMeta int
	fmt.Fscanf(r, "%d %d", &lenChildren, &lenMeta)
	// fmt.Printf("making node with %d children and %d meta entries\n", lenChildren, lenMeta)

	n := &Node{
		Children: make([]*Node, lenChildren),
		Meta:     make([]int, lenMeta),
		value:    -1,
	}

	for i := range n.Children {
		n.Children[i] = ReadNode(r)
	}

	for j := range n.Meta {
		fmt.Fscanf(r, "%d", &(n.Meta[j]))
		n.MetaSum += n.Meta[j]
	}

	AllNodes = append(AllNodes, n)
	return n
}

func (n *Node) Value() int {
	if n.value > 0 {
		return n.value
	}

	if len(n.Children) == 0 {
		n.value = n.MetaSum
	} else {
		n.value = 0
		for _, m := range n.Meta {
			idx := m - 1
			if idx >= 0 && idx < len(n.Children) {
				n.value += n.Children[idx].Value()
			}
		}
	}
	return n.value
}
