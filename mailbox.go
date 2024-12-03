package gopipedrive

import (
	"context"
	"fmt"
	"net/http"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

// GetMailMessage retrieves details of a specific mail message.
func (c *Client) GetMailMessage(ctx context.Context, id int, includeBody int) (*models.MailMessageResponse, error) {
	url := fmt.Sprintf("%s%s/%d?include_body=%d", c.APIBase, utils.MailMessageEndpoint, id, includeBody)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.MailMessageResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetMailThreads retrieves mail threads from a specified folder.
func (c *Client) GetMailThreads(ctx context.Context, folder string, start, limit int) (*models.MailThreadResponse, error) {
	url := fmt.Sprintf("%s%s?folder=%s&start=%d&limit=%d", c.APIBase, utils.MailThreadEndpoint, folder, start, limit)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.MailThreadResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetMailThread retrieves a specific mail thread by ID.
func (c *Client) GetMailThread(ctx context.Context, id int) (*models.MailThreadResponse, error) {
	url := fmt.Sprintf("%s%s/%d", c.APIBase, utils.MailThreadEndpoint, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.MailThreadResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetMailThreadMessages retrieves all mail messages within a specified mail thread.
func (c *Client) GetMailThreadMessages(ctx context.Context, threadID int) (*models.MailMessageResponse, error) {
	url := fmt.Sprintf("%s%s", c.APIBase, fmt.Sprintf(utils.MailThreadMessagesEndpoint, threadID))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.MailMessageResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateMailThread updates properties of a mail thread.
// func (c *Client) UpdateMailThread(ctx context.Context, id int, updateData map[string]interface{}) (*models.MailThreadResponse, error) {
// 	url := fmt.Sprintf("%s%s/%d", c.APIBase, utils.MailThreadEndpoint, id)

// 	// Prepare form data
// 	for key, value := range updateData {
// 		data.Set(key, fmt.Sprintf("%v", value))
// 	}
// 	req, err := http.NewRequest("PUT", url, strings.NewReader(data.Encode()))
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	if err != nil {
// 		return nil, err
// 	}

// 	var response models.MailThreadResponse
// 	err = c.SendWithAccessToken(req, &response)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &response, nil
// }

// DeleteMailThread deletes a mail thread by marking it as deleted.
func (c *Client) DeleteMailThread(ctx context.Context, id int) (*models.MailThreadResponse, error) {
	url := fmt.Sprintf("%s%s/%d", c.APIBase, utils.MailThreadEndpoint, id)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.MailThreadResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
