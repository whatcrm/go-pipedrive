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

// GetAllFiles retrieves data about all files with pagination and sorting.
func (c *Client) GetAllFiles(ctx context.Context, start, limit int, sort string) ([]models.File, error) {
	queries := fmt.Sprintf("?start=%d&limit=%d&sort=%s", start, limit, sort)
	url := c.APIBase + utils.FilesEndPoint + queries

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var filesResponse models.FilesResponse
	err = c.SendWithAccessToken(req, &filesResponse)
	if err != nil {
		return nil, err
	}

	return filesResponse.Data, nil
}

// GetFile retrieves data about a specific file by its ID.
func (c *Client) GetFile(ctx context.Context, fileID int) (models.File, error) {
	fileEndpoint := fmt.Sprintf(utils.OneFileEndpoint, fileID)
	url := c.APIBase + fileEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.File{}, err
	}

	var fileResponse models.FileResponse
	err = c.SendWithAccessToken(req, &fileResponse)
	if err != nil {
		return models.File{}, err
	}

	return fileResponse.Data, nil
}

// DownloadFile initializes a file download and returns the URL for downloading the file.
func (c *Client) DownloadFile(ctx context.Context, fileID int) (string, error) {
	downloadEndpoint := fmt.Sprintf(utils.DownloadOneFileEndpoint, fileID)
	url := c.APIBase + downloadEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	var fileResponse models.FileDownloadResponse
	err = c.SendWithAccessToken(req, &fileResponse)
	if err != nil {
		return "", err
	}

	return fileResponse.Data.URL, nil
}

// AddFile uploads a file and associates it with a deal, person, organization, activity, product, or lead.
func (c *Client) AddFile(ctx context.Context, file *os.File, fileReq models.FileRequest) (models.File, error) {
	url := c.APIBase + utils.FilesEndPoint

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		return models.File{}, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return models.File{}, err
	}

	for key, val := range fileReq.ToMap() {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return models.File{}, err
	}

	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return models.File{}, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	var fileResponse models.FileResponse
	err = c.SendWithAccessToken(req, &fileResponse)
	if err != nil {
		return models.File{}, err
	}

	return fileResponse.Data, nil
}

// AddRemoteFile creates a new empty file in a remote location (e.g., Google Drive) and links it to an item.
func (c *Client) AddRemoteFile(ctx context.Context, remoteFileReq models.RemoteFileRequest) (models.File, error) {
	url := c.APIBase + utils.RemoteFileEndpoint

	requestBodyBytes, err := json.Marshal(remoteFileReq)
	if err != nil {
		return models.File{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.File{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var fileResponse models.FileResponse
	err = c.SendWithAccessToken(req, &fileResponse)
	if err != nil {
		return models.File{}, err
	}

	return fileResponse.Data, nil
}

// LinkRemoteFile links an existing remote file (e.g., Google Drive) to an item.
func (c *Client) LinkRemoteFile(ctx context.Context, remoteLinkReq models.RemoteLinkRequest) (models.File, error) {
	url := c.APIBase + utils.RemoteLinkFileEndpoint

	requestBodyBytes, err := json.Marshal(remoteLinkReq)
	if err != nil {
		return models.File{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.File{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var fileResponse models.FileResponse
	err = c.SendWithAccessToken(req, &fileResponse)
	if err != nil {
		return models.File{}, err
	}

	return fileResponse.Data, nil
}

// UpdateFile updates the properties of a file.
func (c *Client) UpdateFile(ctx context.Context, fileID int, fileUpdateReq models.FileUpdateRequest) (models.File, error) {
	fileEndpoint := fmt.Sprintf(utils.OneFileEndpoint, fileID)
	url := c.APIBase + fileEndpoint

	requestBodyBytes, err := json.Marshal(fileUpdateReq)
	if err != nil {
		return models.File{}, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.File{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var fileResponse models.FileResponse
	err = c.SendWithAccessToken(req, &fileResponse)
	if err != nil {
		return models.File{}, err
	}

	return fileResponse.Data, nil
}

// DeleteFile marks a file as deleted. After 30 days, the file will be permanently deleted.
func (c *Client) DeleteFile(ctx context.Context, fileID int) error {
	fileEndpoint := fmt.Sprintf(utils.OneFileEndpoint, fileID)
	url := c.APIBase + fileEndpoint

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}
