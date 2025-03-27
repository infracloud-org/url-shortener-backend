package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"url-shortener/internal/service"
	"url-shortener/pkg/common"

	"github.com/labstack/echo/v4"
)

var Shortener *service.Shortener

func CreateURLShorten(ctx echo.Context) error {
	reqBody := common.ShortenRequest{}

	bodyBytes, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return ctx.JSONBlob(http.StatusBadRequest, []byte(fmt.Sprintf(`{ "error": %s }`, err.Error())))
	}

	err = json.Unmarshal(bodyBytes, &reqBody)
	if err != nil {
		return ctx.JSONBlob(http.StatusBadRequest, []byte(fmt.Sprintf(`{ "error": %s }`, err.Error())))
	}

	shortURL, err := Shortener.ShortenURL(reqBody.OriginalURL)
	if err != nil {
		return ctx.JSONBlob(http.StatusInternalServerError, []byte(fmt.Sprintf(`{ "error": %s }`, err.Error())))
	}

	respBody := common.ShortenResponse{ShortURL: shortURL}
	respBodyBytes, _ := json.Marshal(respBody)
	return ctx.JSONBlob(http.StatusOK, respBodyBytes)
}
