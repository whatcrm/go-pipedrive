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

// GetActivityTypes retrieves all activity types.
func (c *Client) GetActivityTypes(ctx context.Context) ([]models.ActivityType, error) {
	url := c.APIBase + utils.ActivityTypesEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var typesResponse models.ActivityTypesResponse
	err = c.SendWithAccessToken(req, &typesResponse)
	if err != nil {
		return nil, err
	}

	return typesResponse.Data, nil
}

// AddActivityType creates a new activity type.
func (c *Client) AddActivityType(ctx context.Context, activityTypeReq models.ActivityTypeRequest) (models.ActivityType, error) {
	url := c.APIBase + utils.ActivityTypesEndpoint

	requestBodyBytes, err := json.Marshal(activityTypeReq)
	if err != nil {
		return models.ActivityType{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.ActivityType{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var typeResponse models.ActivityTypeResponse
	err = c.SendWithAccessToken(req, &typeResponse)
	if err != nil {
		return models.ActivityType{}, err
	}

	return typeResponse.Data, nil
}

// UpdateActivityType updates an existing activity type by ID.
func (c *Client) UpdateActivityType(ctx context.Context, typeID int, activityTypeReq models.ActivityTypeRequest) (models.ActivityType, error) {
	url := c.APIBase + fmt.Sprintf("%s/%d", utils.ActivityTypesEndpoint, typeID)

	requestBodyBytes, err := json.Marshal(activityTypeReq)
	if err != nil {
		return models.ActivityType{}, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.ActivityType{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var typeResponse models.ActivityTypeResponse
	err = c.SendWithAccessToken(req, &typeResponse)
	if err != nil {
		return models.ActivityType{}, err
	}

	return typeResponse.Data, nil
}

// DeleteActivityType deletes a specific activity type by ID.
func (c *Client) DeleteActivityType(ctx context.Context, typeID int) error {
	url := c.APIBase + fmt.Sprintf("%s/%d", utils.ActivityTypesEndpoint, typeID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

// DeleteActivityTypesBulk deletes multiple activity types by IDs.
func (c *Client) DeleteActivityTypesBulk(ctx context.Context, typeIDs []int) error {
	url := c.APIBase + utils.ActivityTypesEndpoint
	queryParams := "?"
	for _, id := range typeIDs {
		queryParams += fmt.Sprintf("ids=%d&", id)
	}
	url += queryParams[:len(queryParams)-1]

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}
