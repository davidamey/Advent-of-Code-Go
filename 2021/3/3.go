package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	// data := util.MustReadFileToLines("example")
	data := util.MustReadFileToLines("input")

	g, e := 0, 0
	for i := range data[0] {
		zeros, ones := counts(data, i)
		if ones > zeros {
			g = (g << 1) + 1
			e <<= 1
		} else {
			g <<= 1
			e = (e << 1) + 1
		}
	}
	fmt.Println("p1=", g*e)

	og := rating(data, func(zeros, ones int) byte {
		if ones >= zeros {
			return '1'
		} else {
			return '0'
		}
	})

	co := rating(data, func(zeros, ones int) byte {
		if ones < zeros {
			return '1'
		} else {
			return '0'
		}
	})

	fmt.Println("p2=", og*co)
}

func toInt(s string) (i int) {
	for _, b := range s {
		i <<= 1
		if b == '1' {
			i++
		}
	}
	return
}

func rating(data []string, f func(zeros, ones int) byte) int {
	filtered := make([]string, len(data))
	copy(filtered, data)
	for i := range data[0] {
		if len(filtered) == 1 {
			break
		}
		zeros, ones := counts(filtered, i)
		filtered = filter(filtered, i, f(zeros, ones))
	}
	return toInt(filtered[0])
}

func filter(in []string, pos int, b byte) (out []string) {
	for _, x := range in {
		if x[pos] == b {
			out = append(out, x)
		}
	}
	return
}

func counts(data []string, pos int) (zeros, ones int) {
	for _, d := range data {
		switch d[pos] {
		case '0':
			zeros++
		case '1':
			ones++
		}
	}
	return
}
