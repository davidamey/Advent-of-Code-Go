package intcode

import "fmt"

type Interpreter struct {
	// id      string
	debug   bool
	program []int
}

func New(prog []int, debug bool) *Interpreter {
	return &Interpreter{
		program: append([]int(nil), prog...),
		debug:   debug,
	}
}

func (itp *Interpreter) log(a ...interface{}) {
	if itp.debug {
		fmt.Println(a...)
	}
}

func (itp *Interpreter) Run(input ...int) (result []int) {
	in := make(chan int, len(input))
	out := make(chan int)
	for _, i := range input {
		in <- i
	}
	go func() {
		itp.RunBuf("itp", in, out)
		close(in)
		close(out)
	}()
	for o := range out {
		result = append(result, o)
	}
	return
}

func (itp *Interpreter) RunBuf(id string, in, out chan int) (code int) {
	itp.log(id, "starting")
	mem := append([]int(nil), itp.program...)

	for i := 0; mem[i] != 99; {
		oc, modes := parseInstruct(mem[i])

		switch oc {
		case 1:
			mem[mem[i+3]] = getVal(mem, i+1, modes[0]) + getVal(mem, i+2, modes[1])
			i += 4
		case 2:
			mem[mem[i+3]] = getVal(mem, i+1, modes[0]) * getVal(mem, i+2, modes[1])
			i += 4
		case 3:
			v := <-in
			itp.log(id, "in:", v)
			mem[mem[i+1]] = v
			i += 2
		case 4:
			v := getVal(mem, i+1, modes[0])
			itp.log(id, "out:", v)
			out <- v
			i += 2
		case 5:
			if getVal(mem, i+1, modes[0]) != 0 {
				i = getVal(mem, i+2, modes[1])
			} else {
				i += 3
			}
		case 6:
			if getVal(mem, i+1, modes[0]) == 0 {
				i = getVal(mem, i+2, modes[1])
			} else {
				i += 3
			}
		case 7:
			if getVal(mem, i+1, modes[0]) < getVal(mem, i+2, modes[1]) {
				mem[mem[i+3]] = 1
			} else {
				mem[mem[i+3]] = 0
			}
			i += 4
		case 8:
			if getVal(mem, i+1, modes[0]) == getVal(mem, i+2, modes[1]) {
				mem[mem[i+3]] = 1
			} else {
				mem[mem[i+3]] = 0
			}
			i += 4
		default:
			panic("unknown opcode")
		}
	}

	itp.log(id, "exiting", mem[0])
	return mem[0]
}

func parseInstruct(inst int) (oc int, modes []int) {
	oc = inst % 100 // last 2 digits
	inst /= 100
	for i := 0; i < 4; i++ {
		modes = append(modes, inst%10)
		inst /= 10
	}
	return
}

func getVal(mem []int, ref, mode int) int {
	if mode == 1 {
		return mem[ref]
	}
	return mem[mem[ref]]
}
