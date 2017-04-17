package crs_test

import (
	"testing"

	"github.com/scoutred/geo-tools/geo"
	"github.com/scoutred/geo-tools/geo/crs"
	"github.com/scoutred/geo-tools/geometry"
)

func TestScale(t *testing.T) {
	scale := crs.Scale(2)
	if scale != 1024 {
		t.Error("invalid scaling result")
	}
}

func TestZoom(t *testing.T) {
	zoom := crs.Zoom(2)
	if zoom != -7 {
		t.Error("invalid zoom result")
	}
}

func TestLatLngToPoint(t *testing.T) {
	testcases := []struct {
		crs      crs.ProjectTransformer
		latLng   geo.LatLng
		zoom     float64
		expected geometry.Point
	}{
		{
			crs:      crs.NewEspg3857(),
			latLng:   geo.NewLatLng(32.7305263087481, -117.180183060805),
			zoom:     13.0,
			expected: geometry.NewPoint(365951.95759351854, 846601.1035181411),
		},
	}

	for i, tc := range testcases {
		point := crs.LatLngToPoint(tc.crs, tc.latLng, tc.zoom)
		if point.X != tc.expected.X {
			t.Errorf("test (%v) failed. expected X (%v) does not match output X (%v)", i, tc.expected.X, point.X)
			return
		}

		if point.Y != tc.expected.Y {
			t.Errorf("test (%v) failed. expected Y (%v) does not match output Y (%v)", i, tc.expected.Y, point.Y)
			return
		}
	}
}

func TestPointLatLng(t *testing.T) {
	testcases := []struct {
		crs      crs.ProjectTransformer
		point    geometry.Point
		zoom     float64
		expected geo.LatLng
	}{
		{
			crs:      crs.NewEspg3857(),
			point:    geometry.NewPoint(365951.95759351854, 846601.1035181411),
			zoom:     13.0,
			expected: geo.NewLatLng(32.7305263087481, -117.180183060805),
		},
	}

	for i, tc := range testcases {
		point := crs.PointToLatLng(tc.crs, tc.point, tc.zoom)
		if point.Lat != tc.expected.Lat {
			t.Errorf("test (%v) failed. expected Lat (%v) does not match output Lat (%v)", i, tc.expected.Lat, point.Lat)
			return
		}

		if point.Lng != tc.expected.Lng {
			t.Errorf("test (%v) failed. expected Lng (%v) does not match output Lng (%v)", i, tc.expected.Lng, point.Lng)
			return
		}
	}
}
