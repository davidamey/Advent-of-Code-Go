package intcode

type Program []int

func (p Program) Run(input ...int) (result []int) {
	in := make(chan int, len(input))
	out := make(chan int)
	for _, i := range input {
		in <- i
	}
	close(in)
	go p.RunBuf("c", in, out)
	for o := range out {
		result = append(result, o)
	}
	return
}

func (p Program) RunBuf(id string, in <-chan int, out chan<- int) {
	r := newRunner(p)
	for r.current() != 99 {
		switch r.parseInstruct() {
		case 1:
			r.set(3, r.get(1)+r.get(2))
			r.pos += 4
		case 2:
			r.set(3, r.get(1)*r.get(2))
			r.pos += 4
		case 3:
			r.set(1, <-in)
			r.pos += 2
		case 4:
			out <- r.get(1)
			r.pos += 2
		case 5:
			if r.get(1) != 0 {
				r.pos = r.get(2)
			} else {
				r.pos += 3
			}
		case 6:
			if r.get(1) == 0 {
				r.pos = r.get(2)
			} else {
				r.pos += 3
			}
		case 7:
			if r.get(1) < r.get(2) {
				r.set(3, 1)
			} else {
				r.set(3, 0)
			}
			r.pos += 4
		case 8:
			if r.get(1) == r.get(2) {
				r.set(3, 1)
			} else {
				r.set(3, 0)
			}
			r.pos += 4
		case 9:
			r.relBase += r.get(1)
			r.pos += 2
		default:
			panic("unknown opcode")
		}
	}
	close(out)
}

type runner struct {
	memory  []int
	relBase int
	pos     int
	modes   [4]int
}

func newRunner(p Program) *runner {
	r := &runner{
		memory: make([]int, 10000),
	}
	copy(r.memory, p)
	return r
}

func (r *runner) current() int {
	return r.memory[r.pos]
}

func (r *runner) parseInstruct() (oc int) {
	inst := r.current()
	oc = inst % 100 // last 2 digits
	inst /= 100
	for i := 0; i < 4; i++ {
		r.modes[i] = inst % 10
		inst /= 10
	}
	return
}

func (r *runner) get(i int) int {
	return r.memory[r.getRef(i)]
}

func (r *runner) set(i, val int) {
	r.memory[r.getRef(i)] = val
}

func (r *runner) getRef(i int) int {
	switch r.modes[i-1] {
	case 0: // position
		return r.memory[r.pos+i]
	case 1: // immediate
		return r.pos + i
	case 2: // relative
		return r.relBase + r.memory[r.pos+i]
	default:
		panic("unknown mode")
	}
}
