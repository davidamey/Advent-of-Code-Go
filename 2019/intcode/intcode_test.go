package intcode

import "testing"

func TestParseInstruct(t *testing.T) {
	r := &runner{memory: []int{1002}}
	oc := r.parseInstruct()

	if oc != 2 {
		t.Errorf("opcode should be 2, was %d", oc)
	}
	if r.modes != [4]int{0, 1, 0, 0} {
		t.Errorf("modes should be [0 1 0 0], was %v", r.modes)
	}
}

func TestRunSimple(t *testing.T) {
	r := &runner{memory: []int{1002, 4, 3, 4, 33}}

	oc := r.parseInstruct()
	if oc != 2 {
		t.Errorf("opcode should be 2, was %d", oc)
	}

	a, b := r.get(1), r.get(2)
	if a != 33 || b != 3 {
		t.Errorf("expected 33 and 3, got %d and %d", a, b)
	}

	r.set(3, a*b)
	if r.memory[4] != 99 {
		t.Errorf("memory[4] should be 99, got %d", r.memory[4])
	}

	r.pos += 4
	if c := r.current(); c != 99 {
		t.Errorf("expected pos to have moved and r.current() == 99, got %d", c)
	}
}

func TestRunDay5(t *testing.T) {
	prog := Program([]int{
		3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
	})

	if o := prog.Run(5)[0]; o != 999 {
		t.Errorf("expected run(5) == 999, got %d", o)
	}
	if o := prog.Run(8)[0]; o != 1000 {
		t.Errorf("expected run(8) == 1000, got %d", o)
	}
	if o := prog.Run(10)[0]; o != 1001 {
		t.Errorf("expected run(10) == 1001, got %d", o)
	}
}

func TestRunDay9Quine(t *testing.T) {
	instructs := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	prog := Program(instructs)

	out := prog.Run()
	for i, v := range out {
		if v != instructs[i] {
			t.Errorf("expected out[%d] == %d but got %d", i, instructs[i], v)
		}
	}
}
