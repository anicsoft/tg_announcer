package service

import (
	"anik/internal/model"
	"context"
)

func (s *serv) Create(ctx context.Context, company *model.Company) (string, error) {
	id, err := s.repo.Create(ctx, company)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *serv) Get(ctx context.Context, id string) (*model.Company, error) {
	company, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (s *serv) GetAll(ctx context.Context) ([]model.Company, error) {
	companies, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (s *serv) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return err
}

func (s *serv) Update(ctx context.Context, company *model.Company) error {
	err := s.repo.Update(ctx, company)
	if err != nil {
		return err
	}

	return nil
}
