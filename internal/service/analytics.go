package service

import (
	"fmt"
	"sort"
	"url-shortener/internal/storage"
)

type Analytics struct {
	storage *storage.Storage
}

func NewAnalytics(s *storage.Storage) *Analytics {
	return &Analytics{storage: s}
}

func (a *Analytics) GetTopNDomains(n int) []string {
	domainStats := a.storage.GetDomainStats()

	type domainData struct {
		domain string
		cnt    int
	}

	var domainsList []domainData
	for domain, cnt := range domainStats {
		domainsList = append(domainsList, domainData{domain: domain, cnt: cnt})
	}

	sort.Slice(domainsList, func(idx, jdx int) bool {
		return domainsList[idx].cnt > domainsList[jdx].cnt
	})

	var topDomains []string
	for idx := 0; idx < n && idx < len(domainsList); idx++ {
		topDomains = append(topDomains, fmt.Sprintf("%s: %d", domainsList[idx].domain, domainsList[idx].cnt))
	}

	fmt.Println(topDomains)

	return topDomains
}
