package gopipedrive

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

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

func NewClient(clientID, clientSecret, redirectURI string) (*Client, error) {
	if clientID == "" || clientSecret == "" {
		return nil, errors.New("client id and client secret are required to create a Client")
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
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(data))
	}

	fmt.Println(string(data))

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

	if req.Method != "DELETE" {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err = c.Client.Do(req)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) error {
		return Body.Close()
	}(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		data, err = io.ReadAll(resp.Body)
		fmt.Println(string(data))

		var payload map[string]interface{}
		if err2 := json.Unmarshal(data, &payload); err2 == nil {
			if ad, ok := payload["additional_data"].(map[string]interface{}); ok {
				if code, ok := ad["code"].(string); ok && code == "CHANNEL_ALREADY_EXISTS" {
					return errors.New("CHANNEL_ALREADY_EXISTS")
				}
			}
		}

		return errors.New(string(data))
	}

	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		_, err = io.Copy(w, resp.Body)
		return err
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

func buildQueryParamsString(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}
	var queryString strings.Builder
	queryString.WriteString("?")
	for key, value := range params {
		queryString.WriteString(fmt.Sprintf("%s=%v&", key, value))
	}
	query := queryString.String()
	return query[:len(query)-1] // remove the trailing '&'
}
