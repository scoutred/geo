package geo

func NewLatLngBounds(minLat, minLng, maxLat, maxLng float64) LatLngBounds {
	return LatLngBounds{
		NE: LatLng{
			Lat: minLat,
			Lng: minLng,
		},
		SW: LatLng{
			Lat: maxLat,
			Lng: maxLng,
		},
	}
}

//	represents a rectangular area on the map in geographical coordinates.
type LatLngBounds struct {
	//	north east
	NE LatLng
	//	south west
	SW LatLng
}

//	center of the bounds
func (this *LatLngBounds) Center() LatLng {
	return LatLng{
		Lat: (this.SW.Lat + this.NE.Lat) / 2,
		Lng: (this.SW.Lng + this.NE.Lng) / 2,
	}
}

func (this *LatLngBounds) West() float64 {
	return this.SW.Lng
}

func (this *LatLngBounds) South() float64 {
	return this.SW.Lat
}

func (this *LatLngBounds) East() float64 {
	return this.NE.Lng
}

func (this *LatLngBounds) North() float64 {
	return this.NE.Lat
}

func (this *LatLngBounds) SouthWest() LatLng {
	return LatLng{
		Lat: this.South(),
		Lng: this.West(),
	}
}

func (this *LatLngBounds) NorthEast() LatLng {
	return LatLng{
		Lat: this.North(),
		Lng: this.East(),
	}
}

func (this *LatLngBounds) NorthWest() LatLng {
	return LatLng{
		Lat: this.North(),
		Lng: this.West(),
	}
}

func (this *LatLngBounds) SouthEast() LatLng {
	return LatLng{
		Lat: this.South(),
		Lng: this.East(),
	}
}
