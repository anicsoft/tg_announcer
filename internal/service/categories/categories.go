package categories

import (
	"anik/internal/model"
	"anik/internal/repository"
	"anik/internal/service"
	"context"
)

type serv struct {
	repo repository.CategoriesRepository
}

func New(repo repository.CategoriesRepository) service.CategoriesService {
	return &serv{
		repo: repo,
	}
}

func (s *serv) AddBusinessCategory(ctx context.Context, category *model.Category) (int, error) {
	id, err := s.repo.AddBusinessCategory(ctx, category.Name)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *serv) AddOfferCategory(ctx context.Context, category *model.Category) (int, error) {
	id, err := s.repo.AddOfferCategory(ctx, category.Name)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *serv) GetBusinessCategories(ctx context.Context) ([]model.Category, error) {
	categories, err := s.repo.GetBusinessCategories(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *serv) GetOfficerCategories(ctx context.Context) ([]model.Category, error) {
	categories, err := s.repo.GetOfferCategories(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
