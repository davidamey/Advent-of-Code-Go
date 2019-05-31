package main

import (
	"advent/util"
	"fmt"
	"strconv"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	lines := util.MustReadFileToLines("input")

	fmt.Println("p1=", p1(lines))
	fmt.Println("p2=", p2(lines))
}

func p1(lines []string) (mulCount int) {
	registers := make(map[rune]int)

	for i := 0; i >= 0 && i < len(lines); i++ {
		instruct := lines[i][:3]
		r := rune(lines[i][4])
		v := readVal(registers, lines[i])

		switch instruct {
		case "set":
			registers[r] = v
		case "sub":
			registers[r] -= v
		case "mul":
			registers[r] *= v
			mulCount++
		case "jnz":
			doJump := registers[r] != 0
			// in the case of jnz, we might actually be conditional on a number not a reg, so check.
			if v, err := strconv.Atoi(lines[i][4:5]); err == nil {
				doJump = v != 0
			}
			if doJump {
				i += v
				i-- // counter loop `++`
			}
		}
	}
	return
}

func p2(lines []string) (nonPrimes int) {
	b := 57*100 + 100000

	for i := b; i <= b+17000; i += 17 {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				nonPrimes++
				break
			}
		}
	}

	// search:
	// 	for i := b; i <= b+17000; i += 17 {
	// 		for j := 2; j < i; j++ {
	// 			for k := 2; k < i; k++ {
	// 				if j*k == i {
	// 					nonPrimes++
	// 					continue search
	// 				}
	// 			}
	// 		}

	// 	}

	return
}

func readVal(registers map[rune]int, input string) int {
	const offset = 6
	if offset > len(input) {
		return -1
	}

	if x, err := strconv.Atoi(input[offset:]); err == nil {
		return x
	}

	return registers[rune(input[offset])]
}
