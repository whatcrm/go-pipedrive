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

// GetAllLeads retrieves all leads with optional query parameters
func (c *Client) GetAllLeads(ctx context.Context, queryParams map[string]string) ([]models.Lead, error) {
	url := c.APIBase + utils.LeadsEndpoint
	if len(queryParams) > 0 {
		url += buildQueryParamsString(queryParams)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var leadsResponse models.LeadsResponse
	err = c.SendWithAccessToken(req, &leadsResponse)
	if err != nil {
		return nil, err
	}

	return leadsResponse.Data, nil
}

// GetLead retrieves a specific lead by ID
func (c *Client) GetLead(ctx context.Context, leadID string) (models.Lead, error) {
	leadEndpoint := fmt.Sprintf(utils.OneLeadEndpoint, leadID)
	url := c.APIBase + leadEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Lead{}, err
	}

	var leadResponse models.LeadResponse
	err = c.SendWithAccessToken(req, &leadResponse)
	if err != nil {
		return models.Lead{}, err
	}

	return leadResponse.Data, nil
}

// ListPermittedUsers lists the users permitted to access a lead
func (c *Client) ListPermittedUsers(ctx context.Context, leadID string) ([]int, error) {
	leadEndpoint := fmt.Sprintf(utils.LeadPermittedUsersEndpoint, leadID)
	url := c.APIBase + leadEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var permittedUsersResponse models.PermittedUsersResponse
	err = c.SendWithAccessToken(req, &permittedUsersResponse)
	if err != nil {
		return nil, err
	}

	return permittedUsersResponse.Data, nil
}

// SearchLeads searches leads based on query parameters
func (c *Client) SearchLeads(ctx context.Context, queryParams map[string]string) ([]models.Lead, error) {
	url := utils.LeadSearchURL
	if len(queryParams) > 0 {
		url += buildQueryParamsString(queryParams)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var leadSearchResponse models.LeadSearchResponse
	err = c.SendWithAccessToken(req, &leadSearchResponse)
	if err != nil {
		return nil, err
	}

	return leadSearchResponse.Data, nil
}

// AddLead creates a new lead
func (c *Client) AddLead(ctx context.Context, leadReq models.LeadRequest) (models.Lead, error) {
	url := c.APIBase + utils.LeadsEndpoint

	requestBodyBytes, err := json.Marshal(leadReq)
	if err != nil {
		return models.Lead{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.Lead{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var leadResponse models.LeadResponse
	err = c.SendWithAccessToken(req, &leadResponse)
	if err != nil {
		return models.Lead{}, err
	}

	return leadResponse.Data, nil
}

// UpdateLead updates an existing lead
func (c *Client) UpdateLead(ctx context.Context, leadID string, leadUpdateReq models.LeadUpdateRequest) (models.Lead, error) {
	leadEndpoint := fmt.Sprintf(utils.OneLeadEndpoint, leadID)
	url := c.APIBase + leadEndpoint

	requestBodyBytes, err := json.Marshal(leadUpdateReq)
	if err != nil {
		return models.Lead{}, err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.Lead{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var leadResponse models.LeadResponse
	err = c.SendWithAccessToken(req, &leadResponse)
	if err != nil {
		return models.Lead{}, err
	}

	return leadResponse.Data, nil
}

// DeleteLead deletes a specific lead
func (c *Client) DeleteLead(ctx context.Context, leadID string) error {
	leadEndpoint := fmt.Sprintf(utils.OneLeadEndpoint, leadID)
	url := c.APIBase + leadEndpoint

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}
