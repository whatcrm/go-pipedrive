package gopipedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"slices"

	"github.com/h2non/filetype"
	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

// AddPersonPicture downloads an image from a URL and uploads it as a picture to a person in Pipedrive.
func (c *Client) AddPersonPicture(ctx context.Context, personID int, requestBody models.PersonPictureRequest) (res models.PersonPictureResponse, err error) {
	if requestBody.FileName == "" && requestBody.FilePath == "" {
		return res, errors.New("filename and filepath are required")
	}

	url := fmt.Sprintf("%s%s/%d/%s", c.APIBase, utils.PersonsEndpoint, personID, "picture")

	supportedExts := []string{"image/png", "image/jpg", "image/jpeg"}

	content, fileType, err := c.downloadImage(requestBody.FilePath, supportedExts)
	if err != nil {
		return models.PersonPictureResponse{}, err
	}

	body, writer, err := c.createMultipartFormData(requestBody.FileName, fileType, content)
	if err != nil {
		return models.PersonPictureResponse{}, err
	}

	res, err = c.sendPersonPictureRequest(url, body, writer)
	if err != nil {
		return models.PersonPictureResponse{}, err
	}

	return res, nil
}

func (c *Client) downloadImage(filePath string, supportedExts []string) (content []byte, fileType string, err error) {
	imgResp, err := http.Get(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("error downloading file: %v", err)
	}
	defer imgResp.Body.Close()

	content, err = io.ReadAll(imgResp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error reading file: %v", err)
	}

	fileMatch, err := filetype.Match(content)
	if err != nil {
		return nil, "", fmt.Errorf("unable to detect image type: %v", err)
	}

	if !slices.Contains(supportedExts, fileMatch.MIME.Value) {
		return nil, "", fmt.Errorf("unsupported image type: %s", fileMatch.MIME.Value)
	}

	return content, fileMatch.MIME.Value, nil
}

func (c *Client) createMultipartFormData(fileName, fileType string, content []byte) (body *bytes.Buffer, writer *multipart.Writer, err error) {
	body = &bytes.Buffer{}
	writer = multipart.NewWriter(body)

	part, err := writer.CreatePart(
		textproto.MIMEHeader{
			"Content-Disposition": []string{fmt.Sprintf(`form-data; name="file"; filename="%s"`, fileName)},
			"Content-Type":        []string{fileType},
		},
	)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating multipart form part: %v", err)
	}

	_, err = part.Write(content)
	if err != nil {
		return nil, nil, fmt.Errorf("error writing content to form part: %v", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, nil, fmt.Errorf("error closing multipart writer: %v", err)
	}

	return body, writer, nil
}

func (c *Client) sendPersonPictureRequest(url string, body *bytes.Buffer, writer *multipart.Writer) (res models.PersonPictureResponse, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return models.PersonPictureResponse{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.PersonPictureResponse{}, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.PersonPictureResponse{}, fmt.Errorf("error reading response body: %v", err)
	}

	filePath := "./response_body.json" // Define the file path
	err = os.WriteFile(filePath, respBody, 0644)
	if err != nil {
		return models.PersonPictureResponse{}, fmt.Errorf("error saving response body to file: %v", err)
	}

	err = json.Unmarshal(respBody, &res)
	if err != nil {
		return models.PersonPictureResponse{}, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return res, nil
}
