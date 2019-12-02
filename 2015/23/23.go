package main

import (
	"advent-of-code-go/util"
	"fmt"
)

func main() {
	// file, _ := util.OpenExample()
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	instructions := make([]instruction, len(lines))
	for i := range lines {
		name := lines[i][:3]
		instruct := instruction{name: name}
		switch name {
		case "hlf": // halve
			instruct.exec = func(r registers, i int) (newi int) {
				r[lines[i][4]] /= 2
				return i + 1
			}
		case "tpl": // triple
			instruct.exec = func(r registers, i int) (newi int) {
				r[lines[i][4]] *= 3
				return i + 1
			}
		case "inc": // increment
			instruct.exec = func(r registers, i int) (newi int) {
				r[lines[i][4]] += 1
				return i + 1
			}
		case "jmp": // jump
			instruct.exec = func(r registers, i int) (newi int) {
				var offset int
				if _, err := fmt.Sscanf(lines[i][4:], "%d", &offset); err != nil {
					panic(fmt.Errorf(`unable to parse line "%s"`, lines[i]))
				}
				return i + offset
			}
		case "jie": // jump if even
			instruct.exec = func(r registers, i int) (newi int) {
				if r[lines[i][4]]%2 == 1 {
					return i + 1
				}

				var offset int
				if _, err := fmt.Sscanf(lines[i][7:], "%d", &offset); err != nil {
					panic(fmt.Errorf(`unable to parse line "%s"`, lines[i]))
				}
				return i + offset
			}
		case "jio": // jump if one
			instruct.exec = func(r registers, i int) (newi int) {
				if r[lines[i][4]] != 1 {
					return i + 1
				}

				var offset int
				if _, err := fmt.Sscanf(lines[i][7:], "%d", &offset); err != nil {
					panic(fmt.Errorf(`unable to parse line "%s"`, lines[i]))
				}
				return i + offset
			}
		}

		instructions[i] = instruct
	}

	r := registers{'a': 0, 'b': 0}
	for i := 0; i >= 0 && i < len(instructions); {
		i = instructions[i].exec(r, i)
	}
	fmt.Println("p1=", r['b'])

	r = registers{'a': 1, 'b': 0}
	for i := 0; i >= 0 && i < len(instructions); {
		i = instructions[i].exec(r, i)
	}
	fmt.Println("p2=", r['b'])
}

type registers map[byte]int

func (r registers) dump() {
	fmt.Println("-- registers --")
	for i := 0; i < len(r); i++ {
		b := byte(i + 97)
		fmt.Printf("%c: %2d\n", b, r[b])
	}
	fmt.Println()
}

type instruction struct {
	name string
	exec func(r registers, i int) (newi int)
}
