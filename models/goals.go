package models

// Goal represents a goal object.
type Goal struct {
	ID             string `json:"id"`
	OwnerID        int    `json:"owner_id"`
	Title          string `json:"title"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Interval       string `json:"interval"`
	Start          string `json:"start"`
	End            string `json:"end"`
	Target         int    `json:"target"`
	TrackingMetric string `json:"tracking_metric"`
	IsActive       bool   `json:"is_active"`
}

// GoalResponse represents a response for a single goal.
type GoalResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`
	Data       Goal   `json:"data"`
}

// GoalsListResponse represents a response for a list of goals.
type GoalsListResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`
	Data       []Goal `json:"data"`
}

// GoalRequest represents the request payload for creating or updating a goal.
type GoalRequest struct {
	Title           string   `json:"title"`
	Assignee        Assignee `json:"assignee"`
	Type            GoalType `json:"type"`
	ExpectedOutcome Outcome  `json:"expected_outcome"`
	Duration        Duration `json:"duration"`
	Interval        string   `json:"interval"`
}

// Assignee represents the assignee of a goal.
type Assignee struct {
	ID   string `json:"id"`
	Type string `json:"type"` // Can be "person", "company", or "team"
}

// GoalType represents the type and parameters of the goal.
type GoalType struct {
	Name   string                 `json:"name"`
	Params map[string]interface{} `json:"params"` // Can include "pipeline_id", "stage_id", "activity_type_id"
}

// Outcome represents the expected outcome of a goal.
type Outcome struct {
	Target         int    `json:"target"`
	TrackingMetric string `json:"tracking_metric"`
	CurrencyID     int    `json:"currency_id,omitempty"` // Only for "sum" type goals
}

// Duration represents the start and end date for a goal.
type Duration struct {
	Start string `json:"start"`
	End   string `json:"end,omitempty"`
}

// GoalProgressResponse represents the progress data for a goal.
type GoalProgressResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`
	Data       struct {
		Progress int `json:"progress"`
	} `json:"data"`
}
