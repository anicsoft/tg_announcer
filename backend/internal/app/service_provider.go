package app

import (
	"context"
	"log"
	"tg_announcer/internal/api"
	"tg_announcer/internal/client/db"
	"tg_announcer/internal/client/db/pg"
	"tg_announcer/internal/client/db/transaction"
	"tg_announcer/internal/config"
	"tg_announcer/internal/repository"
	"tg_announcer/internal/repository/announcement"
	"tg_announcer/internal/repository/categories"
	"tg_announcer/internal/repository/companies"
	"tg_announcer/internal/repository/users"
	"tg_announcer/internal/service"
	announcementService "tg_announcer/internal/service/announcements"
	categoriesService "tg_announcer/internal/service/categories"
	companiesService "tg_announcer/internal/service/companies"
	usersService "tg_announcer/internal/service/users"
)

type serviceProvider struct {
	httpConfig config.HTTPConfig
	pgConfig   config.PGConfig
	awsConfig  config.AWSConfig

	dbClient  db.Client
	txManager db.TxManager

	companiesRepo    repository.CompaniesRepository
	announcementRepo repository.AnnouncementRepository
	categoryRepo     repository.CategoriesRepository
	userRepo         repository.UsersRepository
	imageRepo        repository.ImageRepository
	companiesServ    service.CompaniesService
	announcementServ service.AnnouncementService
	categoryServ     service.CategoriesService
	usersServ        service.UsersService
	imageServ        service.ImageService

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

func (p *serviceProvider) AWSConfig() config.AWSConfig {
	if p.awsConfig == nil {
		cfg := config.NewAwsConfig()
		p.awsConfig = cfg
	}

	return p.awsConfig
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

func (p *serviceProvider) UsersRepository(ctx context.Context) repository.UsersRepository {
	if p.userRepo == nil {
		repo := users.New(p.DBClient(ctx))
		p.userRepo = repo
	}

	return p.userRepo
}

func (p *serviceProvider) ImageRepository(ctx context.Context) repository.ImageRepository {
	if p.imageRepo == nil {
		repo := repository.New(p.DBClient(ctx))
		p.imageRepo = repo
	}

	return p.imageRepo
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
			p.UsersRepository(ctx),
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

func (p *serviceProvider) UserService(ctx context.Context) service.UsersService {
	if p.usersServ == nil {
		serv := usersService.New(p.UsersRepository(ctx), p.CompaniesRepository(ctx), p.TxManager(ctx))
		p.usersServ = serv
	}

	return p.usersServ
}

func (p *serviceProvider) ImageService(ctx context.Context) service.ImageService {
	if p.imageServ == nil {
		serv := service.New(
			p.ImageRepository(ctx),
			p.AWSConfig(),
			p.TxManager(ctx),
		)
		p.imageServ = serv
	}

	return p.imageServ
}

func (p *serviceProvider) Api(ctx context.Context) api.Api {
	if p.api == nil {
		a := api.New(
			p.CompaniesService(ctx),
			p.AnnouncementService(ctx),
			p.CategoriesService(ctx),
			p.UserService(ctx),
			p.ImageService(ctx),
		)
		p.api = a
	}

	return p.api
}
