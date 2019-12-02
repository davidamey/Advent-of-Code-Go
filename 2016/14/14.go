package main

import (
	"advent-of-code-go/util"
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

const (
	// salt = "abc" // example
	salt = "ahsbgdzn" // input
)

func main() {
	defer util.Duration(time.Now())

	fmt.Println("p1=", solve(false))
	fmt.Println("p2=", solve(true))
}

func solve(streched bool) int {
	hr := NewHasher(streched)
	keys := 0
	i := 0
	for ; keys < 64; i++ {
		t := hr.Hash(i).triple
		if t == 0x00 {
			continue
		}

		for j := 1; j <= 1000; j++ {
			if hr.Hash(i + j).HasQuintuple(t) {
				keys++
				break
			}
		}
	}
	return i - 1
}

type hash struct {
	hex        string
	triple     rune
	quintuples map[rune]struct{}
}

func NewHash(i int, stretched bool) *hash {
	h := md5.Sum([]byte(salt + strconv.Itoa(i)))
	hex := fmt.Sprintf("%x", h)
	if stretched {
		for i := 0; i < 2016; i++ {
			hex = fmt.Sprintf("%x", md5.Sum([]byte(hex)))
		}
	}
	triple, quintuples := count(hex)
	return &hash{hex, triple, quintuples}
}

func (h *hash) HasQuintuple(r rune) bool {
	_, ok := h.quintuples[r]
	return ok
}

func count(hex string) (triple rune, quintuples map[rune]struct{}) {
	quintuples = make(map[rune]struct{})
start:
	for i := 0; i <= len(hex)-3; i++ {
		for j := 1; j < 5; j++ {
			if i+j >= len(hex) || hex[i] != hex[i+j] {
				continue start
			}
			if j == 2 && triple == 0x00 {
				triple = rune(hex[i])
			}
		}
		quintuples[rune(hex[i])] = struct{}{}
	}
	return
}

type hasher struct {
	cache     map[int]*hash
	stretched bool
}

func NewHasher(stretched bool) *hasher {
	return &hasher{make(map[int]*hash), stretched}
}

func (hr *hasher) Hash(i int) *hash {
	if h, ok := hr.cache[i]; ok {
		return h
	}
	h := NewHash(i, hr.stretched)
	hr.cache[i] = h
	return h
}
