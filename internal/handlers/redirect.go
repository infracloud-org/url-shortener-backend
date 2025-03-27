package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RedirectURL(ctx echo.Context) error {
	shortCode := ctx.Param("shortCode")

	originalURL, err := Shortener.ExpandURL(shortCode)
	if err != nil {
		return ctx.JSONBlob(http.StatusBadRequest, []byte(fmt.Sprintf(`{ "error": %s }`, err.Error())))
	}

	return ctx.Redirect(http.StatusFound, originalURL)
}
