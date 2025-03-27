package rest

import (
	"url-shortener/internal/handlers"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.Group("/api/v1")
	e.POST("/shorten", handlers.CreateURLShorten)
	e.GET("/:shortenCode", handlers.RedirectURL)
	e.GET("/metrics", handlers.GetTop3ShortenedDomains)

	return e
}
