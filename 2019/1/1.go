package main

import (
	"fmt"
	"local/advent/util"
)

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	modules, _ := util.ReadLinesToInts(file)

	fmt.Println("p1=", p1(modules))
	fmt.Println("p2=", p2(modules))
}

func p1(modules []int) (fuel int) {
	for _, m := range modules {
		fuel += m/3 - 2
	}
	return
}

func p2(modules []int) (fuel int) {
	for _, m := range modules {
		fuel += p2m(m)
	}
	return
}

func p2m(m int) int {
	f := m/3 - 2
	if f > 0 {
		return f + p2m(f)
	}
	return 0
}
