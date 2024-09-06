package gopipedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

// AddPerson creates a new lead
func (c *Client) AddPerson(ctx context.Context, leadReq models.PersonsRequest) (models.Person, error) {
	url := c.APIBase + utils.PersonsEndpoint
	requestBodyBytes, err := json.Marshal(leadReq)
	if err != nil {
		return models.Person{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.Person{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var leadResponse models.PersonResponse
	err = c.SendWithAccessToken(req, &leadResponse)
	if err != nil {
		return models.Person{}, err
	}

	return leadResponse.Data, nil
}

// AddPersonPicture uploads a picture to a person in Pipedrive
func (c *Client) AddPersonPicture(ctx context.Context, personID int, req models.PersonPictureRequest) (models.PersonPictureResponse, error) {
	baseUrl := c.APIBase + utils.PersonsEndpoint
	url := fmt.Sprintf("%s/%d/picture", baseUrl, personID)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	resp, err := http.Get(req.FilePath)
	if err != nil {
		return models.PersonPictureResponse{}, fmt.Errorf("failed to download file: %v", err)
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	fmt.Println("Content-Type:", contentType)
	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/gif" {
		return models.PersonPictureResponse{}, fmt.Errorf("unsupported image type: %s", contentType)
	}

	part, err := writer.CreateFormFile("file", "image_from_url")
	if err != nil {
		return models.PersonPictureResponse{}, err
	}

	_, err = io.Copy(part, resp.Body) // Directly copy response body into the multipart form
	if err != nil {
		return models.PersonPictureResponse{}, fmt.Errorf("failed to write file to form: %v", err)
	}

	// Step 2: Add cropping fields if provided
	if req.CropX != 0 {
		_ = writer.WriteField("crop_x", fmt.Sprintf("%d", req.CropX))
	}
	if req.CropY != 0 {
		_ = writer.WriteField("crop_y", fmt.Sprintf("%d", req.CropY))
	}
	if req.CropWidth != 0 {
		_ = writer.WriteField("crop_width", fmt.Sprintf("%d", req.CropWidth))
	}
	if req.CropHeight != 0 {
		_ = writer.WriteField("crop_height", fmt.Sprintf("%d", req.CropHeight))
	}

	// Step 3: Close multipart writer to finalize the body
	err = writer.Close()
	if err != nil {
		return models.PersonPictureResponse{}, err
	}

	// Step 4: Send the POST request with the image data
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return models.PersonPictureResponse{}, err
	}

	// Set the appropriate headers

	request.Header.Set("Content-Type", "multipart/form-data; boundary="+writer.Boundary())
	request.Header.Set("Content-Length", fmt.Sprintf("%d", body.Len()))
	request.Host = request.URL.Host // Set Host header

	// Execute the request
	var pictureResponse models.PersonPictureResponse
	err = c.SendWithAccessTokenFile(request, &pictureResponse)
	if err != nil {
		return models.PersonPictureResponse{}, err
	}

	return pictureResponse, nil
}
