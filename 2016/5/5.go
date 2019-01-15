package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

const pwdLen = 8
const input = "uqwqemis"

func main() {
	eg1, eg2 := password("abc")
	fmt.Println("e.g. \"abc\"=", eg1, eg2) // example

	p1, p2 := password(input)
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func password(doorID string) (string, string) {
	p1 := make([]byte, pwdLen)
	p2 := make([]byte, pwdLen)

	i, j := 0, 0
	for idx := 0; i < pwdLen || j < pwdLen; idx++ {
		if h := hash(doorID, idx); h[:5] == "00000" {
			// p1
			if i < pwdLen {
				p1[i] = h[5]
				i++
			}

			// p2
			if j < pwdLen && h[5] >= '0' && h[5] <= '7' {
				if pos := int(h[5] - '0'); p2[pos] == 0 {
					p2[pos] = h[6]
					j++
				}
			}
		}
	}

	return string(p1), string(p2)
}

func hash(doorID string, idx int) string {
	hash := md5.Sum([]byte(doorID + strconv.Itoa(idx)))
	return fmt.Sprintf("%x\n", hash)
}
