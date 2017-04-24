package crs

import (
	"math"

	"github.com/scoutred/geo-tools/geo"
	"github.com/scoutred/geo-tools/geometry"
)

type ProjectTransformer interface {
	Projecter
	Transformer
}

type Projecter interface {
	Project(geo.LatLng) geometry.Point
	UnProject(geometry.Point) geo.LatLng
}

type Transformer interface {
	Transform(geometry.Point, float64) geometry.Point
	UnTransform(geometry.Point, float64) geometry.Point
}

// converts geo coords to pixel ones
func LatLngToPoint(crs ProjectTransformer, latLng geo.LatLng, zoom float64) geometry.Point {
	projectedPoint := crs.Project(latLng)
	scale := Scale(zoom)

	return crs.Transform(projectedPoint, scale)
}

// converts pixel coords to geo coords
func PointToLatLng(crs ProjectTransformer, point geometry.Point, zoom float64) geo.LatLng {
	scale := Scale(zoom)
	untransformedPoint := crs.UnTransform(point, scale)

	return crs.UnProject(untransformedPoint)
}

// defines how the world scales with zoom
func Scale(zoom float64) float64 {
	return 256 * math.Pow(2, zoom)
}

func Zoom(scale float64) float64 {
	return math.Log(scale/256) / math.Ln2
}

//	https://gist.github.com/perrygeo/4478844
func MetersPerPixel(zoom, lat float64) float64 {
	return (math.Cos(lat*math.Pi/180.0) * 2 * math.Pi * 6378137) / (256 * math.Pow(2, zoom))
}
