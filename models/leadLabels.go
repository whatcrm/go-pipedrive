package models

type LeadLabel struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	AddTime   string `json:"add_time"`
	UpdateTime string `json:"update_time"`
}

type LeadLabelRequest struct {
	Name      string `json:"name"`
	Color     string `json:"color"`
}

type LeadLabelsResponse struct {
	Success bool        `json:"success"`
	Data    []LeadLabel `json:"data"`
}

type LeadLabelResponse struct {
	Success bool      `json:"success"`
	Data    LeadLabel `json:"data"`
}
