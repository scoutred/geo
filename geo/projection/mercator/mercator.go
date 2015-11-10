package mercator

import (
	"math"

	"github.com/scoutred/geo-tools/geometry"
)

const (
	R       = 6378137
	R_MINOR = 6356752.314245179
)

func Project(latLng geo.LatLng) geometry.Point {
	d := math.Pi / 180
	y := latLng.Lat * d
	tmp := R_MINOR / R
	e := math.Sqrt(1 - tmp*tmp)
	con := e * math.Sin(y)

	ts := math.Tan(math.Pi/4-y/2) / math.Pow((1-con)/(1+con), e/2)
	y = -R * math.Log(math.Max(ts, 1E-10))

	return geometry.Point{
		X: int64(latLng.Lng * d * R),
		Y: int64(y),
	}
}
