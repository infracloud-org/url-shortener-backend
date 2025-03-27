package rest

import (
	"url-shortener/internal/handlers"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	api := e.Group("/api/v1")
	api.POST("/shorten", handlers.CreateURLShorten)
	api.GET("/:shortCode", handlers.RedirectURL)
	api.GET("/metrics", handlers.GetTop3ShortenedDomains)

	return e
}
