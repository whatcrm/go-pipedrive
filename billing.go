package gopipedrive

import (
	"context"
	"net/http"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

// GetBillingAddOns retrieves all add-ons for a single company.
func (c *Client) GetBillingAddOns(ctx context.Context) ([]models.AddOn, error) {
	url := c.APIBase + utils.BillingAddOnsEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var addOnsResponse models.BillingAddOnsResponse
	err = c.SendWithAccessToken(req, &addOnsResponse)
	if err != nil {
		return nil, err
	}

	return addOnsResponse.Data, nil
}
