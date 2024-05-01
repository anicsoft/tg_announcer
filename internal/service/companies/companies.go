package companies

import (
	"anik/internal/client/db"
	"anik/internal/model"
	"anik/internal/repository"
	"anik/internal/service"
	"context"
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

func (s *serv) Create(ctx context.Context, company *model.Company) (int, error) {
	var id int
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.companiesRepo.Create(ctx, company)
		if errTx != nil {
			return errTx
		}

		for _, category := range company.Category {
			errTx = s.companiesRepo.AddCategory(ctx, category, id)
			if errTx != nil {
				return errTx
			}
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *serv) Get(ctx context.Context, id string) (*model.Company, error) {
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

func (s *serv) Delete(ctx context.Context, id int) error {
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
