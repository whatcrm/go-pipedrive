package models

type SuccessResponse struct {
	Success        string         `json:"success"`
	Data           any            `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type AdditionalData struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Start                 string `json:"start"`
	Limit                 any    `json:"limit"`
	MoreItemsInCollection bool   `json:"more_items_in_collection"`
}

type ErrorResponse struct {
	Success        bool           `json:"success"`
	Error          string         `json:"error"`
	ErrorInfo      string         `json:"error_info"`
	Data           interface{}    `json:"data"`
	AdditionalData interface{}    `json:"additional_data"`
}
