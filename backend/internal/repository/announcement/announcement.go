package announcement

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	apiModel "tg_announcer/internal/api/model"
	"tg_announcer/internal/client/db"
	"tg_announcer/internal/model"
	"tg_announcer/internal/repository"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

const (
	tableName           = "announcements"
	idColumn            = "announcement_id"
	companyIDColumn     = "company_id"
	titleColumn         = "title"
	contentColumn       = "content"
	pictureUrlColumn    = "picture_url"
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

func (r *repo) Create(ctx context.Context, announcement *model.Announcement) (string, error) {
	const op = "announcement.Create"

	builder := squirrel.Insert(tableName).
		PlaceholderFormat(repository.PlaceHolder).
		Columns(
			companyIDColumn,
			titleColumn,
			contentColumn,
			promoCodeColumn,
			pictureUrlColumn,
			startDateTimeColumn,
			endDateTimeColumn,
			createdAtColumn,
		).
		Values(
			announcement.CompanyID,
			announcement.Title,
			announcement.Content,
			announcement.PromoCode,
			announcement.PictureUrl,
			announcement.StartDateTime,
			announcement.EndDateTime,
			announcement.CreatedAt,
		).
		Suffix("RETURNING " + idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%s: %w", repository.ErrBuildQuery, err)
		log.Println(err)
		return "", err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var announcementId string
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&announcementId); err != nil {
		err := fmt.Errorf("%s: %w", op, err)
		log.Println(err)
		return "", err
	}

	return announcementId, nil
}

func (r *repo) Get(ctx context.Context, announcementID string) (*model.Announcement, error) {
	const op = "announcement.Get"
	builder := squirrel.Select("a.*",
		"p.url AS announcement_picture",
		"array_agg(oc.name ORDER BY oc.name) AS category_names",
		"c.name AS company_name",
		"c.description AS company_description",
		"c.address AS company_address",
		"c.latitude AS company_latitude",
		"c.longitude AS company_longitude",
		"pp.url AS company_logo",
	).
		From("announcements a").
		Where(squirrel.Eq{"a.announcement_id": announcementID}).
		LeftJoin("pictures p ON a.announcement_id = p.announcement_id").
		Join("announcementoffers ao ON a.announcement_id = ao.announcement_id").
		Join("offercategories oc ON ao.offer_category_id = oc.offer_category_id").
		Join("companies c ON a.company_id = c.company_id").
		Join("pictures pp ON a.company_id = pp.company_id").
		GroupBy("a.announcement_id, c.company_id, p.url, pp.url").
		PlaceholderFormat(squirrel.Dollar)

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

	var ann model.Announcement

	for rows.Next() {
		var categories pq.StringArray
		var company model.Company
		var distance sql.NullFloat64 // Use sql.NullFloat64 to handle NULL distance values
		var pictureUrl sql.NullString

		if err = rows.Scan(
			&ann.AnnouncementID,
			&ann.CompanyID,
			&ann.Title,
			&ann.Content,
			&ann.PromoCode,
			&ann.StartDateTime,
			&ann.EndDateTime,
			&ann.CreatedAt,
			&ann.Active,
			&pictureUrl,
			&categories,
			&company.Name,
			&company.Description,
			&company.Address,
			&company.Latitude,
			&company.Longitude,
			&company.LogoUrl,
		); err != nil {
			return nil, err
		}

		ann.Categories = categories
		ann.Company = company
		if pictureUrl.Valid {
			ann.PictureUrl = &pictureUrl.String // Assign the picture URL to the announcement
		}
		if distance.Valid {
			ann.Distance = distance.Float64 // Assign the distance to the announcement
		}
	}

	log.Println("announcement", ann)
	return &ann, nil
}

func (r *repo) GetAll(ctx context.Context, filter apiModel.Filter) ([]model.Announcement, error) {
	const op = "announcement.GetAll"
	builder := squirrel.Select(
		"a.*",
		"p.url AS announcement_picture",
		"array_agg(oc.name ORDER BY oc.name) AS category_names",
		"c.name AS company_name",
		"c.description AS company_description",
		"c.address AS company_address",
		"c.latitude AS company_latitude",
		"c.longitude AS company_longitude",
		"COALESCE(lp.url, cp.url) AS company_logo",
	).
		From("announcements a").
		LeftJoin("pictures p ON a.announcement_id = p.announcement_id").
		Join("announcementoffers ao ON a.announcement_id = ao.announcement_id").
		Join("offercategories oc ON ao.offer_category_id = oc.offer_category_id").
		Join("companies c ON a.company_id = c.company_id").
		LeftJoin("pictures lp ON c.company_id = lp.company_id").
		LeftJoin("pictures cp ON p.announcement_id = cp.announcement_id").
		GroupBy("a.announcement_id, c.company_id, p.url, lp.url, cp.url").
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
		var pictureUrl sql.NullString

		if err = rows.Scan(
			&ann.AnnouncementID,
			&ann.CompanyID,
			&ann.Title,
			&ann.Content,
			&ann.PromoCode,
			&ann.StartDateTime,
			&ann.EndDateTime,
			&ann.CreatedAt,
			&ann.Active,
			&pictureUrl,
			&categories,
			&company.Name,
			&company.Description,
			&company.Address,
			&company.Latitude,
			&company.Longitude,
			&company.LogoUrl,
			&distance,
		); err != nil {
			return nil, err
		}

		ann.Categories = categories
		ann.Company = company
		if pictureUrl.Valid {
			ann.PictureUrl = &pictureUrl.String // Assign the picture URL to the announcement
		}
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
	if filter.CompanyID != "" {
		builder = builder.Where(squirrel.Eq{"a.company_id": filter.CompanyID})
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
	} else {
		builder = builder.Column("0 AS distance")
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

func (r *repo) AddCategory(ctx context.Context, category string, announcementId string) error {
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
