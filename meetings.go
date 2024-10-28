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

// LinkUserProvider links a user with a video calling integration.
func (c *Client) LinkUserProvider(ctx context.Context, linkRequest models.LinkUserProviderRequest) (*models.LinkUserProviderResponse, error) {
	url := c.APIBase + utils.LinkUserProviderEndpoint

	// Encode the request body
	requestBodyBytes, err := json.Marshal(linkRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var response models.LinkUserProviderResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteUserProviderLink removes the link between a user and the video calling integration.
func (c *Client) DeleteUserProviderLink(ctx context.Context, userProviderID string) (*models.DeleteUserProviderLinkResponse, error) {
	url := fmt.Sprintf(c.APIBase+utils.DeleteUserProviderLinkEndpoint, userProviderID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.DeleteUserProviderLinkResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
