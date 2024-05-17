package companies

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
	tableName          = "Companies"
	idColumn           = "company_id"
	nameColumn         = "name"
	descriptionColumn  = "description"
	addressColumn      = "address"
	latitudeColumn     = "latitude"
	longitudeColumn    = "longitude"
	companyCategoryTbl = "CompanyCategories"
	categoryIdColumn   = "category_id"
)

type repo struct {
	db db.Client
}

func New(db db.Client) repository.CompaniesRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, company *model.Company) (int, error) {
	const op = "repository.Create"

	builder := squirrel.Insert(tableName).
		PlaceholderFormat(repository.PlaceHolder).
		Columns(
			nameColumn,
			descriptionColumn,
			addressColumn,
			latitudeColumn,
			longitudeColumn,
		).
		Values(
			company.Name,
			company.Description,
			company.Address,
			company.Latitude,
			company.Longitude,
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

func (r *repo) GetByID(ctx context.Context, id int) (*model.Company, error) {
	const op = "repository.GetByID"

	// TODO Join company categories
	builder := squirrel.Select(
		idColumn,
		nameColumn,
		descriptionColumn,
		addressColumn,
		latitudeColumn,
		longitudeColumn,
	).
		PlaceholderFormat(repository.PlaceHolder).
		From(tableName).
		Where(squirrel.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", repository.ErrBuildQuery, err)
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
		&company.Address,
		&company.Latitude,
		&company.Longitude,
	); err != nil {
		return nil, fmt.Errorf("%w, %v : %v", repository.ErrExecQuery, op, err)
	}

	return &company, nil
}

func (r *repo) Delete(ctx context.Context, id int) error {
	const op = "repository.Delete"
	builder := squirrel.Delete(tableName).
		PlaceholderFormat(repository.PlaceHolder).
		Where(squirrel.Eq{idColumn: id})
	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%w: %v", repository.ErrBuildQuery, err)
		log.Println(err.Error())
		return err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	result, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		err := fmt.Errorf("%w, %v : %v", repository.ErrExecQuery, op, err)
		log.Println(err)
		return err
	}

	rowCount := result.RowsAffected()
	if rowCount == 0 {
		err := fmt.Errorf("no company with such ID %s", id)
		log.Println(err)
		return err
	}

	return nil
}

func (r *repo) GetAll(ctx context.Context) ([]model.Company, error) {
	const op = "repository.GetAll"
	builder := squirrel.Select("*").From(tableName)
	query, _, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", repository.ErrBuildQuery, err)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("%w, %v : %v", repository.ErrExecQuery, op, err)
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
		PlaceholderFormat(repository.PlaceHolder).
		Where(squirrel.Eq{idColumn: company.Id})

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%w: %v", repository.ErrBuildQuery, err)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	result, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("%w, %v : %v", repository.ErrExecQuery, op, err)
	}

	rowCount := result.RowsAffected()
	if rowCount == 0 {
		return fmt.Errorf("no company with such ID %s", company.Id)
	}

	return nil
}

func (r *repo) GetCategoryId(ctx context.Context, categoryName string) (int64, error) {
	const op = "repository.GetCategoryId"
	builder := squirrel.Select("category_id").
		PlaceholderFormat(repository.PlaceHolder).
		From("BusinessCategories").
		Where(squirrel.Eq{"name": categoryName})

	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%s: %v", repository.ErrBuildQuery, err)
		log.Println(err)
		return 0, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var id int64
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		err := fmt.Errorf("%s: %v", "no such category in db", err)
		log.Println(err)
		return 0, err
	}

	return id, nil
}

func (r *repo) AddCategory(ctx context.Context, categories string, companyId int) error {
	const op = "repository.AddCategory"

	categoryId, err := r.GetCategoryId(ctx, categories)
	if err != nil {
		return err
	}

	builder := squirrel.Insert(companyCategoryTbl).
		PlaceholderFormat(repository.PlaceHolder).
		Columns(idColumn, categoryIdColumn).
		Values(companyId, categoryId)

	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%w: %v", repository.ErrBuildQuery, err)
		log.Println(err)
		return err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	if _, err = r.db.DB().ExecContext(ctx, q, args...); err != nil {
		err := fmt.Errorf("%w: %v", repository.ErrExecQuery, err)
		log.Println(err)
		return err
	}

	return nil
}

func (r *repo) DeleteCategory(ctx context.Context, id int) error {
	const op = "repository.DeleteCategory"

	builder := squirrel.Delete(tableName).
		PlaceholderFormat(repository.PlaceHolder).
		Where(squirrel.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%w: %v", repository.ErrBuildQuery, err)
		log.Println(err)
		return err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	result, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		err := fmt.Errorf("%w: %v", repository.ErrBuildQuery, err)
		log.Println(err)
		return err
	}

	rowCount := result.RowsAffected()
	if rowCount == 0 {
		return fmt.Errorf("no company with such ID %s", id)
	}

	return nil
}
