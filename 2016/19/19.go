package main

import (
	"advent/util"
	"fmt"
	"time"
)

const (
	// elfCount = 5 // example
	elfCount = 3014387 // input
)

func main() {
	p1()
	p2()
}

func p1() {
	defer util.Duration(time.Now())

	e, _ := makeElves()
	for i := 0; i < elfCount-1; i++ {
		e.l.remove()
		e = e.l
	}
	fmt.Println("p1=", e.id)
}

func p2() {
	defer util.Duration(time.Now())

	e, mid := makeElves()
	for i := 0; i < elfCount-1; i++ {
		mid.remove()
		mid = mid.l
		if (elfCount-i)%2 == 1 {
			mid = mid.l
		}
		e = e.l
	}
	fmt.Println("p2=", e.id)
}

func makeElves() (first, mid *elf) {
	first = &elf{id: 1}
	e := first
	for i := 1; i < elfCount; i++ {
		e.l = &elf{id: i + 1}
		e.l.r = e
		e = e.l
		if i == elfCount/2 {
			mid = e
		}
	}
	e.l = first
	first.r = e
	return
}

type elf struct {
	id   int
	l, r *elf
}

func (e *elf) remove() {
	e.r.l = e.l
	e.l.r = e.r
}
