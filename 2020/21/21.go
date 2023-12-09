package main

import (
	"advent-of-code-go/util"
	"fmt"
	"sort"
	"strings"
	"time"
)

type allergen struct {
	name        string
	ingredients []string
}

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	var allIngredients []string
	allAllergens := make(map[string]*allergen)

	for _, l := range lines {
		parts := strings.Split(l, " (contains ")

		ingredients := strings.Split(parts[0], " ")
		allergens := strings.Split(strings.TrimSuffix(parts[1], ")"), ", ")

		allIngredients = append(allIngredients, ingredients...)

		for _, a := range allergens {
			if v, exists := allAllergens[a]; exists {
				allAllergens[a].ingredients = util.Intersect(v.ingredients, ingredients)
			} else {
				allAllergens[a] = &allergen{a, ingredients}
			}
		}
	}

	unsafe := make(map[string]bool)
	for _, a := range allAllergens {
		for _, i := range a.ingredients {
			unsafe[i] = true
		}
	}

	p1 := 0
	for _, i := range allIngredients {
		if !unsafe[i] {
			p1++
		}
	}
	fmt.Println("p1=", p1)

	var known []*allergen
	for len(allAllergens) > 0 {
		for k, a := range allAllergens {
			if len(a.ingredients) > 1 {
				continue
			}

			known = append(known, a)
			delete(allAllergens, k)
			for _, x := range allAllergens {
				util.Remove(&x.ingredients, a.ingredients[0])
			}
		}
	}

	sort.Slice(known, func(i, j int) bool { return known[i].name < known[j].name })
	var p2 strings.Builder
	p2.WriteString(known[0].ingredients[0])
	for _, a := range known[1:] {
		p2.WriteString(",")
		p2.WriteString(a.ingredients[0])
	}
	fmt.Println("p2=", p2.String())

}
