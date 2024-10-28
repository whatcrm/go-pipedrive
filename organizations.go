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

func (c *Client) ListOrganizationParticipants(ctx context.Context, orgID int, queryParams map[string]string) ([]models.Person, error) {
	orgEndpoint := fmt.Sprintf(utils.OrganizationParticipantsEndpoint, orgID)
	url := c.APIBase + orgEndpoint
	if len(queryParams) > 0 {
		url += buildQueryParamsString(queryParams)
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

// GetOrganizations fetches all organizations with optional query params.
func (c *Client) GetOrganizations(ctx context.Context, queryParams map[string]string) (*models.OrganizationsResponse, error) {
	url := c.APIBase + utils.OrganizationsEndpoint
	if len(queryParams) > 0 {
		url += buildQueryParamsString(queryParams)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.OrganizationsResponse
	err = c.SendWithAccessToken(req, &response)
	return &response, err
}

// GetOrganizationDetails fetches an organization's details by ID.
func (c *Client) GetOrganizationDetails(ctx context.Context, id int) (*models.OrganizationResponse, error) {
	url := fmt.Sprintf("%s%s/%d", c.APIBase, utils.OrganizationsEndpoint, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var response models.OrganizationResponse
	err = c.SendWithAccessToken(req, &response)
	return &response, err
}

// SearchOrganizations performs a search for organizations by term.
func (c *Client) SearchOrganizations(ctx context.Context, term string, queryParams map[string]string) (*models.OrganizationsResponse, error) {
	url := c.APIBase + utils.OrganizationSearchEndpoint

	if term != "" {
		queryParams["term"] = term
	}

	if len(queryParams) > 0 {
		url += buildQueryParamsString(queryParams)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.OrganizationsResponse
	err = c.SendWithAccessToken(req, &response)
	return &response, err
}

// AddOrganization adds a new organization.
func (c *Client) AddOrganization(ctx context.Context, organization models.Organization) (*models.OrganizationResponse, error) {
	url := c.APIBase + utils.OrganizationsEndpoint
	body, err := json.Marshal(organization)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	var response models.OrganizationResponse
	err = c.SendWithAccessToken(req, &response)
	return &response, err
}

// UpdateOrganization updates an existing organization by ID.
func (c *Client) UpdateOrganization(ctx context.Context, id int, updates map[string]interface{}) (*models.OrganizationResponse, error) {
	url := fmt.Sprintf("%s%s/%d", c.APIBase, utils.OrganizationsEndpoint, id)
	body, err := json.Marshal(updates)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	var response models.OrganizationResponse
	err = c.SendWithAccessToken(req, &response)
	return &response, err
}

// DeleteOrganization deletes an organization by ID.
func (c *Client) DeleteOrganization(ctx context.Context, id int) (*models.FollowerResponse, error) {
	url := fmt.Sprintf("%s%s/%d", c.APIBase, utils.OrganizationsEndpoint, id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	var response models.FollowerResponse
	err = c.SendWithAccessToken(req, &response)
	return &response, err
}

// ListActivities fetches activities for a specific organization.
func (c *Client) ListActivities(ctx context.Context, orgID int, queryParams map[string]string) (*models.ActivitiesResponse, error) {
	url := fmt.Sprintf(c.APIBase+utils.OrganizationActivitiesEndpoint, orgID)
	if len(queryParams) > 0 {
		url += buildQueryParamsString(queryParams)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.ActivitiesResponse
	err = c.SendWithAccessToken(req, &response)
	return &response, err
}

// ListDeals fetches deals associated with an organization.
func (c *Client) ListDeals(ctx context.Context, orgID int, queryParams map[string]string) (*models.DealsResponse, error) {
	url := fmt.Sprintf(c.APIBase+utils.OrganizationDealsEndpoint, orgID)
	if len(queryParams) > 0 {
		url += buildQueryParamsString(queryParams)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.DealsResponse
	err = c.SendWithAccessToken(req, &response)
	return &response, err
}

// AddFollower adds a follower to an organization.
func (c *Client) AddFollower(ctx context.Context, orgID, userID int) (*models.FollowerResponse, error) {
	url := fmt.Sprintf(c.APIBase+utils.OrganizationFollowersEndpoint, orgID)
	body, err := json.Marshal(map[string]int{"user_id": userID})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	var response models.FollowerResponse
	err = c.SendWithAccessToken(req, &response)
	return &response, err
}

// DeleteFollower removes a follower from an organization.
func (c *Client) DeleteFollower(ctx context.Context, orgID, followerID int) (*models.FollowerResponse, error) {
	url := fmt.Sprintf(c.APIBase+utils.OrganizationFollowersEndpoint+"/%d", orgID, followerID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	var response models.FollowerResponse
	err = c.SendWithAccessToken(req, &response)
	return &response, err
}
