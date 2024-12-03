package gopipedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

// GetWebhooks fetches all webhooks for the company.
func (c *Client) GetWebhooks(ctx context.Context) (*[]models.Webhook, error) {
	url := c.APIBase + utils.WebhooksEndPoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var webhooksResponse models.WebhooksResponse
	if err := c.SendWithAccessToken(req, &webhooksResponse); err != nil {
		return nil, err
	}

	return &webhooksResponse.Data, nil
}

// AddWebhook adds a new webhook.
func (c *Client) AddWebhook(ctx context.Context, webhookReq models.CreateWebhookReq) (*models.Webhook, error) {
	url := c.APIBase + utils.WebhooksEndPoint

	requestBodyBytes, err := json.Marshal(webhookReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	var webhookResponse models.WebhookResponse
	if err := c.SendWithAccessToken(req, &webhookResponse); err != nil {
		return nil, err
	}

	return &webhookResponse.Data, nil
}

// DeleteWebhook deletes a specific webhook by ID.
func (c *Client) DeleteWebhook(ctx context.Context, webhookID int) error {
	url := c.APIBase + utils.WebhooksEndPoint + "/" + strconv.Itoa(webhookID)

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
