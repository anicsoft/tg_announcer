package pg

import (
	"anik/internal/client/db"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type pgClient struct {
	masterDBC db.DB
}

func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}
	//migrate(dbc)
	return &pgClient{
		masterDBC: NewDB(dbc),
	}, nil
}

//func migrate() {
//	instance, err := pgx.WithInstance(db, pgx.Config{})
//	if err != nil {
//		return nil, fmt.Errorf("instance: %w", err)
//	}
//
//	fileSource, err := (&file.File{}).Open(migrations)
//	if err != nil {
//		return nil, fmt.Errorf("fileSource: %w", err)
//	}
//
//	m, err := migrate.NewWithInstance("file", fileSource, driver, instance)
//	if err != nil {
//		return nil, fmt.Errorf("migrations new: %w", err)
//	}
//
//	if err = m.Up(); err != nil {
//		if !errors.Is(err, migrate.ErrNoChange) {
//			return nil, fmt.Errorf("migrations run: %w", err)
//		}
//	}
//}

func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
