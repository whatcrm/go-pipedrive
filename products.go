package gopipedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

// GetAllProducts retrieves all products with optional query parameters.
func (c *Client) GetAllProducts(ctx context.Context, queryParams map[string]string) ([]models.Product, error) {
	url := c.APIBase + utils.ProductsEndpoint
	if len(queryParams) > 0 {
		url += buildQueryParamsString(queryParams)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var productsResponse models.ProductsResponse
	err = c.SendWithAccessToken(req, &productsResponse)
	if err != nil {
		return nil, err
	}

	return productsResponse.Data, nil
}

// GetProduct retrieves a specific product by ID.
func (c *Client) GetProduct(ctx context.Context, productID string) (models.Product, error) {
	url := c.APIBase + fmt.Sprintf(utils.OneProductEndpoint, productID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Product{}, err
	}

	var productResponse models.ProductResponse
	err = c.SendWithAccessToken(req, &productResponse)
	if err != nil {
		return models.Product{}, err
	}

	return productResponse.Data, nil
}

// AddProduct creates a new product.
func (c *Client) AddProduct(ctx context.Context, productReq models.ProductRequest) (models.Product, error) {
	url := c.APIBase + utils.ProductsEndpoint

	requestBodyBytes, err := json.Marshal(productReq)
	if err != nil {
		return models.Product{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.Product{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var productResponse models.ProductResponse
	err = c.SendWithAccessToken(req, &productResponse)
	if err != nil {
		return models.Product{}, err
	}

	return productResponse.Data, nil
}

// UpdateProduct updates an existing product by ID.
func (c *Client) UpdateProduct(ctx context.Context, productID string, productUpdateReq models.ProductUpdateRequest) (models.Product, error) {
	url := c.APIBase + fmt.Sprintf(utils.OneProductEndpoint, productID)

	requestBodyBytes, err := json.Marshal(productUpdateReq)
	if err != nil {
		return models.Product{}, err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.Product{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var productResponse models.ProductResponse
	err = c.SendWithAccessToken(req, &productResponse)
	if err != nil {
		return models.Product{}, err
	}

	return productResponse.Data, nil
}

// DeleteProduct deletes a specific product by ID.
func (c *Client) DeleteProduct(ctx context.Context, productID string) error {
	url := c.APIBase + fmt.Sprintf(utils.OneProductEndpoint, productID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

// GetProductDeals retrieves deals associated with a specific product.
func (c *Client) GetProductDeals(ctx context.Context, productID string, queryParams map[string]string) ([]models.Deal, error) {
	url := c.APIBase + fmt.Sprintf(utils.ProductDealsEndpoint, productID)
	if len(queryParams) > 0 {
		url += buildQueryParamsString(queryParams)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var dealsResponse models.DealsResponse
	err = c.SendWithAccessToken(req, &dealsResponse)
	if err != nil {
		return nil, err
	}

	return dealsResponse.Data, nil
}

// AddProductVariation adds a variation to a product.
func (c *Client) AddProductVariation(ctx context.Context, productID string, variationReq models.ProductVariationRequest) (models.ProductVariation, error) {
	url := c.APIBase + fmt.Sprintf(utils.ProductVariationsEndpoint, productID)

	requestBodyBytes, err := json.Marshal(variationReq)
	if err != nil {
		return models.ProductVariation{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.ProductVariation{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var variationResponse models.ProductVariationResponse
	err = c.SendWithAccessToken(req, &variationResponse)
	if err != nil {
		return models.ProductVariation{}, err
	}

	return variationResponse.Data, nil
}

// DeleteProductVariation deletes a product variation by ID.
func (c *Client) DeleteProductVariation(ctx context.Context, productID, variationID string) error {
	url := c.APIBase + fmt.Sprintf(utils.ProductVariationEndpoint, productID, variationID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}
