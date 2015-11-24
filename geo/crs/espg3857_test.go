package crs

import (
	"testing"

	"github.com/scoutred/geo-tools/geo"
	"github.com/scoutred/geo-tools/geometry"
)

func TestLatLngToPointESPG3857(t *testing.T) {
	latLng := geo.NewLatLng(
		32.720064477996,
		-117.16588899493217,
	)

	//	coordinate ref system
	crs := NewEspg3857()

	//	no scale
	p := LatLngToPoint(crs, latLng, 0.0)
	if p.X != 44.68203449249269 || p.Y != 103.35370445251465 {
		t.Errorf("espg3857 failed to convert latLng to point (scale of 0.0): %v", p)
	}

	//	convert to point with 2x scale
	p = LatLngToPoint(crs, latLng, 2.0)
	if p.X != 178.72813796997076 || p.Y != 413.4148178100586 {
		t.Errorf("espg3857 failed to convert latLng to point (scale of 2.0): %v", p)
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
		32.720064477996,
		-117.16588899493217,
	)

	p := NewEspg3857().Project(latLng)

	if p.X != -13042847.101257065 || p.Y != 3858205.880090526 {
		t.Errorf("ESPG3857 project failed: %+v", p)
	}
}

func TestUnProjectESPG3857(t *testing.T) {
	point := geometry.NewPoint(
		-13042847.101257065,
		3858205.880090526,
	)

	latLng := NewEspg3857().UnProject(point)

	if latLng.Lat != 32.720064477996 || latLng.Lng != -117.16588899493217 {
		t.Errorf("ESPG3857 unproject failed: %+v", latLng)
	}
}
