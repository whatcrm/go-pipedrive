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

// GetAllDeals retrieves all deals with optional query parameters.
func (c *Client) GetAllDeals(ctx context.Context, queryParams map[string]string) ([]models.Deal, error) {
	url := c.APIBase + utils.DealsEndpoint
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

// GetDeal retrieves a specific deal by ID.
func (c *Client) GetDeal(ctx context.Context, dealID string) (models.Deal, error) {
	url := c.APIBase + fmt.Sprintf(utils.OneDealEndpoint, dealID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Deal{}, err
	}

	var dealResponse models.DealResponse
	err = c.SendWithAccessToken(req, &dealResponse)
	if err != nil {
		return models.Deal{}, err
	}

	return dealResponse.Data, nil
}

// AddDeal creates a new deal.
func (c *Client) AddDeal(ctx context.Context, dealReq models.DealRequest) (models.Deal, error) {
	url := c.APIBase + utils.DealsEndpoint

	requestBodyBytes, err := json.Marshal(dealReq)
	if err != nil {
		return models.Deal{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.Deal{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var dealResponse models.DealResponse
	err = c.SendWithAccessToken(req, &dealResponse)
	if err != nil {
		return models.Deal{}, err
	}

	return dealResponse.Data, nil
}

// UpdateDeal updates an existing deal.
func (c *Client) UpdateDeal(ctx context.Context, dealID string, dealUpdateReq models.DealUpdateRequest) (models.Deal, error) {
	url := c.APIBase + fmt.Sprintf(utils.OneDealEndpoint, dealID)

	requestBodyBytes, err := json.Marshal(dealUpdateReq)
	if err != nil {
		return models.Deal{}, err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.Deal{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var dealResponse models.DealResponse
	err = c.SendWithAccessToken(req, &dealResponse)
	if err != nil {
		return models.Deal{}, err
	}

	return dealResponse.Data, nil
}

// DeleteDeal deletes a specific deal by ID.
func (c *Client) DeleteDeal(ctx context.Context, dealID string) error {
	url := c.APIBase + fmt.Sprintf(utils.OneDealEndpoint, dealID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

func (c *Client) ListDealParticipants(ctx context.Context, dealID int, queryParams map[string]string) ([]models.Participant, error) {
	dealEndpoint := fmt.Sprintf(utils.DealParticipantsEndpoint, dealID)
	url := c.APIBase + dealEndpoint
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var participantsResponse models.ParticipantsResponse
	err = c.SendWithAccessToken(req, &participantsResponse)
	if err != nil {
		return nil, err
	}

	return participantsResponse.Data, nil
}
