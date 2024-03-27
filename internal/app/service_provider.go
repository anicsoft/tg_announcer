package app

import (
	"anik/internal/api"
	"anik/internal/config"
	"anik/internal/repository"
	"anik/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

type serviceProvider struct {
	httpConfig   config.HTTPConfig
	sqliteConfig config.SQLiteConfig
	db           *sql.DB

	companiesRepo repository.CompaniesRepository
	companiesServ service.CompaniesService

	impl *api.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (p *serviceProvider) DB(ctx context.Context) *sql.DB {
	db, err := sql.Open(p.SQLiteConfig().Driver(), p.SQLiteConfig().Path())
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("failed to create insance: %v", err)
	}

	fileSource, err := (&file.File{}).Open(p.SQLiteConfig().Migrations())
	if err != nil {
		log.Fatalf("failed to get migraions source: %v", err)
	}

	m, err := migrate.NewWithInstance("file", fileSource, p.SQLiteConfig().Driver(), instance)
	if err != nil {
		log.Fatalf("failed to create migrations instance: %v", err)
	}

	if err = m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("failed to run migrations: %v", err)
		}
	}

	if err = db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	return db
}

func (p *serviceProvider) SQLiteConfig() config.SQLiteConfig {
	if p.sqliteConfig == nil {
		cfg := config.NewSQLiteConfig()
		p.sqliteConfig = cfg
	}

	return p.sqliteConfig
}

func (p *serviceProvider) HTTPConfig() config.HTTPConfig {
	if p.httpConfig == nil {
		cfg := config.NewHTTPConfig()
		p.httpConfig = cfg
	}

	return p.httpConfig
}

func (p *serviceProvider) CompaniesRepository(ctx context.Context) repository.CompaniesRepository {
	if p.companiesRepo == nil {
		repo := repository.NewRepository(p.DB(ctx))
		p.companiesRepo = repo
	}

	return p.companiesRepo
}

func (p *serviceProvider) CompaniesService(ctx context.Context) service.CompaniesService {
	if p.companiesServ == nil {
		serv := service.NewService(p.CompaniesRepository(ctx))
		p.companiesServ = serv
	}

	return p.companiesServ
}

func (p *serviceProvider) CompaniesImpl(ctx context.Context) *api.Implementation {
	if p.impl == nil {
		impl := api.NewImplementation(p.CompaniesService(ctx))
		p.impl = impl
	}

	return p.impl
}
