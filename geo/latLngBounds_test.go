package geo_test

import (
	"testing"

	"github.com/scoutred/geo-tools/geo"
)

func TestLatLngBoundsCenter(t *testing.T) {
	testcases := []struct {
		p1             geo.LatLng
		p2             geo.LatLng
		expectedCenter geo.LatLng
	}{
		{
			p1:             geo.NewLatLng(-117.180183060805, 32.6963180459813),
			p2:             geo.NewLatLng(-117.147086641189, 32.7305263087481),
			expectedCenter: geo.NewLatLng(-117.163634850997, 32.7134221773647),
		},
	}

	for i, tc := range testcases {
		llb := geo.NewLatLngBounds(tc.p1, tc.p2)
		center := llb.Center()

		if center.Lat != tc.expectedCenter.Lat || center.Lng != tc.expectedCenter.Lng {
			t.Errorf("test (%v) failed. expected (%+v) does not match output (%+v)", i, tc.expectedCenter, center)
			return
		}
	}
}
