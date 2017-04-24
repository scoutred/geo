package crs

import (
	"testing"

	"github.com/scoutred/geo-tools/geo"
	"github.com/scoutred/geo-tools/geometry"
)

func TestLatLngToPointESPG3857(t *testing.T) {
	testcases := []struct {
		latLng   geo.LatLng
		zoom     float64
		expected geometry.Point
	}{
		{
			latLng: geo.NewLatLng(
				32.720064477996,
				-117.16588899493217,
			),
			zoom: 0.0,
			expected: geometry.NewPoint(
				44.68203449249269,
				103.35370445251465,
			),
		},
		{
			latLng: geo.NewLatLng(
				32.720064477996,
				-117.16588899493217,
			),
			zoom: 2.0,
			expected: geometry.NewPoint(
				178.72813796997076,
				413.4148178100586,
			),
		},
		{
			latLng: geo.NewLatLng(
				32.71342381720106,
				-117.163634850997,
			),
			zoom: 14.0,
			expected: geometry.NewPoint(
				732096.7158053443,
				1693439.0519629498,
			),
		},
	}

	//	coordinate ref system
	crs := NewEspg3857()

	for i, tc := range testcases {
		p := LatLngToPoint(crs, tc.latLng, tc.zoom)
		if p.X != tc.expected.X || p.Y != tc.expected.Y {
			t.Errorf("testcase (%v) failed. expected (%+v) does not match output (%+v)", i, tc.expected, p)
		}
	}
}

func TestPointToLatLngESPG3857(t *testing.T) {
	point := geometry.NewPoint(
		44.68203449249269,
		103.35370445251465,
	)

	//	coordinate ref system
	crs := NewEspg3857()

	//	no scale
	latLng := PointToLatLng(crs, point, 0.0)
	if latLng.Lat != 32.720064477996 || latLng.Lng != -117.16588899493217 {
		t.Errorf("espg3857 failed to convert point to latLng: %v", latLng)
	}

	//	2x scale
	latLng = PointToLatLng(crs, point, 2.0)
	if latLng.Lat != 80.6838883930096 || latLng.Lng != -164.29147224873307 {
		t.Errorf("espg3857 failed to convert point to latLng: %v", latLng)
	}

}

func TestTransformESPG3857(t *testing.T) {
	point := geometry.NewPoint(
		44.68203449249269,
		103.35370445251465,
	)

	p := NewEspg3857().Transform(point, 2.0)

	if p.X != 1.0000022299196951 || p.Y != 0.9999948419882011 {
		t.Errorf("espg3857 failed to transform point: %+v", p)
	}
}

func TestUnTransformESPG3857(t *testing.T) {
	point := geometry.NewPoint(
		1.0000022299196951,
		0.9999948419882011,
	)

	p := NewEspg3857().UnTransform(point, 2.0)

	if p.X != 44.68203449467896 || p.Y != 103.35370445267007 {
		t.Errorf("espg3857 failed to untransform point: %+v", p)
	}
}

func TestProjectESPG3857(t *testing.T) {
	latLng := geo.NewLatLng(
		32.7305263087481,
		-117.180183060805,
	)

	p := NewEspg3857().Project(latLng)

	if p.X != -13044438.309391394 || p.Y != 3859590.2188198487 {
		t.Errorf("ESPG3857 project failed: %+v", p)
	}
}

func TestUnProjectESPG3857(t *testing.T) {
	point := geometry.NewPoint(
		-13044438.309391394,
		3859590.2188198487,
	)

	latLng := NewEspg3857().UnProject(point)

	if latLng.Lat != 32.73052630874809 || latLng.Lng != -117.180183060805 {
		t.Errorf("ESPG3857 unproject failed: %+v", latLng)
	}
}
