package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	opcodes := make(map[string]op)
	opcodes["addr"] = addr
	opcodes["addi"] = addi
	opcodes["mulr"] = mulr
	opcodes["muli"] = muli
	opcodes["banr"] = banr
	opcodes["bani"] = bani
	opcodes["borr"] = borr
	opcodes["bori"] = bori
	opcodes["setr"] = setr
	opcodes["seti"] = seti
	opcodes["gtir"] = gtir
	opcodes["gtri"] = gtri
	opcodes["gtrr"] = gtrr
	opcodes["eqir"] = eqir
	opcodes["eqri"] = eqri
	opcodes["eqrr"] = eqrr

	// part1(opcodes)
	part2(opcodes)
}

func part1(opcodes map[string]op) {
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	var ipReg int
	fmt.Sscanf(lines[0], "#ip %d", &ipReg)
	instructs := lines[1:]

	r := registers{}
	for ip := 0; ip >= 0 && ip < len(instructs); ip++ {
		var op string
		var a, b, c int
		fmt.Sscanf(instructs[ip], "%s %d %d %d", &op, &a, &b, &c)

		// fmt.Printf("ip=%d %v ", ip, r)
		// fmt.Printf("%s %d %d %d ", op, a, b, c)
		r[ipReg] = ip
		opcodes[op](&r, a, b, c)
		ip = r[ipReg]
		// fmt.Println(r)
	}

	fmt.Println("== part1 ==")
	fmt.Println("registers = ", r)
}

func part2(opcodes map[string]op) {
	// file, _ := util.OpenInput()
	// defer file.Close()
	// lines, _ := util.ReadLines(file)

	// var ipReg int
	// fmt.Sscanf(lines[0], "#ip %d", &ipReg)
	// instructs := lines[1:]

	// r := registers{}
	// r[0] = 1
	// for ip := 0; ip >= 0 && ip < len(instructs); ip++ {
	// 	var op string
	// 	var a, b, c int
	// 	fmt.Sscanf(instructs[ip], "%s %d %d %d", &op, &a, &b, &c)

	// 	fmt.Printf("ip=%d %v ", ip, r)
	// 	fmt.Printf("%s %d %d %d ", op, a, b, c)
	// 	r[ipReg] = ip
	// 	opcodes[op](&r, a, b, c)
	// 	ip = r[ipReg]
	// 	fmt.Println(r)
	// }

	fmt.Println("== part2 ==")
	// fmt.Println("registers = ", r)

	// The commented loop runs a very long time. From inspection, we're
	// trying to find the sum of factors of C.
	// C=10551394
	fmt.Println("answer=", 18200448)
}

type registers [6]int
type op func(r *registers, a, b, c int)

func addr(r *registers, a, b, c int) {
	r[c] = r[a] + r[b]
}

func addi(r *registers, a, b, c int) {
	r[c] = r[a] + b
}

func mulr(r *registers, a, b, c int) {
	r[c] = r[a] * r[b]
}

func muli(r *registers, a, b, c int) {
	r[c] = r[a] * b
}

func banr(r *registers, a, b, c int) {
	r[c] = r[a] & r[b]
}

func bani(r *registers, a, b, c int) {
	r[c] = r[a] & b
}

func borr(r *registers, a, b, c int) {
	r[c] = r[a] | r[b]
}

func bori(r *registers, a, b, c int) {
	r[c] = r[a] | b
}

func setr(r *registers, a, b, c int) {
	r[c] = r[a]
}

func seti(r *registers, a, b, c int) {
	r[c] = a
}

func gtir(r *registers, a, b, c int) {
	if a > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

func gtri(r *registers, a, b, c int) {
	if r[a] > b {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

func gtrr(r *registers, a, b, c int) {
	if r[a] > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

func eqir(r *registers, a, b, c int) {
	if a == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

func eqri(r *registers, a, b, c int) {
	if r[a] == b {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

func eqrr(r *registers, a, b, c int) {
	if r[a] == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}
