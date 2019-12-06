package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
)

func main() {
	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	om := make(orbitmap)
	for _, l := range lines {
		parts := strings.Split(l, ")")
		o1 := om.getOrNew(parts[0])
		o2 := om.getOrNew(parts[1])
		o2.orbits = o1
	}

	fmt.Println("p1=", p1(om))
	fmt.Println("p2=", p2(om))
}

func p1(om orbitmap) (count int) {
	for _, o := range om {
		count += len(o.orbit())
	}
	return
}

func p2(om orbitmap) (count int) {
	oYOU := om["YOU"].orbit()
	oSAN := om["SAN"].orbit()

	for i, o := range oYOU {
		for j, p := range oSAN {
			if o == p {
				return i + j
			}
		}
	}

	return -1
}

type obj struct {
	name   string
	orbits *obj
}

func (o *obj) orbit() (objs []string) {
	for p := o.orbits; p != nil; p = p.orbits {
		objs = append(objs, p.name)
	}
	return
}

type orbitmap map[string]*obj

func (om orbitmap) getOrNew(name string) *obj {
	if o, ok := om[name]; ok {
		return o
	}

	om[name] = &obj{name: name}
	return om[name]
}
