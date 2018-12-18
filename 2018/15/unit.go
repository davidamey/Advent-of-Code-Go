package main

import (
	"advent/util"
	"sort"
)

type Unit struct {
	Game *Game
	Rune rune
	Pos  Point
	HP   int
	Atk  int
}

func (u *Unit) Alive() bool {
	return u.HP > 0
}

func (u *Unit) OtherUnits() []*Unit {
	return Filter(u.Game.Units, func(v *Unit) bool {
		return v.Alive() && v != u
	})
}

func (u *Unit) Targets() []*Unit {
	return Filter(u.Game.Units, func(v *Unit) bool {
		return v.Alive() && v.Rune != u.Rune
	})
}

func (u *Unit) TargetsAdjacent() []*Unit {
	return Filter(u.Game.Units, func(v *Unit) bool {
		dist := util.AbsInt(u.Pos.X-v.Pos.X) + util.AbsInt(u.Pos.Y-v.Pos.Y)
		return dist == 1 && v.Alive() && v.Rune != u.Rune
	})
}

func (u *Unit) Opponent() *Unit {
	var op *Unit
	adj := u.TargetsAdjacent()
	for i := range adj {
		if op == nil {
			op = adj[i]
			continue
		}

		if adj[i].HP < op.HP {
			op = adj[i]
			continue
		}

		if adj[i].HP == op.HP {
			if adj[i].Pos.Y == op.Pos.Y && adj[i].Pos.X < op.Pos.X {
				op = adj[i]
				continue
			}

			if adj[i].Pos.Y < op.Pos.Y {
				op = adj[i]
				continue
			}
		}
	}
	return op
}

func (u *Unit) Move() bool {
	var moveTargets []Point
	targets := u.Targets()
	if len(targets) == 0 {
		return false
	}

	for _, t := range targets {
		for _, p := range t.Pos.Adjacent() {
			if u.Game.Grid[p.Y][p.X] == '.' {
				// u.Game.Grid[p.Y][p.X] = '?'
				moveTargets = append(moveTargets, p)
			}
		}
	}

	found := ShortestPath(u.Pos, moveTargets, u.Game.Grid)
	if len(found) == 0 {
		return true
	}

	// Sort in reading order
	sort.Slice(found, func(i, j int) bool {
		if found[i].Pos.Y == found[j].Pos.Y {
			return found[i].Pos.X < found[j].Pos.X
		}
		return found[i].Pos.Y < found[j].Pos.Y
	})

	nextMove := found[0]
	for nextMove.Parent.Pos != u.Pos {
		nextMove = *nextMove.Parent
	}

	u.Game.Grid[u.Pos.Y][u.Pos.X] = '.'
	u.Game.Grid[nextMove.Pos.Y][nextMove.Pos.X] = u.Rune
	u.Pos = nextMove.Pos

	return true
}

func (u *Unit) Attack(v *Unit) {
	v.HP -= u.Atk
	if v.HP <= 0 {
		v.HP = 0
		u.Game.Grid[v.Pos.Y][v.Pos.X] = '.'
	}
}

func (u *Unit) TakeTurn() bool {
	if !u.Alive() {
		return true
	}

	// u.Acting = true
	u.Game.LastActed = u

	if op := u.Opponent(); op != nil {
		u.Attack(op)
		return true
	}

	haveTargets := u.Move()
	if !haveTargets {
		return false
	}

	if op := u.Opponent(); op != nil {
		u.Attack(op)
	}

	return true
}

func Filter(units []*Unit, test func(u *Unit) bool) []*Unit {
	var result []*Unit
	for i := range units {
		if test(units[i]) {
			result = append(result, units[i])
		}
	}
	return result
}
