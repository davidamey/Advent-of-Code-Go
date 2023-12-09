package main

import (
	"advent-of-code-go/util"
	"fmt"
	"math"
)

func main() {
	// file, _ := util.OpenExample()
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	rp := ParseLines(lines)
	sp := util.NewPermuter(rp.Cities)

	minD := math.MaxInt32
	maxD := 0
	for p := range sp.Permutations() {
		d := rp.DistanceForRoute(p)
		if d < minD {
			minD = d
		}
		if d > maxD {
			maxD = d
		}
	}

	fmt.Println("p1", minD)
	fmt.Println("p2", maxD)
}

func ParseLines(lines []string) *RoutePlanner {
	uniqueCities := make(map[string]struct{})
	rp := NewRoutePlanner()

	for _, l := range lines {
		var c1, c2 string
		var d int
		fmt.Sscanf(l, "%s to %s = %d", &c1, &c2, &d)
		rp.AddRoute(c1, c2, d)
		rp.AddRoute(c2, c1, d)
		uniqueCities[c1] = struct{}{}
		uniqueCities[c2] = struct{}{}
	}

	rp.Cities = make([]string, 0, len(uniqueCities))
	for c := range uniqueCities {
		rp.Cities = append(rp.Cities, c)
	}

	return rp
}

type RoutePlanner struct {
	Cities []string
	routes map[string]int
}

func NewRoutePlanner() *RoutePlanner {
	return &RoutePlanner{routes: make(map[string]int)}
}

func (rp *RoutePlanner) AddRoute(c1, c2 string, d int) {
	rp.routes[c1+c2] = d
	rp.routes[c2+c1] = d
}

func (rp *RoutePlanner) GetDistance(c1, c2 string) int {
	return rp.routes[c1+c2]
}

func (rp *RoutePlanner) DistanceForRoute(cities []string) int {
	d := 0
	for i := 0; i < len(cities)-1; i++ {
		d += rp.GetDistance(cities[i], cities[i+1])
	}
	return d
}
