package models

type Deal struct {
	ID                int     `json:"id"`
	Title             string  `json:"title"`
	Value             float64 `json:"value"`
	Currency          string  `json:"currency"`
	UserID            int     `json:"user_id"`
	PersonID          int     `json:"person_id,omitempty"`
	OrganizationID    int     `json:"organization_id,omitempty"`
	ExpectedCloseDate string  `json:"expected_close_date,omitempty"`
	Status            string  `json:"status"`
	AddTime           string  `json:"add_time"`
	UpdateTime        string  `json:"update_time"`
}

type DealRequest struct {
	Title             string  `json:"title"`
	Value             float64 `json:"value,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	UserID            int     `json:"user_id,omitempty"`
	PersonID          int     `json:"person_id,omitempty"`
	OrganizationID    int     `json:"organization_id,omitempty"`
	ExpectedCloseDate string  `json:"expected_close_date,omitempty"`
	Status            string  `json:"status,omitempty"`
}

type DealUpdateRequest struct {
	Title             string  `json:"title,omitempty"`
	Value             float64 `json:"value,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	UserID            int     `json:"user_id,omitempty"`
	PersonID          int     `json:"person_id,omitempty"`
	OrganizationID    int     `json:"organization_id,omitempty"`
	ExpectedCloseDate string  `json:"expected_close_date,omitempty"`
	Status            string  `json:"status,omitempty"`
}

type DealResponse struct {
	Success bool `json:"success"`
	Data    Deal `json:"data"`
}

type DealsResponse struct {
	Success bool   `json:"success"`
	Data    []Deal `json:"data"`
}
