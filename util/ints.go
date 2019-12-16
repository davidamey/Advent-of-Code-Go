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

func IntToDigits(x int) (digits []int) {
	if x == 0 {
		return []int{}
	}
	return append(IntToDigits(x/10), x%10)
}

func DigitsToInt(digits []int) (x int) {
	pow := 1
	for i := len(digits) - 1; i >= 0; i-- {
		x += digits[i] * pow
		pow *= 10
	}
	return
}

func Pow(a, b int) (result int) {
	result = 1
	for 0 != b {
		if 0 != (b & 1) {
			result *= a
		}
		b >>= 1
		a *= a
	}
	return
}

func GCD(a, b int) int {
	d := uint(0)

	for a%2 == 0 && b%2 == 0 {
		a /= 2
		b /= 2
		d++
	}

	for a != b {
		if a%2 == 0 {
			a /= 2
		} else if b%2 == 0 {
			b /= 2
		} else if a > b {
			a = (a - b) / 2
		} else {
			b = (b - a) / 2
		}
	}

	return a * (1 << d)
}

func LCM(nums ...int) (m int) {
	m = 1
	for _, n := range nums {
		m = lcm(m, n)
	}
	return
}

func lcm(a, b int) int {
	// Theoretically more efficient to divide first
	return (a / GCD(a, b)) * b
}
