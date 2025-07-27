package domain

// Product is what we get back from the Scraper service.
type Product struct {
	Title      string  `json:"title"`
	Price      float64 `json:"price"`
	Source     string  `json:"source"`
	ExternalID string  `json:"external_id"`
}

// ProductClient knows how to fetch a Product by its ExternalID.
type ProductClient interface {
	GetByExternalID(externalID string) (*Product, error)
}
