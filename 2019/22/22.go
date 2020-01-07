package main

import (
	"advent-of-code-go/util"
	"fmt"
	"math/big"
	"strings"
)

func main() {
	// d, shuffles := newDeck(10), util.MustReadFileToLines("example")
	d, shuffles := newDeck(10007), util.MustReadFileToLines("input")

	fmt.Println("p1=", p1(d, shuffles))
	fmt.Println("p2=", p2(shuffles))
}

func p1(d deck, shuffles []string) int {
	for _, s := range shuffles {
		switch {
		case strings.HasPrefix(s, "deal into"):
			d = dealNew(d)
		case strings.HasPrefix(s, "deal with"):
			var n int
			fmt.Sscanf(s, "deal with increment %d", &n)
			d = dealInc(d, n)
		case strings.HasPrefix(s, "cut"):
			var n int
			fmt.Sscanf(s, "cut %d", &n)
			d = cut(d, n)
		default:
			panic("unknown shuffle")
		}
	}
	for i, c := range d {
		if c == 2019 {
			return i
		}
	}
	return -1
}

func p2(shuffles []string) int {
	deckSize := big.NewInt(119315717514047)
	repeats := big.NewInt(101741582076661)
	card := big.NewInt(2020)

	one := big.NewInt(1)
	a, b := big.NewInt(1), big.NewInt(0)

	// run once
	for _, s := range shuffles {
		switch {
		case strings.HasPrefix(s, "deal into"):
			a.Neg(a).Mod(a, deckSize)
			b.Neg(b).Sub(b, one).Mod(b, deckSize)
		case strings.HasPrefix(s, "deal with"):
			var n int64
			fmt.Sscanf(s, "deal with increment %d", &n)
			bn := big.NewInt(n)
			a.Mul(a, bn).Mod(a, deckSize)
			b.Mul(b, bn).Mod(b, deckSize)
		case strings.HasPrefix(s, "cut"):
			var n int64
			fmt.Sscanf(s, "cut %d", &n)
			b.Sub(b, big.NewInt(n)).Mod(b, deckSize)
		default:
			panic("unknown shuffle")
		}
	}

	// f(i)   = ai + b
	// f^2(i) = a(ai+b) + b = (a^2)i + ab + b
	// f^3(i) = a(a^2 i + ab + b) + b = (a^3)i + (a^2)b + b
	// ...
	// f^n(i) = a^n i + [a^(n-1) + a^(n-2) + ... + 1] * b
	// f^n(i) = a^n i + [(a^n - 1) / (a - 1)] * b
	// =>
	// f^n(i) = xi + y
	// where
	//   x = a^n
	//   y = [(a^n - 1) / (a - 1)] * b

	// i = the original card pos
	// n = times repeated
	// all modulo deckSize

	x := new(big.Int).Exp(a, repeats, deckSize)
	xInv := new(big.Int).ModInverse(x, deckSize)

	y := new(big.Int).Sub(x, one)
	y = y.Mul(y, b).Mul(y, new(big.Int).ModInverse(a.Sub(a, one), deckSize))

	// i = (f^n(i) - y)/x
	r := card
	r = r.Sub(r, y).Mul(r, xInv).Mod(r, deckSize)
	return int(r.Int64())
}

type deck []int

func newDeck(n int) (d deck) {
	d = make(deck, n)
	for i := range d {
		d[i] = i
	}
	return
}

func cut(d deck, n int) (d2 deck) {
	l := len(d)
	if n < 0 {
		n += l
	}
	d2 = make(deck, 0, l)
	d2 = append(d2, d[n:]...)
	d2 = append(d2, d[:n]...)
	return
}

func dealInc(d deck, n int) (d2 deck) {
	l := len(d)
	d2 = make(deck, l)
	for i, c := range d {
		d2[i*n%l] = c
	}
	return
}

func dealNew(d deck) deck {
	for left, right := 0, len(d)-1; left < right; left, right = left+1, right-1 {
		d[left], d[right] = d[right], d[left]
	}
	return d
}
