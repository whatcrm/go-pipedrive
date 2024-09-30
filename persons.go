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

// AddPerson creates a new lead
func (c *Client) AddPerson(ctx context.Context, leadReq models.PersonsRequest) (models.Person, error) {
	url := c.APIBase + utils.PersonsEndpoint
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

// UpdatePerson updates the properties of an existing person
func (c *Client) UpdatePerson(ctx context.Context, personID int, updateReq models.PersonsRequest) (models.Person, error) {
	url := fmt.Sprintf("%s%s/%d", c.APIBase, utils.PersonsEndpoint, personID)
	requestBodyBytes, err := json.Marshal(updateReq)
	if err != nil {
		return models.Person{}, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.Person{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var personResponse models.PersonResponse
	err = c.SendWithAccessToken(req, &personResponse)
	if err != nil {
		return models.Person{}, err
	}

	return personResponse.Data, nil
}
