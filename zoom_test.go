package geo

import (
	"testing"

	"github.com/scoutred/geo-tools/geo"
	"github.com/scoutred/geo-tools/geo/crs"
	"github.com/scoutred/geo-tools/geometry"
)

func TestCetnterBoundsZoom(t *testing.T) {
	bounds := geo.NewLatLngBounds(
		-117.166048186108, 32.7199686756585,
		-117.16572277941, 32.7201589855535,
	)

	size := geometry.NewPoint(1107, 360)

	proj := crs.NewEspg3857()

	//	first test should max out at 30
	center, zoom := CenterBoundsZoom(proj, bounds, size, 30.0)
	if zoom != 22.0 {
		t.Errorf("boundZoom failed to zoom correctly: %v", zoom)
	}

	if center.Lat != -85.05112877979998 || center.Lng != 32.72006383060597 {
		t.Errorf("center calulated incorrectly: %+v", center)
	}
}

func TestBoundsZoom(t *testing.T) {
	bounds := geo.NewLatLngBounds(
		-117.166048186108, 32.7199686756585,
		-117.16572277941, 32.7201589855535,
	)

	size := geometry.Point{
		X: 1107,
		Y: 360,
	}

	proj := crs.NewEspg3857()

	//	first test should max out at 30
	zoom := BoundsZoom(proj, bounds, size, 30.0)
	if zoom != 22.0 {
		t.Error("boundZoom failed to zoom correctly")
	}

	//	make sure the max zoom works
	zoom2 := BoundsZoom(proj, bounds, size, 19.0)
	if zoom2 != 19.0 {
		t.Error("boundZoom failed to zoom correctly")
	}
}
