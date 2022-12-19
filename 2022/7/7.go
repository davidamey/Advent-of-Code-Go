package main

import (
	"advent-of-code-go/util"
	"fmt"
	"math"
	"strings"
	"time"
)

const (
	spaceTotal    = 70_000_000
	spaceRequired = 30_000_000
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	dirs := createFS(lines)

	spaceFree := spaceTotal - dirs[0].size
	toDelete := spaceRequired - spaceFree

	p1, p2 := 0, math.MaxInt
	for _, d := range dirs {
		if d.size <= 100_000 {
			p1 += d.size
		}

		if d.size >= toDelete && d.size < p2 {
			p2 = d.size
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func createFS(lines []string) (dirs []*item) {
	var cwd *item = nil
	for _, l := range lines {
		switch {
		case strings.HasPrefix(l, "$ cd .."):
			cwd = cwd.parent
		case strings.HasPrefix(l, "$ cd "):
			dir := &item{cwd, 0}
			dirs = append(dirs, dir)
			cwd = dir
		case l[0] >= '0' && l[0] <= '9':
			var size int
			fmt.Sscan(l, &size)
			for d := cwd; d != nil; d = d.parent {
				d.size += size
			}
		}
	}
	return
}

type item struct {
	parent *item
	size   int
}
