package util

import (
	"math"
	"math/big"
	"strconv"
)

func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("can't parse to int: " + s)
	}
	return i
}

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

func IntProduct(inputs ...int) (product int) {
	product = 1
	for _, i := range inputs {
		product *= i
	}
	return
}

func IntSum(inputs ...int) (sum int) {
	for _, i := range inputs {
		sum += i
	}
	return
}

func DigitsToInt(digits []int) (x int) {
	pow := 1
	for i := len(digits) - 1; i >= 0; i-- {
		x += digits[i] * pow
		pow *= 10
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

func IntToDigits(x int) (digits []int) {
	if x == 0 {
		return []int{}
	}
	return append(IntToDigits(x/10), x%10)
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

func Mod(x, m int) int {
	r := x % m
	if r < 0 {
		return r + m
	}
	return r
}

func ModInverse(x, mod int) int {
	bx := big.NewInt(int64(x))
	bx.ModInverse(bx, big.NewInt(int64(mod)))
	return int(bx.Int64())
}

func Pow(base, pow int) int {
	r := 1
	for pow > 0 {
		if pow&1 != 0 {
			r *= base
		}
		pow >>= 1
		base *= base
	}
	return r
}

func PowMod(base, pow, mod int) (r int) {
	base = base % mod
	r = 1
	for pow > 0 {
		if pow&1 != 0 {
			r = (r * base) % mod
		}
		pow >>= 1
		base *= base
	}
	return
}
