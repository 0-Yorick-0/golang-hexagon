package services

import (
	"golang-hexagon/internal/app/core/models"
	"golang-hexagon/internal/app/core/ports"
	"log"
)

type UrlService struct {
	Scraper    ports.ScraperService
	Repository ports.UrlRepository
}

func NewUrlService(scraper ports.ScraperService, repository ports.UrlRepository) *UrlService {
	return &UrlService{
		Scraper:    scraper,
		Repository: repository,
	}
}

func (dcs *UrlService) Analyze(urls []string) error {
	for _, url := range urls {
		metaTags, err := dcs.Scraper.GetMetaTags(url)
		if err != nil {
			log.Printf("Error scraping %s: %v", url, err)
			continue
		}

		analysisResult := models.UrlAnalysisResult{
			Link:     url,
			MetaTags: metaTags,
		}

		if err = dcs.Repository.Create(analysisResult); err != nil {
			log.Printf("Error saving result for %s: %v", url, err)
		}
	}

	return nil
}
