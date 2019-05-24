package main

import "fmt"

// const input = "flqrgnkx" // example
const input = "hwlqcszp"

func main() {
	// example
	// startA, startB := 65, 8921

	// input
	startA, startB := 512, 191

	fmt.Println("p1=", p1(startA, startB))
	fmt.Println("p2=", p2(startA, startB))
}

func p1(startA, startB int) (count int) {
	genA := gen{value: startA, factor: 16807}
	genB := gen{value: startB, factor: 48271}

	for i := 0; i < 40000000; i++ {
		if judge(genA.next(), genB.next()) {
			count++
		}
	}
	return
}

func p2(startA, startB int) (count int) {
	genA := gen{value: startA, factor: 16807, filter: 4}
	genB := gen{value: startB, factor: 48271, filter: 8}

	for i := 0; i < 5000000; i++ {
		if judge(genA.nextFiltered(), genB.nextFiltered()) {
			count++
		}
	}
	return
}

func judge(a, b int) bool {
	mask := (1 << 16) - 1
	return (a & mask) == (b & mask)
}

type gen struct {
	value, factor, filter int
}

func (g *gen) next() int {
	g.value *= g.factor
	g.value %= 2147483647
	return g.value
}

func (g *gen) nextFiltered() int {
	for v := g.next(); v%g.filter != 0; v = g.next() {
	}
	return g.value
}
