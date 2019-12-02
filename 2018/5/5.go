package main

import (
	"advent-of-code-go/util"
	"fmt"
	"io/ioutil"
)

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	poly, _ := ioutil.ReadAll(file)

	part1(poly)
	part2(poly)
}

func part1(origPoly []byte) {
	// copy so we don't alter input
	poly := clonePoly(origPoly)

	reactPoly(&poly)

	fmt.Println("== part1 ==")
	fmt.Printf("length of reacted polymer: %d\n", len(poly))
}

func part2(origPoly []byte) {
	shortest := len(origPoly)
	for _, r := range []byte("abcdefghijklmnopqrstuvwxyz") {
		poly := clonePoly(origPoly)
		removeFromPoly(&poly, r)
		reactPoly(&poly)
		if len(poly) < shortest {
			shortest = len(poly)
		}
	}
	fmt.Println("== part2 ==")
	fmt.Printf("shorted poly = %d", shortest)
}

func reactPoly(poly *[]byte) {
	var offset int
	for {
		p := *poly
		reacted := false
		offset = 0
		for i := 0; i < len(p); i++ {
			if i < len(p)-1 {
				a := p[i]
				b := p[i+1]
				if a-b == 32 || b-a == 32 {
					// fmt.Println("found reaction", i, string(a), string(b))
					reacted = true
					i++
					continue
				}
			}

			p[offset] = p[i]
			offset++
		}

		*poly = p[:offset]

		if !reacted {
			break
		}
	}
}

func removeFromPoly(poly *[]byte, unit byte) {
	p := *poly
	offset := 0
	for i := 0; i < len(p); i++ {
		if p[i] == unit || p[i] == unit-32 {
			continue
		}

		p[offset] = p[i]
		offset++
	}
	*poly = p[:offset]
}

func clonePoly(origPoly []byte) []byte {
	poly := make([]byte, len(origPoly))
	copy(poly, origPoly)
	return poly
}
