package geometry

import "math"

type Bounds struct {
	min Point
	max Point
}

func NewBounds(a, b Point) Bounds {
	var bounds Bounds

	bounds.extend(a)
	bounds.extend(b)

	return bounds
}

func (b *Bounds) extend(p Point) {
	if b.min.IsZero() && b.max.IsZero() {
		b.min = p
		b.max = p
	} else {
		b.min.X = math.Min(p.X, b.min.X)
		b.max.X = math.Max(p.X, b.max.X)
		b.min.Y = math.Min(p.Y, b.min.Y)
		b.max.Y = math.Max(p.Y, b.max.Y)
	}
}

func (b *Bounds) Center() Point {
	return Point{
		X: (b.min.X + b.max.X) / 2,
		Y: (b.min.Y + b.max.Y) / 2,
	}
}

func (b *Bounds) Size() Point {
	return b.max.Subtract(b.min)
}
