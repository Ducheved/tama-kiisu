package app

import (
	"github.com/Ducheved/tama-kiisu/internal/config"
	"github.com/Ducheved/tama-kiisu/internal/handlers"
	"github.com/Ducheved/tama-kiisu/internal/middleware"
	"github.com/Ducheved/tama-kiisu/internal/repositories"
	"github.com/Ducheved/tama-kiisu/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	config *config.Config
	router *gin.Engine
}

func New(cfg *config.Config) (*App, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	catRepo := repositories.NewCatRepository(db)
	catService := services.NewCatService(catRepo)
	h := handlers.NewHandlers(catService)

	router := gin.Default()
	router.Use(middleware.IPLimiter())
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/static", "./web/static")

	router.GET("/", h.IndexHandler)
	router.POST("/action", h.ActionHandler)

	return &App{
		config: cfg,
		router: router,
	}, nil
}

func (a *App) Run() error {
	return a.router.Run(":" + a.config.Port)
}
