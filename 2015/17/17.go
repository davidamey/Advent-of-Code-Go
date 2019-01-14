package main

import (
	"advent/util"
	"fmt"
)

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	containers, _ := util.ReadLinesToInts(file)

	// containers := []int{20, 15, 10, 5, 5}
	p1 := 0
	lengths := make(map[int]int)
	for c := range combos(containers) {
		if sum(c) == 150 {
			p1++
			lengths[len(c)]++
		}
	}
	fmt.Println("p1=", p1)
	for i := 1; i <= len(containers); i++ {
		if count, ok := lengths[i]; ok {
			fmt.Println("p2=", count)
			break
		}
	}
}

func combos(containers []int) <-chan []int {
	ch := make(chan []int)

	go func() {
		defer close(ch)
		clen := len(containers)
		choices := 1 << uint(clen)

		for i := 1; i < choices; i++ {
			choice := make([]int, 0, clen)
			for j := 0; j < clen; j++ {
				if (1<<uint(j))&i > 0 {
					choice = append(choice, containers[j])
				}
			}
			ch <- choice
		}
	}()

	return ch
}

func sum(containers []int) (sum int) {
	for _, c := range containers {
		sum += c
	}
	return
}
