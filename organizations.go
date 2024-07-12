package gopipedrive

import (
	"context"
	"fmt"
	"net/http"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

func (c *Client) ListOrganizationParticipants(ctx context.Context, orgID int, queryParams map[string]string) ([]models.Person, error) {
	orgEndpoint := fmt.Sprintf(utils.OrganizationParticipantsEndpoint, orgID)
	url := c.APIBase + orgEndpoint
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var participantsResponse models.PersonsResponse
	err = c.SendWithAccessToken(req, &participantsResponse)
	if err != nil {
		return nil, err
	}

	return participantsResponse.Data, nil
}

func (c *Client) GetPersonByID(ctx context.Context, personID int) (*models.Person, error) {
	personEndpoint := fmt.Sprintf(utils.PersonEndpoint, personID)
	url := c.APIBase + personEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var personResponse models.PersonResponse
	err = c.SendWithAccessToken(req, &personResponse)
	if err != nil {
		return nil, err
	}

	return &personResponse.Data, nil
}
