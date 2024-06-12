package repository

/*
CREATE TABLE Image (
id serial PRIMARY KEY UNIQUE NOT NULL,
url text NOT NULL,
announcement_id integer,
company_id integer,
created_at timestamp NOT NULL,
updated_at timestamp NOT NULL,
FOREIGN KEY (announcement_id) REFERENCES Announcements(announcement_id) ON DELETE CASCADE,
FOREIGN KEY (company_id) REFERENCES Companies(company_id)  ON DELETE CASCADE
);
*/
/*const (
	tableName            = "image"
	idColumn             = "id"
	urlColumn            = "url"
	announcementIdColumn = "announcement_id"
	companyIdColumn      = "company_id"
	createdAtColumn      = "created_at"
	updatedAtColumn      = "updated_at"
)

type repo struct {
	db db.Client
}

func New(db db.Client) ImageRepository {
	return &repo{
		db: db,
	}
}

func (r repo) Get(ctx context.Context, parentId int) ([]string, error) {
	const op = "repository.Get image"

	builder := squirrel.Select(urlColumn).
		From(tableName).
		Where(squirrel.Eq{announcementIdColumn: parentId}).
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

func (r repo) Add(ctx context.Context, parentId int, path string) (int, error) {
	const op = "repository.Add image"

	builder := squirrel.Insert(tableName).
		PlaceholderFormat(PlaceHolder).
		Columns(
			urlColumn,
			announcementIdColumn,
		).
		Values(
			path,
			parentId,
		).Suffix("RETURNING " + idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		err := fmt.Errorf("%w: %v", ErrBuildQuery, err)
		log.Println(err)
		return 0, err
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}

	var id int
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		err := fmt.Errorf("%w, %v : %v", ErrExecQuery, op, err)
		log.Println(err)
		return 0, err
	}

	return id, nil
}
*/
