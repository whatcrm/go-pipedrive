package gopipedrive

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

type Client struct {
	Client         *http.Client
	ClientID       string
	ClientSecret   string
	APIBase        string
	Token          *models.TokenResponse
	tokenExpiresAt time.Time
	RedirectURI    string
}

func NewClient(domain, clientID, clientSecret, redirectURI string) (*Client, error) {
	if domain == "" || clientID == "" || clientSecret == "" {
		return nil, errors.New("domain, client id and client secret are required to create a Client")
	}

	apiBase := fmt.Sprintf(utils.DomainBaseURL, domain)

	return &Client{
		Client:       &http.Client{},
		ClientID:     clientID,
		ClientSecret: clientSecret,
		APIBase:      apiBase,
		RedirectURI:  redirectURI,
	}, nil
}

func (c *Client) Send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err = c.Client.Do(req)
	if err != nil {
		return err
	}
	
	defer func(Body io.ReadCloser) error {
		return Body.Close()
	}(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errResp := &models.ErrorResponse{}
		data, err = io.ReadAll(resp.Body)

		if err == nil && len(data) > 0 {
			err := json.Unmarshal(data, errResp)
			if err != nil {
				return err
			}
		}

		return errors.New(errResp.Error)
	}
	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		_, err := io.Copy(w, resp.Body)
		return err
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

