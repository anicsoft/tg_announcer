package announcement

import (
	apiModel "anik/internal/api/model"
	"anik/internal/client/db"
	"anik/internal/model"
	"anik/internal/repository"
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"log"
)

const (
	tableName           = "announcements"
	idColumn            = "announcement_id"
	companyIDColumn     = "company_id"
	titleColumn         = "title"
	startDateTimeColumn = "start_date_time"
	endDateTimeColumn   = "end_date_time"
	promoCodeColumn     = "promo_code"
	createdAtColumn     = "created_at"
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
			startDateTimeColumn,
			endDateTimeColumn,
			promoCodeColumn,
			createdAtColumn,
		).
		Values(
			announcement.CompanyID,
			announcement.Title,
			announcement.PromoCode,
			announcement.CreatedAt,
			announcement.StartDateTime,
			announcement.EndDateTime,
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

func (r *repo) Get(ctx context.Context, announcementID int) (*model.Announcement, error) {
	const op = "announcement.Get"
	builder := squirrel.Select("a.*", "oc.name AS category_name").
		From("Announcements a").
		Join("AnnouncementOffers ao ON a.announcement_id = ao.announcement_id").
		Join("OfferCategories oc ON ao.offer_category_id = oc.offer_category_id").
		Where(squirrel.Eq{"a.announcement_id": announcementID}).
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

	var announcement *model.Announcement
	var categories []string

	for rows.Next() {
		var annID int
		var ann model.Announcement
		var category string

		if err = rows.Scan(
			&annID,
			&ann.CompanyID,
			&ann.Title,
			&ann.PromoCode,
			&ann.CreatedAt,
			&ann.StartDateTime,
			&ann.EndDateTime,
			&category,
		); err != nil {
			return nil, err
		}

		if announcement == nil {
			ann.AnnouncementID = annID
			announcement = &ann
		}
		categories = append(categories, category)
	}

	if announcement != nil {
		announcement.Categories = categories
	}

	return announcement, nil
}

func (r *repo) GetAll(ctx context.Context, filter apiModel.Filter) ([]model.Announcement, error) {
	const op = "announcement.GetAll"
	builder := squirrel.Select(
		"a.*",
		"array_agg(oc.name ORDER BY oc.name) AS category_names",
		"c.name AS company_name",
		"c.address AS company_address",
		"c.description AS company_description",
		"c.latitude AS company_latitude",
		"c.longitude AS company_longitude",
	).
		From("Announcements a").
		Join("AnnouncementOffers ao ON a.announcement_id = ao.announcement_id").
		Join("OfferCategories oc ON ao.offer_category_id = oc.offer_category_id").
		Join("Companies c ON a.company_id = c.company_id").
		GroupBy("a.announcement_id, c.company_id").
		PlaceholderFormat(squirrel.Dollar)

	builder = applyFilters(builder, filter)
	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%s: %w", "Error building query", err)
		log.Println(err)
		return nil, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		err := fmt.Errorf("%s: %w", "Error executing query", err)
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var announcements []model.Announcement

	for rows.Next() {
		var ann model.Announcement
		var categories pq.StringArray
		var company model.Company
		var distance sql.NullFloat64 // Use sql.NullFloat64 to handle NULL distance values

		if err = rows.Scan(
			&ann.AnnouncementID,
			&ann.CompanyID,
			&ann.Title,
			&ann.PromoCode,
			&ann.CreatedAt,
			&ann.StartDateTime,
			&ann.EndDateTime,
			&categories,
			&company.Name,
			&company.Description,
			&company.Address,
			&company.Latitude,
			&company.Longitude,
			&distance, // Scan the distance
		); err != nil {
			return nil, err
		}

		ann.Categories = categories
		ann.Company = company
		if distance.Valid {
			ann.Distance = distance.Float64 // Assign the distance to the announcement
		}

		announcements = append(announcements, ann)
	}

	return announcements, nil
}

func applyFilters(builder squirrel.SelectBuilder, filter apiModel.Filter) squirrel.SelectBuilder {
	if len(filter.Categories) > 0 {
		builder = builder.Where(squirrel.Eq{"oc.name": filter.Categories})
	}
	if filter.StartDate != "" {
		builder = builder.Where(squirrel.GtOrEq{"a.start_date_time": filter.StartDate})
	}
	if filter.EndDate != "" {
		builder = builder.Where(squirrel.LtOrEq{"a.end_date_time": filter.EndDate})
	}
	if filter.PromoCode {
		builder = builder.Where(squirrel.NotEq{"a.promo_code": nil})
	}
	if filter.CreatedAt != "" {
		builder = builder.Where(squirrel.GtOrEq{"a.created_at": filter.CreatedAt})
	}
	if filter.Latitude != 0 && filter.Longitude != 0 {
		haversineFormula := fmt.Sprintf(`(
            6371 * acos(
                cos(radians(%f)) * cos(radians(c.latitude)) *
                cos(radians(c.longitude) - radians(%f)) +
                sin(radians(%f)) * sin(radians(c.latitude))
            )
        )`, filter.Latitude, filter.Longitude, filter.Latitude)

		builder = builder.Column(haversineFormula + " AS distance")
	}
	if filter.SortBy != "" {
		order := "ASC"
		if filter.SortOrder == "desc" {
			order = "DESC"
		}
		builder = builder.OrderBy(filter.SortBy + " " + order)
	}
	if filter.PageSize > 0 {
		builder = builder.Limit(uint64(filter.PageSize))
	}
	if filter.Offset > 0 {
		builder = builder.Offset(uint64(filter.Offset))
	}
	return builder
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
