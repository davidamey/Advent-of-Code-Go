package main

import (
	"advent/util"
	"fmt"
	"image"
	"image/color"
	"math"
)

func main() {
	// file, _ := util.OpenExample()
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	points := make([]Point, len(lines))
	for i, l := range lines {
		var x, y, vx, vy int
		fmt.Sscanf(l, "position=<%d, %d> velocity=<%d, %d>", &x, &y, &vx, &vy)
		points[i] = Point{x, y, vx, vy}
	}

	rect := RectForPoints(points)
	i := 0
	for {
		i++
		Move(points)
		next := RectForPoints(points)
		if next.Dx()*next.Dy() > rect.Dx()*rect.Dy() {
			MoveBack(points)
			break
		}
		rect = next
		// fmt.Println("new area", rect.Dx()*rect.Dy())
	}

	// fmt.Printf("final size of %dx%d after %d turns\n", rect.Dx(), rect.Dy(), i)
	fmt.Printf("Message found after %d seconds:\n\n", i-1)
	Print(rect, points)
}

type Point struct {
	X, Y, VX, VY int
}

func Move(points []Point) {
	for i := range points {
		points[i].X += points[i].VX
		points[i].Y += points[i].VY
	}
}

func MoveBack(points []Point) {
	for i := range points {
		points[i].X -= points[i].VX
		points[i].Y -= points[i].VY
	}
}

func RectForPoints(points []Point) image.Rectangle {
	minX := math.MaxInt16
	minY := math.MaxInt16
	maxX := math.MinInt16
	maxY := math.MinInt16
	for _, p := range points {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X + 1
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y + 1
		}
	}
	return image.Rect(minX, minY, maxX, maxY)
}

func Print(rect image.Rectangle, points []Point) {
	im := image.NewGray(rect)
	for _, p := range points {
		im.SetGray(p.X, p.Y, color.Gray{Y: 255})
	}

	fmt.Print(" ")
	for i, p := range im.Pix {
		if p == 255 {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
		if (i+1)%rect.Dx() == 0 {
			fmt.Print("\n ")
		}
	}
	fmt.Println()
}
