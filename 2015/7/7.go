package main

import (
	"advent/util"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// file, _ := util.OpenExample()
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	circuit := NewCircuit()
	for _, l := range lines {
		parts := strings.Split(l, " -> ")
		if len(parts) != 2 {
			panic(fmt.Errorf(`error parsing instruction "%s"`, l))
		}

		circuit.SetWire(parts[1], ParseWire(parts[0]))
	}
	// circuit.Dump()

	p1 := circuit.Eval("a")

	circuit.Reset()
	circuit.SetWire("b", func(c *Circuit) uint16 {
		return p1
	})

	p2 := circuit.Eval("a")

	fmt.Println("p1", p1)
	fmt.Println("p2", p2)
}

type Wire func(*Circuit) uint16

func ParseWire(raw string) Wire {
	for _, i := range Instructs {
		if i.rgx.MatchString(raw) {
			return i.ToWire(i.rgx.FindStringSubmatch(raw))
		}
	}
	panic(fmt.Errorf(`unable to parse "%s"`, raw))

}

type Circuit struct {
	wires map[string]Wire
	cache map[string]uint16
}

func NewCircuit() Circuit {
	return Circuit{
		make(map[string]Wire),
		make(map[string]uint16),
	}
}

func (c *Circuit) Reset() {
	c.cache = make(map[string]uint16)
}

func (c *Circuit) SetWire(name string, w Wire) {
	c.wires[name] = w
}

func (c *Circuit) Eval(s string) uint16 {
	if v, exists := c.cache[s]; exists {
		return v
	}
	v := c.wires[s](c)
	c.cache[s] = v
	return v
}

func (c *Circuit) Dump() {
	keys := make([]string, 0, len(c.wires))
	for k := range c.wires {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println("wires:")
	for _, k := range keys {
		v := c.wires[k](c)
		fmt.Printf("%3s = %5d\n", k, v)
	}
}

type Instruct struct {
	rgx    *regexp.Regexp
	ToWire func([]string) Wire
}

var Instructs []Instruct = []Instruct{
	Instruct{ // Set
		regexp.MustCompile(`^(\w+)$`),
		func(m []string) Wire {
			return func(c *Circuit) uint16 {
				if val, err := strconv.Atoi(m[1]); err == nil {
					return uint16(val)
				} else {
					return c.Eval(m[1])
				}
			}
		},
	},
	Instruct{ // AND
		regexp.MustCompile(`^(\w+) AND (\w+)$`),
		func(m []string) Wire {
			return func(c *Circuit) uint16 {
				if val, err := strconv.Atoi(m[1]); err == nil {
					return uint16(val) & c.Eval(m[2])
				} else {
					return c.Eval(m[1]) & c.Eval(m[2])
				}
			}
		},
	},
	Instruct{ // OR
		regexp.MustCompile(`^(\w+) OR (\w+)$`),
		func(m []string) Wire {
			return func(c *Circuit) uint16 {
				if val, err := strconv.Atoi(m[1]); err == nil {
					return uint16(val) | c.Eval(m[2])
				} else {
					return c.Eval(m[1]) | c.Eval(m[2])
				}
			}
		},
	},
	Instruct{ // LSHIFT
		regexp.MustCompile(`^(\w+) LSHIFT (\d+)$`),
		func(m []string) Wire {
			return func(c *Circuit) uint16 {
				i, _ := strconv.Atoi(m[2])
				return c.Eval(m[1]) << uint(i)
			}
		},
	},
	Instruct{ // RSHIFT
		regexp.MustCompile(`^(\w+) RSHIFT (\d+)$`),
		func(m []string) Wire {
			return func(c *Circuit) uint16 {
				i, _ := strconv.Atoi(m[2])
				return c.Eval(m[1]) >> uint(i)
			}
		},
	},
	Instruct{
		regexp.MustCompile(`^NOT (\w+)$`),
		func(m []string) Wire {
			return func(c *Circuit) uint16 {
				return ^c.Eval(m[1])
			}
		},
	},
}
