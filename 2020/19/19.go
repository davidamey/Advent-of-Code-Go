package main

import (
	"advent-of-code-go/util"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	lines := util.MustReadFileToLines("input")
	rules := make(map[string]string)
	for _, l := range lines {
		if l == "" {
			break
		}
		parts := strings.Split(l, ": ")
		rules[parts[0]] = parts[1]
	}
	cases := lines[len(rules)+1:]

	p1, p2 := solve(rules, cases)
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func solve(rules map[string]string, cases []string) (p1, p2 int) {
	parsed := make(map[string]string)
	found0 := false
	for !found0 {
		parseRules(rules, parsed)
		_, found0 = parsed["0"]
	}

	rgx := regexp.MustCompile("^" + parsed["0"] + "$")
	for _, c := range cases {
		if rgx.MatchString(c) {
			p1++
		}
	}

	// Now modify with loops for p2
	// Add a 'loop layer' each time until nothing further matched
	rules["8"] = "42 | 42 8"
	rules["11"] = "42 31 | 42 11 31"

	buffer := 1
	matchingCases := make(map[string]struct{})
	for {
		parseRules(rules, parsed)

		rgx = regexp.MustCompile("^" + parsed["0"] + "$")

		matchedNew := false
		for _, c := range cases {
			if !rgx.MatchString(c) {
				continue
			}

			if _, alreadyMatching := matchingCases[c]; !alreadyMatching {
				matchingCases[c] = struct{}{}
				matchedNew = true
			}
		}
		if !matchedNew {
			if buffer == 0 {
				break
			}
			buffer--
		}
	}

	return p1, len(matchingCases)
}

func parseRules(rules, parsed map[string]string) {
parsing:
	for rid, rule := range rules {
		if rule[0] == '"' {
			parsed[rid] = rule[1:2]
			continue
		}

		parts := strings.Split(rule, " ")

		var sb strings.Builder
		sb.WriteString("(")
		for _, p := range parts {
			if p == "|" {
				sb.WriteString("|")
			} else {
				v, exists := parsed[p]
				if !exists {
					continue parsing
				}
				sb.WriteString(v)
			}
		}
		sb.WriteString(")")
		parsed[rid] = sb.String()
	}
}
