package knothash

import (
	"bytes"
	"fmt"
)

func Compute(input []byte) string {
	input = bytes.ReplaceAll(input, []byte(" "), []byte(""))

	lengths := make([]int, len(input), len(input)+5)
	for i, b := range input {
		lengths[i] = int(b)
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)

	l := make(List, 256)
	for i := range l {
		l[i] = i
	}

	pos := 0
	skip := 0

	for i := 0; i < 64; i++ {
		for _, ln := range lengths {
			l.Reverse(pos, pos+ln-1)
			pos += ln + skip
			skip++
		}
	}

	dense := make([]byte, 16)
	for i := range dense {
		h := 0
		for _, v := range l[i*16 : (i+1)*16] {
			h ^= v
		}
		dense[i] = byte(h)
	}

	return fmt.Sprintf("%x", dense)
}

type List []int

func (l List) Reverse(i, j int) {
	for ; j > i; i, j = i+1, j-1 {
		si, sj := l.makeSafe(i, j)
		l[si], l[sj] = l[sj], l[si]
	}
}

func (l List) makeSafe(i, j int) (safeI, safeJ int) {
	size := len(l)
	return (i + size) % size, (j + size) % size
}
