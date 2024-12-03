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

// GetNotes retrieves all notes based on query parameters.
func (c *Client) GetNotes(ctx context.Context, queryParams map[string]interface{}) (*models.NotesResponse, error) {
	url := c.APIBase + utils.NotesEndpoint
	url += buildQueryParams(queryParams)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.NotesResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetNote retrieves details of a specific note by ID.
func (c *Client) GetNote(ctx context.Context, noteID int) (*models.SingleNoteResponse, error) {
	url := fmt.Sprintf(c.APIBase+utils.SingleNoteEndpoint, noteID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.SingleNoteResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// AddNote creates a new note.
func (c *Client) AddNote(ctx context.Context, noteReq models.AddNoteRequest) (*models.SingleNoteResponse, error) {
	url := c.APIBase + utils.NotesEndpoint

	requestBodyBytes, err := json.Marshal(noteReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var response models.SingleNoteResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateNote updates an existing note.
func (c *Client) UpdateNote(ctx context.Context, noteID int, noteReq models.AddNoteRequest) (*models.SingleNoteResponse, error) {
	url := fmt.Sprintf(c.APIBase+utils.SingleNoteEndpoint, noteID)

	requestBodyBytes, err := json.Marshal(noteReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var response models.SingleNoteResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteNote deletes a specific note by ID.
func (c *Client) DeleteNote(ctx context.Context, noteID int) (bool, error) {
	url := fmt.Sprintf(c.APIBase+utils.SingleNoteEndpoint, noteID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return false, err
	}

	var response map[string]bool
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return false, err
	}

	return response["success"], nil
}

// GetComments retrieves all comments for a specific note.
func (c *Client) GetComments(ctx context.Context, noteID int) (*models.CommentsResponse, error) {
	url := fmt.Sprintf(c.APIBase+utils.CommentsEndpoint, noteID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var response models.CommentsResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// AddComment adds a new comment to a note.
func (c *Client) AddComment(ctx context.Context, noteID int, commentReq models.AddCommentRequest) (*models.SingleCommentResponse, error) {
	url := fmt.Sprintf(c.APIBase+utils.CommentsEndpoint, noteID)

	requestBodyBytes, err := json.Marshal(commentReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var response models.SingleCommentResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateComment updates an existing comment for a note.
func (c *Client) UpdateComment(ctx context.Context, noteID int, commentID string, commentReq models.AddCommentRequest) (*models.SingleCommentResponse, error) {
	url := fmt.Sprintf(c.APIBase+utils.SingleCommentEndpoint, noteID, commentID)

	requestBodyBytes, err := json.Marshal(commentReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var response models.SingleCommentResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteComment deletes a comment associated with a note.
func (c *Client) DeleteComment(ctx context.Context, noteID int, commentID string) (bool, error) {
	url := fmt.Sprintf(c.APIBase+utils.SingleCommentEndpoint, noteID, commentID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return false, err
	}

	var response map[string]bool
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return false, err
	}

	return response["success"], nil
}
