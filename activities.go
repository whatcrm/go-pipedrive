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

// GetActivities retrieves activities with optional filters.
func (c *Client) GetActivities(ctx context.Context, queryParams map[string]string) ([]models.Activity, error) {
	url := c.APIBase + utils.ActivitiesEndpoint
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

	var activitiesResponse models.ActivitiesResponse
	err = c.SendWithAccessToken(req, &activitiesResponse)
	if err != nil {
		return nil, err
	}

	return activitiesResponse.Data, nil
}

// GetActivity retrieves a specific activity by ID.
func (c *Client) GetActivity(ctx context.Context, activityID int) (models.Activity, error) {
	url := c.APIBase + fmt.Sprintf(utils.OneActivityEndpoint, activityID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Activity{}, err
	}

	var activityResponse models.ActivityResponse
	err = c.SendWithAccessToken(req, &activityResponse)
	if err != nil {
		return models.Activity{}, err
	}

	return activityResponse.Data, nil
}

// AddActivity creates a new activity.
func (c *Client) AddActivity(ctx context.Context, activityReq models.ActivityRequest) (models.Activity, error) {
	url := c.APIBase + utils.ActivitiesEndpoint

	requestBodyBytes, err := json.Marshal(activityReq)
	if err != nil {
		return models.Activity{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.Activity{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var activityResponse models.ActivityResponse
	err = c.SendWithAccessToken(req, &activityResponse)
	if err != nil {
		return models.Activity{}, err
	}

	return activityResponse.Data, nil
}

// UpdateActivity updates an existing activity by ID.
func (c *Client) UpdateActivity(ctx context.Context, activityID int, activityUpdateReq models.ActivityUpdateRequest) (models.Activity, error) {
	url := c.APIBase + fmt.Sprintf(utils.OneActivityEndpoint, activityID)

	requestBodyBytes, err := json.Marshal(activityUpdateReq)
	if err != nil {
		return models.Activity{}, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.Activity{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var activityResponse models.ActivityResponse
	err = c.SendWithAccessToken(req, &activityResponse)
	if err != nil {
		return models.Activity{}, err
	}

	return activityResponse.Data, nil
}

// DeleteActivity deletes a specific activity by ID.
func (c *Client) DeleteActivity(ctx context.Context, activityID int) error {
	url := c.APIBase + fmt.Sprintf(utils.OneActivityEndpoint, activityID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

// DeleteActivitiesBulk deletes multiple activities by IDs.
func (c *Client) DeleteActivitiesBulk(ctx context.Context, activityIDs []int) error {
	url := c.APIBase + utils.ActivitiesEndpoint
	queryParams := "?"
	for _, id := range activityIDs {
		queryParams += fmt.Sprintf("ids=%d&", id)
	}
	url += queryParams[:len(queryParams)-1]

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

// GetActivitiesBeta retrieves all activities using the BETA endpoint with cursor-based pagination.
func (c *Client) GetActivitiesBeta(ctx context.Context, queryParams map[string]string) ([]models.Activity, string, error) {
	url := c.APIBase + utils.ActivitiesCollectionEndpoint
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", err
	}

	var activitiesResponse models.ActivitiesBetaResponse
	err = c.SendWithAccessToken(req, &activitiesResponse)
	if err != nil {
		return nil, "", err
	}

	return activitiesResponse.Data, activitiesResponse.AdditionalData.NextCursor, nil
}
