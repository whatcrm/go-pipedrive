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

// GetAllLeadLabels retrieves all lead labels.
func (c *Client) GetAllLeadLabels(ctx context.Context) ([]models.LeadLabel, error) {
	url := c.APIBase + utils.LeadLabelsEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var leadLabelsResponse models.LeadLabelsResponse
	err = c.SendWithAccessToken(req, &leadLabelsResponse)
	if err != nil {
		return nil, err
	}

	return leadLabelsResponse.Data, nil
}

// AddLeadLabel creates a new lead label.
func (c *Client) AddLeadLabel(ctx context.Context, leadLabel models.LeadLabelRequest) (models.LeadLabel, error) {
	url := c.APIBase + utils.LeadLabelsEndpoint

	requestBodyBytes, err := json.Marshal(leadLabel)
	if err != nil {
		return models.LeadLabel{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.LeadLabel{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	var leadLabelResponse models.LeadLabelResponse
	err = c.SendWithAccessToken(req, &leadLabelResponse)
	if err != nil {
		return models.LeadLabel{}, err
	}

	return leadLabelResponse.Data, nil
}

// UpdateLeadLabel updates the properties of a lead label.
func (c *Client) UpdateLeadLabel(ctx context.Context, leadLabelID string, leadLabel models.LeadLabelRequest) (models.LeadLabel, error) {
	url := fmt.Sprintf("%s%s/%s", c.APIBase, utils.LeadLabelsEndpoint, leadLabelID)

	requestBodyBytes, err := json.Marshal(leadLabel)
	if err != nil {
		return models.LeadLabel{}, err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.LeadLabel{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	var leadLabelResponse models.LeadLabelResponse
	err = c.SendWithAccessToken(req, &leadLabelResponse)
	if err != nil {
		return models.LeadLabel{}, err
	}

	return leadLabelResponse.Data, nil
}

// DeleteLeadLabel deletes a specific lead label.
func (c *Client) DeleteLeadLabel(ctx context.Context, leadLabelID string) error {
	url := fmt.Sprintf("%s%s/%s", c.APIBase, utils.LeadLabelsEndpoint, leadLabelID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}
