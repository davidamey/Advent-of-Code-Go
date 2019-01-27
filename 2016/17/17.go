package main

import (
	"advent/util/vector"
	"crypto/md5"
	"fmt"
)

var (
	// passcode []byte = []byte("ihgpwlah") // example
	passcode []byte = []byte("yjjvjgan") // input

	min  vector.Vec = vector.New(0, 0)
	max  vector.Vec = vector.New(3, 3)
	dirs [4]byte    = [4]byte{'U', 'D', 'L', 'R'}
)

func main() {
	nodes := AllPaths(min, max)

	var min, max *Node
	for _, n := range nodes {
		if min == nil || n.Length() < min.Length() {
			min = n
		}
		if max == nil || n.Length() > min.Length() {
			max = n
		}
	}

	fmt.Println("p1=", string(min.Path))
	fmt.Println("p2=", len(max.Path))
}

type Node struct {
	Pos  vector.Vec
	Path []byte
}

func (n Node) String() string {
	return fmt.Sprintf("%s: %s", n.Pos, string(n.Path))
}

func (n *Node) Length() int {
	return len(n.Path)
}

func (n *Node) Doors() (open [4]bool) {
	hash := fmt.Sprintf("%x", md5.Sum(append(passcode, n.Path...)))
	for i := 0; i < 4; i++ {
		open[i] = hash[i] >= 'b' && hash[i] <= 'f'
	}
	return
}

func AllPaths(start, end vector.Vec) (finish []*Node) {
	queue := []*Node{&Node{Pos: start}}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		// Have we reached the end?
		if n.Pos == end {
			finish = append(finish, n)
			continue
		}

		// Add all valid routes
		routes := [4]vector.Vec{n.Pos.Up(), n.Pos.Down(), n.Pos.Left(), n.Pos.Right()}
		doors := n.Doors()
		for i := 0; i < 4; i++ {
			if doors[i] && routes[i].Within(min, max) {
				queue = append(queue, &Node{
					Pos:  routes[i],
					Path: append(append([]byte{}, n.Path...), dirs[i]),
				})
			}
		}
	}

	return
}
