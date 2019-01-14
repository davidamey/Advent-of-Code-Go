package main

import (
	"advent/util"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	// lines := []string{
	// 	"Sue 1: goldfish: 6, trees: 9, akitas: 0",
	// 	"Sue 3: cars: 10, akitas: 6, perfumes: 7",
	// }

	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	rgx := regexp.MustCompile(`([a-z]+): (\d+)`)

	sues := make([]MFCSAM, len(lines))
	for i, l := range lines {
		sues[i] = make(MFCSAM)

		matches := rgx.FindAllStringSubmatch(l, -1)
		for _, m := range matches {
			v, _ := strconv.Atoi(m[2])
			sues[i][m[1]] = v
		}
	}

	wanted := GetWanted()
	var p1, p2 int
	for i, s := range sues {
		if s.MatchesP1(wanted) {
			p1 = i + 1
		}
		if s.MatchesP2(wanted) {
			p2 = i + 1
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

type MFCSAM map[string]int

func GetWanted() MFCSAM {
	wanted := make(MFCSAM)
	wanted["children"] = 3
	wanted["cats"] = 7
	wanted["samoyeds"] = 2
	wanted["pomeranians"] = 3
	wanted["akitas"] = 0
	wanted["vizslas"] = 0
	wanted["goldfish"] = 5
	wanted["trees"] = 3
	wanted["cars"] = 2
	wanted["perfumes"] = 1
	return wanted
}

func (m1 MFCSAM) MatchesP1(wanted MFCSAM) bool {
	for k, wv := range wanted {
		if av, ok := m1[k]; ok {
			if av != wv {
				return false
			}
		}
	}
	return true
}

func (m1 MFCSAM) MatchesP2(wanted MFCSAM) bool {
	for k, wv := range wanted {
		if av, ok := m1[k]; ok {
			switch k {
			case "cats", "trees":
				if av <= wv {
					return false
				}
			case "pomeranians", "goldfish":
				if av >= wv {
					return false
				}
			default:
				if av != wv {
					return false
				}
			}
		}
	}
	return true
}
