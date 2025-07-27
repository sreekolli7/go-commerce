package usecase

import (
	"strings"

	"github.com/sreekolli7/scraper-service/internal/domain"
)

// I'm creating a scraper usecase that handles all the business logic
// This is where I put all the rules for scraping products
type ScraperUsecase struct {
	Repo domain.ScrapedProductRepository
}

// NewScraperUsecase constructs a new ScraperUsecase.
func NewScraperUsecase(repo domain.ScrapedProductRepository) *ScraperUsecase {
	return &ScraperUsecase{Repo: repo}
}

// I'm scraping products from a URL and saving them
// This is my main business logic for the scraper service
func (u *ScraperUsecase) ScrapeAndSave(url string) error {
	// TODO: I need to implement actual web scraping here
	// For now, I'm just returning nil to make the API work
	return nil
}

// I'm getting a product by its external ID
// This is how I retrieve specific scraped products
func (u *ScraperUsecase) GetByExternalID(externalID string) (*domain.ScrapedProduct, error) {
	// TODO: I need to implement database lookup here
	// For now, I'm returning a mock product
	return &domain.ScrapedProduct{
		ID:         1,
		Title:      "Sample Product",
		Price:      9.99,
		Source:     "example.com",
		ExternalID: externalID,
	}, nil
}

// I'm getting all scraped products
// This is how I retrieve all the products I've scraped
func (u *ScraperUsecase) GetAll() ([]domain.ScrapedProduct, error) {
	// TODO: I need to implement database lookup here
	// For now, I'm returning mock data
	products := []domain.ScrapedProduct{
		{ID: 1, Title: "Sample Product", Price: 9.99, Source: "example.com", ExternalID: "abc123"},
	}
	return products, nil
}

// extractExternalID pulls out the product identifier from an Amazon URL.
func extractExternalID(url string) string {
	parts := strings.Split(url, "/dp/")
	if len(parts) > 1 {
		idParts := strings.Split(parts[1], "/")
		return idParts[0]
	}
	return ""
}
