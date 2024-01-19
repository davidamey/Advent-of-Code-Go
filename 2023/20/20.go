package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines, presses := util.MustReadFileToLines("example"), 1000
	// lines, presses := util.MustReadFileToLines("example2"), 1000
	lines, presses := util.MustReadFileToLines("input"), 10_000

	modules := parseModules(lines)
	rxFeeder := findRxFeeder(modules)

	pulses := []pulse{}
	low, high := 0, 0
loop:
	for i := 0; i < presses; i++ {
		pulses = append(pulses, pulse{from: "button", target: "broadcaster"})
		for len(pulses) > 0 {
			p := pulses[0]
			pulses = pulses[1:]

			// p1
			if i < 1000 {
				if p.val == 0 {
					low++
				} else {
					high++
				}
			}

			// p2
			if rxFeeder != nil && p.val == 1 && p.target == rxFeeder.name {
				if checkFeeder(rxFeeder, p, i) {
					break loop
				}
			}

			m, exists := modules[p.target]
			if exists {
				pulses = append(pulses, m.apply(p)...)
			}
		}
	}

	p2Indexes := []int{}
	if rxFeeder != nil {
		for _, x := range rxFeeder.p2Indexes {
			p2Indexes = append(p2Indexes, x)
		}
	}

	fmt.Println("p1=", low*high)
	fmt.Println("p2=", util.LCM(p2Indexes...))
}

func checkFeeder(feeder *conjunction, p pulse, i int) bool {
	if feeder.p2Indexes[p.from] > 0 {
		return false
	}
	feeder.p2Indexes[p.from] = i + 1
	for _, c := range feeder.p2Indexes {
		if c == 0 {
			return false
		}
	}
	return true
}

func findRxFeeder(modules map[string]module) *conjunction {
	for _, m := range modules {
		for _, t := range m.getTargets() {
			if t == "rx" {
				return m.(*conjunction) // assume feeder is conjuction from inspection
			}
		}
	}
	return nil
}

func parseModules(lines []string) map[string]module {
	modules := map[string]module{}
	for _, l := range lines {
		name, targetsRaw, _ := strings.Cut(l, " -> ")

		targets := strings.Split(targetsRaw, ", ")

		switch {
		case name == "broadcaster":
			modules[name] = &broadcaster{targets}
		case name[0] == '%':
			modules[name[1:]] = &flipFlop{targets: targets}
		case name[0] == '&':
			modules[name[1:]] = &conjunction{name[1:], targets, map[string]int{}, map[string]int{}}
		}
	}

	// Pre-populate conjunctions
	for src, m := range modules {
		for _, t := range m.getTargets() {
			if c, ok := modules[t].(*conjunction); ok {
				c.inputs[src] = 0
				c.p2Indexes[src] = 0
				continue
			}
		}
	}

	return modules
}

type pulse struct {
	from   string
	target string
	val    int
}

func (p pulse) String() string {
	arr := " -low-> "
	if p.val == 1 {
		arr = " -high-> "
	}
	return fmt.Sprint(p.from, arr, p.target)
}

type module interface {
	getTargets() []string
	apply(in pulse) []pulse
}

type broadcaster struct {
	targets []string
}

func (b *broadcaster) apply(_ pulse) []pulse {
	pulses := make([]pulse, len(b.targets))
	for i, t := range b.targets {
		pulses[i] = pulse{from: "broadcaster", target: t, val: 0}
	}
	return pulses
}

func (b *broadcaster) getTargets() []string {
	return b.targets
}

type flipFlop struct {
	targets []string
	on      bool
}

func (ff *flipFlop) apply(in pulse) []pulse {
	if in.val == 1 {
		return nil
	}

	ff.on = !ff.on
	x := 0
	if ff.on {
		x = 1
	}

	pulses := make([]pulse, len(ff.targets))
	for i, t := range ff.targets {
		pulses[i] = pulse{from: in.target, target: t, val: x}
	}
	return pulses
}

func (ff *flipFlop) getTargets() []string {
	return ff.targets
}

type conjunction struct {
	name      string
	targets   []string
	inputs    map[string]int
	p2Indexes map[string]int
}

func (c *conjunction) apply(in pulse) []pulse {
	c.inputs[in.from] = in.val

	x := 0
	for _, v := range c.inputs {
		if v == 0 {
			x = 1
			break
		}
	}

	pulses := make([]pulse, len(c.targets))
	for i, t := range c.targets {
		pulses[i] = pulse{from: in.target, target: t, val: x}
	}
	return pulses
}

func (c *conjunction) getTargets() []string {
	return c.targets
}
