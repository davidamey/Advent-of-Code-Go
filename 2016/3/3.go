package main

import (
	"advent/util"
	"fmt"
)

func main() {
	lines := util.MustReadFileToLines("input")
	p1 := 0
	for _, l := range lines {
		var s1, s2, s3 int
		fmt.Sscanf(l, "%d %d %d", &s1, &s2, &s3)
		if test(s1, s2, s3) {
			p1++
		}
	}
	fmt.Println("p1=", p1)

	p2 := 0
	for i := 0; i < len(lines); i += 3 {
		var t1, t2, t3 [3]int
		fmt.Sscanf(lines[i], "%d %d %d", &t1[0], &t2[0], &t3[0])
		fmt.Sscanf(lines[i+1], "%d %d %d", &t1[1], &t2[1], &t3[1])
		fmt.Sscanf(lines[i+2], "%d %d %d", &t1[2], &t2[2], &t3[2])

		if test(t1[0], t1[1], t1[2]) {
			p2++
		}
		if test(t2[0], t2[1], t2[2]) {
			p2++
		}
		if test(t3[0], t3[1], t3[2]) {
			p2++
		}
	}
	fmt.Println("p2=", p2)
}

func test(s1, s2, s3 int) bool {
	if s1+s2 > s3 && s2+s3 > s1 && s3+s1 > s2 {
		return true
	}
	return false
}
