package announcements

import (
	apiModel "anik/internal/api/model"
	"anik/internal/client/db"
	"anik/internal/model"
	"anik/internal/repository"
	"anik/internal/service"
	"context"
	"math"
	"sort"
)

const earthRadius = 6371

type serv struct {
	announcementRepo repository.AnnouncementRepository
	txManager        db.TxManager
}

type DistanceCalculator interface {
	Distance(loc1, loc2 *model.Location) float64
}

type DistCalculator struct {
}

func New(
	announcementRepo repository.AnnouncementRepository,
	txManager db.TxManager,
) service.AnnouncementService {
	return &serv{
		announcementRepo: announcementRepo,
		txManager:        txManager,
	}
}

func (s *serv) Create(ctx context.Context, announcement *model.Announcement) (int, error) {
	var id int
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error
		id, txErr = s.announcementRepo.Create(ctx, announcement)
		if txErr != nil {
			return txErr
		}

		for _, category := range announcement.Categories {
			txErr = s.announcementRepo.AddCategory(ctx, category, id)
			if txErr != nil {
				return txErr
			}
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *serv) Get(ctx context.Context, id int) (*model.Announcement, error) {
	announcement, err := s.announcementRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return announcement, nil
}

func (s *serv) GetAll(ctx context.Context) ([]model.Announcement, error) {
	announcements, err := s.announcementRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return announcements, nil
}

func (s *serv) GetFiltered(ctx context.Context, filter apiModel.Filter) ([]model.Announcement, error) {
	announcements, err := s.announcementRepo.GetFiltered(ctx, filter)
	if err != nil {
		return nil, err
	}

	if filter.Latitude != 0 && filter.Longitude != 0 {
		//lat, err := strconv.ParseFloat(filter.Latitude, 64)
		//if err != nil {
		//	log.Println(err)
		//}
		//
		//long, err := strconv.ParseFloat(filter.Longitude, 64)
		//if err != nil {
		//	log.Println(err)
		//}

		userLoc := model.NewLocation(filter.Latitude, filter.Longitude)

		distanceCalculator := DistCalculator{}
		for i := range announcements {
			companyLoc := model.NewLocation(announcements[i].Company.Latitude, announcements[i].Company.Longitude)
			dist := distanceCalculator.Distance(userLoc, companyLoc)
			//log.Println("DISTANCE:", dist)
			announcements[i].Company.DistToUser = dist
		}

		sort.Slice(announcements, func(i, j int) bool {
			return announcements[i].Company.DistToUser < announcements[j].Company.DistToUser
		})
	}

	return announcements, nil
}

func (s *serv) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s *serv) Update(ctx context.Context, announcement *model.Announcement) error {
	//TODO implement me
	panic("implement me")
}

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
