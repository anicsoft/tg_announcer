package announcement

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
	tableName       = "announcements"
	idColumn        = "announcement_id"
	companyIDColumn = "company_id"
	titleColumn     = "title"
	startDateColumn = "start_date"
	endDateColumn   = "end_date"
	startTimeColumn = "start_time"
	endTimeColumn   = "end_time"
	promoCodeColumn = "promo_code"
	createdAtColumn = "created_at"
)

type repo struct {
	db db.Client
}

func New(db db.Client) repository.AnnouncementRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, announcement *model.Announcement) (int, error) {
	const op = "announcement.Create"

	builder := squirrel.Insert(tableName).
		PlaceholderFormat(repository.PlaceHolder).
		Columns(
			companyIDColumn,
			titleColumn,
			startDateColumn,
			endDateColumn,
			promoCodeColumn,
			startTimeColumn,
			endTimeColumn,
			createdAtColumn,
		).
		Values(
			announcement.CompanyID,
			announcement.Title,
			announcement.StartDate,
			announcement.EndDate,
			announcement.PromoCode,
			announcement.StartTime,
			announcement.EndTime,
			announcement.CreatedAt,
		).
		Suffix("RETURNING " + idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%s: %w", repository.ErrBuildQuery, err)
		log.Println(err)
		return 0, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var announcementId int
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&announcementId); err != nil {
		err := fmt.Errorf("%s: %w", op, err)
		log.Println(err)
		return 0, err
	}

	return announcementId, nil
}

func (r *repo) Get(ctx context.Context, id string) (*model.Announcement, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repo) GetAll(ctx context.Context) ([]model.Announcement, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repo) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (r *repo) Update(ctx context.Context, company *model.Announcement) error {
	//TODO implement me
	panic("implement me")
}

func (r *repo) GetCategoryId(ctx context.Context, categoryName string) (int, error) {
	const op = "announcement.GetCategoryId"

	builder := squirrel.Select("offer_category_id").
		PlaceholderFormat(repository.PlaceHolder).
		From("OfferCategories").
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

	var id int
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		err := fmt.Errorf("%s: %v", "no such category in db", err)
		log.Println(err)
		return 0, err
	}

	return id, nil
}

func (r *repo) AddCategory(ctx context.Context, category string, announcementId int) error {
	const op = "announcement.AddCategory"
	categoryId, err := r.GetCategoryId(ctx, category)
	if err != nil {
		return err
	}

	builder := squirrel.Insert("AnnouncementOffers").
		PlaceholderFormat(repository.PlaceHolder).
		Columns(idColumn, "offer_category_id").
		Values(announcementId, categoryId)

	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%s: %v", repository.ErrBuildQuery, err)
		log.Println(err)
		return err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	if _, err := r.db.DB().ExecContext(ctx, q, args...); err != nil {
		err := fmt.Errorf("%s: %v", "no such category in db", err)
		log.Println(err)
		return err
	}

	return nil
}
