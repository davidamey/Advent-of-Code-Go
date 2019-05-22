package main

import (
	"advent/util"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// examples
	// fmt.Println("ex p1=", p1(5, "3, 4, 1, 5"))
	// fmt.Println("ex p2=", p2([]byte("")))
	// fmt.Println("ex p2=", p2([]byte("AoC 2017")))
	// fmt.Println("ex p2=", p2([]byte("1,2,3")))
	// fmt.Println("ex p2=", p2([]byte("1,2,4")))

	// actual
	input := util.MustReadFile("input")
	fmt.Println("p1=", p1(256, string(input)))
	fmt.Println("p2=", p2(input))
}

func p1(listSize int, input string) int {
	var lengths []int
	for _, s := range strings.Split(input, ", ") {
		l, _ := strconv.Atoi(s)
		lengths = append(lengths, l)
	}

	l := make(list, listSize)
	for i := range l {
		l[i] = i
	}

	pos := 0
	skip := 0

	for _, ln := range lengths {
		l.reverse(pos, pos+ln-1)
		pos += ln + skip
		skip++
	}

	return l[0] * l[1]
}

func p2(input []byte) string {
	input = bytes.ReplaceAll(input, []byte(" "), []byte(""))

	lengths := make([]int, len(input), len(input)+5)
	for i, b := range input {
		lengths[i] = int(b)
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)

	l := make(list, 256)
	for i := range l {
		l[i] = i
	}

	pos := 0
	skip := 0

	for i := 0; i < 64; i++ {
		for _, ln := range lengths {
			l.reverse(pos, pos+ln-1)
			pos += ln + skip
			skip++
		}
	}

	dense := make([]byte, 16)
	for i := range dense {
		h := 0
		for _, v := range l[i*16 : (i+1)*16] {
			h ^= v
		}
		dense[i] = byte(h)
	}

	return fmt.Sprintf("%x", dense)
}

type list []int

func (l list) reverse(i, j int) {
	for ; j > i; i, j = i+1, j-1 {
		si, sj := l.makeSafe(i, j)
		l[si], l[sj] = l[sj], l[si]
	}
}

func (l list) makeSafe(i, j int) (safeI, safeJ int) {
	size := len(l)
	return (i + size) % size, (j + size) % size
}
