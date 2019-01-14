package main

import (
	"advent/util"
	"fmt"
)

const teaspoons = 100

func main() {
	// file, _ := util.OpenExample()
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	ingredients := make([]*Ingredient, len(lines))
	for i, l := range lines {
		var name string
		var cap, dur, flv, txt, cal int
		fmt.Sscanf(l, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
			&name, &cap, &dur, &flv, &txt, &cal)
		ingredients[i] = &Ingredient{name[:len(name)-1], cap, dur, flv, txt, cal}
	}

	p1 := 0
	p2 := 0
	for c := range combos(len(ingredients), 100) {
		s, cal := score(ingredients, c)
		if s > p1 {
			p1 = s
		}
		if cal == 500 && s > p2 {
			p2 = s
		}
	}
	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func score(ingredients []*Ingredient, combo []int) (int, int) {
	totals := Ingredient{}
	for i, amt := range combo {
		ing := ingredients[i]
		totals.Capacity += amt * ing.Capacity
		totals.Durability += amt * ing.Durability
		totals.Flavour += amt * ing.Flavour
		totals.Texture += amt * ing.Texture
		totals.Calories += amt * ing.Calories
	}
	score := util.MaxInt(0, totals.Capacity) *
		util.MaxInt(0, totals.Durability) *
		util.MaxInt(0, totals.Flavour) *
		util.MaxInt(0, totals.Texture)
	return score, totals.Calories
}

func combos(ingredientCount, teaspoons int) <-chan []int {
	ch := make(chan []int)
	go func() {
		defer close(ch)
		combo(ch, nil, ingredientCount, teaspoons)
	}()
	return ch
}

func combo(ch chan []int, base []int, itemCount, total int) {
	start := 0
	if itemCount == 1 {
		start = total
	}
	for i := start; i <= total; i++ {
		newBase := append(base, i)
		if itemCount-1 > 0 {
			combo(ch, newBase, itemCount-1, total-i)
		} else {
			ch <- newBase
		}
	}
}

type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavour    int
	Texture    int
	Calories   int
}
