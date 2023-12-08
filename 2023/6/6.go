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

	rawTimes, rawDistances := util.Divide(raw, "\n")

	var times []int
	for _, f := range strings.Fields(rawTimes)[1:] {
		times = append(times, util.Atoi(f))
	}
	var distances []int
	for _, f := range strings.Fields(rawDistances)[1:] {
		distances = append(distances, util.Atoi(f))
	}

	p1 := 1
	for i := range times {
		p1 *= len(race(times[i], distances[i]))
	}
	fmt.Println("p1=", p1)

	p2Time := util.Atoi(strings.Join(strings.Fields(rawTimes)[1:], ""))
	p2Distance := util.Atoi(strings.Join(strings.Fields(rawDistances)[1:], ""))
	fmt.Println("p2=", len(race(p2Time, p2Distance)))
}

func race(raceTime, recordDistance int) (winningTimes []int) {
	for t := 0; t <= raceTime; t++ {
		d := (raceTime - t) * t
		if d > recordDistance {
			winningTimes = append(winningTimes, t)
		}
	}
	return
}
