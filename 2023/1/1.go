package main

import (
	"advent-of-code-go/util"
	"fmt"
	"regexp"
	"time"
	"unicode"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	// lines := util.MustReadFileToLines("example2")
	lines := util.MustReadFileToLines("input")

	p1, p2 := 0, 0
	for _, l := range lines {
		p1 += recover(l)
		p2 += recoverP2(l)
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func recover(s string) int {
	var digits []int
	for _, r := range s {
		if unicode.IsDigit(r) {
			digits = append(digits, int(r-'0'))
		}
	}
	return 10*digits[0] + digits[len(digits)-1]
}

var rgx = regexp.MustCompile(`[1-9]|one|two|three|four|five|six|seven|eight|nine`)

func recoverP2(s string) int {
	leftMost := toDigit(rgx.FindString(s))

	for i := len(s) - 1; i >= 0; i-- {
		if m := rgx.FindString(s[i:]); m != "" {
			return 10*leftMost + toDigit(m)
		}
	}
	panic(fmt.Sprintf("Couldn't find two digits in %s", s))
}

func toDigit(s string) int {
	switch s {
	case "1", "one":
		return 1
	case "2", "two":
		return 2
	case "3", "three":
		return 3
	case "4", "four":
		return 4
	case "5", "five":
		return 5
	case "6", "six":
		return 6
	case "7", "seven":
		return 7
	case "8", "eight":
		return 8
	case "9", "nine":
		return 9
	}
	panic("invalid digit")
}
