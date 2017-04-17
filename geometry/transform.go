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

func (this *Transformation) Transform(p Point, scale float64) Point {
	if scale == 0.0 {
		scale = 1.0
	}

	return Point{
		X: scale * (this.A*p.X + this.B),
		Y: scale * (this.C*p.Y + this.D),
	}
}

func (this *Transformation) UnTransform(p Point, scale float64) Point {
	if scale == 0.0 {
		scale = 1.0
	}

	return Point{
		X: (p.X/scale - this.B) / this.A,
		Y: (p.Y/scale - this.D) / this.C,
	}
}
