package app

import (
	"anik/internal/api"
	"anik/internal/client/db"
	"anik/internal/client/db/pg"
	"anik/internal/config"
	"anik/internal/repository"
	"anik/internal/service"
	"context"
	"log"
)

type serviceProvider struct {
	httpConfig config.HTTPConfig
	/*sqliteConfig config.SQLiteConfig*/
	pgConfig config.PGConfig

	dbClient db.Client

	companiesRepo repository.CompaniesRepository
	companiesServ service.CompaniesService

	impl *api.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (p *serviceProvider) DBClient(ctx context.Context) db.Client {
	if p.dbClient == nil {
		cl, err := pg.New(ctx, p.PGConfig().DNS())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		p.dbClient = cl
	}

	return p.dbClient
}

//func (p *serviceProvider) SQLiteConfig() config.SQLiteConfig {
//	if p.sqliteConfig == nil {
//		cfg := config.NewSQLiteConfig()
//		p.sqliteConfig = cfg
//	}
//
//	return p.sqliteConfig
//}

func (p *serviceProvider) PGConfig() config.PGConfig {
	if p.pgConfig == nil {
		cfg := config.NewPGConfig()
		p.pgConfig = cfg
	}

	return p.pgConfig
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
		repo := repository.NewRepository(p.DBClient(ctx))
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
