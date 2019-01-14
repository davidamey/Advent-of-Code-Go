package util

import "math"

type Vec struct {
	X, Y int
}

func NewVec(x, y int) Vec {
	return Vec{x, y}
}
func NewMaxVec() Vec {
	return Vec{math.MaxInt32, math.MaxInt32}
}
func NewMinVec() Vec {
	return Vec{math.MinInt32, math.MinInt32}
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

func (v *Vec) ManhattanTo(to Vec) int {
	return AbsInt(to.X-v.X) + AbsInt(to.Y-v.Y)
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
