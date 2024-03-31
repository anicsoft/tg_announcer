package service

import (
	"anik/internal/model"
	"anik/internal/repository"
	"context"
)

type CompaniesService interface {
	Create(ctx context.Context, company *model.Company) (string, error)
	Get(ctx context.Context, id string) (*model.Company, error)
	GetAll(ctx context.Context) ([]model.Company, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, company *model.Company) error
	NearbyLocations(ctx context.Context, location *model.Location) ([]model.CompanyWithDist, error)
}

type serv struct {
	repo repository.CompaniesRepository
}

func NewService(repo repository.CompaniesRepository) CompaniesService {
	return &serv{repo: repo}
}
