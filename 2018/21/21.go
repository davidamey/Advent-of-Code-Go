package main

import (
	"advent/util"
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

	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	var ipReg int
	fmt.Sscanf(lines[0], "#ip %d", &ipReg)

	// We actually care about performance now so best not Sscanf for each instruct!
	var instructs []instruct
	for _, l := range lines[1:] {
		var op string
		var a, b, c int
		fmt.Sscanf(l, "%s %d %d %d", &op, &a, &b, &c)
		instructs = append(instructs, instruct{opcodes[op], a, b, c})
	}

	r := registers{}
	r3s := make(map[int]bool)
	firstR3 := -1
	lastR3 := -1
	for ip := 0; ip >= 0 && ip < len(instructs); ip++ {
		if ip == 28 {
			if firstR3 == -1 {
				firstR3 = r[3]
			}
			if _, ok := r3s[r[3]]; ok {
				break
			}
			r3s[r[3]] = true
			lastR3 = r[3]
		}

		r[ipReg] = ip
		i := instructs[ip]
		i.op(&r, i.a, i.b, i.c)
		ip = r[ipReg]
	}

	fmt.Println("part1=", firstR3)
	fmt.Println("part2=", lastR3)
}

type registers [6]int
type op func(r *registers, a, b, c int)
type instruct struct {
	op      op
	a, b, c int
}

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
