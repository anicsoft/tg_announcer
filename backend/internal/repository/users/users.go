package users

import (
	"anik/internal/client/db"
	"anik/internal/model"
	"anik/internal/repository"
	"context"
	"database/sql"
	"errors"
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
	companyIdColumn    = "company_id"
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

func (r repo) Update(ctx context.Context, user *model.User) error {
	const op = "repository.Update user"

	builder := squirrel.Update(tableName).
		Where(squirrel.Eq{idColumn: user.Id}).
		Set(firstNameColumn, user.FirstName).
		Set(lastNameColumn, user.LastName).
		Set(usernameColumn, user.Username).
		Set(latitudeColumn, user.Latitude).
		Set(longitudeColumn, user.Longitude).
		Set(languageCodeColumn, user.LanguageCode).
		Set(userTypeColumn, user.UserType).
		Set(companyIdColumn, user.CompanyId)

	query, args, err := builder.PlaceholderFormat(repository.PlaceHolder).ToSql()
	if err != nil {
		err := fmt.Errorf("%w, %v : %v", repository.ErrBuildQuery, op, err)
		log.Println(err)
		return err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	if _, err := r.db.DB().ExecContext(ctx, q, args...); err != nil {
		err := fmt.Errorf("%w, %v : %v", repository.ErrExecQuery, op, err)
		log.Println(err)
		return err
	}

	return nil
}

func (r repo) Exists(ctx context.Context, id int) (bool, error) {
	const op = "repository.Exists"

	builder := squirrel.Select("COUNT(*)").
		From(tableName).
		Where(squirrel.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%w, %v : %v", repository.ErrBuildQuery, op, err)
		log.Println(err)
		return false, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var count int
	if err := r.db.DB().QueryRowContext(ctx, q, args...).Scan(&count); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil // User does not exist
		}
		err := fmt.Errorf("%w, %v : %v", repository.ErrExecQuery, op, err)
		log.Println(err)
		return false, err
	}

	return count > 0, nil // User exists if count is greater than 0
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
			companyIdColumn,
		).
		Values(
			user.Id,
			user.FirstName,
			user.LastName,
			user.Username,
			user.Latitude,
			user.Longitude,
			user.LanguageCode,
			user.UserType,
			user.CompanyId,
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

func (r repo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	const op = "repository.GetBy username user"

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
		Where(squirrel.Eq{usernameColumn: username})

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

func (r repo) GetByID(ctx context.Context, id int) (*model.User, error) {
	const op = "repository.GetByID user"

	builder := squirrel.Select(
		idColumn,
		firstNameColumn,
		lastNameColumn,
		usernameColumn,
		latitudeColumn,
		longitudeColumn,
		languageCodeColumn,
		userTypeColumn,
		companyIdColumn,
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
		&user.CompanyId,
		&user.CreatedAt,
	); err != nil {
		err := fmt.Errorf("%w, %v : %v", repository.ErrExecQuery, op, err)
		log.Println(err)
		return nil, err
	}

	return &user, nil
}
