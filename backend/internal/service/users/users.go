package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"tg_announcer/internal/client/db"
	"tg_announcer/internal/model"
	"tg_announcer/internal/repository"
	"tg_announcer/internal/service"

	"github.com/jackc/pgx/v4"
)

type serv struct {
	usersRepo     repository.UsersRepository
	companiesRepo repository.CompaniesRepository
	txManager     db.TxManager
}

func New(
	usersRepo repository.UsersRepository,
	companiesRepo repository.CompaniesRepository,
	txManager db.TxManager,
) service.UsersService {
	return &serv{
		usersRepo:     usersRepo,
		companiesRepo: companiesRepo,
		txManager:     txManager,
	}
}

func (s *serv) Update(ctx context.Context, user *model.User) error {
	err := s.usersRepo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *serv) GetByID(ctx context.Context, id int) (*model.User, error) {
	user, err := s.usersRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no such user with such id %d", id)
		}
		return nil, err
	}

	return user, nil
}

func (s *serv) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user, err := s.usersRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *serv) AddUser(ctx context.Context, user *model.User) (int, error) {
	existingUser, err := s.usersRepo.GetByID(ctx, user.ID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return 0, err
	}

	if existingUser != nil {
		return existingUser.ID, fmt.Errorf("user with id %d already exists", existingUser.ID)
	}

	id, err := s.usersRepo.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *serv) Exists(ctx context.Context, id int) (bool, error) {
	user, err := s.usersRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	if user != nil {
		return true, nil
	}

	return false, nil
}

func (s *serv) AddFavorite(ctx context.Context, userId int, companyId string) error {
	if err := s.usersRepo.AddFavoriteCompany(ctx, userId, companyId); err != nil {
		return errors.New("failed to add favorite company")
	}

	return nil
}

func (s *serv) DeleteFavorite(ctx context.Context, userId int, companyId string) error {
	if err := s.usersRepo.DeleteFavoriteCompany(ctx, userId, companyId); err != nil {
		return errors.New("failed to delete favorite company")
	}

	return nil
}

func (s *serv) Favorites(ctx context.Context, userId int) ([]model.Company, error) {
	var companiesList []model.Company
	if err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error
		companies, txErr := s.usersRepo.GetFavoriteCompanies(ctx, userId)
		if txErr != nil {
			return txErr
		}

		for _, companyId := range companies {
			company, txErr := s.companiesRepo.Get(ctx, companyId)
			if txErr != nil {
				return txErr
			}

			companiesList = append(companiesList, *company)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return companiesList, nil

}

func (s *serv) UserList(ctx context.Context) ([]model.User, error) {
	users, err := s.usersRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
