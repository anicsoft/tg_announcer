package repository

import (
	"anik/internal/client/db"
	"anik/internal/model"
	"context"
)

type CompaniesRepository interface {
	Create(ctx context.Context, company *model.Company) (string, error)
	Get(ctx context.Context, id string) (*model.Company, error)
	GetAll(ctx context.Context) ([]model.Company, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, company *model.Company) error
}

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) CompaniesRepository {
	return &repo{db: db}
}
