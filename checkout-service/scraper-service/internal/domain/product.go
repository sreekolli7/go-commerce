package domain

// I'm defining what a scraped product looks like in my system
// This is the data structure for products I scrape from other websites
type ScrapedProduct struct {
	ID         int64   `json:"id" db:"id"`                   // The unique ID from my database
	Title      string  `json:"title" db:"title"`             // The product title from the website
	Price      float64 `json:"price" db:"price"`             // The price I scraped
	Source     string  `json:"source" db:"source"`           // Which website I scraped it from (amazon, ebay, etc.)
	ExternalID string  `json:"external_id" db:"external_id"` // The ID from the original website
}

type ScrapedProductRepository interface {
	Save(product *ScrapedProduct) error
	GetByExternalID(externalID string) (*ScrapedProduct, error)
	GetAll() ([]ScrapedProduct, error)
}
