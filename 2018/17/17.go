package main

import (
	"advent-of-code-go/util"
	"advent-of-code-go/util/vector"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"time"
)

func main() {
	start := time.Now()
	defer func() { fmt.Printf("** completed in %fs\n", time.Since(start).Seconds()) }()

	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	g := NewGrid()
	g.Set(vector.Vec{X: 500, Y: 0}, '+')
	minY := math.MaxInt32
	maxY := math.MinInt32
	for _, l := range lines {
		var r1, r2 rune
		var v1, v2, v3 int

		fmt.Sscanf(l, "%c=%d, %c=%d..%d", &r1, &v1, &r2, &v2, &v3)

		if r1 == 'x' {
			if v2 < minY {
				minY = v2
			}
			if v3 > maxY {
				maxY = v3
			}

			for y := v2; y <= v3; y++ {
				g.Set(vector.Vec{X: v1, Y: y}, '#')
			}
		} else {
			if v1 < minY {
				minY = v1
			}
			if v1 > maxY {
				maxY = v1
			}

			for x := v2; x <= v3; x++ {
				g.Set(vector.Vec{X: x, Y: v1}, '#')
			}
		}
	}

	// g.Print(true)
	// fmt.Println(g.minY)
	// os.Exit(0)

	flow := []vector.Vec{
		vector.Vec{X: 500, Y: 1},
	}

	for ; len(flow) > 0; dump(g) {
		w := flow[0]
		flow = flow[1:]
		g.current = &w

		if g.Blocked(w) || g.OutOfBounds(w) {
			continue
		}
		g.Set(w, '|')

		d := w.Down()
		if !g.Blocked(d) {
			flow = append(flow, d)
			continue
		}

		// Flow left and right as far as possible
		// fmt.Printf("flowing left from (%d, %d)=%c\n", w.X, w.Y, g.Get(w))
		var row []vector.Vec
		contained := true
		l := w.Left()
		for !g.Blocked(l) {
			// fmt.Printf("(%d, %d)=%c not blocked\n", l.X, l.Y, g.Get(l))
			g.Set(l, '|')
			row = append(row, l)
			if !g.Blocked(l.Down()) {
				// fmt.Printf("(%d, %d)=%c not blocked below\n", l.X, l.Y, g.Get(l))
				flow = append(flow, l.Down())
				contained = false
				break
			}
			l = l.Left()
		}
		// fmt.Printf("found block at (%d, %d)=%c\n", l.X, l.Y, g.Get(l))
		// fmt.Scanln()
		r := w.Right()
		for !g.Blocked(r) {
			g.Set(r, '|')
			row = append(row, r)
			if !g.Blocked(r.Down()) {
				flow = append(flow, r.Down())
				contained = false
				break
			}
			r = r.Right()
		}

		if !contained {
			continue
		}

		g.Set(w, '~')
		for _, p := range row {
			g.Set(p, '~')
		}

		// g.Print(true)
		// fmt.Printf("appending up from (%d, %d)=%c to (%d, %d)=%c\n", w.X, w.Y, g.Get(w), w.Up().X, w.Up().Y, g.Get(w.Up()))
		// fmt.Scanln()
		flow = append(flow, w.Up())
	}

	waterCount := 0
	standingCount := 0
	for y := minY; y <= maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			w := g.Get(vector.Vec{X: x, Y: y})
			if w == '~' || w == '|' {
				waterCount++
			}
			if w == '~' {
				standingCount++
			}
		}
	}

	g.SaveToFile()
	fmt.Println("part1=", waterCount)
	fmt.Println("part2=", standingCount)
}

func dump(g *Grid) {
	// if pngCount > 99 && pngCount < 200 {
	// 	g.SaveToFile()
	// 	return
	// }
	// pngCount++

	g.Print(true)
	fmt.Scanln()
}

type Grid struct {
	current                *vector.Vec
	minX, minY, maxX, maxY int
	entries                map[string]rune
}

func NewGrid() *Grid {
	return &Grid{
		minX:    math.MaxInt32,
		minY:    math.MaxInt32,
		maxX:    math.MinInt32,
		maxY:    math.MinInt32,
		entries: make(map[string]rune),
	}
}

func (g *Grid) Get(p vector.Vec) rune {
	if val, ok := g.entries[fmt.Sprintf("%d,%d", p.X, p.Y)]; ok {
		return val
	}
	return '.'
}

func (g *Grid) Set(p vector.Vec, val rune) {
	// fmt.Printf("setting (%d, %d) to %c\n", x, y, val)
	if p.X < g.minX {
		g.minX = p.X
	}
	if p.Y < g.minY {
		g.minY = p.Y
	}
	if p.X > g.maxX {
		g.maxX = p.X
	}
	if p.Y > g.maxY {
		g.maxY = p.Y
	}
	g.entries[fmt.Sprintf("%d,%d", p.X, p.Y)] = val
}

func (g *Grid) Sand(p vector.Vec) bool {
	r := g.Get(p)
	return r == '.' // || r == '|'
}

func (g *Grid) Flowing(p vector.Vec) bool {
	return g.Get(p) == '|'
}

func (g *Grid) Blocked(p vector.Vec) bool {
	return g.Get(p) == '~' || g.Get(p) == '#'
}

func (g *Grid) OutOfBounds(p vector.Vec) bool {
	// Cheeky <= as the water source is at 0
	return p.Y <= g.minY || p.Y > g.maxY
}

func (g *Grid) Print(clear bool) {
	if clear {
		fmt.Printf("\033[0;0H")
		fmt.Printf("\033[2J")
	}

	fmt.Println()
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX - 1; x <= g.maxX+1; x++ {
			p := vector.Vec{X: x, Y: y}
			if g.current != nil && x == g.current.X && y == g.current.Y {
				fmt.Printf("\033[%d;%dm%c\033[0m", 92, 49, g.Get(p))
				continue
			}
			fmt.Printf("%c", g.Get(p))
		}
		fmt.Println()
	}
}

var pngCount = 0

func (g *Grid) SaveToFile() {
	img := image.NewRGBA(image.Rect(g.minX-1, g.minY, g.maxX+1, g.maxY))

	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX - 1; x <= g.maxX+1; x++ {
			p := vector.Vec{X: x, Y: y}
			switch g.Get(p) {
			case '+':
				img.Set(x, y, color.RGBA{0, 255, 0, 255})
			case '|':
				img.Set(x, y, color.RGBA{255, 255, 255, 255})
			case '~':
				img.Set(x, y, color.RGBA{0, 0, 255, 255})
			case '#':
				img.Set(x, y, color.RGBA{255, 0, 0, 255})
			}
		}
	}

	// f, _ := os.OpenFile(fmt.Sprintf("21_png/%d.png", pngCount), os.O_WRONLY|os.O_CREATE, 0600)
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
	pngCount++
}
