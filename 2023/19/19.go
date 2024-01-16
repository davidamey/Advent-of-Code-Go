package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// raw := string(util.MustReadFile("example"))
	raw := string(util.MustReadFile("input"))

	rawFlows, rawParts, _ := strings.Cut(raw, "\n\n")

	workflows := make(map[string]workflow)
	for _, l := range strings.Split(rawFlows, "\n") {
		name, rules, _ := strings.Cut(l, "{")
		w := workflow{}
		for _, r := range strings.Split(rules[:len(rules)-1], ",") {
			if !strings.ContainsRune(r, ':') {
				w = append(w, rule{target: r})
				continue
			}

			attrib, op := r[0], r[1]
			var cond int
			var target string
			fmt.Sscanf(r[2:], "%d:%s", &cond, &target)

			w = append(w, rule{
				attrib, op, cond, target,
			})
		}
		workflows[name] = w
	}

	var parts []part
	for _, l := range strings.Split(rawParts, "\n") {
		var x, m, a, s int
		fmt.Sscanf(l, "{x=%d,m=%d,a=%d,s=%d", &x, &m, &a, &s)
		parts = append(parts, part{'x': x, 'm': m, 'a': a, 's': s})
	}

	p1 := 0
	for _, p := range parts {
		if runWorkflow(workflows, "in", p) == "A" {
			p1 += p['x'] + p['m'] + p['a'] + p['s']
		}
	}
	fmt.Println("p1=", p1)

	combinations := findCombinations(workflows, "in", map[byte][2]int{
		'x': {1, 4000}, 'm': {1, 4000}, 'a': {1, 4000}, 's': {1, 4000},
	})
	fmt.Println("p2=", combinations)
}

func runWorkflow(workflows map[string]workflow, name string, p part) string {
	var target string
	for _, r := range workflows[name] {
		if r.op == 0 { // constant
			target = r.target
			break
		}

		if r.op == '>' && p[r.attrib] > r.cond {
			target = r.target
			break
		}

		if r.op == '<' && p[r.attrib] < r.cond {
			target = r.target
			break
		}

	}
	if target == "A" || target == "R" {
		return target
	}
	return runWorkflow(workflows, target, p)
}

func findCombinations(workflows map[string]workflow, name string, ranges map[byte][2]int) int {
	total := 0

	for _, r := range workflows[name] {
		newRange := ranges[r.attrib]
		if r.op == '<' {
			newRange[1] = r.cond - 1
			ranges[r.attrib] = [2]int{r.cond, ranges[r.attrib][1]}
		} else {
			newRange[0] = r.cond + 1
			ranges[r.attrib] = [2]int{ranges[r.attrib][0], r.cond}
		}

		newRanges := make(map[byte][2]int)
		for k, v := range ranges {
			if k == r.attrib {
				newRanges[k] = newRange
			} else {
				newRanges[k] = v
			}
		}

		if r.target == "R" {
			continue
		}

		if r.target == "A" {
			total += (newRanges['x'][1] - newRanges['x'][0] + 1) *
				(newRanges['m'][1] - newRanges['m'][0] + 1) *
				(newRanges['a'][1] - newRanges['a'][0] + 1) *
				(newRanges['s'][1] - newRanges['s'][0] + 1)
			continue
		}

		total += findCombinations(workflows, r.target, newRanges)
	}

	return total
}

type part map[byte]int
type workflow []rule

type rule struct {
	attrib byte
	op     byte
	cond   int
	target string
}
