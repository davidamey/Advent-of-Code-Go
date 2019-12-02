package main

import (
	"advent-of-code-go/util"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const DEBUG = false

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	winningTeam, remainingUnits := Solve(lines, 0)
	fmt.Println("== part1 ==")
	fmt.Printf("%s wins with %d units\n", winningTeam.ToString(), remainingUnits)

	lowerBoost := 1
	upperBoost := 10000
	for lowerBoost < upperBoost {
		boost := (upperBoost + lowerBoost) / 2
		// fmt.Println(boost, lowerBoost, upperBoost)
		var t Team
		t, remainingUnits = Solve(lines, boost)
		if t == immune {
			upperBoost = boost - 1
		} else {
			lowerBoost = boost + 1
		}
	}

	fmt.Println("== part2 ==")
	fmt.Printf("With a boost of +%d, the Immune system survives with %d units\n",
		lowerBoost, remainingUnits)
}

func Solve(lines []string, immuneBoost int) (Team, int) {
	var groups []*Group
	var t Team
	groupIDs := make(map[Team]int)
	groupIDs[immune] = 1
	groupIDs[infection] = 1
	for _, l := range lines {
		switch l {
		case "Immune System:":
			t = immune
		case "Infection:":
			t = infection
		case "":
			// nothing
		default:
			id := groupIDs[t]
			groupIDs[t]++
			groups = append(groups, ParseGroup(id, t, immuneBoost, l))
		}
	}

	for {
		DumpGroups(groups)

		// *** Phase: Target selection
		sort.Slice(groups, func(i, j int) bool {
			if groups[i].Power() == groups[j].Power() {
				return groups[i].UnitInitiative > groups[j].UnitInitiative
			}
			return groups[i].Power() > groups[j].Power()
		})

		targets := make(map[*Group]struct{})
		for _, g := range groups {
			if g.UnitCount == 0 {
				continue
			}

			var bestTarget *Group
			targetDmg := 0
			for _, g2 := range groups {
				// Same group or same team
				if g2 == g || g2.Team == g.Team || g2.UnitCount == 0 {
					continue
				}

				// Already targetted
				if _, ok := targets[g2]; ok {
					continue
				}

				if d := g.DamageTo(g2); d > targetDmg {
					bestTarget = g2
					targetDmg = d
				}
			}

			if targetDmg > 0 {
				g.Target = bestTarget
				targets[bestTarget] = struct{}{}
			}
		}

		DumpTargetting(groups)

		// *** Phase: Attack

		// Groups attack in decreasing order of initiative
		sort.Slice(groups, func(i, j int) bool {
			return groups[i].UnitInitiative > groups[j].UnitInitiative
		})

		for _, g := range groups {
			if g.Target == nil {
				continue
			}

			dmg := g.DamageTo(g.Target)
			killedUnits := util.MinInt(dmg/g.Target.UnitHP, g.Target.UnitCount)

			if DEBUG {
				// fmt.Println(dmg, g.Target.UnitHP, dmg/g.Target.UnitHP, g.Target.UnitCount)
				fmt.Printf("%s group %d attacks defending group %d, killing %d units\n",
					t.ToString(), g.ID, g.Target.ID, killedUnits)
			}

			g.Target.UnitCount -= killedUnits
			g.Target = nil
		}

		// Do both armies have units
		remainingUnits := make(map[Team]int)
		for _, g := range groups {
			remainingUnits[g.Team] += g.UnitCount
		}

		if remainingUnits[immune] == 0 {
			return infection, remainingUnits[infection]
		}
		if remainingUnits[infection] == 0 {
			return immune, remainingUnits[immune]
		}
	}
}

type (
	DmgType int8
	Team    int
)

const (
	none        DmgType = -1
	bludgeoning DmgType = iota
	cold
	fire
	radiation
	slashing
)

const (
	immune Team = iota
	infection
)

func (t Team) ToString() string {
	if t == immune {
		return "Immune System"
	}
	return "Infection"
}

