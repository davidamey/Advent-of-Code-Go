package main

import (
	"advent-of-code-go/2019/intcode"
	"advent-of-code-go/util"
	"fmt"
	"math"
	"strconv"
	"sync"
)

func main() {
	// data := util.MustReadCSInts("example")
	// data := util.MustReadCSInts("example2") // p2
	// data := util.MustReadCSInts("example3") // p2
	data := util.MustReadCSInts("input")

	fmt.Println("p1=", p1(data))
	fmt.Println("p2=", p2(data))
}

func p1(data []int) int {
	prog := intcode.New(data, false)
	maxE := math.MinInt16
	// var maxEperm []int
	for p := range util.NewIntPermuter([]int{0, 1, 2, 3, 4}).Permutations() {
		a := prog.Run(p[0], 0)[0]
		b := prog.Run(p[1], a)[0]
		c := prog.Run(p[2], b)[0]
		d := prog.Run(p[3], c)[0]
		e := prog.Run(p[4], d)[0]
		if e > maxE {
			maxE = e
			// maxEperm = p
		}
	}

	// fmt.Println(maxE, maxEperm)
	return maxE
}

func p2(data []int) (maxV int) {
	count := 5
	chanSize := 2

	prog := intcode.New(data, false)
	for p := range util.NewIntPermuter([]int{5, 6, 7, 8, 9}).Permutations() {
		in := make(chan int, chanSize)
		in <- p[0]
		in <- 0

		pipes := make([]chan int, count)
		for i := range pipes {
			pipe := make(chan int, chanSize)
			pipes[i] = pipe

			if i+1 < len(pipes) {
				pipe <- p[i+1]
			}
		}

		var wg sync.WaitGroup
		for i := range pipes {
			go func(i int) {
				if i == 0 {
					prog.RunBuf(strconv.Itoa(i), in, pipes[i])
				} else {
					prog.RunBuf(strconv.Itoa(i), pipes[i-1], pipes[i])
				}
				close(pipes[i])
			}(i)
		}

		lastV := 0
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range pipes[count-1] {
				lastV = v
				in <- v
			}
		}()
		wg.Wait()

		if lastV > maxV {
			maxV = lastV
		}
	}

	return maxV
}
