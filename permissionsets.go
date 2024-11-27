package gopipedrive

import (
	"context"
	"fmt"
	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
	"net/http"
)

func (c *Client) ListUserPermissions(ctx context.Context, userID int) (models.Permission, error) {
	url := c.APIBase + fmt.Sprintf(utils.UserEndPoint+utils.UserPermissionsEndPoint, userID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Permission{}, err
	}

	var permissionResponse models.PermissionResponse
	err = c.SendWithAccessToken(req, &permissionResponse)
	if err != nil {
		return models.Permission{}, err
	}

	return permissionResponse.Data, nil
}

func (c *Client) ListRoleAssignments(ctx context.Context, userID int, startDate, limit int) ([]models.RoleAssignment, error) {
	url := c.APIBase + fmt.Sprintf(utils.UserEndPoint+utils.UserRoleAssignmentsEndPoint, userID)

	query := fmt.Sprintf("start=%d&limit=%d", startDate, limit)
	url = fmt.Sprintf("%s?%s", url, query)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var roleAssignmentResponse models.RoleAssignmentResponse
	err = c.SendWithAccessToken(req, &roleAssignmentResponse)
	if err != nil {
		return nil, err
	}

	return roleAssignmentResponse.Data, nil
}

func (c *Client) ListUserRoleSettings(ctx context.Context, userID int) (models.UserRoleSettings, error) {
	url := c.APIBase + fmt.Sprintf(utils.UserEndPoint+utils.UserRoleSettingsEndPoint, userID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.UserRoleSettings{}, err
	}

	var userRoleSettingsResponse models.UserRoleSettingsResponse
	err = c.SendWithAccessToken(req, &userRoleSettingsResponse)
	if err != nil {
		return models.UserRoleSettings{}, err
	}

	return userRoleSettingsResponse.Data, nil

}
