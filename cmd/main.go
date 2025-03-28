package main

import (
	"log"
	"url-shortener/internal/handlers"
	"url-shortener/internal/rest"
	"url-shortener/internal/service"
	"url-shortener/internal/storage"
)

func main() {
	storage := storage.NewStorage()

	handlers.Shortener = service.NewShortener(storage)
	handlers.Analytics = service.NewAnalytics(storage)

	e := rest.NewRouter()
	log.Fatal(e.Start("0.0.0.0:9090"))
}
