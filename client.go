package gopipedrive

import (
	"context"
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
	Token          string
	tokenExpiresAt time.Time
	RedirectURI    string
}

func NewClient(domain, clientID, clientSecret, redirectURI string) (*Client, error) {
	if domain == "" || clientID == "" || clientSecret == "" {
		return nil, errors.New("domain, client id and client secret are required to create a Client")
	}

	return &Client{
		Client:       &http.Client{},
		ClientID:     clientID,
		ClientSecret: clientSecret,
		APIBase:      utils.BaseURL,
		RedirectURI:  redirectURI,
	}, nil
}

func (c *Client) SetToken(ctx context.Context, token string, tokenExpiresAt time.Time) error {
	if token == "" {
		return utils.ErrTokenRequired
	}

	if time.Now().After(tokenExpiresAt) {
		return utils.ErrTokenExpired
	}

	c.Token = token
	c.tokenExpiresAt = tokenExpiresAt

	return nil
}

func (c *Client) Send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	

	resp, err = c.Client.Do(req)
	if err != nil {
		fmt.Println("Do client error: ", err)
		return err
	}

	defer func(Body io.ReadCloser) error {
		return Body.Close()
	}(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errResp := &models.TokenErrorResponse{}
		data, err = io.ReadAll(resp.Body)

		if err == nil && len(data) > 0 {
			err := json.Unmarshal(data, errResp)
			if err != nil {
				fmt.Println("unmarshall error: ", err)	
				return err
			}
		}

		return errors.New(errResp.Error.Message)
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

func (c *Client) SendWithAccessToken(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + c.Token)

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
