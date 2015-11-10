package geo

type LatLng struct {
	Lat float64
	Lng float64
}

func NewLatLng(lat, lng float64) LatLng {
	return LatLng{
		Lat: lat,
		Lng: lng,
	}
}
