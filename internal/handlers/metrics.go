package handlers

import (
	"net/http"
	"url-shortener/internal/service"

	"github.com/labstack/echo/v4"
)

var Analytics *service.Analytics

func GetTop3ShortenedDomains(ctx echo.Context) error {
	topDomains := Analytics.GetTopNDomains(3)
	return ctx.JSON(http.StatusOK, topDomains)
}
