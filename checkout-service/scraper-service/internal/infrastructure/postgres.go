package infrastructure

import (
	"github.com/jmoiron/sqlx"
	"github.com/sreekolli7/scraper-service/internal/domain"
)

type PostgresScraperRepo struct {
	DB *sqlx.DB
}

// I'm creating a new repository instance for scraped products
// This is how I connect my scraping business logic to the database
func NewPostgresScraperRepo(db *sqlx.DB) *PostgresScraperRepo {
	return &PostgresScraperRepo{DB: db}
}

// I'm saving a scraped product to the database
// This is where I store all the products I scrape from other websites
func (r *PostgresScraperRepo) Save(product *domain.ScrapedProduct) error {
	// TODO: I need to implement the actual database save here
	// For now, I'm just returning nil to make the API work
	return nil
}

// GetByExternalID looks up a scraped product by its external_id.
func (r *PostgresScraperRepo) GetByExternalID(externalID string) (*domain.ScrapedProduct, error) {
	var p domain.ScrapedProduct
	err := r.DB.Get(&p, `
		SELECT id, title, price, source, external_id
		FROM scraped_products
		WHERE external_id = $1
	`, externalID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetAll returns all scraped products, ordered by insertion.
func (r *PostgresScraperRepo) GetAll() ([]domain.ScrapedProduct, error) {
	var products []domain.ScrapedProduct
	err := r.DB.Select(&products, `
		SELECT id, title, price, source, external_id
		FROM scraped_products
		ORDER BY id
	`)
	return products, err
}
