package gopipedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

// GetDealFields fetches all deal fields for the company.
func (c *Client) GetDealFields(ctx context.Context) (*[]models.DealField, error) {
	url := c.APIBase + utils.DealFieldsEndPoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var dealFieldsResponse models.DealFieldsResponse
	if err := c.SendWithAccessToken(req, &dealFieldsResponse); err != nil {
		return nil, err
	}

	return &dealFieldsResponse.Data, nil
}

// GetDealField fetches details of a specific deal field by ID.
func (c *Client) GetDealField(ctx context.Context, fieldID int) (*models.DealField, error) {
	url := fmt.Sprintf("%s%s/%d", c.APIBase, utils.DealFieldsEndPoint, fieldID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var dealFieldResponse models.DealFieldResponse
	if err := c.SendWithAccessToken(req, &dealFieldResponse); err != nil {
		return nil, err
	}

	return &dealFieldResponse.Data, nil
}

// AddDealField adds a new custom deal field.
func (c *Client) AddDealField(ctx context.Context, field models.DealFieldReq) (*models.DealField, error) {
	url := c.APIBase + utils.DealFieldsEndPoint

	requestBodyBytes, err := json.Marshal(field)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	var dealFieldResponse models.DealFieldResponse
	if err := c.SendWithAccessToken(req, &dealFieldResponse); err != nil {
		return nil, err
	}

	return &dealFieldResponse.Data, nil
}

// UpdateDealField updates a specific deal field by ID.
func (c *Client) UpdateDealField(ctx context.Context, fieldID int, field models.DealFieldReq) (*models.DealField, error) {
	url := fmt.Sprintf("%s%s/%d", c.APIBase, utils.DealFieldsEndPoint, fieldID)

	requestBodyBytes, err := json.Marshal(field)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	var dealFieldResponse models.DealFieldResponse
	if err := c.SendWithAccessToken(req, &dealFieldResponse); err != nil {
		return nil, err
	}

	return &dealFieldResponse.Data, nil
}

// DeleteDealField deletes a specific deal field by ID.
func (c *Client) DeleteDealField(ctx context.Context, fieldID int) error {
	url := fmt.Sprintf("%s%s/%d", c.APIBase, utils.DealFieldsEndPoint, fieldID)

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

// DeleteMultipleDealFields deletes multiple deal fields by their IDs.
func (c *Client) DeleteMultipleDealFields(ctx context.Context, fieldIDs []int) error {
	url := c.APIBase + utils.DealFieldsEndPoint

	ids := make([]string, len(fieldIDs))
	for i, id := range fieldIDs {
		ids[i] = fmt.Sprintf("%d", id)
	}

	req, err := http.NewRequest("DELETE", url+"?ids="+strings.Join(ids, ","), nil)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	if err := c.SendWithAccessToken(req, &result); err != nil {
		return err
	}

	return nil
}
