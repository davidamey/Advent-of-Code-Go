package main

import (
	"advent-of-code-go/util"
	"fmt"
	"math/big"
	"time"
)

const mod = 20201227

// example
// const cardPub = 5764801
// const doorPub = 17807724

// input
const cardPub = 19774466
const doorPub = 7290641

func main() {
	defer util.Duration(time.Now())

	cardLoop := 0
	for x := uint64(1); x != cardPub; cardLoop++ {
		x = (x * 7) % mod
	}

	fmt.Println("p1=", transform(doorPub, cardLoop))
	fmt.Println("p2=", "free!")
}

func transform(subject, loop int) int {
	s := big.NewInt(int64(subject))
	l := big.NewInt(int64(loop))
	s.Exp(s, l, big.NewInt(mod))
	return int(s.Int64())
}
