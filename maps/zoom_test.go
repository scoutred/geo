package maps_test

import (
	"testing"

	"github.com/scoutred/geo-tools/geo"
	"github.com/scoutred/geo-tools/geo/crs"
	"github.com/scoutred/geo-tools/geometry"
	"github.com/scoutred/geo-tools/maps"
)

func TestBoundsCenterZoom(t *testing.T) {
	testcases := []struct {
		bounds         geo.LatLngBounds
		size           geometry.Point
		maxZoom        float64
		expectedZoom   float64
		expectedCenter geo.LatLng
	}{
		{
			bounds: geo.NewLatLngBounds(
				geo.NewLatLng(32.6963180459813, -117.180183060805),
				geo.NewLatLng(32.7305263087481, -117.147086641189),
			),
			size:           geometry.NewPoint(1107, 360),
			maxZoom:        30.0,
			expectedZoom:   13.0,
			expectedCenter: geo.NewLatLng(32.71342381720108, -117.163634850997),
		},
		{
			bounds: geo.NewLatLngBounds(
				geo.NewLatLng(32.7204706651118, -117.16439634561537),
				geo.NewLatLng(32.71965828903011, -117.1673735976219),
			),
			size:           geometry.NewPoint(1107, 360),
			maxZoom:        30.0,
			expectedZoom:   18.0,
			expectedCenter: geo.NewLatLng(32.720064477996, -117.16588497161865),
		},
	}

	proj := crs.NewEspg3857()

	for i, tc := range testcases {
		//	first test should max out at 30
		center, zoom := maps.BoundsCenterZoom(proj, tc.bounds, tc.size, tc.maxZoom)
		if zoom != tc.expectedZoom {
			t.Errorf("testcase (%v) failed. zoom (%v) does not match expectedZoom (%v)", i, zoom, tc.expectedZoom)
			return
		}

		if center.Lat != tc.expectedCenter.Lat {
			t.Errorf("testcase (%v) failed. center.Lat (%v) does not match expectedCenter.Lat (%v)", i, center.Lat, tc.expectedCenter.Lat)
			return
		}

		if center.Lng != tc.expectedCenter.Lng {
			t.Errorf("testcase (%v) failed. center.Lng (%v) does not match expectedCenter.Lng (%v)", i, center.Lng, tc.expectedCenter.Lng)
			return
		}
	}
}

func TestBoundsZoom(t *testing.T) {
	testcases := []struct {
		bounds   geo.LatLngBounds
		size     geometry.Point
		maxZoom  float64
		expected float64
	}{
		{
			bounds: geo.NewLatLngBounds(
				geo.NewLatLng(32.7204706651118, -117.16439634561537),
				geo.NewLatLng(32.71965828903011, -117.1673735976219),
			),
			size:     geometry.NewPoint(862, 300),
			maxZoom:  30.0,
			expected: 18.0,
		},
	}

	proj := crs.NewEspg3857()

	for i, tc := range testcases {
		zoom := maps.BoundsZoom(proj, tc.bounds, tc.size, tc.maxZoom)

		if tc.expected != zoom {
			t.Errorf("test (%v) failed. expected (%v) does not match output (%v)", i, tc.expected, zoom)
			return
		}
	}
}

func TestZoomScale(t *testing.T) {
	testcases := []struct {
		scale    float64
		zoom     float64
		expected float64
	}{
		{
			scale:    0.04749943415844548,
			zoom:     18.0,
			expected: 13.604054137528504,
		},
	}

	proj := crs.NewEspg3857()

	for i, tc := range testcases {
		z := maps.ScaleZoom(proj, tc.scale, tc.zoom)
		if tc.expected != z {
			t.Errorf("test (%v) failed. expected (%v) does not match output (%v)", i, tc.expected, z)
			return
		}
	}
}
