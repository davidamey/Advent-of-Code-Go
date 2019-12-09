package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	// img := imageFromFile("example", 3, 2)
	// img := imageFromFile("example2", 2, 2)
	img := imageFromFile("input", 25, 6)

	fmt.Println("p1=", img.checksum())
	fmt.Println("p2=")
	img.print()
}

type image struct {
	w, h   int
	layers [][]int
	counts []map[int]int
}

func (img *image) checksum() int {
	minZeroLayer := img.counts[0]
	for _, c := range img.counts[1:] {
		if c[0] < minZeroLayer[0] {
			minZeroLayer = c
		}
	}
	return minZeroLayer[1] * minZeroLayer[2]
}

func (img *image) print() {
	merged := img.layers[0]
	for _, l := range img.layers[1:] {
		for i, v := range l {
			if merged[i] == 2 {
				merged[i] = v
			}
		}
	}

	for y := 0; y < img.h; y++ {
		for x := 0; x < img.w; x++ {
			if merged[y*img.w+x] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func count(s []int) map[int]int {
	counts := make(map[int]int)
	for _, v := range s {
		counts[v]++
	}
	return counts
}

func imageFromFile(file string, w, h int) (img *image) {
	img = &image{w: w, h: h}
	raw := util.MustReadFile(file)

	var l []int
	counts := make(map[int]int)
	for i, b := range raw {
		v := (int)(b - '0')
		l = append(l, v)
		counts[v]++

		if (i+1)%(w*h) == 0 {
			img.layers = append(img.layers, l)
			img.counts = append(img.counts, counts)
			l = []int{}
			counts = make(map[int]int)
		}
	}
	return
}
