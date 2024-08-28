package gopipedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

// GetCallLogs fetches all call logs assigned to a particular user.
func (c *Client) GetCallLogs(ctx context.Context) (*[]models.CallLog, error) {
	url := c.APIBase + utils.CallLogsEndPoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var callLogsResponse models.CallLogsResponse
	if err := c.SendWithAccessToken(req, &callLogsResponse); err != nil {
		return nil, err
	}

	return &callLogsResponse.Data, nil
}

// GetCallLog fetches details of a specific call log by ID.
func (c *Client) GetCallLog(ctx context.Context, callLogID string) (*models.CallLog, error) {
	url := c.APIBase + utils.CallLogsEndPoint + "/" + callLogID

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var callLogResponse models.CallLogResponse
	if err := c.SendWithAccessToken(req, &callLogResponse); err != nil {
		return nil, err
	}

	return &callLogResponse.Data, nil
}

// AddCallLog adds a new call log.
func (c *Client) AddCallLog(ctx context.Context, callLog models.CallLogReq) (*models.CallLog, error) {
	url := c.APIBase + utils.CallLogsEndPoint

	requestBodyBytes, err := json.Marshal(callLog)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	var callLogResponse models.CallLogResponse
	if err := c.SendWithAccessToken(req, &callLogResponse); err != nil {
		return nil, err
	}

	return &callLogResponse.Data, nil
}

// DeleteCallLog deletes a specific call log by ID.
func (c *Client) DeleteCallLog(ctx context.Context, callLogID string) error {
	url := c.APIBase + utils.CallLogsEndPoint + "/" + callLogID

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	if err := c.SendWithAccessToken(req, &result); err != nil {
		return err
	}

	return nil
}
