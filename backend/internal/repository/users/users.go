package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"tg_announcer/internal/client/db"
	"tg_announcer/internal/model"
	"tg_announcer/internal/repository"

	"github.com/Masterminds/squirrel"
)

const (
	usersTable         = "users"
	favoriteTable      = "favorites"
	userIdColumn       = "id"
	firstNameColumn    = "first_name"
	lastNameColumn     = "last_name"
	usernameColumn     = "username"
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

func (r *repo) Update(ctx context.Context, user *model.User) error {
	const op = "repository.Update user"

	builder := squirrel.Update(usersTable).
		Where(squirrel.Eq{userIdColumn: user.ID}).
		Set(firstNameColumn, user.FirstName).
		Set(lastNameColumn, user.LastName).
		Set(usernameColumn, user.Username).
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

func (r *repo) Exists(ctx context.Context, id int) (bool, error) {
	const op = "repository.Exists"

	builder := squirrel.Select("COUNT(*)").
		From(usersTable).
		Where(squirrel.Eq{userIdColumn: id})

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

func (r *repo) Create(ctx context.Context, user *model.User) (int, error) {
	const op = "repository.Create user"

	builder := squirrel.Insert(usersTable).
		PlaceholderFormat(repository.PlaceHolder).
		Columns(
			userIdColumn,
			firstNameColumn,
			lastNameColumn,
			usernameColumn,
			languageCodeColumn,
			userTypeColumn,
			companyIdColumn,
		).
		Values(
			user.ID,
			user.FirstName,
			user.LastName,
			user.Username,
			user.LanguageCode,
			user.UserType,
			user.CompanyId,
		).
		Suffix("RETURNING " + userIdColumn)

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

func (r *repo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	const op = "repository.GetBy username user"

	builder := squirrel.Select(
		userIdColumn,
		firstNameColumn,
		lastNameColumn,
		usernameColumn,
		languageCodeColumn,
		userTypeColumn,
		createdAtColumn,
	).
		PlaceholderFormat(repository.PlaceHolder).
		From(usersTable).
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
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
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

func (r *repo) GetByID(ctx context.Context, id int) (*model.User, error) {
	const op = "repository.Get user"

	builder := squirrel.Select(
		userIdColumn,
		firstNameColumn,
		lastNameColumn,
		usernameColumn,
		languageCodeColumn,
		userTypeColumn,
		companyIdColumn,
		createdAtColumn,
	).
		PlaceholderFormat(repository.PlaceHolder).
		From(usersTable).
		Where(squirrel.Eq{userIdColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		log.Println(errors.Join(err, repository.ErrBuildQuery))
		return nil, errors.Join(err, repository.ErrBuildQuery)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var user model.User
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.LanguageCode,
		&user.UserType,
		&user.CompanyId,
		&user.CreatedAt,
	); err != nil {
		log.Println(errors.Join(err, repository.ErrExecQuery))
		return nil, errors.Join(err, repository.ErrExecQuery)
	}

	return &user, nil
}

func (r *repo) AddFavoriteCompany(ctx context.Context, userId int, companyId string) error {
	const op = "repository.AddFavoriteCompany"

	builder := squirrel.Insert(favoriteTable).
		PlaceholderFormat(repository.PlaceHolder).
		Columns(
			"user_id",
			companyIdColumn,
		).
		Values(
			userId,
			companyId,
		).
		Suffix("RETURNING " + "id")

	query, args, err := builder.ToSql()
	if err != nil {
		log.Println(errors.Join(err, repository.ErrBuildQuery))
		return errors.Join(err, repository.ErrBuildQuery)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var id string
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		log.Println(errors.Join(err, repository.ErrExecQuery))
		return errors.Join(err, repository.ErrExecQuery)
	}

	return nil
}

func (r *repo) GetAll(ctx context.Context) ([]model.User, error) {
	const op = "repository.GetAll users"

	builder := squirrel.Select(
		userIdColumn,
		firstNameColumn,
		lastNameColumn,
		usernameColumn,
		languageCodeColumn,
		userTypeColumn,
		companyIdColumn,
		createdAtColumn,
	).
		PlaceholderFormat(repository.PlaceHolder).
		From(usersTable)

	query, args, err := builder.ToSql()
	if err != nil {
		log.Println(errors.Join(err, repository.ErrBuildQuery))
		return nil, errors.Join(err, repository.ErrBuildQuery)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var users []model.User
	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		log.Println(errors.Join(err, repository.ErrExecQuery))
		return nil, errors.Join(err, repository.ErrExecQuery)
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.LanguageCode,
			&user.UserType,
			&user.CompanyId,
			&user.CreatedAt,
		); err != nil {
			log.Println(errors.Join(err, repository.ErrExecQuery))
			return nil, errors.Join(err, repository.ErrExecQuery)
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *repo) GetFavoriteCompanies(ctx context.Context, userId int) ([]string, error) {
	const op = "repository.GetFavoriteCompanies"

	builder := squirrel.Select(companyIdColumn).
		From(favoriteTable).
		Where(squirrel.Eq{"user_id": userId}).
		PlaceholderFormat(repository.PlaceHolder)

	query, args, err := builder.ToSql()
	if err != nil {
		log.Println(errors.Join(err, repository.ErrBuildQuery))
		return nil, errors.Join(err, repository.ErrBuildQuery)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		log.Println(errors.Join(err, repository.ErrExecQuery))
		return nil, errors.Join(err, repository.ErrExecQuery)
	}

	defer rows.Close()

	var companies []string
	for rows.Next() {
		var companyId string
		if err = rows.Scan(&companyId); err != nil {
			log.Println(errors.Join(err, repository.ErrExecQuery))
			return nil, errors.Join(err, repository.ErrExecQuery)
		}

		companies = append(companies, companyId)
	}

	return companies, nil
}

func (r *repo) DeleteFavoriteCompany(ctx context.Context, userId int, companyId string) error {
	const op = "repository.DeleteFavoriteCompany"

	builder := squirrel.Delete(favoriteTable).
		Where(squirrel.Eq{"user_id": userId}).
		Where(squirrel.Eq{companyIdColumn: companyId}).
		PlaceholderFormat(repository.PlaceHolder)

	query, args, err := builder.ToSql()
	if err != nil {
		log.Println(errors.Join(err, repository.ErrBuildQuery))
		return errors.Join(err, repository.ErrBuildQuery)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	if _, err = r.db.DB().ExecContext(ctx, q, args...); err != nil {
		log.Println(errors.Join(err, repository.ErrExecQuery))
		return errors.Join(err, repository.ErrExecQuery)
	}

	return nil
}

func (r *repo) IsFavoriteCompany(ctx context.Context, userId int, companyId string) (bool, error) {
	const op = "repository.IsFavoriteCompany"

	builder := squirrel.Select("COUNT(*)").
		From(favoriteTable).
		Where(squirrel.Eq{"user_id": userId}).
		Where(squirrel.Eq{"company_id": companyId}).
		PlaceholderFormat(repository.PlaceHolder)

	query, args, err := builder.ToSql()
	if err != nil {
		log.Println(errors.Join(err, repository.ErrBuildQuery))
		return false, errors.Join(err, repository.ErrBuildQuery)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var count int
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&count); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		log.Println(errors.Join(err, repository.ErrExecQuery))
		return false, errors.Join(err, repository.ErrExecQuery)
	}

	return count > 0, nil
}
