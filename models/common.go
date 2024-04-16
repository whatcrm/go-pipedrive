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
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type TokenErrorResponse struct {
	Error   struct {
		Message string `json:"message"` // Assuming "error" is the correct field name inside the nested object
	} `json:"error"`
}
