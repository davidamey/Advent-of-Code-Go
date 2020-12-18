package main

import (
	"advent-of-code-go/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	tileSize        = 10
	trimmedTileSize = tileSize - 2
)

type dir uint8

const (
	dN dir = iota
	dE
	dS
	dW
)

func (d dir) String() string {
	return []string{"N", "E", "S", "W"}[d]
}

func main() {
	defer util.Duration(time.Now())

	// raw := string(util.MustReadFile("example"))
	raw := string(util.MustReadFile("input"))

	var tiles []*tile
	for _, tileRaw := range strings.Split(raw, "\n\n") {
		tiles = append(tiles, newTile(tileRaw))
	}

	// borders will have count of 1
	edgeCount := make(map[edge]int)
	for _, t := range tiles {
		for d := range t.edges {
			edgeCount[t.edges[d]]++
			edgeCount[t.segde[d]]++
		}
	}

	// find corners = 2 border edges
	var centres, corners, edges []*tile
	for _, t := range tiles {
		for d := range t.edges {
			if b := t.edges[d]; edgeCount[b] == 1 {
				t.borders = append(t.borders, b)
			} else if b := t.segde[d]; edgeCount[b] == 1 {
				t.borders = append(t.borders, b)
			}
		}
		switch len(t.borders) {
		case 0:
			centres = append(centres, t)
		case 1:
			edges = append(edges, t)
		case 2:
			corners = append(corners, t)
		default:
			panic("unexpected edge number")
		}
	}

	// fmt.Printf("tiles: %d\ncentres: %d\nedges: %d\ncorners: %d\n", len(tiles), len(centres), len(edges), len(corners))

	p1 := 1
	for _, c := range corners {
		p1 *= c.id
	}

	g, seaCount := assembleGrid(centres, edges, corners)

	var monCount int
	for orient := 0; orient < 8; orient++ { // 4 rotations + flipped 4 rotations
		monCount = findMonster(g, orient)
		if monCount > 0 {
			break
		}
	}

	// fmt.Println()
	// for y := range g {
	// 	for x, c := range g[y] {
	// 		if x > 0 && x%(tileSize-2) == 0 {
	// 			// fmt.Print(" ")
	// 		}
	// 		fmt.Printf("%c", c)
	// 	}
	// 	fmt.Println()
	// 	if y > 0 && y%(tileSize-2) == 0 {
	// 		// fmt.Println()
	// 	}
	// }

	fmt.Println(seaCount, monCount, seaCount-(monCount*15))

	fmt.Println("p1=", p1)
	fmt.Println("p2=", seaCount-(monCount*len(monsterOffsets)))
}

type point struct{ x, y int }

//                   #
// #    ##    ##    ###
//  #  #  #  #  #  #
var monsterOffsets = [15]point{
	{18, 0},
	{0, 1}, {5, 1}, {6, 1}, {11, 1}, {12, 1}, {17, 1}, {18, 1}, {19, 1},
	{1, 2}, {4, 2}, {7, 2}, {10, 2}, {13, 2}, {16, 2},
}

func findMonster(g [][]byte, orient int) (monCount int) {
	for y := range g {
	search:
		for x := range g[y] {
			for _, m := range monsterOffsets {
				dx, dy := m.x, m.y

				if orient >= 4 {
					dy = -dy
				}

				switch orient % 4 {
				case 0: // noop
				case 1: // ccw90
					dx, dy = -dy, dx
				case 2: // ccw180
					dx, dy = -dx, -dy
				case 3: // ccw270
					dx, dy = dy, -dx
				}

				gx, gy := x+dx, y+dy

				if gx < 0 || gy < 0 || gy >= len(g) || gx >= len(g[gy]) || g[gy][gx] != '#' {
					// No monster, keep looking
					continue search
				}
			}

			// Found a monster!
			monCount++
		}
	}
	return
}

func assembleGrid(centre, edges, corners []*tile) (g [][]byte, seaCount int) {
	gridSize := 2 + len(edges)/4

	// fmt.Println("gridSize=", gridSize)

	tiles := make([][]*tile, gridSize)
	for ty := range tiles {
		tiles[ty] = make([]*tile, gridSize)

		for tx := range tiles[ty] {
			if ty == 0 && tx == 0 {
				tiles[ty][tx] = corners[0].orientAsTopLeft()
				corners[0].placed = true
				continue
			}

			var e edge
			var td dir
			pool := centre

			if tx == 0 {
				td, e = dN, tiles[ty-1][tx].edges[dS]
			} else {
				td, e = dW, tiles[ty][tx-1].edges[dE]
			}

			if tx == 0 || tx == gridSize-1 || ty == 0 || ty == gridSize-1 {
				pool = edges
			}

			if (tx == 0 && ty == gridSize-1) ||
				(tx == gridSize-1 && ty == 0) ||
				(tx == gridSize-1 && ty == gridSize-1) {
				pool = corners
			}

			tiles[ty][tx] = findAndOrientTile(pool, td, e)
		}
	}

	g = make([][]byte, gridSize*trimmedTileSize)
	for y := range g {
		g[y] = make([]byte, gridSize*trimmedTileSize)

		for x := range g[y] {
			ty, dy := divmod(y, trimmedTileSize)
			tx, dx := divmod(x, trimmedTileSize)

			g[y][x] = tiles[ty][tx].lines[dy+1][dx+1]
			if g[y][x] == '#' {
				seaCount++
			}
		}

	}

	return
}

