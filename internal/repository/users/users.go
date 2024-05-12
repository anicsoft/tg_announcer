package users

import (
	"anik/internal/client/db"
	"anik/internal/model"
	"anik/internal/repository"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"log"
)

const (
	tableName          = "users"
	idColumn           = "id"
	firstNameColumn    = "first_name"
	lastNameColumn     = "last_name"
	usernameColumn     = "username"
	latitudeColumn     = "latitude"
	longitudeColumn    = "longitude"
	languageCodeColumn = "language_code"
	userTypeColumn     = "user_type"
	createdAtColumn    = "created_at"
)

type repo struct {
	db db.Client
}

func New(db db.Client) repository.UsersRepository {
	return &repo{
		db: db,
	}
}

func (r repo) Create(ctx context.Context, user *model.User) (int, error) {
	const op = "repository.Create user"

	builder := squirrel.Insert(tableName).
		PlaceholderFormat(repository.PlaceHolder).
		Columns(
			idColumn,
			firstNameColumn,
			lastNameColumn,
			usernameColumn,
			latitudeColumn,
			longitudeColumn,
			languageCodeColumn,
			userTypeColumn,
		).
		Values(
			user.Id,
			user.FirstName,
			user.LastName,
			user.Username,
			user.Latitude,
			user.Longitude,
			user.LanguageCode,
			"user",
		).
		Suffix("RETURNING " + idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%w: %v", repository.ErrBuildQuery, err)
		log.Println(err)
		return 0, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var id int
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		err := fmt.Errorf("%w, %v : %v", repository.ErrExecQuery, op, err)
		log.Println(err)
		return 0, err
	}

	return id, nil
}

func (r repo) Get(ctx context.Context, id int) (*model.User, error) {
	const op = "repository.Get user"

	builder := squirrel.Select(
		idColumn,
		firstNameColumn,
		lastNameColumn,
		usernameColumn,
		latitudeColumn,
		longitudeColumn,
		languageCodeColumn,
		userTypeColumn,
		createdAtColumn,
	).
		PlaceholderFormat(repository.PlaceHolder).
		From(tableName).
		Where(squirrel.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%w, %v : %v", repository.ErrBuildQuery, op, err)
		log.Println(err)
		return nil, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var user model.User
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Latitude,
		&user.Longitude,
		&user.LanguageCode,
		&user.UserType,
		&user.CreatedAt,
	); err != nil {
		err := fmt.Errorf("%w, %v : %v", repository.ErrExecQuery, op, err)
		log.Println(err)
		return nil, err
	}

	return &user, nil
}
