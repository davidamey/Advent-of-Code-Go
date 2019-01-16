package main

import (
	"advent/util"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	// examples p1
	// for _, s := range [][]byte{
	// 	[]byte("ADVENT"),
	// 	[]byte("A(1x5)BC"),
	// 	[]byte("(3x3)XYZ"),
	// 	[]byte("A(2x2)BCD(2x2)EFG"),
	// 	[]byte("(6x1)(1x3)A"),
	// 	[]byte("X(8x2)(3x3)ABCY"),
	// } {
	// 	d := decompress(s)
	// 	fmt.Printf("%s => %s (%d)\n", s, d, len(d))
	// }

	// examples p2
	// for _, s := range [][]byte{
	// 	[]byte("(3x3)XYZ"),
	// 	[]byte("X(8x2)(3x3)ABCY"),
	// 	[]byte("(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN"),
	// } {
	// 	fmt.Printf("%s => %d\n", s, fullLength(s))
	// }

	input := util.MustReadFile("input")
	fmt.Println("p1=", len(decompress(input)))
	fmt.Println("p2=", fullLength(input))
}

func fullLength(compressed []byte) (length int) {
	weights := make([]int, len(compressed))
	for i := range weights {
		weights[i] = 1
	}

	reader := bytes.NewReader(compressed)
	for pos := 0; ; {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		pos = getPos(reader)

		weight := weights[pos-1]

		if r != '(' {
			length += weight
			continue
		}

		var a, b int
		fmt.Fscanf(reader, "%dx%d)", &a, &b)
		pos = getPos(reader)

		for i := pos; i < pos+a && i < len(weights); i++ {
			weights[i] = weight * b
		}
	}

	return
}

func getPos(r *bytes.Reader) int {
	pos, err := r.Seek(0, os.SEEK_CUR)
	if err != nil {
		panic(err)
	}
	return int(pos)
}

func decompress(compressed []byte) string {
	out := make([]byte, 0, len(compressed))
	for i := 0; i < len(compressed); i++ {
		if b := compressed[i]; b != '(' {
			out = append(out, b)
			continue
		}

		j := i + 1
		var bytes int
		for s := make([]byte, 0, 2); ; j++ {
			if compressed[j] == 'x' {
				j++
				bytes, _ = strconv.Atoi(string(s))
				break
			}
			s = append(s, compressed[j])
		}

		var repeat int
		for s := make([]byte, 0, 2); ; j++ {
			if compressed[j] == ')' {
				j++
				repeat, _ = strconv.Atoi(string(s))
				break
			}
			s = append(s, compressed[j])
		}

		for k := 0; k < repeat; k++ {
			out = append(out, compressed[j:j+bytes]...)
		}

		i = j + bytes - 1
	}

	return string(out)
}
