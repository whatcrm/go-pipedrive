package gopipedrive

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

func (c *Client) GetAccessToken(ctx context.Context, authorizationCode string) (*models.TokenResponse, error) {
	authHeaderValue := "Basic " + base64.StdEncoding.EncodeToString([]byte(c.ClientID + ":" + c.ClientSecret))

	data := url.Values{}
	data.Set("grant_type", utils.AccessTokenGrantTypeString)
	data.Set("code", authorizationCode)
	data.Set("redirect_uri", c.RedirectURI)

	req, err := http.NewRequest("POST", utils.TokenEndPoint, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("[GetAccessToken][NewRequest] err: ", err)
		return &models.TokenResponse{}, err
	}

	req.Header.Set("Authorization", authHeaderValue)

	response := &models.TokenResponse{}

	err = c.Send(req, response)
	if err != nil {
		fmt.Println("[GetAccessToken][Send request] err: ", err)
		
	}

	if response.AccessToken != "" {
		c.Token = response.AccessToken
		c.tokenExpiresAt = time.Now().Add(time.Duration(response.ExpiresIn) * time.Second)
	}
	
	return response, err
}

func (c *Client) RefreshToken(ctx context.Context, refreshToken string) (*models.TokenResponse, error) {
	authHeaderValue := "Basic " + base64.StdEncoding.EncodeToString([]byte(c.ClientID + ":" + c.ClientSecret))

	data := url.Values{}
	data.Set("grant_type", utils.RefreshTokenGrantTypeString)
	data.Set("refresh_token", refreshToken)
	data.Set("redirect_uri", c.RedirectURI)

	req, err := http.NewRequest("POST", utils.TokenEndPoint, strings.NewReader(data.Encode()))
	if err != nil {
		return &models.TokenResponse{}, err
	}

	req.Header.Set("Authorization", authHeaderValue)

	response := &models.TokenResponse{}

	err = c.Send(req, response)

	if response.AccessToken != "" {
		c.Token = response.AccessToken
		c.tokenExpiresAt = time.Now().Add(time.Duration(response.ExpiresIn) * time.Second)
	}
	
	return response, err
}
