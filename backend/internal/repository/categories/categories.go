package categories

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
	businessCategoryTbl      = "BusinessCategories"
	offerCategoryTbl         = "OfferCategories"
	businessCategoryIdColumn = "category_id"
	offerCategoryIdColumn    = "offer_category_id"
	categoryNameTableColumn  = "name"
)

type repo struct {
	db db.Client
}

func New(client db.Client) repository.CategoriesRepository {
	return &repo{
		db: client,
	}
}

func (r repo) AddOfferCategory(ctx context.Context, category string) (int, error) {
	const op = "repository.addOfferCategory"

	builder := squirrel.Insert(offerCategoryTbl).
		PlaceholderFormat(squirrel.Dollar).
		Columns(categoryNameTableColumn).
		Values(category).
		Suffix("RETURNING " + offerCategoryIdColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		err = fmt.Errorf("%s: %w", repository.ErrBuildQuery, err)
		log.Println(err)
		return 0, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var id int
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		err = fmt.Errorf("%w, %v : %v", repository.ErrExecQuery, op, err)
		log.Println(err)
		return 0, err
	}

	return id, nil
}

func (r repo) AddBusinessCategory(ctx context.Context, category string) (int, error) {
	const op = "repository.AddBusinessCategory"

	builder := squirrel.Insert(businessCategoryTbl).
		PlaceholderFormat(squirrel.Dollar).
		Columns(categoryNameTableColumn).
		Values(category).
		Suffix("RETURNING " + businessCategoryIdColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		err = fmt.Errorf("%s: %w", repository.ErrBuildQuery, err)
		log.Println(err)
		return 0, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var id int
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		err = fmt.Errorf("%w, %v : %v", repository.ErrExecQuery, op, err)
		log.Println(err)
		return 0, err
	}

	return id, nil
}

func (r repo) GetOfferCategories(ctx context.Context) ([]model.Category, error) {
	const op = "repository.GetOfferCategories"

	builder := squirrel.Select(offerCategoryIdColumn, categoryNameTableColumn).
		PlaceholderFormat(repository.PlaceHolder).
		From(offerCategoryTbl)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", repository.ErrBuildQuery, err)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("%w, %v : %v", repository.ErrExecQuery, op, err)
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		if err = rows.Scan(
			&category.ID,
			&category.Name,
		); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (r repo) GetBusinessCategories(ctx context.Context) ([]model.Category, error) {
	const op = "repository.GetBusinessCategories"

	builder := squirrel.Select(businessCategoryIdColumn, categoryNameTableColumn).
		PlaceholderFormat(repository.PlaceHolder).
		From(businessCategoryTbl)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", repository.ErrBuildQuery, err)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("%w, %v : %v", repository.ErrExecQuery, op, err)
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		if err = rows.Scan(
			&category.ID,
			&category.Name,
		); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}
