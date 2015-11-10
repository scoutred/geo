//	utility class to perform simple point transformations through a 2d-matrix.
package geometry

func NewTransformation(a, b, c, d float64) Transformation {
	return Transformation{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

type Transformation struct {
	A, B, C, D float64
}

func (this *Transformation) Transform(point Point, scale float64) Point {
	if scale == 0.0 {
		scale = 1.0
	}

	return Point{
		X: scale * (this.A*point.X + this.B),
		Y: scale * (this.C*point.Y + this.D),
	}
}

func (this *Transformation) UnTransform(point Point, scale float64) Point {
	if scale == 0.0 {
		scale = 1.0
	}

	return Point{
		X: (point.X/scale - this.B) / this.A,
		Y: (point.Y/scale - this.D) / this.C,
	}
}
