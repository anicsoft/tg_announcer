package image

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"tg_announcer/internal/client/db"
	"tg_announcer/internal/config"
	"tg_announcer/internal/repository"
	"tg_announcer/internal/service"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type serv struct {
	imageRepo repository.ImageRepository
	awsConfig config.AWSConfig
	txManager db.TxManager
}

func New(
	imageRepo repository.ImageRepository,
	awsConfig config.AWSConfig,
	txManager db.TxManager,
) service.ImageService {
	return &serv{
		imageRepo: imageRepo,
		awsConfig: awsConfig,
		txManager: txManager,
	}
}

func (s *serv) UploadLogo(ctx context.Context, companyId string, header *multipart.FileHeader) (string, error) {
	s3URL, err := s.uploadToS3(header)
	if err != nil {
		return "", err
	}

	id, err := s.imageRepo.AddLogo(ctx, companyId, s3URL)
	if err != nil {
		return "", err
	}

	log.Println("id of the uploaded logo: ", id)
	return s3URL, nil
}

func (s *serv) GetAnnouncPictures(ctx context.Context, announcementId string) ([]string, error) {
	_, err := s.imageRepo.GetAnnouncementPictures(ctx, announcementId)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *serv) UploadAnnouncPictures(ctx context.Context, announcementId string, header *multipart.FileHeader) (string, error) {
	s3URL, err := s.uploadToS3(header)
	if err != nil {
		return "", err
	}

	urls, err := s.imageRepo.AddAnnouncementPictures(ctx, announcementId, []string{s3URL})
	if err != nil {
		return "", err
	}

	return urls, nil
}

func (s *serv) GetLogo(ctx context.Context, parentId string) (string, error) {
	url, err := s.imageRepo.GetLogo(ctx, parentId)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *serv) uploadToS3(header *multipart.FileHeader) (string, error) {
	file, err := header.Open()
	if err != nil {
		return "", err
	}

	awsSession := s.awsConfig.CreateSession()
	s3Client := s.awsConfig.CreateS3Session(awsSession)

	fileKey := "uploads/" + header.Filename
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(s.awsConfig.Bucket()),
		Key:                aws.String(fileKey),
		Body:               file,
		ContentType:        aws.String(header.Header.Get("Content-Type")),
		ContentDisposition: aws.String(fmt.Sprintf("%s; %s", "inline", "filename="+header.Filename)),
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", os.Getenv("S3_BUCKET_NAME"), fileKey), nil
}
