package main

import (
	"advent/util"
	"fmt"
	"strings"
)

const (
	// // example
	// seekLow = 2
	// seekHigh = 5

	seekLow  = 17
	seekHigh = 61
)

func main() {
	// lines := util.MustReadFileToLines("example")
	lines := util.MustReadFileToLines("input")

	bots := make(BotCollection)
	outputs := make(OutputCollection)
	for _, l := range lines {
		if l[:5] == "value" {
			var v, bID int
			fmt.Sscanf(l, "value %d goes to bot %d", &v, &bID)
			bots.Bot(bID).AddVal(v)
			continue
		}

		var bID, lowID, highID int
		var lowT, highT string
		fmt.Sscanf(l, "bot %d gives low to %s %d and high to %s %d", &bID, &lowT, &lowID, &highT, &highID)
		b := bots.Bot(bID)
		if lowT == "output" {
			b.LowTarget = outputs.Output(lowID)
		} else {
			b.LowTarget = bots.Bot(lowID)
		}
		if highT == "output" {
			b.HighTarget = outputs.Output(highID)
		} else {
			b.HighTarget = bots.Bot(highID)
		}
	}

	p1 := -1
	for {
		acted := false
		for _, b := range bots {
			if b.HasTwoVals() {
				if b.LowVal == seekLow && b.HighVal == seekHigh {
					p1 = b.ID
				}

				b.LowTarget.AddVal(b.LowVal)
				b.HighTarget.AddVal(b.HighVal)
				b.ClearVals()

				acted = true
			}
		}
		if !acted {
			break
		}
	}

	fmt.Printf("p1= bot %d compared microchips %d and %d\n", p1, seekLow, seekHigh)
	fmt.Println("p2=", outputs[0].Val*outputs[1].Val*outputs[2].Val)
}

type ValueHolder interface {
	Name() string
	AddVal(int)
}

type Bot struct {
	ID         int
	LowVal     int
	HighVal    int
	LowTarget  ValueHolder
	HighTarget ValueHolder
}

func (b *Bot) Name() string {
	return fmt.Sprintf("Bot %d", b.ID)
}

func (b *Bot) AddVal(v int) {
	if b.HasTwoVals() {
		panic(fmt.Errorf("bot %d already has 2 values", b.ID))
	}
	switch {
	case b.LowVal == 0:
		b.LowVal = v
	case v > b.LowVal:
		b.HighVal = v
	default:
		b.HighVal = b.LowVal
		b.LowVal = v
	}

}

func (b *Bot) Vals() []int {
	return []int{b.LowVal, b.HighVal}
}

func (b *Bot) ClearVals() {
	b.LowVal = 0
	b.HighVal = 0
}

func (b *Bot) HasTwoVals() bool {
	return b.LowVal > 0 && b.HighVal > 0
}

func (b Bot) String() string {
	lowName, highName := "none", "none"
	if b.LowTarget != nil {
		lowName = b.LowTarget.Name()
	}
	if b.HighTarget != nil {
		highName = b.HighTarget.Name()
	}
	return fmt.Sprintf(`Bot %d
  Vals: %v
  Low:  %s
  High: %s`, b.ID, b.Vals(), lowName, highName)
}

type BotCollection map[int]*Bot

func (bc BotCollection) Bot(id int) *Bot {
	if b, ok := bc[id]; ok {
		return b
	}

	bc[id] = &Bot{ID: id}
	return bc[id]
}

func (bc BotCollection) String() string {
	parts := make([]string, len(bc))
	for i := 0; i < len(bc); i++ {
		parts[i] = bc[i].String()
	}
	return strings.Join(parts, "\n")
}

type Output struct {
	ID  int
	Val int
}

func (o *Output) Name() string {
	return fmt.Sprintf("Output %d", o.ID)
}

func (o *Output) AddVal(v int) {
	if o.Val != 0 {
		panic(fmt.Errorf("output %d already has a value", o.ID))
	}
	o.Val = v
}

func (o Output) String() string {
	return fmt.Sprintf(`Output %d: %d`, o.ID, o.Val)
}

type OutputCollection map[int]*Output

func (oc OutputCollection) Output(id int) *Output {
	if o, ok := oc[id]; ok {
		return o
	}

	oc[id] = &Output{ID: id}
	return oc[id]
}

func (oc OutputCollection) String() string {
	parts := make([]string, len(oc))
	for i := 0; i < len(oc); i++ {
		parts[i] = oc[i].String()
	}
	return strings.Join(parts, "\n")
}
