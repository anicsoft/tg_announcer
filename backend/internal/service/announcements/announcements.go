package announcements

import (
	"anik/internal/client/db"
	"anik/internal/model"
	"anik/internal/repository"
	"anik/internal/service"
	"context"
	"net/url"
)

type serv struct {
	announcementRepo repository.AnnouncementRepository
	txManager        db.TxManager
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

func (s *serv) GetFiltered(ctx context.Context, query url.Values) ([]model.Announcement, error) {
	var categories []string
	for _, category := range query["category"] {
		categories = append(categories, category)
	}

	announcements, err := s.announcementRepo.GetByCategory(ctx, categories)
	if err != nil {
		return nil, err
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
