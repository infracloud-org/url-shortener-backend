package service

import (
	"fmt"
	"net/url"
	"strings"
	"url-shortener/internal/storage"
	"url-shortener/pkg/utils"
)

type Shortener struct {
	storage *storage.Storage
}

func NewShortener(s *storage.Storage) *Shortener {
	return &Shortener{storage: s}
}

func (s *Shortener) ShortenURL(originalURL string) (string, error) {
	if !utils.IsValidURL(originalURL) {
		return "", fmt.Errorf("invalid URL provided")
	}

	if shortURL, ok := s.storage.GetShortURL(originalURL); ok {
		return shortURL, nil
	}

	shortKey := utils.GenerateRandomShortKey(6)
	s.storage.SaveURL(originalURL, shortKey)

	parsedURL, _ := url.Parse(originalURL)
	domain := strings.TrimPrefix(parsedURL.Hostname(), "www.")
	s.storage.SaveDomain(domain)

	return shortKey, nil
}

func (s *Shortener) ExpandURL(shortKey string) (string, error) {
	if !utils.IsValidShortKey(shortKey) {
		return "", fmt.Errorf("invaid short key")
	}

	originalURL, ok := s.storage.GetOriginalURL(shortKey)
	if !ok {
		return "", fmt.Errorf("shortURL does not exists")
	}

	s.storage.IncrementClick(shortKey)
	return originalURL, nil
}
