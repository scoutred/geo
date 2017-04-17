package geo

import "math"

func NewLatLngBounds(a, b LatLng) LatLngBounds {
	var bounds LatLngBounds

	bounds.extend(a)
	bounds.extend(b)

	return bounds
}

//	represents a rectangular area on the map in geographical coordinates.
type LatLngBounds struct {
	//	north east
	ne LatLng
	//	south west
	sw LatLng
}

func (llb *LatLngBounds) extend(p LatLng) {
	if llb.sw.IsZero() && llb.ne.IsZero() {
		llb.sw = p
		llb.ne = p
	} else {
		llb.sw.Lat = math.Min(p.Lat, llb.sw.Lat)
		llb.sw.Lng = math.Min(p.Lng, llb.sw.Lng)
		llb.ne.Lat = math.Max(p.Lat, llb.ne.Lat)
		llb.ne.Lng = math.Max(p.Lng, llb.ne.Lng)
	}
}

//	center of the bounds
func (llb *LatLngBounds) Center() LatLng {
	return LatLng{
		Lat: (llb.sw.Lat + llb.ne.Lat) / 2,
		Lng: (llb.sw.Lng + llb.ne.Lng) / 2,
	}
}

func (llb *LatLngBounds) West() float64 {
	return llb.sw.Lng
}

func (llb *LatLngBounds) South() float64 {
	return llb.sw.Lat
}

func (llb *LatLngBounds) East() float64 {
	return llb.ne.Lng
}

func (llb *LatLngBounds) North() float64 {
	return llb.ne.Lat
}

func (llb *LatLngBounds) SouthWest() LatLng {
	return LatLng{
		Lat: llb.South(),
		Lng: llb.West(),
	}
}

func (llb *LatLngBounds) NorthEast() LatLng {
	return LatLng{
		Lat: llb.North(),
		Lng: llb.East(),
	}
}

func (llb *LatLngBounds) NorthWest() LatLng {
	return LatLng{
		Lat: llb.North(),
		Lng: llb.West(),
	}
}

func (llb *LatLngBounds) SouthEast() LatLng {
	return LatLng{
		Lat: llb.South(),
		Lng: llb.East(),
	}
}
