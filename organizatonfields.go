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

// GetOrganizationFields fetches all organization fields.
func (c *Client) GetOrganizationFields(ctx context.Context, queryParams map[string]string) (*models.OrganizationFieldsResponse, error) {
	url := c.APIBase + utils.OrganizationFieldsEndpoint
	if len(queryParams) > 0 {
		url += buildQueryParamsString(queryParams)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.OrganizationFieldsResponse
	err = c.SendWithAccessToken(req, &response)
	return &response, err
}

// GetOrganizationField fetches a specific organization field by ID.
func (c *Client) GetOrganizationField(ctx context.Context, id int) (*models.OrganizationFieldResponse, error) {
	url := fmt.Sprintf("%s%s/%d", c.APIBase, utils.OrganizationFieldsEndpoint, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var response models.OrganizationFieldResponse
	err = c.SendWithAccessToken(req, &response)
	return &response, err
}

// AddOrganizationField adds a new custom organization field.
func (c *Client) AddOrganizationField(ctx context.Context, field models.OrganizationField) (*models.OrganizationFieldResponse, error) {
	url := c.APIBase + utils.OrganizationFieldsEndpoint
	body, err := json.Marshal(field)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	var response models.OrganizationFieldResponse
	err = c.SendWithAccessToken(req, &response)
	return &response, err
}

// UpdateOrganizationField updates an existing organization field by ID.
func (c *Client) UpdateOrganizationField(ctx context.Context, id int, updates map[string]interface{}) (*models.OrganizationFieldResponse, error) {
	url := fmt.Sprintf("%s%s/%d", c.APIBase, utils.OrganizationFieldsEndpoint, id)
	body, err := json.Marshal(updates)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	var response models.OrganizationFieldResponse
	err = c.SendWithAccessToken(req, &response)
	return &response, err
}

// DeleteOrganizationField deletes a specific organization field by ID.
func (c *Client) DeleteOrganizationField(ctx context.Context, id int) (*models.FollowerResponse, error) {
	url := fmt.Sprintf("%s%s/%d", c.APIBase, utils.OrganizationFieldsEndpoint, id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	var response models.FollowerResponse
	err = c.SendWithAccessToken(req, &response)
	return &response, err
}

// DeleteMultipleOrganizationFields deletes multiple organization fields in bulk.
func (c *Client) DeleteMultipleOrganizationFields(ctx context.Context, ids []int) (*models.FollowerResponse, error) {
	url := c.APIBase + utils.OrganizationFieldsEndpoint
	query := ""
	for i, id := range ids {
		if i == 0 {
			query += fmt.Sprintf("ids=%d", id)
		} else {
			query += fmt.Sprintf(",%d", id)
		}
	}
	req, err := http.NewRequest("DELETE", url+"?"+query, nil)
	if err != nil {
		return nil, err
	}
	var response models.FollowerResponse
	err = c.SendWithAccessToken(req, &response)
	return &response, err
}
