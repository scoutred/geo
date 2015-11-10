package geometry

import (
	"testing"
)

func TestTransform(t *testing.T) {
	point := NewPoint(
		178.7281379699707,
		413.41481781005854,
	)

	trans := NewTransformation(
		2.495320233665337e-08,
		0.5,
		-2.495320233665337e-08,
		0.5,
	)

	p := trans.Transform(point, 2.0)

	if p.X != 1.00000891967878 || p.Y != 0.9999793679528044 {
		t.Errorf("transform failed: %v", p)
	}
}

func TestUnTransform(t *testing.T) {
	point := NewPoint(
		1.00000891967878,
		0.9999793679528044,
	)

	trans := NewTransformation(
		2.495320233665337e-08,
		0.5,
		-2.495320233665337e-08,
		0.5,
	)

	p := trans.UnTransform(point, 2.0)

	if p.X != 178.72813796981737 || p.Y != 413.4148178106803 {
		t.Errorf("untransform failed: %v", p)
	}
}
