package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// // examples
	// for _, s := range [][]string{
	// 	[]string{"1", "11"},
	// 	[]string{"11", "21"},
	// 	[]string{"21", "1211"},
	// 	[]string{"1211", "111221"},
	// 	[]string{"111221", "312211"},
	// } {
	// 	fmt.Println(s[0], LookAndSay(s[0]), LookAndSay(s[0]) == s[1])
	// }

	p1 := "3113322113"
	for i := 0; i < 40; i++ {
		p1 = LookAndSay(p1)
	}
	fmt.Println("p1", len(p1))

	p2 := "3113322113"
	for i := 0; i < 50; i++ {
		p2 = LookAndSay(p2)
	}
	fmt.Println("p2", len(p2))
}

func LookAndSay(in string) string {
	digits := strings.Split(in, "")

	var out []string
	for i := 0; i < len(digits); {
		d := digits[i]
		count := 1
		for i+count < len(digits) && d == digits[i+count] {
			count++
		}
		out = append(out, strconv.Itoa(count), d)
		i = i + count
	}

	return strings.Join(out, "")
}
