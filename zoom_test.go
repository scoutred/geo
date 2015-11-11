package geo

import (
	"testing"

	"github.com/scoutred/geo-tools/geo"
	"github.com/scoutred/geo-tools/geo/crs"
	"github.com/scoutred/geo-tools/geometry"
)

func TestCetnterBoundsZoom(t *testing.T) {
	bounds := geo.NewLatLngBounds(
		32.7204706651118, -117.16439634561537,
		32.71965828903011, -117.1673735976219,
	)

	size := geometry.NewPoint(1107, 360)

	proj := crs.NewEspg3857()

	//	first test should max out at 30
	center, zoom := CenterBoundsZoom(proj, bounds, size, 30.0)
	if zoom != 18.0 {
		t.Errorf("boundZoom failed to zoom correctly: %v", zoom)
	}

	if center.Lat != 32.720064477996 || center.Lng != -117.16588497161865 {
		t.Errorf("center calulated incorrectly: %+v", center)
	}
}

func TestBoundsZoom(t *testing.T) {
	bounds := geo.NewLatLngBounds(
		32.7204706651118, -117.16439634561537,
		32.71965828903011, -117.1673735976219,
	)

	size := geometry.NewPoint(1107, 360)

	proj := crs.NewEspg3857()

	//	first test should max out at 30
	zoom := BoundsZoom(proj, bounds, size, 30.0)
	if zoom != 18.0 {
		t.Error("boundZoom failed to zoom correctly")
	}

	//	make sure the max zoom works
	zoom = BoundsZoom(proj, bounds, size, 16.0)
	if zoom != 16.0 {
		t.Error("boundZoom failed to zoom correctly")
	}
}
