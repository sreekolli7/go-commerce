package _interface

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sreekolli7/scraper-service/internal/usecase"
)

type ScraperHandler struct {
	Usecase *usecase.ScraperUsecase
}

func NewScraperHandler(uc *usecase.ScraperUsecase) *ScraperHandler {
	return &ScraperHandler{Usecase: uc}
}

// I'm handling POST requests to scrape products from a URL
// This is how other services can trigger scraping
func (h *ScraperHandler) Scrape(c *gin.Context) {
	// I'm reading the request body to get the URL to scrape
	var req struct {
		URL string `json:"url"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// I'm calling my business logic to scrape and save the product
	// This is where the actual scraping happens
	if err := h.Usecase.ScrapeAndSave(req.URL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Scraping failed",
			"details": err.Error(),
		})
		return
	}

	// I'm sending back a success message
	c.JSON(http.StatusOK, gin.H{"message": "Scraped product saved"})
}

// I'm handling GET requests to get a specific product by external ID
// This is how other services can retrieve scraped product data
func (h *ScraperHandler) GetProduct(c *gin.Context) {
	// I'm getting the external ID from the URL parameter
	id := c.Param("external_id")

	// I'm calling my business logic to get the product
	p, err := h.Usecase.GetByExternalID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found", "details": err.Error()})
		return
	}

	// I'm sending back the product data as JSON
	c.JSON(http.StatusOK, p)
}

// I'm handling GET requests to list all scraped products
// This is how other services can see all the products I've scraped
func (h *ScraperHandler) ListProducts(c *gin.Context) {
	// I'm calling my business logic to get all products
	all, err := h.Usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products", "details": err.Error()})
		return
	}

	// I'm sending back all the products as JSON
	c.JSON(http.StatusOK, all)
}
