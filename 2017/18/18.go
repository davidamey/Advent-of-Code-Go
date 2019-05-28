package main

import (
	"advent/util"
	"fmt"
	"strconv"
)

func main() {
	// lines := util.MustReadFileToLines("example")
	// lines := util.MustReadFileToLines("exampleP2")
	lines := util.MustReadFileToLines("input")

	p1(lines)
	p2(lines)
}

func p1(lines []string) {
	registers := make(map[rune]int)
	freq := -1
run:
	for i := 0; i >= 0 && i < len(lines); i++ {
		instruct := lines[i][:3]
		r := rune(lines[i][4])
		v := readVal(registers, lines[i])

		switch instruct {
		case "snd":
			freq = registers[r]
		case "set":
			registers[r] = v
		case "add":
			registers[r] += v
		case "mul":
			registers[r] *= v
		case "mod":
			registers[r] %= v
		case "rcv":
			if registers[r] != 0 {
				break run
			}
		case "jgz":
			if registers[r] > 0 {
				i += v
				i-- // counter loop `++`
			}
		}
	}

	fmt.Println("p1=", freq)
}

func p2(lines []string) {
	p0, p1 := newProg(0), newProg(1)
	quit := make(chan struct{})
	go p0.run(lines, quit)
	go p1.run(lines, quit)

loop:
	for {
		select {
		case v := <-p0.out:
			p1.in <- v
		case v := <-p1.out:
			p0.in <- v
		default:
			if p0.waiting && p1.waiting && len(p0.in) == 0 && len(p1.in) == 0 {
				close(quit)
				break loop
			}
		}
	}

	fmt.Println("p2=", p1.sendCount)
}

type prog struct {
	id        int
	registers map[rune]int
	in, out   chan int
	waiting   bool
	sendCount int
}

func newProg(id int) *prog {
	return &prog{
		id:  id,
		in:  make(chan int, 100),
		out: make(chan int, 100),
	}
}

func (p *prog) run(lines []string, quit chan struct{}) {
	registers := make(map[rune]int)
	registers['p'] = p.id

run:
	for i := 0; i >= 0 && i < len(lines); i++ {
		instruct := lines[i][:3]
		r := rune(lines[i][4])
		v := readVal(registers, lines[i])

		switch instruct {
		case "snd":
			p.sendCount++
			p.out <- registers[r]
		case "set":
			registers[r] = v
		case "add":
			registers[r] += v
		case "mul":
			registers[r] *= v
		case "mod":
			registers[r] %= v
		case "rcv":
			for {
				select {
				case <-quit:
					return
				case v := <-p.in:
					registers[r] = v
					p.waiting = false
					continue run
				default:
					p.waiting = true
				}
			}
		case "jgz":
			doJump := registers[r] > 0
			// in the case of jgz, we might actually be conditional on a number not a reg, so check.
			if v, err := strconv.Atoi(lines[i][4:5]); err == nil {
				doJump = v > 0
			}
			if doJump {
				i += v
				i-- // counter loop `++`
			}
		}
	}
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
