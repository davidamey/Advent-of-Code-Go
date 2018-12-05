package main

import (
	"advent/util"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

type guard struct {
	ID        int
	SleepTime int
	SleepMap  map[int]int
}

type entry struct {
	Time time.Time
	Text string
}

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	inputs, _ := util.ReadLines(file)

	var entries []entry
	r, _ := regexp.Compile("\\[(.+)\\]\\s*(.+)")
	for _, s := range inputs {
		matches := r.FindStringSubmatch(s)
		t, _ := time.Parse("2006-01-02 15:04", matches[1])
		entries = append(entries, entry{
			t,
			matches[2],
		})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Time.Before(entries[j].Time)
	})

	guards := make(map[int]*guard)
	var g *guard
	for i := 0; i < len(entries); i++ {
		if strings.Contains(entries[i].Text, "Guard") {
			g = processGuard(guards, entries[i].Text)
			continue
		}

		// Assume we're at a sleep row and the sleep/wake come in pairs
		total := 0
		for m := entries[i].Time.Minute(); m < entries[i+1].Time.Minute(); m++ {
			g.SleepMap[m]++
			total++
		}
		g.SleepTime += total
		i++ // Skip the already-processed line
	}

	part1(guards)
	part2(guards)
}

func part1(guards map[int]*guard) {
	var longestSleeper *guard
	for i := range guards {
		if longestSleeper == nil {
			longestSleeper = guards[i]
			continue
		}

		if guards[i].SleepTime > longestSleeper.SleepTime {
			longestSleeper = guards[i]
		}
	}

	mostSleptMinute := 0
	for m, c := range longestSleeper.SleepMap {
		if c > longestSleeper.SleepMap[mostSleptMinute] {
			mostSleptMinute = m
		}
	}

	fmt.Println("== part1 ==")
	fmt.Printf("guard %d slept the most (%d minutes)\n", longestSleeper.ID, longestSleeper.SleepTime)
	fmt.Printf("he slept mostly on minute %d (%d times)\n", mostSleptMinute, longestSleeper.SleepMap[mostSleptMinute])
	fmt.Printf("answer = guardID * longestMinute = %d\n", longestSleeper.ID*mostSleptMinute)
}

func part2(guards map[int]*guard) {
	var g *guard
	mostSleptMinute := 0
	for i := range guards {
		if g == nil {
			g = guards[i]
		}

		for m, c := range guards[i].SleepMap {
			if c > g.SleepMap[mostSleptMinute] {
				g = guards[i]
				mostSleptMinute = m
			}
		}
	}

	fmt.Println("== part2 ==")
	fmt.Printf("guard %d slept on minute %d longer than any other guard/minute combo (%d times)\n", g.ID, mostSleptMinute, g.SleepMap[mostSleptMinute])
	fmt.Printf("answer = guardID * mostSleptMinute = %d\n", g.ID*mostSleptMinute)
}

func processGuard(guards map[int]*guard, s string) *guard {
	var id int
	fmt.Sscanf(s, "Guard #%d begins shift", &id)

	if guards[id] == nil {
		guards[id] = &guard{id, 0, make(map[int]int)}
	}

	return guards[id]
}
