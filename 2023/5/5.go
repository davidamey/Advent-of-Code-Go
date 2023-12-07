package main

import (
	"advent-of-code-go/util"
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// raw := string(util.MustReadFile("example"))
	raw := string(util.MustReadFile("input"))

	parts := strings.Split(raw, "\n\n")

	var p1Input, p2Input [][2]int
	for i, f := range strings.Fields(parts[0])[1:] {
		x := util.Atoi(f)
		p1Input = append(p1Input, [2]int{x, x})
		if i%2 == 0 {
			p2Input = append(p2Input, [2]int{x, x})
		} else {
			j := len(p2Input) - 1
			p2Input[j][1] = p2Input[j][0] + x
		}
	}

	mappers := make([]*mapper, len(parts)-1)
	for i, p := range parts[1:] {
		mappers[i] = newMapper(p)
	}

	fmt.Println("p1=", solve(p1Input, mappers))
	fmt.Println("p2=", solve(p2Input, mappers))
}

func solve(input [][2]int, mappers []*mapper) int {
	output := input
	for _, m := range mappers {
		output = m.run(output)
	}

	min := math.MaxInt
	for _, o := range output {
		if o[0] < min {
			min = o[0]
		}
	}
	return min
}

type mapRange struct {
	end, start, length int
}

type mapper struct {
	name   string
	ranges []*mapRange
}

func newMapper(s string) *mapper {
	parts := strings.Fields(s)
	m := &mapper{name: parts[0]}
	for i := 2; i < len(parts); i += 3 {
		m.ranges = append(m.ranges, &mapRange{
			end:    util.Atoi(parts[i]),
			start:  util.Atoi(parts[i+1]),
			length: util.Atoi(parts[i+2]),
		})
	}
	return m
}

func (m *mapper) run(inputs [][2]int) (outputs [][2]int) {
process:
	for len(inputs) > 0 {
		input := util.Pop(&inputs)

		for _, r := range m.ranges {
			delta := r.end - r.start

			// fully contained
			if input[0] >= r.start && input[1] < r.start+r.length {
				outputs = append(outputs, [2]int{
					input[0] + delta,
					input[1] + delta,
				})
				continue process
			}

			// partially contained
			if input[0] >= r.start && input[0] < r.start+r.length {
				outputs = append(outputs, [2]int{
					input[0] + delta,
					r.start + r.length - 1 + delta,
				})

				// remaining input
				inputs = append(inputs, [2]int{r.start + r.length, input[1]})
				continue process
			}
		}

		// no ranges matched, so pass on
		outputs = append(outputs, input)
	}
	return
}
