package gopipedrive

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "strings"

    "github.com/whatcrm/go-pipedrive/models"
    "github.com/whatcrm/go-pipedrive/utils"
)

// GetAllFilters retrieves all filters, optionally filtered by type (deals, leads, etc.).
func (c *Client) GetAllFilters(ctx context.Context, filterType string) ([]models.Filter, error) {
    url := c.APIBase + utils.FiltersEndpoint
    if filterType != "" {
        url += fmt.Sprintf("?type=%s", filterType)
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    var resp models.FiltersListResponse
    err = c.SendWithAccessToken(req, &resp)
    if err != nil {
        return nil, err
    }
    return resp.Data, nil
}

// GetFilter retrieves a specific filter by ID.
func (c *Client) GetFilter(ctx context.Context, filterID int) (models.Filter, error) {
    url := fmt.Sprintf("%s%s/%d", c.APIBase, utils.FiltersEndpoint, filterID)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return models.Filter{}, err
    }

    var resp models.FilterResponse
    err = c.SendWithAccessToken(req, &resp)
    if err != nil {
        return models.Filter{}, err
    }
    return resp.Data, nil
}

// GetAllFilterHelpers retrieves available filter helpers (conditions and operators).
func (c *Client) GetAllFilterHelpers(ctx context.Context) (map[string]string, error) {
    url := c.APIBase + utils.FilterHelpersEndpoint
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    var resp models.FilterHelpersResponse
    err = c.SendWithAccessToken(req, &resp)
    if err != nil {
        return nil, err
    }
    return resp.Data, nil
}

// CreateFilter creates a new filter with specified name, conditions, and type.
func (c *Client) CreateFilter(ctx context.Context, name string, conditions map[string]interface{}, filterType string) (models.Filter, error) {
    url := c.APIBase + utils.FiltersEndpoint

    payload := map[string]interface{}{
        "name":       name,
        "conditions": conditions,
        "type":       filterType,
    }
    requestBodyBytes, err := json.Marshal(payload)
    if err != nil {
        return models.Filter{}, err
    }

    req, err := http.NewRequest("POST", url, strings.NewReader(string(requestBodyBytes)))
    if err != nil {
        return models.Filter{}, err
    }
    req.Header.Set("Content-Type", "application/json")

    var resp models.FilterResponse
    err = c.SendWithAccessToken(req, &resp)
    if err != nil {
        return models.Filter{}, err
    }
    return resp.Data, nil
}

// UpdateFilter updates an existing filter by ID.
func (c *Client) UpdateFilter(ctx context.Context, filterID int, name string, conditions map[string]interface{}) (models.Filter, error) {
    url := fmt.Sprintf("%s%s/%d", c.APIBase, utils.FiltersEndpoint, filterID)

    payload := map[string]interface{}{
        "name":       name,
        "conditions": conditions,
    }
    requestBodyBytes, err := json.Marshal(payload)
    if err != nil {
        return models.Filter{}, err
    }

    req, err := http.NewRequest("PUT", url, strings.NewReader(string(requestBodyBytes)))
    if err != nil {
        return models.Filter{}, err
    }
    req.Header.Set("Content-Type", "application/json")

    var resp models.FilterResponse
    err = c.SendWithAccessToken(req, &resp)
    if err != nil {
        return models.Filter{}, err
    }
    return resp.Data, nil
}

// DeleteFilter deletes a specific filter by ID.
func (c *Client) DeleteFilter(ctx context.Context, filterID int) error {
    url := fmt.Sprintf("%s%s/%d", c.APIBase, utils.FiltersEndpoint, filterID)

    req, err := http.NewRequest("DELETE", url, nil)
    if err != nil {
        return err
    }

    return c.SendWithAccessToken(req, nil)
}

// DeleteFiltersBulk deletes multiple filters in bulk by passing a comma-separated list of IDs.
func (c *Client) DeleteFiltersBulk(ctx context.Context, filterIDs []int) error {
    idsStr := strings.Trim(strings.Replace(fmt.Sprint(filterIDs), " ", ",", -1), "[]")
    url := fmt.Sprintf("%s%s?ids=%s", c.APIBase, utils.FiltersEndpoint, idsStr)

    req, err := http.NewRequest("DELETE", url, nil)
    if err != nil {
        return err
    }

    return c.SendWithAccessToken(req, nil)
}
