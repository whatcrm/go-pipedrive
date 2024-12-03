package gopipedrive

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "strings"

    "github.com/whatcrm/go-pipedrive/models"
    "github.com/whatcrm/go-pipedrive/utils"
)

// FindGoals searches for goals based on the provided query parameters.
func (c *Client) FindGoals(ctx context.Context, queryParams map[string]string) ([]models.Goal, error) {
    var queryStrings []string
    for key, value := range queryParams {
        queryStrings = append(queryStrings, fmt.Sprintf("%s=%s", key, value))
    }
    query := strings.Join(queryStrings, "&")
    url := fmt.Sprintf("%s%s?%s", c.APIBase, utils.FindGoalsEndpoint, query)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    var resp models.GoalsListResponse
    err = c.SendWithAccessToken(req, &resp)
    if err != nil {
        return nil, err
    }
    return resp.Data, nil
}

// GetGoalProgress retrieves the progress of a goal by ID and specified period.
func (c *Client) GetGoalProgress(ctx context.Context, goalID string, startDate, endDate string) (int, error) {
    url := fmt.Sprintf("%s%s", c.APIBase, fmt.Sprintf(utils.GoalResultsEndpoint, goalID))
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return 0, err
    }

    q := req.URL.Query()
    q.Add("period.start", startDate)
    q.Add("period.end", endDate)
    req.URL.RawQuery = q.Encode()

    var resp models.GoalProgressResponse
    err = c.SendWithAccessToken(req, &resp)
    if err != nil {
        return 0, err
    }
    return resp.Data.Progress, nil
}

// CreateGoal creates a new goal with the specified parameters.
func (c *Client) CreateGoal(ctx context.Context, goalReq models.GoalRequest) (models.Goal, error) {
    url := c.APIBase + utils.GoalsEndpoint

    requestBodyBytes, err := json.Marshal(goalReq)
    if err != nil {
        return models.Goal{}, err
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
    if err != nil {
        return models.Goal{}, err
    }
    req.Header.Set("Content-Type", "application/json")

    var resp models.GoalResponse
    err = c.SendWithAccessToken(req, &resp)
    if err != nil {
        return models.Goal{}, err
    }
    return resp.Data, nil
}

// UpdateGoal updates an existing goal by ID with the provided details.
func (c *Client) UpdateGoal(ctx context.Context, goalID string, goalReq models.GoalRequest) (models.Goal, error) {
    url := fmt.Sprintf("%s%s/%s", c.APIBase, utils.GoalsEndpoint, goalID)

    requestBodyBytes, err := json.Marshal(goalReq)
    if err != nil {
        return models.Goal{}, err
    }

    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBodyBytes))
    if err != nil {
        return models.Goal{}, err
    }
    req.Header.Set("Content-Type", "application/json")

    var resp models.GoalResponse
    err = c.SendWithAccessToken(req, &resp)
    if err != nil {
        return models.Goal{}, err
    }
    return resp.Data, nil
}

// DeleteGoal deletes an existing goal by ID.
func (c *Client) DeleteGoal(ctx context.Context, goalID string) error {
    url := fmt.Sprintf("%s%s/%s", c.APIBase, utils.GoalsEndpoint, goalID)

    req, err := http.NewRequest("DELETE", url, nil)
    if err != nil {
        return err
    }

    return c.SendWithAccessToken(req, nil)
}
