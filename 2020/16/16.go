package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	// lines := util.MustReadFileToLines("example2")
	lines := util.MustReadFileToLines("input")

	rules, myTicket, nearbyTickets := parseInfo(lines)

	p1 := 0
	validTickets := nearbyTickets[:0]
	for _, t := range nearbyTickets {
		var valid bool
		for _, v := range t {
			valid = false
			for _, r := range rules {
				valid = valid || r.check(v)
			}
			if !valid {
				p1 += v
				break
			}
		}
		if valid {
			validTickets = append(validTickets, t)
		}
	}

	////
	// part2
	////

	// initially all rules have all possible positions
	for _, r := range rules {
		for i := range myTicket {
			r.possiblePositions[i] = struct{}{}
		}
	}

	// remove positions that don't validate
	for _, t := range validTickets {
		for i, v := range t {
			for _, r := range rules {
				if !r.check(v) {
					delete(r.possiblePositions, i)
				}
			}
		}
	}

	// iteratively remove found positions from all rules until 1 position each
	for {
		finished := true

		for _, r := range rules {
			if len(r.possiblePositions) > 1 {
				continue
			}

			// remove this rule's position from all other rules
			for _, rr := range rules {
				if rr == r || len(rr.possiblePositions) == 1 {
					continue
				}
				finished = false
				delete(rr.possiblePositions, r.pos())
			}
		}

		if finished {
			break
		}
	}

	p2 := 1
	for _, r := range rules {
		if !strings.HasPrefix(r.name, "departure") {
			continue
		}

		for p := range r.possiblePositions {
			p2 *= myTicket[p]
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func parseInfo(lines []string) (rules []*rule, myTicket []int, nearbyTickets [][]int) {
	i := 0
	for _, l := range lines {
		i++
		if l == "" {
			break
		}

		rules = append(rules, newRule(l))
	}

	i++
	myTicket = util.ParseInts(lines[i], ",")

	i += 3
	for _, l := range lines[i:] {
		nearbyTickets = append(nearbyTickets, util.ParseInts(l, ","))
	}

	return
}

type rule struct {
	name                   string
	min1, max1, min2, max2 int
	possiblePositions      map[int]struct{}
}

func newRule(s string) *rule {
	parts := strings.Split(s, ": ")
	r := &rule{name: parts[0]}
	fmt.Sscanf(parts[1], "%d-%d or %d-%d", &r.min1, &r.max1, &r.min2, &r.max2)
	r.possiblePositions = make(map[int]struct{})
	return r
}

func (r *rule) check(i int) bool {
	return (i >= r.min1 && i <= r.max1) || (i >= r.min2 && i <= r.max2)
}

func (r *rule) pos() int {
	if len(r.possiblePositions) != 1 {
		panic("unsure on position")
	}
	for p := range r.possiblePositions {
		return p
	}
	panic("shouldn't get here")
}
