package repository

import (
	"context"
	"fmt"
	"log"
	"tg_announcer/internal/client/db"

	"github.com/Masterminds/squirrel"
)

const (
	tableName            = "pictures"
	idColumn             = "picture_id"
	urlColumn            = "url"
	announcementIdColumn = "announcement_id"
	companyIdColumn      = "company_id"
)

type repo struct {
	db db.Client
}

func New(db db.Client) ImageRepository {
	return &repo{
		db: db,
	}
}

func (r repo) AddLogo(ctx context.Context, companyId string, path string) (string, error) {
	const op = "repository.AddLogo image"

	builder := squirrel.Insert(tableName).
		PlaceholderFormat(PlaceHolder).
		Columns(
			urlColumn,
			companyIdColumn,
		).
		Values(
			path,
			companyId,
		).Suffix("RETURNING " + idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%w: %v", ErrBuildQuery, err)
		log.Println(err)
		return "", err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var id string
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		err := fmt.Errorf("%w, %v : %v", ErrExecQuery, op, err)
		log.Println(err)
		return "", err
	}

	return id, nil
}

func (r repo) GetLogo(ctx context.Context, companyId string) (string, error) {
	const op = "repository.GetLogo image"

	builder := squirrel.Select(urlColumn).
		From(tableName).
		Where(squirrel.Eq{companyIdColumn: companyId}).
		PlaceholderFormat(PlaceHolder)

	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%w: %v", ErrBuildQuery, err)
		log.Println(err)
		return "", err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var path string
	if err := r.db.DB().QueryRowContext(ctx, q, args...).Scan(&path); err != nil {
		err := fmt.Errorf("%w: %v", ErrExecQuery, err)
		log.Println(err)
		return "", err
	}

	return path, nil
}

func (r repo) AddAnnouncementPictures(ctx context.Context, announcementId string, paths []string) (string, error) {
	const op = "repository.AddAnnouncementPictures"

	builder := squirrel.Insert(tableName).
		Columns(announcementIdColumn, urlColumn).
		PlaceholderFormat(PlaceHolder)

	for _, url := range paths {
		builder = builder.Values(announcementId, url)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%w: %v", ErrBuildQuery, err)
		log.Println(err)
		return "", err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	if _, err := r.db.DB().ExecContext(ctx, q, args...); err != nil {
		err := fmt.Errorf("%w: %v", ErrExecQuery, err)
		log.Println(err)
		return "", err
	}

	return "", nil
}

func (r repo) GetAnnouncementPictures(ctx context.Context, announcementId string) ([]string, error) {
	const op = "repository.Get image"

	builder := squirrel.Select(urlColumn).
		From(tableName).
		Where(squirrel.Eq{announcementIdColumn: announcementId}).
		PlaceholderFormat(PlaceHolder)

	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%w: %v", ErrBuildQuery, err)
		log.Println(err)
		return nil, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		err := fmt.Errorf("%w: %v", ErrExecQuery, err)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var paths []string
	for rows.Next() {
		var path string
		if err := rows.Scan(&path); err != nil {
			err := fmt.Errorf("%w", err)
			log.Println(err)
			return nil, err
		}

		paths = append(paths, path)
	}

	return paths, nil
}
