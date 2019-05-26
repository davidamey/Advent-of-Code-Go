package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// b := newBuffer(3, 0)
	b := newBuffer(304, 0)

	// p1
	for i := 1; i <= 2017; i++ {
		b.insert(i)
	}
	fmt.Println("p1=", b.current.next.val)

	// p2
	fmt.Println("p2=", p2(304, 0))
}

func p2(step, start int) (p2 int) {
	idx := 0
	length := 1
	for i := 1; i <= 50000000; i++ {
		idx = (idx + step) % length
		idx++
		length++
		if idx == 1 {
			p2 = i
		}
	}
	return
}

type node struct {
	val        int
	next, prev *node
}

type buffer struct {
	step    int
	current *node
	zero    *node
}

func newBuffer(step, val int) *buffer {
	b := &buffer{step: step, zero: &node{val: val}}
	b.current = b.zero
	b.current.next = b.current
	b.current.prev = b.current
	return b
}

func (b *buffer) String() string {
	var sb strings.Builder
	for n := b.current; ; n = n.next {
		sb.WriteString(strconv.Itoa(n.val))
		if n.next == b.current {
			break
		}
	}
	return sb.String()
}

func (b *buffer) insert(i int) {
	for i := 0; i < b.step; i++ {
		b.current = b.current.next
	}
	n := &node{
		val:  i,
		prev: b.current,
		next: b.current.next,
	}
	n.prev.next = n
	n.next.prev = n
	b.current = n
}
