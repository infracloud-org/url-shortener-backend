package rest

import (
	"url-shortener/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	api := e.Group("/api/v1")
	api.POST("/shorten", handlers.CreateURLShorten)
	api.GET("/:shortCode", handlers.RedirectURL)
	api.GET("/metrics", handlers.GetTop3ShortenedDomains)

	return e
}
