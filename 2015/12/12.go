package main

import (
	"advent/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	input, _ := ioutil.ReadAll(file)

	// part1
	rgx := regexp.MustCompile(`-?\d+`)
	numbers := rgx.FindAllString(string(input), -1)

	p1 := 0
	for _, n := range numbers {
		i, _ := strconv.Atoi(n)
		p1 += i
	}
	fmt.Println("p1=", p1)

	// part2
	var v interface{}
	err := json.Unmarshal(input, &v)

	if err != nil {
		panic(err)
	}

	p2, _ := ParseValue(v)
	fmt.Println("p2=", p2)
}

func ParseMap(m map[string]interface{}) (sum int) {
	for _, v := range m {
		vsum, red := ParseValue(v)
		sum += vsum

		// part2: discount objects with "red" in
		if red {
			return 0
		}
	}
	return
}

func ParseSlice(s []interface{}) (sum int) {
	for _, v := range s {
		vsum, _ := ParseValue(v)
		sum += vsum
	}
	return
}

func ParseValue(v interface{}) (sum int, red bool) {
	switch x := v.(type) {
	case int:
		sum = x
	case float64:
		sum = int(x)
	case string:
		if x == "red" {
			red = true
		}
	case []interface{}:
		sum = ParseSlice(x)
	case map[string]interface{}:
		sum = ParseMap(x)
	default:
		panic("unknown type")
	}
	return
}
