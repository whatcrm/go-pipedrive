package models

type Pipeline struct {
	ID                       int    `json:"id"`
	Name                     string `json:"name"`
	OrderNr                  int    `json:"order_nr"`
	IsDeleted                bool   `json:"is_deleted"`
	IsDealProbabilityEnabled bool   `json:"is_deal_probability_enabled"`
	AddTime                  string `json:"add_time"`
	UpdateTime               string `json:"update_time"`
	IsSelected               bool   `json:"is_selected"`
}

type PipelinesResponse struct {
	Success        bool       `json:"success"`
	Data           []Pipeline `json:"data"`
	AdditionalData struct {
		NextCursor string `json:"next_cursor"`
	} `json:"additional_data"`
}

type PipelineResponse struct {
	Success        bool     `json:"success"`
	Data           Pipeline `json:"data"`
	AdditionalData struct {
		NextCursor string `json:"next_cursor"`
	} `json:"additional_data"`
}

type Stage struct {
	ID               int    `json:"id"`
	OrderNr          int    `json:"order_nr"`
	Name             string `json:"name"`
	IsDeleted        bool   `json:"is_deleted"`
	DealProbability  int    `json:"deal_probability"`
	PipelineID       int    `json:"pipeline_id"`
	IsDealNotEnabled bool   `json:"is_deal_rot_enabled"`
	DaysToRotted     int    `json:"days_to_rotten"`
	AddTime          string `json:"add_time"`
	UpdateTime       string `json:"update_time"`
}

type StagesResponse struct {
	Success        bool    `json:"success"`
	Data           []Stage `json:"data"`
	AdditionalData struct {
		NextCursor string `json:"next_cursor"`
	} `json:"additional_data"`
}

type StageResponse struct {
	Success        bool  `json:"success"`
	Data           Stage `json:"data"`
	AdditionalData struct {
		NextCursor string `json:"next_cursor"`
	} `json:"additional_data"`
}
