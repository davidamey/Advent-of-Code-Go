package main

import (
	"advent-of-code-go/2019/intcode"
	"advent-of-code-go/util"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// solution found by manually playing:
// south
// take space heater
// east
// north
// west
// south
// take antenna
// north
// east
// south
// east
// north
// take klein bottle
// north
// take spool of cat6
// west
// north

func main() {
	prog := intcode.Program(util.MustReadCSInts("input"))
	// manual(prog)
	auto(prog)
}

func manual(prog intcode.Program) {
	in, out := make(chan int), make(chan int)
	go prog.RunBuf(in, out)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			for _, b := range s.Bytes() {
				in <- int(b)
			}
			in <- 10
		}
	}()

	for o := range out {
		fmt.Print(string(o))
	}
}

func auto(prog intcode.Program) {
	in, out := make(chan int), make(chan int)
	go prog.RunBuf(in, out)

	go func() {
		solution := []string{
			"south", "take space heater",
			"east", "north", "west", "south", "take antenna",
			"north", "east", "south", "east", "north", "take klein bottle",
			"north", "take spool of cat6",
			"west", "north",
		}
		for _, s := range solution {
			for _, b := range s {
				in <- int(b)
			}
			in <- 10
		}
	}()

	var bs []byte
	for o := range out {
		bs = append(bs, byte(o))
	}

	rgx := regexp.MustCompile(`typing (\d+) on the keypad`)
	if r := rgx.FindSubmatch(bs); len(r) > 0 {
		fmt.Println("passcode=", string(r[1]))
		return
	}
}
