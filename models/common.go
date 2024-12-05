package models

type AdditionalData struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Start                 any  `json:"start"`
	Limit                 any  `json:"limit"`
	MoreItemsInCollection bool `json:"more_items_in_collection"`
}
