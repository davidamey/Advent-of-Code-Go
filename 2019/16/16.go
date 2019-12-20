package main

import (
	"advent-of-code-go/util"
	"fmt"
	"time"
)

var phase = [4]int{0, 1, 0, -1}

func main() {
	defer util.Duration(time.Now())
	// raw := util.MustReadFile("example")
	raw := util.MustReadFile("input")

	ints := make([]int, len(raw))
	for i, b := range raw {
		ints[i] = (int)(b - '0')
	}
	fmt.Println("p1=", p1(ints))
	fmt.Println("p2=", p2(ints))
}

func p1(input []int) int {
	in := make([]int, len(input))
	copy(in, input)
	out := make([]int, len(in))
	for x := 0; x < 100; x++ {
		for o := range out {
			sum := 0
			for i, v := range in {
				p := phase[((i+1)/(o+1))%4]
				sum += v * p
			}
			out[o] = util.AbsInt(sum) % 10
		}
		copy(in, out)
	}

	return util.DigitsToInt(out[:8])
}

func p2(input []int) int {
	offset := util.DigitsToInt(input[:7])

	in := make([]int, 10000*len(input))
	for i := range in {
		in[i] = input[i%len(input)]
	}

	// can ignore first offset chars as phase = 0
	in = in[offset:]

	// process from the end
	reverse(in)

	for x := 0; x < 100; x++ {
		var out []int
		sum := 0
		for _, v := range in {
			sum += v
			out = append(out, util.AbsInt(sum)%10)
		}
		copy(in, out)
	}

	// undo earlier reverse and take the first 8
	return util.DigitsToInt(reverse(in)[:8])
}

func reverse(a []int) []int {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return a
}
