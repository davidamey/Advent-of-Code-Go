package main

import (
	"advent/util"
	"fmt"
	"math"
)

func main() {
	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	registers := make(map[string]int)
	maxValEver := math.MinInt16

	for _, l := range lines {
		var reg, op, condReg, cond string
		var val, condVal int
		fmt.Sscanf(l, "%s %s %d if %s %s %d", &reg, &op, &val, &condReg, &cond, &condVal)

		if checkCondition(registers[condReg], cond, condVal) {
			if op == "inc" {
				registers[reg] += val
			} else {
				registers[reg] -= val
			}
		}

		if registers[reg] > maxValEver {
			maxValEver = registers[reg]
		}
	}

	maxVal := math.MinInt16
	for _, v := range registers {
		if v > maxVal {
			maxVal = v
		}
	}

	fmt.Println("p1=", maxVal)
	fmt.Println("p2=", maxValEver)
}

func checkCondition(val int, cond string, condVal int) bool {
	switch cond {
	case ">":
		return val > condVal
	case "<":
		return val < condVal
	case ">=":
		return val >= condVal
	case "<=":
		return val <= condVal
	case "==":
		return val == condVal
	case "!=":
		return val != condVal
	}

	panic(fmt.Errorf("unknown cond %s", cond))
}
