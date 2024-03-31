package model

type Location struct {
	Latitude  float64
	Longitude float64
}

func NewLocation(lat, long float64) *Location {
	return &Location{
		Latitude:  lat,
		Longitude: long,
	}
}
