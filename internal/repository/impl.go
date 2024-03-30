package repository

import (
	"anik/internal/client/db"
	"anik/internal/model"
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
)

const (
	tableName         = "companies"
	idColumn          = "id"
	nameColumn        = "name"
	descriptionColumn = "description"
	addressColumn     = "address"
	latitudeColumn    = "latitude"
	longitudeColumn   = "longitude"
	whoColumn         = "who"
	createdAtColumn   = "create_at"
)

var (
	placeHolder   = squirrel.Dollar
	errBuildQuery = errors.New("error building query")
	errExecQuery  = errors.New("error executing")
)

func (r *repo) Create(ctx context.Context, company *model.Company) (string, error) {
	const op = "repository.Create"
	builder := squirrel.Insert(tableName).
		PlaceholderFormat(placeHolder).
		Columns(
			idColumn,
			nameColumn,
			descriptionColumn,
			addressColumn,
			latitudeColumn,
			longitudeColumn,
			whoColumn,
			createdAtColumn,
		).
		Values(
			company.Id,
			company.Name,
			company.Description,
			company.Address,
			company.Latitude,
			company.Longitude,
			company.Who,
			company.CreatedAt,
		).
		Suffix("RETURNING id")
	query, args, err := builder.ToSql()
	if err != nil {
		return "", fmt.Errorf("%w: %v", errBuildQuery, err)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var id string
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		return "", fmt.Errorf("%w, %v : %v", errExecQuery, op, err)
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id string) (*model.Company, error) {
	const op = "repository.Get"
	builder := squirrel.Select(
		idColumn,
		nameColumn,
		descriptionColumn,
		addressColumn,
		latitudeColumn,
		longitudeColumn,
		whoColumn,
		createdAtColumn,
	).
		PlaceholderFormat(placeHolder).
		From(tableName).
		Where(squirrel.Eq{idColumn: id})
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errBuildQuery, err)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var company model.Company
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(
		&company.Id,
		&company.Name,
		&company.Description,
		&company.CreatedAt,
	); err != nil {
		return nil, fmt.Errorf("%w, %v : %v", errExecQuery, op, err)
	}

	return &company, nil
}

func (r *repo) Delete(ctx context.Context, id string) error {
	const op = "repository.Delete"
	builder := squirrel.Delete(tableName).
		PlaceholderFormat(placeHolder).
		Where(squirrel.Eq{idColumn: id})
	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%w: %v", errBuildQuery, err)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	result, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("%w, %v : %v", errExecQuery, op, err)
	}

	rowCount := result.RowsAffected()
	if rowCount == 0 {
		return fmt.Errorf("no company with such ID %s", id)
	}

	return nil
}

func (r *repo) GetAll(ctx context.Context) ([]model.Company, error) {
	const op = "repository.GetAll"
	builder := squirrel.Select("*").From(tableName)
	query, _, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errBuildQuery, err)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("%w, %v : %v", errExecQuery, op, err)
	}
	defer rows.Close()

	var companies []model.Company
	for rows.Next() {
		var company model.Company
		if err = rows.Scan(
			&company.Id,
			&company.Name,
			&company.Description,
			&company.Address,
			&company.Latitude,
			&company.Longitude,
			&company.Who,
			&company.CreatedAt,
		); err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	return companies, nil
}

func (r *repo) Update(ctx context.Context, company *model.Company) error {
	const op = "repository.Update"
	builder := squirrel.Update(tableName).
		Set(nameColumn, company.Name).
		Set(descriptionColumn, company.Description).
		PlaceholderFormat(placeHolder).
		Where(squirrel.Eq{idColumn: company.Id})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%w: %v", errBuildQuery, err)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	result, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("%w, %v : %v", errExecQuery, op, err)
	}

	rowCount := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowCount == 0 {
		return fmt.Errorf("no company with such ID %s", company.Id)
	}

	return nil
}
