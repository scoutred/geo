package geo

import (
	"testing"
)

func TestLatLngBoundsCenter(t *testing.T) {
	llb := NewLatLngBounds(
		-117.166048186108, 32.7199686756585,
		-117.16572277941, 32.7201589855535,
	)

	center := llb.Center()

	if center.Lat != -117.165885482759 || center.Lng != 32.720063830606 {
		t.Errorf("center failed: %v", center)
	}
}
