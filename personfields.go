package gopipedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

// GetPersonFields fetches all person fields for the company.
func (c *Client) GetPersonFields(ctx context.Context) (*[]models.PersonField, error) {
	url := c.APIBase + utils.PersonFieldsEndPoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var personFieldsResponse models.PersonFieldsResponse
	if err := c.SendWithAccessToken(req, &personFieldsResponse); err != nil {
		return nil, err
	}

	return &personFieldsResponse.Data, nil
}

// GetPersonField fetches details of a specific person field by ID.
func (c *Client) GetPersonField(ctx context.Context, fieldID int) (*models.PersonField, error) {
	url := c.APIBase + utils.PersonFieldsEndPoint + "/" + string(fieldID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var personFieldResponse models.PersonFieldResponse
	if err := c.SendWithAccessToken(req, &personFieldResponse); err != nil {
		return nil, err
	}

	return &personFieldResponse.Data, nil
}

// AddPersonField adds a new custom person field.
func (c *Client) AddPersonField(ctx context.Context, field models.PersonFieldReq) (*models.PersonField, error) {
	url := c.APIBase + utils.PersonFieldsEndPoint

	requestBodyBytes, err := json.Marshal(field)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	var personFieldResponse models.PersonFieldResponse
	if err := c.SendWithAccessToken(req, &personFieldResponse); err != nil {
		return nil, err
	}

	return &personFieldResponse.Data, nil
}

// UpdatePersonField updates a specific person field by ID.
func (c *Client) UpdatePersonField(ctx context.Context, fieldID int, field models.PersonFieldReq) (*models.PersonField, error) {
	url := c.APIBase + utils.PersonFieldsEndPoint + "/" + string(fieldID)

	requestBodyBytes, err := json.Marshal(field)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	var personFieldResponse models.PersonFieldResponse
	if err := c.SendWithAccessToken(req, &personFieldResponse); err != nil {
		return nil, err
	}

	return &personFieldResponse.Data, nil
}

// DeletePersonField deletes a specific person field by ID.
func (c *Client) DeletePersonField(ctx context.Context, fieldID int) error {
	url := c.APIBase + utils.PersonFieldsEndPoint + "/" + string(fieldID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	if err := c.SendWithAccessToken(req, &result); err != nil {
		return err
	}

	return nil
}
