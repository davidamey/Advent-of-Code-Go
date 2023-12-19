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

	// raw := string(util.MustReadFile("example"))
	raw := string(util.MustReadFile("input"))

	p1 := 0
	boxes := [256]box{}
	for _, step := range strings.Split(raw, ",") {
		p1 += hash(step)

		label, op, fLen := parseStep(step)
		boxNum := hash(label)

		if boxes[boxNum] == nil {
			boxes[boxNum] = box{}
		}

		switch op {
		case '=':
			slot := len(boxes[boxNum]) + 1
			if l, exists := boxes[boxNum][label]; exists {
				slot = l.slot
			}
			boxes[boxNum][label] = lens{slot, fLen}
		case '-':
			if toRemove, exists := boxes[boxNum][label]; exists {
				for lbl, l := range boxes[boxNum] {
					if l.slot > toRemove.slot {
						boxes[boxNum][lbl] = lens{
							slot: l.slot - 1,
							fLen: l.fLen,
						}
					}
				}
				delete(boxes[boxNum], label)
			}
		}
	}

	p2 := 0
	for i, b := range boxes {
		for _, l := range b {
			p2 += (i + 1) * l.slot * l.fLen
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func hash(s string) (h int) {
	for i := range s {
		h += int(s[i])
		h *= 17
		h %= 256
	}
	return
}

var stepRgx = regexp.MustCompile(`(\w+)([=-])(\d*)`)

func parseStep(step string) (label string, op rune, fLen int) {
	matches := stepRgx.FindAllStringSubmatch(step, -1)
	if matches[0][2] == "=" {
		return matches[0][1], '=', util.Atoi(matches[0][3])
	}
	return matches[0][1], '-', 0
}

type box map[string]lens

type lens struct {
	slot, fLen int
}
