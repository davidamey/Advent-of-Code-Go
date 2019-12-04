package main

import "fmt"

const (
	MIN = 231832
	MAX = 767346
)

func main() {
	p1, p2 := solve()
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)

}

func solve() (p1, p2 int) {
	for a := 2; a <= 7; a++ {
		for b := a; b <= 9; b++ {
			for c := b; c <= 9; c++ {
				for d := c; d <= 9; d++ {
					for e := d; e <= 9; e++ {
						for f := e; f <= 9; f++ {
							if checkP1(a, b, c, d, e, f) {
								p1++
							}
							if checkP2(a, b, c, d, e, f) {
								p2++
							}
						}
					}
				}
			}
		}
	}
	return
}

func checkP1(a, b, c, d, e, f int) bool {
	if !(a == b || b == c || c == d || d == e || e == f) {
		return false
	}

	i := f + 10*e + 100*d + 1000*c + 10000*b + 100000*a
	return i > MIN && i < MAX
}

func checkP2(a, b, c, d, e, f int) bool {
	aDbl := a == b && a != c && a != d && a != e && a != f
	bDbl := b == c && b != a && b != d && b != e && b != f
	cDbl := c == d && c != a && c != b && c != e && c != f
	dDbl := d == e && d != a && d != b && d != c && d != f
	eDbl := e == f && e != a && e != b && e != c && e != d

	if !(aDbl || bDbl || cDbl || dDbl || eDbl) {
		return false
	}

	i := f + 10*e + 100*d + 1000*c + 10000*b + 100000*a
	return i > MIN && i < MAX
}
