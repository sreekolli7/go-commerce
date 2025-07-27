package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sreekolli7/scraper-service/internal/domain"
)

type HTTPProductClient struct {
	BaseURL string
}

func NewHTTPProductClient(baseURL string) *HTTPProductClient {
	return &HTTPProductClient{BaseURL: baseURL}
}

func (c *HTTPProductClient) GetByExternalID(externalID string) (*domain.Product, error) {
	// e.g. http://localhost:8082/products/ABC123
	url := fmt.Sprintf("%s/products/%s", c.BaseURL, externalID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("scraper returned %s", resp.Status)
	}

	var prod domain.Product
	if err := json.NewDecoder(resp.Body).Decode(&prod); err != nil {
		return nil, err
	}
	return &prod, nil
}
