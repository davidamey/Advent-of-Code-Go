package assembunny

import (
	"fmt"
	"strconv"
)

type inputType uint8

const (
	inputTypeReg inputType = iota
	inputTypeVal
)

type Registers [4]int

func (r Registers) String() string {
	return fmt.Sprintf("[%d, %d, %d, %d]", r[0], r[1], r[2], r[3])
}

type Program struct {
	instructs []Instruction
	Out       []int
}

type Instruction func(r *Registers, i int) (newi int)

func Compile(lines []string) *Program {
	p := &Program{
		instructs: make([]Instruction, len(lines)),
	}
	for i := range lines {
		switch lines[i][:3] {
		case "cpy":
			p.instructs[i] = compileCPY(lines[i][4:])
		case "inc":
			p.instructs[i] = compileINC(lines[i][4:])
		case "dec":
			p.instructs[i] = compileDEC(lines[i][4:])
		case "jnz":
			p.instructs[i] = compileJNZ(lines[i][4:])
		case "out":
			p.instructs[i] = compileOUT(lines[i][4:], p)
		}
	}
	return p
}

func (p *Program) Run(r *Registers, debug bool) {
	for i := 0; i >= 0 && i < len(p.instructs); {
		if debug {
			fmt.Printf("i: %2d, registers: %s\n", i, r)
		}
		i = p.instructs[i](r, i)
	}
}

func (p *Program) RunWithOut(r *Registers, debug bool, maxOut int) {
	p.Out = make([]int, 0, maxOut)
	for i := 0; i >= 0 && i < len(p.instructs); {
		if debug {
			fmt.Printf("i: %2d, registers: %s\n", i, r)
		}
		i = p.instructs[i](r, i)
		if len(p.Out) >= maxOut {
			return
		}
	}
}

func compileCPY(raw string) Instruction {
	var x, y string
	fmt.Sscan(raw, &x, &y)

	yr := toRegister(y)

	if t, z := parseInput(x); t == inputTypeVal {
		return func(r *Registers, i int) (newi int) {
			r[yr] = z
			return i + 1
		}
	} else {
		return func(r *Registers, i int) (newi int) {
			r[yr] = r[z]
			return i + 1
		}
	}
}

func compileINC(raw string) Instruction {
	x := toRegister(raw)
	return func(r *Registers, i int) (newi int) {
		r[x]++
		return i + 1
	}
}

func compileDEC(raw string) Instruction {
	x := toRegister(raw)
	return func(r *Registers, i int) (newi int) {
		r[x]--
		return i + 1
	}
}

func compileJNZ(raw string) Instruction {
	var x string
	var y int
	fmt.Sscan(raw, &x, &y)

	if t, z := parseInput(x); t == inputTypeVal {
		return func(r *Registers, i int) (newi int) {
			if z != 0 {
				return i + y
			}
			return i + 1
		}
	} else {
		return func(r *Registers, i int) (newi int) {
			if r[z] != 0 {
				return i + y
			}
			return i + 1
		}
	}
}

func compileOUT(raw string, p *Program) Instruction {
	if t, x := parseInput(raw); t == inputTypeVal {
		return func(r *Registers, i int) (newi int) {
			if p.Out != nil {
				p.Out = append(p.Out, x)
			}
			return i + 1
		}
	} else {
		return func(r *Registers, i int) (newi int) {
			if p.Out != nil {
				p.Out = append(p.Out, r[x])
			}
			return i + 1
		}
	}
}

func parseInput(x string) (inputType, int) {
	if i, err := strconv.Atoi(x); err == nil {
		return inputTypeVal, i
	}

	return inputTypeReg, toRegister(x)
}

func toRegister(x string) int {
	if len(x) > 1 {
		panic(fmt.Errorf("invalid register: \"%s\"", x))
	}
	y := x[0]
	if y < 'a' && y > 'd' {
		panic(fmt.Errorf("invalid register \"%c\"", y))
	}
	return int(y - 'a')
}
