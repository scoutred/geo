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
			expected: geo.NewLatLng(32.73052630874809, -117.18018306080502),
		},
	}

	for i, tc := range testcases {
		point := crs.PointToLatLng(tc.crs, tc.point, tc.zoom)
		if point.Lat != tc.expected.Lat {
			t.Errorf("test (%v) failed. expected Lat (%v) does not match output Lat (%v)", i, tc.expected.Lat, point.Lat)
		}

		if point.Lng != tc.expected.Lng {
			t.Errorf("test (%v) failed. expected Lng (%v) does not match output Lng (%v)", i, tc.expected.Lng, point.Lng)
		}
	}
}

func TestMetersPerPixel(t *testing.T) {
	testcases := []struct {
		zoom     float64
		lat      float64
		expected float64
	}{
		{
			zoom:     5.0,
			lat:      45.0,
			expected: 3459.1450261885484,
		},
		{
			zoom:     12.0,
			lat:      0.0,
			expected: 38.21851414258813,
		},
		{
			zoom:     12.0,
			lat:      79.0,
			expected: 7.292436288331513,
		},
		{
			zoom:     13.0,
			lat:      0.0,
			expected: 19.109257071294063,
		},
	}

	for i, tc := range testcases {
		output := crs.MetersPerPixel(tc.zoom, tc.lat)

		if output != tc.expected {
			t.Errorf("testcase (%v) failed. output (%v) does not match expected (%v)", i, output, tc.expected)
		}
	}
}
