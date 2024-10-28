package gopipedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

// GetCallLogs retrieves all call logs assigned to a particular user with pagination.
func (c *Client) GetCallLogs(ctx context.Context, start, limit int) ([]models.CallLog, error) {
	url := fmt.Sprintf("%s%s?start=%d&limit=%d", c.APIBase, utils.CallLogsEndpoint, start, limit)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var logsResponse models.CallLogsResponse
	err = c.SendWithAccessToken(req, &logsResponse)
	if err != nil {
		return nil, err
	}

	return logsResponse.Data, nil
}

// GetCallLog retrieves details of a specific call log by ID.
func (c *Client) GetCallLog(ctx context.Context, logID string) (models.CallLog, error) {
	url := fmt.Sprintf("%s%s/%s", c.APIBase, utils.CallLogsEndpoint, logID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.CallLog{}, err
	}

	var logResponse models.CallLogResponse
	err = c.SendWithAccessToken(req, &logResponse)
	if err != nil {
		return models.CallLog{}, err
	}

	return logResponse.Data, nil
}

// AddCallLog creates a new call log.
func (c *Client) AddCallLog(ctx context.Context, callLogReq models.CallLogRequest) (models.CallLog, error) {
	url := c.APIBase + utils.CallLogsEndpoint

	requestBodyBytes, err := json.Marshal(callLogReq)
	if err != nil {
		return models.CallLog{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.CallLog{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var logResponse models.CallLogResponse
	err = c.SendWithAccessToken(req, &logResponse)
	if err != nil {
		return models.CallLog{}, err
	}

	return logResponse.Data, nil
}

// AttachRecordingToCallLog attaches an audio recording to a call log.
func (c *Client) AttachRecordingToCallLog(ctx context.Context, logID string, filePath string) (bool, error) {
	url := fmt.Sprintf("%s%s", c.APIBase, fmt.Sprintf(utils.CallLogRecordingEndpoint, logID))

	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		return false, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return false, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	var response models.CallLogRecordingResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return false, err
	}

	return response.Success, nil
}

// DeleteCallLog deletes a specific call log by ID.
func (c *Client) DeleteCallLog(ctx context.Context, logID string) error {
	url := fmt.Sprintf("%s%s/%s", c.APIBase, utils.CallLogsEndpoint, logID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}
