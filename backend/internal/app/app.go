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
	"os"
	"time"

	_ "anik/docs" // http-swagger middleware
	"github.com/swaggo/http-swagger"
)

type App struct {
	httpServer *http.Server

	serviceProvider *serviceProvider

	router *chi.Mux
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{
		router: chi.NewRouter(),
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
		Addr:    fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT")),
		Handler: corsMiddleware.Handler(a.router),
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
	a.router.Use(middleware.RequestID)
	a.router.Use(middleware.RealIP)
	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Recoverer)
	a.router.Use(middleware.Timeout(60 * time.Second))

	//swagUrl := fmt.Sprintf("http://%s/swagger/doc.json", a.serviceProvider.HTTPConfig().Address())

	a.router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8888/swagger/doc.json"),
	))

	a.router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))

	})
	a.router.Post("/notify", a.serviceProvider.api.Notify(ctx))

	a.router.Group(func(r chi.Router) {
		// r.Use(api.AuthMiddleware)

		r.Patch("/users", a.serviceProvider.api.Update(ctx))
		r.Post("/announcements", a.serviceProvider.api.AddAnnouncement(ctx))
		r.Post("/companies", a.serviceProvider.api.AddCompany(ctx))
		r.Post("/categories/business", a.serviceProvider.api.AddBusinessCategory(ctx))
		r.Post("/categories/offer", a.serviceProvider.api.AddOfferCategory(ctx))

	})

	a.router.Post("/companies/{id}/logo", a.serviceProvider.api.UploadLogo(ctx))
	a.router.Get("/companies/{id}/logo", a.serviceProvider.api.FetchLogo(ctx))
	a.router.Post("/announcement/{id}/image", a.serviceProvider.api.UploadImage(ctx))
	a.router.Get("/announcement/{id}/image", a.serviceProvider.api.FetchImage(ctx))
	a.router.Get("/users/{id}", a.serviceProvider.api.GetUser(ctx))
	a.router.Post("/announcements/filter", a.serviceProvider.api.Announcements(ctx))
	a.router.Get("/announcements/{id}", a.serviceProvider.api.GetAnnouncement(ctx))
	a.router.Get("/companies/{id}", a.serviceProvider.api.GetCompanyByID(ctx))
	a.router.Get("/categories/business", a.serviceProvider.api.BusinessCategories(ctx))
	a.router.Get("/categories/offer", a.serviceProvider.api.OfferCategories(ctx))
}
