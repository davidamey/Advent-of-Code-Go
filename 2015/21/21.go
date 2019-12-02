package main

import (
	"advent-of-code-go/util"
	"fmt"
	"math"
)

func main() {
	weapons := []item{
		item{"Dagger", 8, 4, 0},
		item{"Shortsword", 10, 5, 0},
		item{"Warhammer", 25, 6, 0},
		item{"Longsword", 40, 7, 0},
		item{"Greataxe", 74, 8, 0},
	}

	armour := []item{
		item{"None", 0, 0, 0},
		item{"Leather", 13, 0, 1},
		item{"Chainmail", 31, 0, 2},
		item{"Splintmail", 53, 0, 3},
		item{"Bandedmail", 75, 0, 4},
		item{"Platemail", 102, 0, 5},
	}

	rings := []item{
		item{"None 1", 0, 0, 0},
		item{"None 2", 0, 0, 0},
		item{"Damage +1", 25, 1, 0},
		item{"Damage +2", 50, 2, 0},
		item{"Damage +3", 100, 3, 0},
		item{"Defense +1", 20, 0, 1},
		item{"Defense +2", 40, 0, 2},
		item{"Defense +3", 80, 0, 3},
	}

	minCost := math.MaxInt32
	maxCost := math.MinInt32
	for _, w := range weapons {
		for _, a := range armour {
			for _, r1 := range rings {
				for _, r2 := range rings {
					if r1 == r2 {
						continue
					}

					cost := w.cost + a.cost + r1.cost + r2.cost
					if simulate(w, a, r1, r2) {
						if cost < minCost {
							minCost = cost
						}
					} else {
						if cost > maxCost {
							maxCost = cost
						}
					}
				}
			}
		}
	}

	fmt.Println("p1=", minCost)
	fmt.Println("p2=", maxCost)
}

func simulate(w, a, r1, r2 item) bool {
	boss := being{103, 9, 2}
	plyr := being{
		100,
		w.atk + a.atk + r1.atk + r2.atk,
		w.def + a.def + r1.def + r2.def,
	}

	attacker, defender := &plyr, &boss
	for {
		dmg := attacker.atk - defender.def
		defender.hp -= util.MaxInt(dmg, 1)
		if defender.hp <= 0 {
			break
		}
		attacker, defender = defender, attacker
	}

	return plyr.hp > 0
}

type being struct {
	hp, atk, def int
}

type item struct {
	name     string
	cost     int
	atk, def int
}