type Group struct {
	ID             int
	Team           Team
	UnitCount      int
	UnitHP         int
	UnitAtk        int
	UnitAtkType    DmgType
	UnitWeakness   map[DmgType]struct{}
	UnitImmunity   map[DmgType]struct{}
	UnitInitiative int
	Enemies        *[]*Group
	Target         *Group
}

func ParseGroup(id int, team Team, immuneBoost int, raw string) *Group {
	rgx := regexp.MustCompile(`(\d+) units each with (\d+) hit points(?: \(([^)]+)\))? with an attack that does (\d+) (\w+) damage at initiative (\d+)`)
	m := rgx.FindStringSubmatch(raw)
	if len(m) == 0 {
		panic(fmt.Errorf(`unable to parse "%s"`, raw))
	}
	g := &Group{
		ID:             id,
		Team:           team,
		UnitCount:      ToInt(m[1]),
		UnitHP:         ToInt(m[2]),
		UnitAtk:        ToInt(m[4]),
		UnitAtkType:    ToDmgType(m[5]),
		UnitWeakness:   ParseWeakness(m[3]),
		UnitImmunity:   ParseImmunity(m[3]),
		UnitInitiative: ToInt(m[6]),
	}
	if team == immune {
		g.UnitAtk += immuneBoost
	}
	return g
}

func (g *Group) Power() int {
	return g.UnitAtk * g.UnitCount
}

func (g *Group) DamageTo(g2 *Group) int {
	// Immune = no damage
	if _, ok := g2.UnitImmunity[g.UnitAtkType]; ok {
		return 0
	}

	// Weak = double damage
	if _, ok := g2.UnitWeakness[g.UnitAtkType]; ok {
		return 2 * g.Power()
	}

	return g.Power()
}

func ToInt(in string) int {
	if out, err := strconv.Atoi(in); err == nil {
		return out
	}
	panic("unable to convert " + in)
}

func ToDmgType(in string) DmgType {
	switch strings.TrimSpace(in) {
	case "bludgeoning":
		return bludgeoning
	case "cold":
		return cold
	case "fire":
		return fire
	case "radiation":
		return radiation
	case "slashing":
		return slashing
	}
	panic("unknown dmg type " + in)
}

func ParseWeakness(in string) map[DmgType]struct{} {
	weak := make(map[DmgType]struct{})

	rgx := regexp.MustCompile(`weak to ([^;)]+)`)
	m := rgx.FindStringSubmatch(in)
	if len(m) == 0 {
		return weak
	}

	for _, w := range strings.Split(m[1], ",") {
		weak[ToDmgType(w)] = struct{}{}
	}

	return weak
}

func ParseImmunity(in string) map[DmgType]struct{} {
	immune := make(map[DmgType]struct{})

	rgx := regexp.MustCompile(`immune to ([^;)]+)`)
	m := rgx.FindStringSubmatch(in)
	if len(m) == 0 {
		return immune
	}

	for _, w := range strings.Split(m[1], ",") {
		immune[ToDmgType(w)] = struct{}{}
	}

	return immune
}

func DumpGroups(groups []*Group) {
	if !DEBUG {
		return
	}

	var immuneOut, infectOut string
	for _, g := range groups {
		if g.UnitCount == 0 {
			continue
		}

		if g.Team == immune {
			immuneOut += fmt.Sprintf("Group %d contains %d units\n", g.ID, g.UnitCount)
		} else {
			infectOut += fmt.Sprintf("Group %d contains %d units\n", g.ID, g.UnitCount)
		}
	}

	fmt.Println("Immune System:")
	fmt.Print(immuneOut)
	fmt.Println("Infection:")
	fmt.Print(infectOut)
	fmt.Println()
}

func DumpTargetting(groups []*Group) {
	if !DEBUG {
		return
	}

	var immuneOut, infectOut string
	for _, g := range groups {
		if g.Target == nil {
			continue
		}

		immuneOut += fmt.Sprintf("%s group %d would deal defending group %d %d damage\n",
			g.Team.ToString(), g.ID, g.Target.ID, g.DamageTo(g.Target))
	}

	fmt.Print(infectOut)
	fmt.Print(immuneOut)
	fmt.Println()
}
