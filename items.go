package gopipedrive

import (
	"context"
	"fmt"
	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
	"net/http"
)

// ItemSearch is searching for an item by field's value and filters by item_types
func (c *Client) ItemSearch(ctx context.Context, queryString map[string]string) ([]models.Item, error) {
	// Required data: https://developers.pipedrive.com/docs/api/v1/ItemSearch

	url := utils.ItemSearchEndpoint
	if len(queryString) > 0 {
		url += "?"
		for key, value := range queryString {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var leadSearchResponse models.ItemSearchResponse
	err = c.SendWithAccessToken(req, &leadSearchResponse)
	if err != nil {
		return nil, err
	}

	return leadSearchResponse.Data, nil
}
