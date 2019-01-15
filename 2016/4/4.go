package main

import (
	"advent/util"
	"fmt"
	"strconv"
)

func main() {
	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	p1 := 0
	p2 := 0
	for _, l := range lines {
		name := l[:len(l)-11]
		sid, _ := strconv.Atoi(l[len(l)-10 : len(l)-7])
		cksm := l[len(l)-6 : len(l)-1]

		// fmt.Println(name, sid, cksm, checksum(name), decrypt(name, sid))
		if checksum(name) == cksm {
			p1 += sid
		}

		if decrypt(name, sid) == "northpole object storage" {
			p2 = sid
		}
	}
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func decrypt(name string, sid int) string {
	decrypted := make([]rune, len(name))
	for i, r := range name {
		if r == '-' {
			decrypted[i] = ' '
		} else {
			decrypted[i] = rune(((int(r-'a') + sid) % 26)) + 'a'
		}
	}
	return string(decrypted)
}

func checksum(name string) string {
	maxCount := 0
	runeToCount := make(map[rune]int)
	for _, r := range name {
		if r < 'a' || r > 'z' {
			continue
		}
		runeToCount[r]++
		if runeToCount[r] > maxCount {
			maxCount = runeToCount[r]
		}
	}

	countToRunes := make(map[int][]rune)
	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		countToRunes[runeToCount[r]] = append(countToRunes[runeToCount[r]], r)
	}

	cksm := make([]rune, 5)
	for i, c := 0, maxCount; i < len(cksm); c-- {
		for _, r := range countToRunes[c] {
			cksm[i] = r
			i++
			if i >= len(cksm) {
				break
			}
		}
	}

	return string(cksm)
}
