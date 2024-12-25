package gopipedrive

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

// ItemSearchV2 performs a search for items based on a query string using API v2.
func (c *Client) ItemSearchV2(ctx context.Context, queryParams map[string]interface{}) (*models.ItemSearchResponse, error) {
	url := c.APIBase + utils.ItemSearchV2Endpoint
	url += buildQueryParams(queryParams)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.ItemSearchResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ItemSearchFieldV2 performs a field-specific search using API v2.
func (c *Client) ItemSearchFieldV2(ctx context.Context, searchParams map[string]interface{}) (*models.ItemSearchResponse, error) {
	url := c.APIBase + utils.ItemSearchFieldV2Endpoint
	url += buildQueryParams(searchParams)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.ItemSearchResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// buildQueryParams builds a query string from a map of parameters.
func buildQueryParams(params map[string]interface{}) string {
	if len(params) == 0 {
		return ""
	}
	var queryString strings.Builder
	queryString.WriteString("?")
	for key, value := range params {
		queryString.WriteString(fmt.Sprintf("%s=%v&", key, value))
	}
	query := queryString.String()
	return query[:len(query)-1] // remove the trailing '&'
}
