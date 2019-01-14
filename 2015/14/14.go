package main

import (
	"advent/util"
	"fmt"
)

const raceSecs = 2503

func main() {
	// file, _ := util.OpenExample()
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	rdeer := make([]*Reindeer, len(lines))
	for i, l := range lines {
		var name string
		var speed, stamina, rest int
		fmt.Sscanf(l, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds",
			&name, &speed, &stamina, &rest)
		rdeer[i] = NewReindeer(name, speed, stamina, rest)
	}

	// part1
	maxDist := 0
	var maxDeer *Reindeer
	for _, rd := range rdeer {
		if d := rd.DistanceAfterTime(raceSecs); d > maxDist {
			maxDist = d
			maxDeer = rd
		}
	}
	fmt.Println("p1=", maxDist, maxDeer.Name)

	// part2
	for s := 1; s <= raceSecs; s++ {
		leadingDist := 0
		var leadingDeer []*Reindeer
		for _, rd := range rdeer {
			if d := rd.DistanceAfterTime(s); d > leadingDist {
				leadingDist = d
				leadingDeer = []*Reindeer{rd}
			} else if d == leadingDist {
				leadingDeer = append(leadingDeer, rd)
			}
		}
		for _, rd := range leadingDeer {
			rd.Points++
		}
	}

	var winningDeer *Reindeer
	for _, rd := range rdeer {
		if winningDeer == nil {
			winningDeer = rd
			continue
		}

		if rd.Points > winningDeer.Points {
			winningDeer = rd
		}
	}

	fmt.Println("p2=", winningDeer.Points, winningDeer.Name)
}

type Reindeer struct {
	Name      string
	Speed     int
	Stamina   int
	Rest      int
	CycleSecs int
	CycleDist int
	Points    int
}

func NewReindeer(name string, speed, stamina, rest int) *Reindeer {
	return &Reindeer{
		name,
		speed,
		stamina,
		rest,
		stamina + rest,
		stamina * speed,
		0,
	}
}

func (rd *Reindeer) DistanceAfterTime(secs int) int {
	fullCycles := secs / rd.CycleSecs
	remainingSecs := secs % rd.CycleSecs

	return fullCycles*rd.CycleDist +
		util.MinInt(remainingSecs, rd.Stamina)*rd.Speed
}
