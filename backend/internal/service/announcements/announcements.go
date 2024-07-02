package announcements

import (
	"context"
	"log"
	"tg_announcer/internal/api"
	apiModel "tg_announcer/internal/api/model"
	"tg_announcer/internal/client/db"
	"tg_announcer/internal/model"
	"tg_announcer/internal/repository"
	"tg_announcer/internal/service"
)

type serv struct {
	announcementRepo repository.AnnouncementRepository
	userRepo         repository.UsersRepository
	txManager        db.TxManager
}

func New(
	announcementRepo repository.AnnouncementRepository,
	userRepo repository.UsersRepository,
	txManager db.TxManager,
) service.AnnouncementService {
	return &serv{
		announcementRepo: announcementRepo,
		userRepo:         userRepo,
		txManager:        txManager,
	}
}

func (s *serv) Create(ctx context.Context, announcement *model.Announcement) (string, error) {
	var id string
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
		return "", err
	}

	return id, nil
}

func (s *serv) Get(ctx context.Context, id string) (*model.Announcement, error) {
	announcement, err := s.announcementRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return announcement, nil
}

func (s *serv) GetAll(ctx context.Context, filter apiModel.Filter) ([]model.Announcement, error) {
	announcements, err := s.announcementRepo.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}

	initData, ok, isGuest := api.GetInitData(ctx)
	log.Println("isGuest", isGuest)
	log.Println("ok", ok)
	log.Println("initData", initData)
	if !ok || isGuest {
		for i := range announcements {
			announcements[i].Company.IsFavorite = false
		}

		return announcements, nil
	}
	log.Println("initData.User.ID", initData.User.ID)
	userId := initData.User.ID
	for i := range announcements {
		isFavorite, err := s.userRepo.IsFavoriteCompany(ctx, int(userId), announcements[i].CompanyID)
		if err != nil {
			log.Println("Error checking if company is favorite:", err)
			announcements[i].Company.IsFavorite = false
		} else {
			announcements[i].Company.IsFavorite = isFavorite
		}
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
