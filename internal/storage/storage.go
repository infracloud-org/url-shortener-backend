package storage

import "sync"

type Storage struct {
	urls        sync.Map
	clickCounts sync.Map
	domainCount sync.Map
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) SaveURL(originalURL, shortURL string) {
	s.urls.Store(originalURL, shortURL)
}

func (s *Storage) GetShortURL(originalURL string) (string, bool) {
	shortURL, ok := s.urls.Load(originalURL)
	if !ok {
		return "", false
	}
	return shortURL.(string), true
}

func (s *Storage) GetOriginalURL(shortURL string) (string, bool) {
	var originalURL string
	isExists := false

	s.urls.Range(func(key, value any) bool {
		if value.(string) == shortURL {
			originalURL = key.(string)
			isExists = true
			return true
		}
		return true
	})

	return originalURL, isExists
}

func (s *Storage) IncrementClick(shortURL string) {
	val, _ := s.clickCounts.LoadOrStore(shortURL, 0)
	s.clickCounts.Store(shortURL, val.(int)+1)
}

func (s *Storage) GetClickCount(shortURL string) int {
	val, ok := s.clickCounts.Load(shortURL)
	if !ok {
		return 0
	}
	return val.(int)
}

func (s *Storage) SaveDomain(domain string) {
	val, _ := s.domainCount.LoadOrStore(domain, 0)
	s.domainCount.Store(domain, val.(int)+1)
}

func (s *Storage) GetDomainStats() map[string]int {
	domainStats := make(map[string]int)
	s.domainCount.Range(func(key, value any) bool {
		domainStats[key.(string)] = value.(int)
		return true
	})
	return domainStats
}
