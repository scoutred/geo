//	Spherical Mercator is the most common CRS for web mapping.
package sphericalMercator

import (
	"math"

	"github.com/scoutred/geo-tools/geo"
	"github.com/scoutred/geo-tools/geometry"
)

const (
	R           = 6378137
	MaxLatitude = 85.0511287798
)

type SphericalMercator struct{}

func (this *SphericalMercator) Project(latLng geo.LatLng) geometry.Point {
	d := math.Pi / 180
	lat := math.Max(math.Min(MaxLatitude, latLng.Lat), -MaxLatitude)
	sin := math.Sin(lat * d)

	return geometry.Point{
		X: R * latLng.Lng * d,
		Y: R * math.Log((1+sin)/(1-sin)) / 2,
	}
}

func (this *SphericalMercator) UnProject(point geometry.Point) geo.LatLng {
	d := 180 / math.Pi

	return geo.LatLng{
		Lat: (2*math.Atan(math.Exp(point.Y/R)) - (math.Pi / 2)) * d,
		Lng: point.X * d / R,
	}
}
