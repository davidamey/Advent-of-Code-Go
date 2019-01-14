package main

import (
	"advent/util"
	"fmt"
	"math"
)

const Debug = false

var spells []*spell = []*spell{
	&spell{"Magic Missile", 53, func(g *game) { g.boss.hp -= 4 }},           // 0
	&spell{"Drain", 73, func(g *game) { g.boss.hp -= 2; g.player.hp += 2 }}, // 1
	&spell{"Shield", 113, func(g *game) { g.effects.shield = 6 }},           // 2
	&spell{"Poison", 173, func(g *game) { g.effects.poison = 6 }},           // 3
	&spell{"Recharge", 229, func(g *game) { g.effects.recharge = 5 }},       // 4
}

func main() {
	// input := NewGame(10, 250, 14, 8) // example
	input := NewGame(50, 500, 71, 10, false)
	playerWon, manaUsed := input.clone().simulate(nil)
	if !playerWon {
		panic("p1: player never won!")
	}
	fmt.Println("p1=", manaUsed)

	input.hardMode = true
	playerWon, manaUsed = input.simulate(nil)
	if !playerWon {
		panic("p2: player never won!")
	}
	fmt.Println("p2=", manaUsed)

	// Example 2
	// input.playerTurn(spells[4]) // recharge
	// input.bossTurn()
	// input.playerTurn(spells[2]) // shield
	// input.bossTurn()
	// input.playerTurn(spells[1]) // drain
	// input.bossTurn()
	// input.playerTurn(spells[3]) // poison
	// input.bossTurn()
	// input.playerTurn(spells[0]) // magic missile
	// input.bossTurn()
}

func (g *game) simulate(s *spell) (playerWon bool, manaUsed int) {
	if s != nil {
		if done := g.playerTurn(s); done {
			return g.player.hp > 0, g.manaUsed
		}
		if done := g.bossTurn(); done {
			return g.player.hp > 0, g.manaUsed
		}
	}

	manaUsed = math.MaxInt32
	for _, s := range spells {
		if g.manaUsed+s.cost >= manaUsed {
			continue
		}

		if win, mana := g.clone().simulate(s); win && mana < manaUsed {
			playerWon = true
			manaUsed = mana
		}
	}

	return
}

type game struct {
	player, boss being
	effects      effects
	hardMode     bool
	manaUsed     int
	spellChain   []string
}

func NewGame(pHp, pMana, bHp, bAtk int, hardMode bool) *game {
	return &game{
		being{pHp, 0, 0, pMana},
		being{bHp, bAtk, 0, 0},
		effects{0, 0, 0},
		hardMode,
		0,
		nil,
	}
}

func (g *game) printStatus() {
	def := g.player.def
	if g.effects.shield > 0 {
		def = 7
	}
	logf("- Player has %d hit points, %d armor, %d mana\n", g.player.hp, def, g.player.mana)
	logf("- Boss has %d hit points, %d attack\n", g.boss.hp, g.boss.atk)
}

func (g *game) clone() *game {
	return &game{
		g.player,
		g.boss,
		g.effects,
		g.hardMode,
		g.manaUsed,
		g.spellChain,
	}
}

func (g *game) playerTurn(s *spell) (done bool) {
	defer logln()

	logln("-- Player turn --")
	g.printStatus()

	if g.hardMode { // part 2
		g.player.hp -= 1
		if g.shouldEnd() {
			return true
		}
	}

	g.applyEffects()
	if g.shouldEnd() {
		return true
	}

	switch {
	case s.cost > g.player.mana,
		s.name == "Shield" && g.effects.shield > 0,
		s.name == "Poison" && g.effects.poison > 0,
		s.name == "Recharge" && g.effects.recharge > 0:
		g.player.hp = 0 // pretend we've died
		return true
	}

	logf("Player casts %s\n", s.name)
	g.spellChain = append(g.spellChain, s.name)
	g.player.mana -= s.cost
	g.manaUsed += s.cost
	s.cast(g)
	if g.shouldEnd() {
		return true
	}

	return false
}

func (g *game) bossTurn() (done bool) {
	defer logln()

	logln("-- Boss turn --")
	g.printStatus()

	g.applyEffects()
	if g.shouldEnd() {
		return true
	}

	dmg := util.MaxInt(g.boss.atk-g.player.def, 1)
	logf("Boss attacks for %d damage.\n", dmg)
	g.player.hp -= dmg
	if g.shouldEnd() {
		return true
	}

	return false
}

func (g *game) applyEffects() {
	if g.effects.shield > 0 {
		g.player.def = 7
		g.effects.shield--
		logf("Shield's timer is now %d.\n", g.effects.shield)
		if g.effects.shield == 0 {
			g.player.def = 0
			logln("Shield wears off, decreasing armor by 7.")
		}
	}

	if g.effects.poison > 0 {
		g.boss.hp -= 3
		g.effects.poison--
		logf("Poison deals 3 damage; its timer is now %d.\n", g.effects.poison)
		if g.effects.poison == 0 {
			logln("Poison wears off")
		}
	}

	if g.effects.recharge > 0 {
		g.player.mana += 101
		g.effects.recharge--
		logf("Recharge provides 101 mana; its timer is now %d.\n", g.effects.recharge)
		if g.effects.recharge == 0 {
			logln("Recharge wears off")
		}
	}
}

func (g *game) shouldEnd() bool {
	return g.player.hp <= 0 || g.boss.hp <= 0
}

type being struct {
	hp, atk, def, mana int
}

type spell struct {
	name string
	cost int
	cast func(g *game)
}

type effects struct {
	shield   int
	poison   int
	recharge int
}

func logln(a ...interface{}) {
	if Debug {
		fmt.Println(a...)
	}
}

func logf(format string, a ...interface{}) {
	if Debug {
		fmt.Printf(format, a...)
	}
}
