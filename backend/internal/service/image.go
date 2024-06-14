package service

import (
	"context"
	"log"
	"tg_announcer/internal/client/db"
	"tg_announcer/internal/repository"
)

type serv struct {
	imageRepo repository.ImageRepository
	txManager db.TxManager
}

func New(
	imageRepo repository.ImageRepository,
	txManager db.TxManager,
) ImageService {
	return &serv{
		imageRepo: imageRepo,
		txManager: txManager,
	}
}

func (s serv) UploadLogo(ctx context.Context, companyId string, paths string) error {
	id, err := s.imageRepo.AddLogo(ctx, companyId, paths)
	if err != nil {
		return err
	}

	log.Println("id of the uploaded logo: ", id)
	return nil
}

func (s serv) GetAnnouncPictures(ctx context.Context, announcementId string) ([]string, error) {
	_, err := s.imageRepo.GetAnnouncementPictures(ctx, announcementId)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s serv) UploadAnnouncPictures(ctx context.Context, announcementId string, paths []string) error {
	_, err := s.imageRepo.AddAnnouncementPictures(ctx, announcementId, paths)
	if err != nil {
		return err
	}

	return nil
}

func (s serv) GetLogo(ctx context.Context, parentId string) (string, error) {
	url, err := s.imageRepo.GetLogo(ctx, parentId)
	if err != nil {
		return "", err
	}

	return url, nil
}
