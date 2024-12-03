package models

// Filter represents a filter with various conditions.
type Filter struct {
    ID           int    `json:"id"`
    Name         string `json:"name"`
    ActiveFlag   bool   `json:"active_flag"`
    Type         string `json:"type"`
    TemporaryFlag bool   `json:"temporary_flag,omitempty"`
    UserID       int    `json:"user_id"`
    AddTime      string `json:"add_time"`
    UpdateTime   string `json:"update_time"`
    VisibleTo    int    `json:"visible_to"`
    CustomViewID int    `json:"custom_view_id,omitempty"`
}

// FilterResponse represents a response for single filter data.
type FilterResponse struct {
    Success bool   `json:"success"`
    Data    Filter `json:"data"`
}

// FiltersListResponse represents the response for a list of filters.
type FiltersListResponse struct {
    Success bool     `json:"success"`
    Data    []Filter `json:"data"`
}

// FilterHelpersResponse represents the response containing filter helper options.
type FilterHelpersResponse struct {
    Success bool              `json:"success"`
    Data    map[string]string `json:"data"`
}
