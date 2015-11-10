package geo

import (
	"github.com/scoutred/geo-tools/geo"
	"github.com/scoutred/geo-tools/geo/crs"
	"github.com/scoutred/geo-tools/geometry"
)

func CenterBoundsZoom(proj crs.ProjectTransformer, bounds geo.LatLngBounds, size geometry.Point, maxZoom float64) (geo.LatLng, float64) {

	//	calculate our zoom
	zoom := BoundsZoom(proj, bounds, size, maxZoom)

	//	convert to points
	swPoint := crs.LatLngToPoint(proj, bounds.SouthWest(), zoom)
	nePoint := crs.LatLngToPoint(proj, bounds.NorthEast(), zoom)

	//	find center
	center := crs.PointToLatLng(proj, swPoint.Add(nePoint).DivideBy(2), zoom)

	return center, zoom
}

//	returns the zoom level for supplied bounds
//	useful when rendering static map images
//	pass in 0.0 for maxzoom for no maxzoom
//
//	TODO: add padding support
func BoundsZoom(proj crs.ProjectTransformer, bounds geo.LatLngBounds, size geometry.Point, maxZoom float64) float64 {
	zoom := 0.0
	nw := bounds.NorthWest()
	se := bounds.SouthEast()

	zoomNotFound := true

	for {
		zoom++
		boundsSize := crs.LatLngToPoint(proj, se, zoom).Subtract(crs.LatLngToPoint(proj, nw, zoom)).Floor()
		zoomNotFound = size.Contains(boundsSize)

		if zoomNotFound && zoom <= maxZoom {
			continue
		} else {
			break
		}
	}

	//	if a maxZoom was set and zoom is greater than maxZoom, return maxZoom
	if maxZoom != 0.0 && zoom > maxZoom {
		return maxZoom
	}

	//	return the calculated zoom
	return zoom - 1.0
}
