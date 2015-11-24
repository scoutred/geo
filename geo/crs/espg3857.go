//	Spherical Mercator is the most common CRS for web mapping.
package crs

import (
	"math"

	"github.com/scoutred/geo-tools/geo"
	"github.com/scoutred/geo-tools/geometry"
)

//	create a new ESPG3587 CRS
func NewEspg3857() espg3857 {
	return espg3857{
		r:           6378137,
		maxLatitude: 85.0511287798,
	}
}

type espg3857 struct {
	r           float64
	maxLatitude float64
}

//	fetch the value of R
func (this espg3857) R() float64 {
	return this.r
}

func (this espg3857) MaxLatitude() float64 {
	return this.maxLatitude
}

func (this espg3857) Transform(point geometry.Point, scale float64) geometry.Point {
	//	transfomration scale
	tScale := 0.5 / (math.Pi * this.r)

	//	new transformation
	t := geometry.NewTransformation(tScale, 0.5, -tScale, 0.5)

	//	transform our point with the provided scale
	return t.Transform(point, scale)
}

//	Spherical Mercator
func (this espg3857) Project(latLng geo.LatLng) geometry.Point {
	d := math.Pi / 180
	max := this.maxLatitude
	lat := math.Max(math.Min(max, latLng.Lat), -max)
	sin := math.Sin(lat * d)

	return geometry.Point{
		X: this.r * latLng.Lng * d,
		Y: this.r * math.Log((1+sin)/(1-sin)) / 2,
	}
}

func (this espg3857) UnTransform(point geometry.Point, scale float64) geometry.Point {
	//	transfomration scale
	tScale := 0.5 / (math.Pi * this.r)

	//	new transformation
	t := geometry.NewTransformation(tScale, 0.5, -tScale, 0.5)

	//	transform our point with the provided scale
	return t.UnTransform(point, scale)
}

//	Spherical Mercator
func (this espg3857) UnProject(point geometry.Point) geo.LatLng {
	d := 180 / math.Pi

	return geo.LatLng{
		Lat: (2*math.Atan(math.Exp(point.Y/this.r)) - (math.Pi / 2)) * d,
		Lng: point.X * d / this.r,
	}
}
