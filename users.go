package gopipedrive

import (
	"context"
	"net/http"
	"strconv"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

func (c *Client) GetUsers(ctx context.Context) (*[]models.User, error) {
	url := c.APIBase + utils.UserEndPoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var userResponse models.UsersResponse
	if err := c.SendWithAccessToken(req, &userResponse); err != nil {
		return nil, err
	}

	return &userResponse.Data, nil
}

func (c *Client) GetCurrentUser(ctx context.Context) (models.User, error) {
	url := c.APIBase + utils.UserEndPoint + utils.UserMeEndPoint

	var userResponse models.UserResponse

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return userResponse.Data, err
	}

	if err := c.SendWithAccessToken(req, &userResponse); err != nil {
		return userResponse.Data, err
	}

	return userResponse.Data, nil
}

func (c *Client) GetUser(ctx context.Context, userID int) (*models.User, error) {
	url := c.APIBase + "/v1/users/" + strconv.Itoa(userID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var userResponse models.UserResponse
	if err := c.SendWithAccessToken(req, &userResponse); err != nil {
		return nil, err
	}

	return &userResponse.Data, nil
}
