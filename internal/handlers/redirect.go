package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func RedirectURL(ctx echo.Context) error {
	shortCode := ctx.Param("shortCode")

	originalURL, err := Shortener.ExpandURL(shortCode)
	if err != nil {
		return ctx.JSONBlob(http.StatusBadRequest, []byte(fmt.Sprintf(`{ "error": %s }`, err.Error())))
	}

	originalURL = strings.Trim(originalURL, `"`)
	return ctx.JSON(http.StatusOK, originalURL)
}
