package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	_, ex := fillDisk("10000", 20)
	fmt.Println("example=", ex)

	_, p1 := fillDisk("10111100110001111", 272)
	fmt.Println("p1=", p1)

	_, p2 := fillDisk("10111100110001111", 35651584)
	fmt.Println("p2=", p2)
}

func fillDisk(initial string, length int) (string, string) {
	data := initial
	for len(data) < length {
		data = dragon(data)
	}
	data = data[:length]
	return data, checksum(data)
}

func checksum(data string) (cksm string) {
	cksm = data
	for {
		var sb strings.Builder
		for i := 0; i < len(cksm)-1; i += 2 {
			if cksm[i] == cksm[i+1] {
				sb.WriteRune('1')
			} else {
				sb.WriteRune('0')
			}
		}
		cksm = sb.String()

		if len(cksm)%2 == 1 {
			break
		}
	}
	return
}

func dragon(in string) string {
	var sb strings.Builder
	sb.WriteString(in)
	sb.WriteRune('0')
	for i := len(in) - 1; i >= 0; i-- {
		if in[i] == '0' {
			sb.WriteRune('1')
		} else {
			sb.WriteRune('0')
		}
	}
	return sb.String()
}
