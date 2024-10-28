package gopipedrive

import (
	"context"
	"fmt"
	"net/http"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

// GetCurrencies retrieves all supported currencies, optionally filtered by a search term.
func (c *Client) GetCurrencies(ctx context.Context, term string) ([]models.Currency, error) {
	url := c.APIBase + utils.CurrenciesEndpoint
	if term != "" {
		url += fmt.Sprintf("?term=%s", term)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var currenciesResponse models.CurrenciesResponse
	err = c.SendWithAccessToken(req, &currenciesResponse)
	if err != nil {
		return nil, err
	}

	return currenciesResponse.Data, nil
}