func (t *tile) orientAsTopLeft() *tile {
	if len(t.borders) != 2 {
		panic("not a valid corner")
	}

	// rotate until one edge is north and the other east or west
	rotated := func() bool {
		a := t.edges[dN] == t.borders[0] && (t.edges[dW] == t.borders[1] || t.segde[dW] == t.borders[1])
		b := t.edges[dN] == t.borders[1] && (t.edges[dW] == t.borders[0] || t.segde[dW] == t.borders[0])
		c := t.segde[dN] == t.borders[0] && (t.edges[dE] == t.borders[1] || t.segde[dE] == t.borders[1])
		d := t.segde[dN] == t.borders[1] && (t.edges[dE] == t.borders[0] || t.segde[dE] == t.borders[0])
		return a || b || c || d
	}

	for !rotated() {
		t.cw90()
	}

	if t.segde[dN] == t.borders[0] || t.segde[dN] == t.borders[1] {
		t.flipEW()
	}

	return t
}

func findAndOrientTile(pool []*tile, td dir, e edge) (t *tile) {
	f := 0
	var d dir
search:
	for _, t = range pool {
		if t.placed {
			continue
		}

		for d = 0; d < 4; d++ {
			if t.edges[d] == e || t.segde[d] == e {
				f++
				break search
			}
		}
	}

	if f > 1 {
		panic("FOUND MULTIPLE")
	}

	if d == 4 {
		panic(fmt.Sprintf("could not find a match for ", td, string(e[:])))
	}

	for i := 0; dir(i) < (td-d)&3; i++ {
		t.cw90()
	}

	if t.segde[td] == e {
		switch td {
		case dN, dS:
			t.flipEW()
		case dE, dW:
			t.flipNS()
		}
	}

	t.placed = true

	return
}

type edge [tileSize]byte

func (e edge) String() string {
	return string(e[:])
}

type tile struct {
	id           int
	lines        [tileSize][tileSize]byte
	edges, segde [4]edge
	borders      []edge
	n, s, e, w   *tile
	placed       bool
}

func newTile(raw string) *tile {
	lines := strings.Split(raw, "\n")
	id, _ := strconv.Atoi(lines[0][5:9])
	t := &tile{id: id}

	lines = lines[1:]

	for i := 0; i < tileSize; i++ {
		for j := 0; j < tileSize; j++ {
			t.lines[i][j] = lines[i][j]
		}
	}

	t.edges[dN] = t.lines[0]
	t.edges[dS] = t.lines[tileSize-1]

	for i := 0; i < tileSize; i++ {
		t.edges[dE][i] = t.lines[i][tileSize-1]
		t.edges[dW][i] = t.lines[i][0]
	}

	t.segde = [4]edge{
		rev(t.edges[dN]),
		rev(t.edges[dE]),
		rev(t.edges[dS]),
		rev(t.edges[dW]),
	}

	return t
}

func (t *tile) String() string {
	var sb strings.Builder
	// sb.WriteString("Tile: ")
	// sb.WriteString(strconv.Itoa(t.id))
	// sb.WriteString("\n")

	for _, l := range t.lines {
		for _, b := range l {
			sb.WriteByte(b)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (t *tile) cw90() {
	var rlines [tileSize][tileSize]byte
	for i := 0; i < tileSize; i++ {
		for j := 0; j < tileSize; j++ {
			rlines[i][j] = t.lines[tileSize-j-1][i]
		}
	}

	t.lines = rlines

	t.edges[dN], t.edges[dS], t.edges[dE], t.edges[dW],
		t.segde[dN], t.segde[dS], t.segde[dE], t.segde[dW] =
		t.segde[dW], t.segde[dE], t.edges[dN], t.edges[dS],
		t.edges[dW], t.edges[dE], t.segde[dN], t.segde[dS]

}

func (t *tile) flipNS() {
	for i := 0; i < tileSize/2; i++ {
		t.lines[i], t.lines[tileSize-1-i] = t.lines[tileSize-1-i], t.lines[i]
	}

	t.edges[dN], t.edges[dS] = t.edges[dS], t.edges[dN]
	t.segde[dN], t.segde[dS] = t.segde[dS], t.segde[dN]

	t.edges[dE], t.edges[dW], t.segde[dE], t.segde[dW] =
		t.segde[dE], t.segde[dW], t.edges[dE], t.edges[dW]
}

func (t *tile) flipEW() {
	for r := range t.lines {
		for i := 0; i < tileSize/2; i++ {
			t.lines[r][i], t.lines[r][tileSize-1-i] = t.lines[r][tileSize-1-i], t.lines[r][i]
		}
	}

	t.edges[dE], t.edges[dW] = t.edges[dW], t.edges[dE]
	t.segde[dE], t.segde[dW] = t.segde[dW], t.segde[dE]

	t.edges[dN], t.edges[dS], t.segde[dN], t.segde[dS] =
		t.segde[dN], t.segde[dS], t.edges[dN], t.edges[dS]
}

func rev(a edge) (b edge) {
	for i := 0; i < tileSize/2; i++ {
		b[i], b[tileSize-1-i] = a[tileSize-1-i], a[i]
	}
	return
}

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return
}
