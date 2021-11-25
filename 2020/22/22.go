package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// input := string(util.MustReadFile("example"))
	// input := string(util.MustReadFile("exampleRecursive"))
	input := string(util.MustReadFile("input"))

	players := strings.Split(input, "\n\n")
	d1 := newDeck(strings.Split(players[0], "\n")[1:])
	d2 := newDeck(strings.Split(players[1], "\n")[1:])

	p1 := play(d1, d2)
	_, p2 := playRecursive(d1, d2)

	fmt.Println("p1=", p1)
	fmt.Println("p2=", p2)
}

func play(d1, d2 deck) (score int) {
	for !d1.empty() && !d2.empty() {
		c1, c2 := d1.takeCard(), d2.takeCard()
		if c1 > c2 {
			d1.addCards(c1, c2)
		} else {
			d2.addCards(c2, c1)
		}
	}

	if d1.empty() {
		return d2.score()
	}
	return d1.score()
}

func playRecursive(d1, d2 deck) (winner, score int) {
	seenBefore := make(map[string]bool)
	for !d1.empty() && !d2.empty() {
		deckHash := d1.String() + "|" + d2.String()
		if seenBefore[deckHash] {
			return 1, d1.score()
		}
		seenBefore[deckHash] = true

		c1, c2 := d1.takeCard(), d2.takeCard()
		if len(d1) >= c1 && len(d2) >= c2 {
			d1c, d2c := d1.copy(c1), d2.copy(c2)
			roundWinner, _ := playRecursive(d1c, d2c)
			if roundWinner == 1 {
				d1.addCards(c1, c2)
			} else {
				d2.addCards(c2, c1)
			}
			continue
		}

		if c1 > c2 {
			d1.addCards(c1, c2)
		} else {
			d2.addCards(c2, c1)
		}
	}

	if d1.empty() {
		return 2, d2.score()
	}
	return 1, d1.score()
}

type deck []int

func newDeck(cards []string) deck {
	d := make(deck, len(cards))
	for i, l := range cards {
		c, _ := strconv.Atoi(l)
		d[i] = c
	}
	return d
}

func (d *deck) copy(n int) deck {
	d2 := make(deck, n)
	for i := 0; i < n; i++ {
		d2[i] = (*d)[i]
	}
	return d2
}

func (d *deck) String() string {
	switch len(*d) {
	case 0:
		return ""
	case 1:
		return strconv.Itoa((*d)[0])
	}

	var sb strings.Builder
	sb.WriteString(strconv.Itoa((*d)[0]))
	for _, c := range (*d)[1:] {
		sb.WriteString(",")
		sb.WriteString(strconv.Itoa(c))
	}
	return sb.String()
}

func (d deck) score() (s int) {
	for i, c := range d {
		s += c * (len(d) - i)
	}
	return
}

func (d deck) empty() bool {
	return len(d) == 0
}

func (d *deck) takeCard() (c int) {
	c, *d = (*d)[0], (*d)[1:]
	return
}

func (d *deck) addCards(cards ...int) {
	*d = append(*d, cards...)
}
