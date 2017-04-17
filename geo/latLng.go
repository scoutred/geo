package geo

func NewLatLng(lat, lng float64) LatLng {
	return LatLng{
		Lat: lat,
		Lng: lng,
	}
}

type LatLng struct {
	Lat float64
	Lng float64
}

func (ll *LatLng) IsZero() bool {
	return ll.Lat == 0.0 && ll.Lng == 0.0
}
