package service

import (
	"anik/internal/model"
	"context"
	"math"
	"sort"
)

const earthRadius = 6371

type DistanceCalculator interface {
	Distance(loc1, loc2 *model.Location) float64
}

type DistCalculator struct {
}

func (s *serv) NearbyLocations(
	ctx context.Context,
	location *model.Location,
) ([]model.CompanyWithDist, error) {
	all, err := s.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	userLocation := model.NewLocation(location.Latitude, location.Longitude)

	var companiesWithDist []model.CompanyWithDist
	distanceCalculator := DistCalculator{}

	for _, company := range all {
		companyLocation := model.NewLocation(company.Latitude, company.Longitude)
		dist := distanceCalculator.Distance(userLocation, companyLocation)
		companiesWithDist = append(companiesWithDist, model.NewCompanyWithDistance(&company, dist))
	}

	sort.Slice(companiesWithDist, func(i, j int) bool {
		return companiesWithDist[i].DistToUser < companiesWithDist[j].DistToUser
	})

	return companiesWithDist, nil
}

// Distance returns distance between two points in meters with three decimal places.
func (d *DistCalculator) Distance(loc1, loc2 *model.Location) float64 {
	// Convert latitude and longitude from degrees to radians
	lat1Rad := degreesToRadians(loc1.Latitude)
	lon1Rad := degreesToRadians(loc1.Longitude)
	lat2Rad := degreesToRadians(loc2.Latitude)
	lon2Rad := degreesToRadians(loc2.Longitude)

	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(dLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := earthRadius * c * 1000

	return roundFloat(distance)
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// roundFloat rounds float to three digits after the decimal point.
func roundFloat(value float64) float64 {
	roundedValue := math.Round(value*1000) / 1000
	return roundedValue
}
