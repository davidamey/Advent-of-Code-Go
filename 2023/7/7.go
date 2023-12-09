package main

import (
	"advent-of-code-go/util"
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {
	defer util.Duration(time.Now())

	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	hands := make([]*hand, len(lines))
	for i, l := range lines {
		hands[i] = newHand(l)
	}

	fmt.Println("p1=", score(hands, func(i, j int) bool {
		if hands[i].kind == hands[j].kind {
			for x := range hands[i].cards {
				ci, cj := hands[i].cards[x], hands[j].cards[x]
				if ci == cj {
					continue
				}
				return strength(ci, false) < strength(cj, false)
			}
		}
		return hands[i].kind < hands[j].kind
	}))

	fmt.Println("p2=", score(hands, func(i, j int) bool {
		if hands[i].bestKind == hands[j].bestKind {
			for x := range hands[i].cards {
				ci, cj := hands[i].cards[x], hands[j].cards[x]
				if ci == cj {
					continue
				}
				return strength(ci, true) < strength(cj, true)
			}
		}
		return hands[i].bestKind < hands[j].bestKind
	}))
}

func score(hands []*hand, less func(i, j int) bool) (winnings int) {
	sort.Slice(hands, less)
	for i, h := range hands {
		winnings += (i + 1) * h.bid
	}
	return
}

type handType uint8

const (
	HighCard handType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type hand struct {
	cards    [5]rune
	bid      int
	kind     handType
	bestKind handType
}

func newHand(s string) *hand {
	h := &hand{}
	parts := strings.Fields(s)
	for i, r := range parts[0] {
		h.cards[i] = r
	}
	h.bid = util.Atoi(parts[1])

	counts := make(map[rune]int)
	for _, c := range h.cards {
		counts[c]++
	}

	switch len(counts) {
	case 1: // Five of a kind
		h.kind = FiveOfAKind
	case 2: // Four of a kind or full house
		hasFourTheSame := false
		for _, v := range counts {
			if v == 4 {
				hasFourTheSame = true
			}
		}
		if hasFourTheSame {
			h.kind = FourOfAKind
		} else {
			h.kind = FullHouse
		}
	case 3: // Three of a kind or Two pair
		hasThreeTheSame := false
		for _, v := range counts {
			if v == 3 {
				hasThreeTheSame = true
			}
		}
		if hasThreeTheSame {
			h.kind = ThreeOfAKind
		} else {
			h.kind = TwoPair
		}
	case 4: // One pair
		h.kind = OnePair
	case 5: // High card
		h.kind = HighCard
	}

	h.bestKind = h.kind

	jokers := counts['J']
	if jokers == 0 { // No Jokers, can't improve
		return h
	}

	switch h.kind {
	case FiveOfAKind: // Can't get better
	case FourOfAKind: // 1 or 4 jokers
		h.bestKind = FiveOfAKind
	case FullHouse: // 2 or 3 jokers
		h.bestKind = FiveOfAKind
	case ThreeOfAKind: // 1 or 3 jokers
		h.bestKind = FourOfAKind
	case TwoPair: // 1 or 2 jokers
		if jokers == 2 {
			h.bestKind = FourOfAKind
		} else {
			h.bestKind = FullHouse
		}
	case OnePair:
		h.bestKind = ThreeOfAKind
	case HighCard:
		h.bestKind = OnePair
	}

	return h
}

func strength(r rune, p2 bool) int {
	switch r {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		if p2 {
			return 1
		} else {
			return 11
		}
	case 'T':
		return 10
	default:
		return int(r - '0')
	}
}

func (h hand) String() string {
	var sb strings.Builder
	sb.WriteString(string(h.cards[:]))
	sb.WriteRune(' ')
	switch h.kind {
	case 6:
		sb.WriteString("Five of a kind")
	case 5:
		sb.WriteString("Four of a kind")
	case 4:
		sb.WriteString("Full house")
	case 3:
		sb.WriteString("Three of a kind")
	case 2:
		sb.WriteString("Two pair")
	case 1:
		sb.WriteString("One pair")
	case 0:
		sb.WriteString("High card")
	}
	return sb.String()
}
