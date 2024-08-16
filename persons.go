package gopipedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
	"net/http"
)

// AddPerson creates a new lead
func (c *Client) AddPerson(ctx context.Context, leadReq models.PersonsRequest) (models.Person, error) {
	url := c.APIBase + utils.AddPersonEndpoint
	requestBodyBytes, err := json.Marshal(leadReq)
	if err != nil {
		return models.Person{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.Person{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var leadResponse models.PersonResponse
	err = c.SendWithAccessToken(req, &leadResponse)
	if err != nil {
		return models.Person{}, err
	}

	return leadResponse.Data, nil
}
