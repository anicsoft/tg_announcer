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
)

type serv struct {
	usersRepo repository.UsersRepository
	txManager db.TxManager
}

func New(
	usersRepo repository.UsersRepository,
	txManager db.TxManager,
) service.UsersService {
	return &serv{
		usersRepo: usersRepo,
		txManager: txManager,
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
	create, err := s.usersRepo.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return create, nil
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
