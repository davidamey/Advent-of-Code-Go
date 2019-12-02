package main

import (
	"advent-of-code-go/util"
	"fmt"
	"regexp"
)

func main() {
	// lines := util.MustReadFileToLines("example")  // p1
	// lines := util.MustReadFileToLines("example2") // p2
	lines := util.MustReadFileToLines("input")

	// fmt.Println(hasTLS("ababbacd[efgh]ijkl[mnop]qrst"))
	// fmt.Println(hasSSL("ababbacd[efgh]ijkl[mnbabop]qrst"))

	p1, p2 := 0, 0
	for _, l := range lines {
		if hasTLS(l) {
			p1++
		}
		if hasSSL(l) {
			p2++
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func hasTLS(ip string) bool {
	hasABBA := false
	hn := false
	for i := 0; i < len(ip)-3; i++ {
		switch {
		case ip[i] == '[':
			hn = true
			continue
		case ip[i] == ']':
			hn = false
			continue
		case ip[i+1] == '[':
			hn = true
			i += 1
			continue
		case ip[i+1] == ']':
			hn = false
			i += 1
			continue
		case ip[i+2] == '[':
			hn = true
			i += 2
			continue
		case ip[i+2] == ']':
			hn = false
			i += 2
			continue
		case ip[i+3] == '[':
			hn = true
			i += 3
			continue
		case ip[i+3] == ']':
			hn = false
			i += 3
			continue
		}

		// check ABBA
		if ip[i] == ip[i+3] && ip[i+1] == ip[i+2] && ip[i] != ip[i+1] {
			if hn {
				return false
			}
			hasABBA = true
		}
	}

	// fmt.Println(ip, b1, hn, b2)
	// return (hasABBA(b1) || hasABBA(b2)) && !hasABBA(hn)
	return hasABBA
}

func hasSSL(ip string) bool {
	rgx := regexp.MustCompile(`\[|\]`)
	var bases, HNs []string
	for i, p := range rgx.Split(ip, -1) {
		if i%2 == 0 {
			bases = append(bases, p)
		} else {
			HNs = append(HNs, p)
		}
	}

	ABAs := make(map[string]struct{})
	for _, b := range bases {
		for i := 0; i < len(b)-2; i++ {
			if b[i] != b[i+1] && b[i] == b[i+2] {
				ABAs[string(b[i:i+3])] = struct{}{}
			}
		}
	}

	for _, hn := range HNs {
		for i := 0; i < len(hn)-2; i++ {
			if hn[i] != hn[i+1] && hn[i] == hn[i+2] {
				aba := string([]byte{hn[i+1], hn[i], hn[i+1]})
				if _, ok := ABAs[aba]; ok {
					return true
				}
			}
		}
	}

	return false
}
