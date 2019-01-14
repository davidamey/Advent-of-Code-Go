package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("== part1 ==")
	for _, key := range []string{"abcdef", "pqrstuv", "iwrupvqb"} {
		fmt.Printf("key %s gives answer: %d\n", key, SolveForKey(key, 5))
	}

	fmt.Println("== part2 ==")
	for _, key := range []string{"abcdef", "pqrstuv", "iwrupvqb"} {
		fmt.Printf("key %s gives answer: %d\n", key, SolveForKey(key, 6))
	}
}

func SolveForKey(key string, zeros int) int {
	i := 0
	for {
		hash := md5.Sum([]byte(key + strconv.Itoa(i)))
		hex := fmt.Sprintf("%x", hash)[:zeros]

		allZero := true
		for _, h := range hex {
			if h != '0' {
				allZero = false
				break
			}
		}

		if allZero {
			return i
		}

		i++
	}
}
