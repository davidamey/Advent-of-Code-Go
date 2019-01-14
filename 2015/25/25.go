package main

import "fmt"

const FirstCode = 20151125

func main() {
	code := FirstCode
	for i := 2; i <= codeNoAt(2947, 3029); i++ {
		code = nextCode(code)
	}
	fmt.Println("p1=", code)
}

func codeNoAt(row, col int) int {
	i := col + row - 1
	n := i * (i + 1) / 2 // triangle #
	n -= (row - 1)
	return n
}

func nextCode(code int) int {
	return code * 252533 % 33554393
}
