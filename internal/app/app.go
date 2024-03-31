package app

import (
	"anik/internal/config"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type App struct {
	httpServer *http.Server

	serviceProvider *serviceProvider

	r *chi.Mux
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{
		r: chi.NewRouter(),
	}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runHttpServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHttpServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	a.configureRoutes(ctx)

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHttpServer(_ context.Context) error {

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Address(),
		Handler: corsMiddleware.Handler(a.r),
	}

	return nil
}

func (a *App) runHttpServer() error {
	log.Printf("Companies HTTP service is running on %s", a.serviceProvider.HTTPConfig().Address())
	return a.httpServer.ListenAndServe()
}

func (a *App) configureRoutes(ctx context.Context) {
	a.r.Use(middleware.Logger)
	a.r.Post("/api/v1/companies", a.serviceProvider.CompaniesImpl(ctx).Create(ctx))
	a.r.Put("/api/v1/companies", a.serviceProvider.CompaniesImpl(ctx).Update(ctx))
	a.r.Get("/api/v1/companies", a.serviceProvider.CompaniesImpl(ctx).GetAll(ctx))
	a.r.Get("/api/v1/companies/{id}", a.serviceProvider.CompaniesImpl(ctx).Get(ctx))
	a.r.Delete("/api/v1/companies/{id}", a.serviceProvider.CompaniesImpl(ctx).Delete(ctx))
	a.r.Post("/api/v1/nearby-locations", a.serviceProvider.CompaniesImpl(ctx).NearbyLocations(ctx))
}
