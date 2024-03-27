package repository

import (
	"anik/internal/model"
	"context"
	"database/sql"
)

type CompaniesRepository interface {
	Create(ctx context.Context, company *model.Company) (string, error)
	Get(ctx context.Context, id string) (*model.Company, error)
	GetAll(ctx context.Context) ([]model.Company, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, company *model.Company) error
}

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) CompaniesRepository {
	return &repo{db: db}
}
