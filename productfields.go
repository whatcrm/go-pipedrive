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

// GetAllProductFields retrieves all product fields with optional pagination.
func (c *Client) GetAllProductFields(ctx context.Context, queryParams map[string]string) ([]models.ProductField, error) {
	url := c.APIBase + utils.ProductFieldsEndpoint
	if len(queryParams) > 0 {
		url += buildQueryParamsString(queryParams)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var productFieldsResponse models.ProductFieldsResponse
	err = c.SendWithAccessToken(req, &productFieldsResponse)
	if err != nil {
		return nil, err
	}

	return productFieldsResponse.Data, nil
}

// GetProductField retrieves a specific product field by ID.
func (c *Client) GetProductField(ctx context.Context, fieldID string) (models.ProductField, error) {
	url := c.APIBase + fmt.Sprintf(utils.OneProductFieldEndpoint, fieldID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.ProductField{}, err
	}

	var productFieldResponse models.ProductFieldResponse
	err = c.SendWithAccessToken(req, &productFieldResponse)
	if err != nil {
		return models.ProductField{}, err
	}

	return productFieldResponse.Data, nil
}

// AddProductField creates a new custom product field.
func (c *Client) AddProductField(ctx context.Context, fieldReq models.ProductFieldRequest) (models.ProductField, error) {
	url := c.APIBase + utils.ProductFieldsEndpoint

	requestBodyBytes, err := json.Marshal(fieldReq)
	if err != nil {
		return models.ProductField{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.ProductField{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var productFieldResponse models.ProductFieldResponse
	err = c.SendWithAccessToken(req, &productFieldResponse)
	if err != nil {
		return models.ProductField{}, err
	}

	return productFieldResponse.Data, nil
}

// UpdateProductField updates an existing product field by ID.
func (c *Client) UpdateProductField(ctx context.Context, fieldID string, fieldUpdateReq models.ProductFieldUpdateRequest) (models.ProductField, error) {
	url := c.APIBase + fmt.Sprintf(utils.OneProductFieldEndpoint, fieldID)

	requestBodyBytes, err := json.Marshal(fieldUpdateReq)
	if err != nil {
		return models.ProductField{}, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.ProductField{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var productFieldResponse models.ProductFieldResponse
	err = c.SendWithAccessToken(req, &productFieldResponse)
	if err != nil {
		return models.ProductField{}, err
	}

	return productFieldResponse.Data, nil
}

// DeleteProductField deletes a specific product field by ID.
func (c *Client) DeleteProductField(ctx context.Context, fieldID string) error {
	url := c.APIBase + fmt.Sprintf(utils.OneProductFieldEndpoint, fieldID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

// DeleteProductFields deletes multiple product fields in bulk by their IDs.
func (c *Client) DeleteProductFields(ctx context.Context, fieldIDs []string) error {
	url := c.APIBase + utils.ProductFieldsEndpoint + "?ids=" 

	if len(fieldIDs) > 0 {
		url += "?ids="
		for _, id := range fieldIDs {
			url += id + ","
		}
	}

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}
