package users

import (
	"anik/internal/client/db"
	"anik/internal/model"
	"anik/internal/repository"
	"anik/internal/service"
	"context"
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

func (s *serv) AddUser(ctx context.Context, user *model.User) (int, error) {
	create, err := s.usersRepo.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return create, nil
}

func (s *serv) Exists(ctx context.Context, id int) (bool, error) {
	user, err := s.usersRepo.Get(ctx, id)
	if err != nil {
		return false, err
	}

	if user != nil {
		return true, nil
	}

	return false, nil
}
