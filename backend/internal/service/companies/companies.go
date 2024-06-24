package companies

import (
	"context"
	"tg_announcer/internal/client/db"
	"tg_announcer/internal/model"
	"tg_announcer/internal/repository"
	"tg_announcer/internal/service"
)

type serv struct {
	companiesRepo repository.CompaniesRepository
	txManager     db.TxManager
}

func New(
	companiesRepo repository.CompaniesRepository,
	txManager db.TxManager,
) service.CompaniesService {
	return &serv{
		companiesRepo: companiesRepo,
		txManager:     txManager,
	}
}

func (s *serv) Create(ctx context.Context, company *model.Company) (string, error) {
	var id string
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.companiesRepo.Create(ctx, company)
		if errTx != nil {
			return errTx
		}

		for _, category := range company.Categories {
			errTx = s.companiesRepo.AddCategory(ctx, category, id)
			if errTx != nil {
				return errTx
			}
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *serv) GetByID(ctx context.Context, id string) (*model.Company, error) {
	company, err := s.companiesRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (s *serv) GetAll(ctx context.Context) ([]model.Company, error) {
	companies, err := s.companiesRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (s *serv) Delete(ctx context.Context, id string) error {
	err := s.companiesRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *serv) Update(ctx context.Context, company *model.Company) error {
	err := s.companiesRepo.Update(ctx, company)
	if err != nil {
		return err
	}

	return nil
}
