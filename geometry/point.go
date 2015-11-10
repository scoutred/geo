package geometry

import (
	"math"
)

func NewPoint(x, y float64) Point {
	return Point{
		X: x,
		Y: y,
	}
}

type Point struct {
	X float64
	Y float64
}

func (this Point) Add(p Point) Point {
	return Point{
		X: this.X + p.X,
		Y: this.Y + p.Y,
	}
}

func (this Point) Floor() Point {
	return Point{
		X: math.Floor(this.X),
		Y: math.Floor(this.Y),
	}
}

func (this Point) Subtract(p Point) Point {
	return Point{
		X: this.X - p.X,
		Y: this.Y - p.Y,
	}
}

func (this Point) DivideBy(num float64) Point {
	return Point{
		X: this.X / num,
		Y: this.Y / num,
	}
}

func (this Point) MultiplyBy(num float64) Point {
	return Point{
		X: this.X * num,
		Y: this.Y * num,
	}
}

func (this Point) Contains(p Point) bool {
	return math.Abs(p.X) <= math.Abs(this.X) &&
		math.Abs(p.Y) <= math.Abs(this.Y)
}
