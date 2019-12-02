package main

import (
	"advent-of-code-go/util"
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	opcodes := make(map[int]op)
	opcodes[0] = addr
	opcodes[1] = addi
	opcodes[2] = mulr
	opcodes[3] = muli
	opcodes[4] = banr
	opcodes[5] = bani
	opcodes[6] = borr
	opcodes[7] = bori
	opcodes[8] = setr
	opcodes[9] = seti
	opcodes[10] = gtir
	opcodes[11] = gtri
	opcodes[12] = gtrr
	opcodes[13] = eqir
	opcodes[14] = eqri
	opcodes[15] = eqrr

	part1(opcodes)
	part2(opcodes)
}

func part1(opcodes map[int]op) {
	file, _ := util.OpenFile("samples")
	defer file.Close()
	lines, _ := util.ReadLines(file)
	testcases := readTestCases(lines)

	moreThanThree := 0
	for _, tc := range testcases {
		count := 0
		for _, oc := range opcodes {
			r := registers{}
			copy(r[:], tc.before[:])
			oc(&r, tc.instruct[1], tc.instruct[2], tc.instruct[3])
			// fmt.Println(tc.before, r, tc.after)
			if r == tc.after {
				count++
			}
		}
		if count >= 3 {
			moreThanThree++
		}
	}

	fmt.Println("== part1 ==")
	fmt.Printf("%d samples match >= 3 opcodes\n", moreThanThree)
}

func part2(opcodes map[int]op) {
	samplesFile, _ := util.OpenFile("samples")
	defer samplesFile.Close()
	samples, _ := util.ReadLines(samplesFile)
	testcases := readTestCases(samples)

	// map of opcode -> instruction code
	ocref := make(map[int]int)
	for len(ocref) < len(opcodes) {
		for _, tc := range testcases {
			count := 0
			matched := -1
			for i, oc := range opcodes {
				if _, ok := ocref[i]; ok {
					continue
				}

				r := registers{}
				copy(r[:], tc.before[:])
				oc(&r, tc.instruct[1], tc.instruct[2], tc.instruct[3])
				if r == tc.after {
					count++
					matched = i
				}
			}
			if count == 1 {
				// fmt.Println("single match", tc.instruct[0], matched)
				ocref[matched] = tc.instruct[0]
			}
		}
	}

	// Need a reverse oc map
	instructref := make(map[int]int)
	for oc, inst := range ocref {
		instructref[inst] = oc
	}

	// Load program
	programFile, _ := util.OpenFile("program")
	defer programFile.Close()
	program, _ := util.ReadLines(programFile)

	// execute
	r := registers{}
	for _, l := range program {
		var i, a, b, c int
		fmt.Sscanf(l, "%d %d %d %d", &i, &a, &b, &c)
		opcodes[instructref[i]](&r, a, b, c)
	}

	fmt.Println("== part2 ==")
	fmt.Println("r=", r)
}

func readTestCases(lines []string) []testcase {
	var testcases []testcase
	for i := 0; i < len(lines); i += 4 {
		tc := testcase{}
		fmt.Sscanf(lines[i], "Before: [%d, %d, %d, %d]", &tc.before[0], &tc.before[1], &tc.before[2], &tc.before[3])
		fmt.Sscanf(lines[i+1], "%d %d %d %d", &tc.instruct[0], &tc.instruct[1], &tc.instruct[2], &tc.instruct[3])
		fmt.Sscanf(lines[i+2], "After: [%d, %d, %d, %d]", &tc.after[0], &tc.after[1], &tc.after[2], &tc.after[3])
		testcases = append(testcases, tc)
	}
	return testcases
}

type testcase struct {
	before, after registers
	instruct      [4]int
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

type registers [4]int
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
