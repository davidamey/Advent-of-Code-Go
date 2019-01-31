package main

import "fmt"

func main() {
	p1()
	p2()
}

func p1() {
	a, b, c, d := 7, 0, 0, 0
	run(&a, &b, &c, &d)
	fmt.Println("p1=", a)
}

func p2() {
	fmt.Println("p2=", factorialPlus(12))
}

// ^_^
func run(a, b, c, d *int) {
	toggled := false

	*b = *a // cpy a b
	*b--    // dec b
D:
	*d = *a // cpy a d
	*a = 0  // cpy 0 a
B:
	*c = *b // cpy b c
A:
	*a++         // inc a
	*c--         // dec c
	if *c != 0 { // jnz c -2
		goto A
	}
	*d--         // dec d
	if *d != 0 { // jnz d -5
		goto B
	}
	*b--    // dec b
	*c = *b // cpy b c
	*d = *c // cpy c d
C:
	*d--         // dec d
	*c++         // inc c
	if *d != 0 { // jnz d -2
		goto C
	}
	switch *c { // tgl c (called with c = 10, 8, 6, 4, 2, 0)
	case 10:
		// out of bounds
	case 8, 6, 4:
		// toggle currently un-called instructions
	case 2:
		// change the `goto D` below into cpy
		toggled = true
	case 0:
		// the tgl would become `inc c` but this won't be called
		fmt.Println("uh oh, TGL called again")
		*c++
	}
	*c = -16      // cpy -16 c
	if !toggled { // jnz 1 c => cpy 1 c
		goto D
	} else {
		*c = 1
	}
	*c = 94 // cpy 94 c
F:
	*d = 80 // jnz 80 d => cpy 80 d
E:
	*a++         // inc a
	*d--         // inc d => dec d (tgl 6)
	if *d != 0 { // jnz d -2
		goto E
	}
	*c--         // inc c => dec c (tgl 8)
	if *c != 0 { // jnz c -5
		goto F
	}
}

// Copied p1 + gradually optimised until this loop was obvious
func factorialPlus(a int) int {
	for b := a - 1; b > 1; b-- {
		a = b * a
	}
	return a + 7520
}
