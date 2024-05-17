package app

import (
	"anik/internal/api"
	"anik/internal/config"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"

	_ "anik/docs" // http-swagger middleware
	"github.com/swaggo/http-swagger"
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

	a.configureRoutes(ctx)
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
		a.initApi,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

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
	log.Printf("Backend is running on %s", a.serviceProvider.HTTPConfig().Address())
	return a.httpServer.ListenAndServe()
}

func (a *App) initApi(ctx context.Context) error {
	a.serviceProvider.Api(ctx)
	return nil
}

func (a *App) configureRoutes(ctx context.Context) {
	a.r.Use(middleware.RequestID)
	a.r.Use(middleware.RealIP)
	a.r.Use(middleware.Logger)
	a.r.Use(middleware.Recoverer)
	a.r.Use(middleware.Timeout(60 * time.Second))

	swagUrl := fmt.Sprintf("http://%s/swagger/doc.json", a.serviceProvider.HTTPConfig().Address())

	a.r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swagUrl),
	))

	a.r.Post("/notify", a.serviceProvider.api.Notify(ctx))

	a.r.Group(func(r chi.Router) {
		r.Use(api.AuthMiddleware)

		r.Patch("/user", a.serviceProvider.api.Update(ctx))
		r.Post("/announcement", a.serviceProvider.api.AddAnnouncement(ctx))
		r.Post("/companies", a.serviceProvider.api.AddCompany(ctx))
		r.Post("/categories/business", a.serviceProvider.api.AddBusinessCategory(ctx))
		r.Post("/categories/offer", a.serviceProvider.api.AddOfferCategory(ctx))

	})

	a.r.Get("/announcement", a.serviceProvider.api.Announcements(ctx))
	a.r.Get("/companies", a.serviceProvider.api.GetCompanyByID(ctx))
	a.r.Get("/categories/business", a.serviceProvider.api.BusinessCategories(ctx))
	a.r.Get("/categories/offer", a.serviceProvider.api.OfferCategories(ctx))
}
