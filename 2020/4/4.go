package main

import (
	"advent-of-code-go/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var rgxFields = regexp.MustCompile(`(byr|iyr|eyr|hgt|hcl|ecl|pid):(\S*)`)
var rgxHexColour = regexp.MustCompile(`#[a-z0-9]{6}`)
var rgxEyeColour = regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)

func main() {
	// raw := string(util.MustReadFile("example"))
	raw := string(util.MustReadFile("input"))

	p1, p2 := 0, 0
	for _, pass := range strings.Split(raw, "\n\n") {
		valid1, valid2 := validate(pass)
		if valid1 {
			p1++
		}
		if valid2 {
			p2++
		}
	}

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func validate(pass string) (p1, p2 bool) {
	fields := make(map[string]string)
	for _, f := range rgxFields.FindAllStringSubmatch(pass, -1) {
		fields[f[1]] = f[2]
	}

	if len(fields) != 7 {
		return false, false
	}

	for k, v := range fields {
		valid := false
		switch k {
		case "byr":
			valid = isNumberBetween(v, 1920, 2002)
		case "iyr":
			valid = isNumberBetween(v, 2010, 2020)
		case "eyr":
			valid = isNumberBetween(v, 2020, 2030)
		case "hgt":
			switch v[len(v)-2:] {
			case "cm":
				valid = isNumberBetween(v[:len(v)-2], 150, 193)
			case "in":
				valid = isNumberBetween(v[:len(v)-2], 59, 76)
			}
		case "hcl":
			valid = rgxHexColour.MatchString(v)
		case "ecl":
			valid = rgxEyeColour.MatchString(v)
		case "pid":
			valid = len(v) == 9 && isNumberBetween(v, 0, 0)
		}

		if !valid {
			return true, false
		}
	}

	return true, true
}

func isNumberBetween(num string, min, max int) bool {
	n, err := strconv.Atoi(num)
	if err != nil {
		return false
	}

	if min == 0 && max == 0 {
		return true
	}

	return n >= min && n <= max
}
