package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "tg_announcer/docs" // http-swagger middleware
	"tg_announcer/internal/config"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	httpServer *http.Server

	serviceProvider *serviceProvider

	router *gin.Engine
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{
		router: gin.Default(),
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
	a.router.Use(requestid.New())

	a.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := a.router.Group("/backend")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("time: %v port: %v", time.Now(), os.Getenv("BACKEND_PORT")),
			})
		})

		api.POST("/notify", a.serviceProvider.api.Notify)

		companies := api.Group("/companies")
		{
			companies.POST("", a.serviceProvider.api.AddCompany)
			companies.POST("/:id/logo", a.serviceProvider.api.UploadLogo)
			companies.GET("/:id", a.serviceProvider.api.GetCompanyByID)
			companies.GET("", a.serviceProvider.api.ListCompanies)
			companies.PATCH("/:id", a.serviceProvider.api.UpdateCompany)
			companies.DELETE("/:id", a.serviceProvider.api.DeleteCompany)
		}

		announcements := api.Group("/announcements")
		{
			announcements.POST("/", a.serviceProvider.api.AddAnnouncement)
			announcements.POST("/:id/image", a.serviceProvider.api.UploadImage)
			announcements.POST("/filter", a.serviceProvider.api.Announcements)
			announcements.GET("/:id", a.serviceProvider.api.GetAnnouncement)
		}

		categories := api.Group("/categories")
		{
			categories.POST("/business", a.serviceProvider.api.AddBusinessCategory)
			categories.POST("/offer", a.serviceProvider.api.AddOfferCategory)
			categories.GET("/business", a.serviceProvider.api.BusinessCategories)
			categories.GET("/offer", a.serviceProvider.api.OfferCategories)
		}

		users := api.Group("/users")
		{
			users.PATCH("/", a.serviceProvider.api.Update)
			users.GET("/:id", a.serviceProvider.api.GetUser)
		}
	}
}
