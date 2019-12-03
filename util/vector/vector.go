package vector

import (
	"advent-of-code-go/util"
	"fmt"
	"math"
)

type Vec struct {
	X, Y int
}

func New(x, y int) Vec {
	return Vec{x, y}
}
func NewMax() Vec {
	return Vec{math.MaxInt32, math.MaxInt32}
}
func NewMin() Vec {
	return Vec{math.MinInt32, math.MinInt32}
}

func (v Vec) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}

func (v *Vec) IsOrigin() bool {
	return v.X == 0 && v.Y == 0
}

func (v *Vec) EqualTo(v2 Vec) bool {
	return v.X == v2.X && v.Y == v2.Y
}

func (v *Vec) Left() Vec {
	return Vec{v.X - 1, v.Y}
}
func (v *Vec) Right() Vec {
	return Vec{v.X + 1, v.Y}
}
func (v *Vec) Up() Vec {
	return Vec{v.X, v.Y - 1}
}
func (v *Vec) Down() Vec {
	return Vec{v.X, v.Y + 1}
}

func (v *Vec) Manhattan() int {
	return util.AbsInt(v.X) + util.AbsInt(v.Y)
}

func (v *Vec) ManhattanTo(to Vec) int {
	return util.AbsInt(to.X-v.X) + util.AbsInt(to.Y-v.Y)
}

func (v *Vec) Adjacent(withCorners bool) []Vec {
	if withCorners {
		return []Vec{
			Vec{v.X - 1, v.Y - 1},
			Vec{v.X, v.Y - 1},
			Vec{v.X + 1, v.Y - 1},
			Vec{v.X - 1, v.Y},
			Vec{v.X + 1, v.Y},
			Vec{v.X - 1, v.Y + 1},
			Vec{v.X, v.Y + 1},
			Vec{v.X + 1, v.Y + 1},
		}
	}
	return []Vec{
		Vec{v.X, v.Y - 1},
		Vec{v.X - 1, v.Y},
		Vec{v.X + 1, v.Y},
		Vec{v.X, v.Y + 1},
	}
}

func (v *Vec) Within(b1, b2 Vec) bool {
	return v.X >= b1.X && v.X <= b2.X &&
		v.Y >= b1.Y && v.Y <= b2.Y
}

func (v1 *Vec) Add(v2 Vec) {
	v1.X += v2.X
	v1.Y += v2.Y
}

func (v1 *Vec) Subtract(v2 Vec) {
	v1.X -= v2.X
	v1.Y -= v2.Y
}
