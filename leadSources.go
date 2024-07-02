package gopipedrive

import (
	"context"
	"net/http"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

// GetLeadSources retrieves all lead sources
func (c *Client) GetLeadSources(ctx context.Context) ([]models.LeadSource, error) {
	url := c.APIBase + utils.LeadSourcesEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var leadSourcesResponse models.LeadSourcesResponse
	err = c.SendWithAccessToken(req, &leadSourcesResponse)
	if err != nil {
		return nil, err
	}

	return leadSourcesResponse.Data, nil
}
