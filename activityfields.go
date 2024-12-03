package gopipedrive

import (
	"context"
	"net/http"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

// GetActivityFields retrieves all activity fields.
func (c *Client) GetActivityFields(ctx context.Context) ([]models.ActivityField, error) {
	url := c.APIBase + utils.ActivityFieldsEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var fieldsResponse models.ActivityFieldsResponse
	err = c.SendWithAccessToken(req, &fieldsResponse)
	if err != nil {
		return nil, err
	}

	return fieldsResponse.Data, nil
}
