package main

import (
	"advent-of-code-go/util"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// file, _ := util.OpenExample()
	// file, _ := util.OpenFile("example2")
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	rcount := len(lines) - 2
	replacements := make([]Replacement, rcount)
	for i, l := range lines[:rcount] {
		var in, out string
		fmt.Sscanf(l, "%s => %s", &in, &out)
		replacements[i] = Replacement{in, out}
	}

	base := lines[len(lines)-1]

	fmt.Println("p1=", P1(replacements, base))
	fmt.Println("p2=", P2(replacements, base))
}

func P1(replacements []Replacement, base string) int {
	counts := make(map[string]int)
	for _, r := range replacements {
		i := 0
		for {
			s := base[i:]
			j := strings.Index(s, r.In)
			if j == -1 {
				break
			}

			mole := base[:i] + strings.Replace(s, r.In, r.Out, 1)
			counts[mole]++

			i += j + 1
		}
	}
	return len(counts)
}

func P2(replacements []Replacement, target string) int {
	// From inspection there are three types of replacement
	// t1: X => XX
	// t2: X => X Rn X An
	// t3: X => X Rn X Y X Ar or X => X Rn X Y X Y X Ar

	// t1 only would take count(X) - 1 to reduce to single token
	// t2 need to be removed and number count(An or Rn)
	// t3 give two tokens extra (YX) so 2*count(Y)

	t1rgx := regexp.MustCompile(`[A-Z][a-z]?`)
	t1 := len(t1rgx.FindAllString(target, -1)) - 1

	t2rgx := regexp.MustCompile(`(Rn|Ar)`)
	t2 := len(t2rgx.FindAllString(target, -1))

	t3 := 2 * strings.Count(target, "Y")

	// fmt.Printf("%d - %d - %d\n", t1, t2, t3)

	return t1 - t2 - t3
}

type Replacement struct {
	In, Out string
}
