package util

import "math"

func MinInt(inputs ...int) int {
	min := math.MaxInt32
	for _, i := range inputs {
		if i < min {
			min = i
		}
	}
	return min
}

func MaxInt(inputs ...int) int {
	max := -math.MaxInt32
	for _, i := range inputs {
		if i > max {
			max = i
		}
	}
	return max
}

func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
