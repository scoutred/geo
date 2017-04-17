package geometry_test

import (
	"testing"

	"github.com/scoutred/geo-tools/geometry"
)

func TestBoundsSize(t *testing.T) {
	testcases := []struct {
		bounds   geometry.Bounds
		expected geometry.Point
	}{
		{
			bounds: geometry.NewBounds(
				geometry.NewPoint(366144.7582118256, 846837.9484448085),
				geometry.NewPoint(365951.95759351854, 846601.1035181411),
			),
			expected: geometry.NewPoint(192.8006183070829, 236.8449266673997),
		},
	}

	for i, tc := range testcases {
		size := tc.bounds.Size()

		if size.X != tc.expected.X || size.Y != tc.expected.Y {
			t.Errorf("test (%v) failed. expected (%v) does not match output (%v)", i, tc.expected, size)
			return
		}
	}
}
