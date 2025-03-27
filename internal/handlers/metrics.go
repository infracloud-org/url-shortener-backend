package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/internal/service"

	"github.com/labstack/echo/v4"
)

var Analytics *service.Analytics

func GetTop3ShortenedDomains(ctx echo.Context) error {
	topDomains := Analytics.GetTopNDomains(3)

	bodyBytes, err := json.Marshal(topDomains)
	if err != nil {
		return ctx.JSONBlob(http.StatusInternalServerError, []byte(fmt.Sprintf(`{ "error": %s }`, err.Error())))
	}

	return ctx.JSONBlob(http.StatusOK, bodyBytes)
}
