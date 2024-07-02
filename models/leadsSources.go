package models

type LeadSource struct {
	Name string `json:"name"`
}

type LeadSourcesResponse struct {
	Success bool        `json:"success"`
	Data    []LeadSource `json:"data"`
}
