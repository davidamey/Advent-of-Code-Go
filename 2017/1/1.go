package main

import (
	"advent/util"
	"fmt"
)

func main() {
	// for _, digits := range []string{
	// 	"1122",
	// 	"1111",
	// 	"1234",
	// 	"91212129",
	// } {
	// 	fmt.Println(digits, captcha(digits))
	// }

	digits := string(util.MustReadFile("input"))
	fmt.Println("p1=", captcha(digits, 1))
	fmt.Println("p2=", captcha(digits, len(digits)/2))
}

func captcha(digits string, offset int) (sum int) {
	for i, d := range digits {
		if digits[i] == digits[(i+offset)%len(digits)] {
			sum += int(d - '0')
		}
	}
	return
}
