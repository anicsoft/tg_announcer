package announcement

import (
	"anik/internal/client/db"
	"anik/internal/model"
	"anik/internal/repository"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"log"
	"strconv"
	"strings"
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
	const op = "announcement.GetAll"
	builder := squirrel.Select("a.*", "oc.name AS category_name").
		From("Announcements a").
		Join("AnnouncementOffers ao ON a.announcement_id = ao.announcement_id").
		Join("OfferCategories oc ON ao.offer_category_id = oc.offer_category_id").
		PlaceholderFormat(repository.PlaceHolder)

	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%s: %w", repository.ErrBuildQuery, err)
		log.Println(err)
		return nil, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		err := fmt.Errorf("%s: %w", repository.ErrExecQuery, err)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	announcements := make(map[int]model.Announcement)

	for rows.Next() {
		var annID int
		var ann model.Announcement
		var category string

		if err = rows.Scan(
			&annID,
			&ann.CompanyID,
			&ann.Title,
			&ann.StartDate,
			&ann.EndDate,
			&ann.StartTime,
			&ann.EndTime,
			&ann.PromoCode,
			&ann.CreatedAt,
			&category,
		); err != nil {
			return nil, err
		}

		if _, ok := announcements[annID]; !ok {
			ann.Categories = []string{}
			announcements[annID] = ann
		}

		existingAnn := announcements[annID]
		existingAnn.AnnouncementID = annID
		existingAnn.Categories = append(existingAnn.Categories, category)

		announcements[annID] = existingAnn
	}

	var announcementList []model.Announcement
	for _, ann := range announcements {
		announcementList = append(announcementList, ann)
	}

	return announcementList, nil
}

func (r *repo) GetByCategory(ctx context.Context, categories []string) ([]model.Announcement, error) {
	const op = "announcement.GetByCategory"

	placeholders := "("
	for i := range categories {
		placeholders += "$" + strconv.Itoa(i+1) + ","
	}
	placeholders = strings.TrimSuffix(placeholders, ",") + ")"

	query := fmt.Sprintf(`SELECT a.*, oc.name AS category_name
				FROM Announcements a
				JOIN AnnouncementOffers ao ON a.announcement_id = ao.announcement_id
				JOIN OfferCategories oc ON ao.offer_category_id = oc.offer_category_id
				WHERE a.announcement_id IN (
					SELECT a.announcement_id
					FROM Announcements a
					JOIN AnnouncementOffers ao ON a.announcement_id = ao.announcement_id
					WHERE ao.offer_category_id IN (
						SELECT offer_category_id 
						FROM OfferCategories 
						WHERE name IN %s
					)
				)
				ORDER BY a.start_date DESC;`, placeholders)

	args := make([]interface{}, len(categories))
	for i, category := range categories {
		args[i] = category
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		err := fmt.Errorf("%s: %w", repository.ErrExecQuery, err)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	announcements := make(map[int]model.Announcement)

	for rows.Next() {
		var annID int
		var ann model.Announcement
		var category string

		if err = rows.Scan(
			&annID,
			&ann.CompanyID,
			&ann.Title,
			&ann.StartDate,
			&ann.EndDate,
			&ann.StartTime,
			&ann.EndTime,
			&ann.PromoCode,
			&ann.CreatedAt,
			&category,
		); err != nil {
			return nil, err
		}

		if _, ok := announcements[annID]; !ok {
			ann.Categories = []string{}
			announcements[annID] = ann
		}

		existingAnn := announcements[annID]
		existingAnn.AnnouncementID = annID
		existingAnn.Categories = append(existingAnn.Categories, category)

		announcements[annID] = existingAnn
	}

	var announcementList []model.Announcement
	for _, ann := range announcements {
		announcementList = append(announcementList, ann)
	}

	return announcementList, nil
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
