package main

import (
	"advent-of-code-go/util"
	"bytes"
	"fmt"
)

func main() {
	// examples := []string{
	// 	"{}",
	// 	"{{{}}}",
	// 	"{{},{}}",
	// 	"{{{},{},{{}}}}",
	// 	"{<a>,<a>,<a>,<a>}",
	// 	"{{<ab>},{<ab>},{<ab>},{<ab>}}",
	// 	"{{<!!>},{<!!>},{<!!>},{<!!>}}",
	// 	"{{<a!>},{<a!>},{<a!>},{<ab>}}",
	// }
	// for _, ex := range examples {
	// 	p1, p2 := parseStream([]byte(ex))
	// 	fmt.Printf("%s => %d, %d\n", ex, p1, p2)
	// }

	input := util.MustReadFile("input")
	p1, p2 := parseStream(input)
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func parseStream(raw []byte) (totalScore, garbageCount int) {
	r := bytes.NewReader(raw)
	// streams should start with a group
	if ch, _, _ := r.ReadRune(); ch != '{' {
		panic("invalid input")
	}
	return parseGroup(r, 1)
}

func parseGroup(r *bytes.Reader, score int) (totalScore, garbageCount int) {
	totalScore = score
	for {
		ch, _, err := r.ReadRune()
		if err != nil {
			panic(err)
		}

		switch ch {
		case '{':
			s, gc := parseGroup(r, score+1)
			totalScore += s
			garbageCount += gc
		case '<':
			garbageCount += parseGarbage(r)
		case '}':
			return
		}
	}
}

func parseGarbage(r *bytes.Reader) (count int) {
	for {
		ch, _, err := r.ReadRune()
		if err != nil {
			panic(err)
		}

		switch ch {
		case '!':
			// ignore the next char
			r.ReadRune()
		case '>':
			return
		default:
			count++
		}
	}
}
