package gopipedrive

import (
	"context"
	"net/http"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

// GetNoteFields retrieves all note fields.
func (c *Client) GetNoteFields(ctx context.Context) (*models.NoteFieldsResponse, error) {
	url := c.APIBase + utils.NoteFieldsEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.NoteFieldsResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
