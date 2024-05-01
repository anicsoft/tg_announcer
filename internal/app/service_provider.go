package app

import (
	"anik/internal/api"
	"anik/internal/client/db"
	"anik/internal/client/db/pg"
	"anik/internal/client/db/transaction"
	"anik/internal/config"
	"anik/internal/repository"
	"anik/internal/repository/announcement"
	"anik/internal/repository/categories"
	"anik/internal/repository/companies"
	"anik/internal/service"
	announcementService "anik/internal/service/announcements"
	categoriesService "anik/internal/service/categories"
	companiesService "anik/internal/service/companies"
	"context"
	"log"
)

type serviceProvider struct {
	httpConfig config.HTTPConfig
	pgConfig   config.PGConfig

	dbClient  db.Client
	txManager db.TxManager

	companiesRepo    repository.CompaniesRepository
	announcementRepo repository.AnnouncementRepository
	categoryRepo     repository.CategoriesRepository
	companiesServ    service.CompaniesService
	announcementServ service.AnnouncementService
	categoryServ     service.CategoriesService

	api api.Api
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

func (p *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if p.txManager == nil {
		p.txManager = transaction.NewTransactionManager(p.DBClient(ctx).DB())
	}

	return p.txManager
}

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
		repo := companies.New(p.DBClient(ctx))
		p.companiesRepo = repo
	}

	return p.companiesRepo
}

func (p *serviceProvider) CategoriesRepository(ctx context.Context) repository.CategoriesRepository {
	if p.categoryRepo == nil {
		repo := categories.New(p.DBClient(ctx))
		p.categoryRepo = repo
	}

	return p.categoryRepo
}

func (p *serviceProvider) AnnouncementRepository(ctx context.Context) repository.AnnouncementRepository {
	if p.announcementRepo == nil {
		repo := announcement.New(p.DBClient(ctx))
		p.announcementRepo = repo
	}

	return p.announcementRepo
}

func (p *serviceProvider) CompaniesService(ctx context.Context) service.CompaniesService {
	if p.companiesServ == nil {
		serv := companiesService.New(
			p.CompaniesRepository(ctx),
			p.TxManager(ctx))
		p.companiesServ = serv
	}

	return p.companiesServ
}

func (p *serviceProvider) AnnouncementService(ctx context.Context) service.AnnouncementService {
	if p.announcementServ == nil {
		serv := announcementService.New(
			p.AnnouncementRepository(ctx),
			p.TxManager(ctx))
		p.announcementServ = serv
	}

	return p.announcementServ
}

func (p *serviceProvider) CategoriesService(ctx context.Context) service.CategoriesService {
	if p.categoryServ == nil {
		serv := categoriesService.New(p.CategoriesRepository(ctx))
		p.categoryServ = serv
	}

	return p.categoryServ
}

func (p *serviceProvider) Api(ctx context.Context) api.Api {
	if p.api == nil {
		a := api.New(
			p.CompaniesService(ctx),
			p.AnnouncementService(ctx),
			p.CategoriesService(ctx),
		)
		p.api = a
	}

	return p.api
}
