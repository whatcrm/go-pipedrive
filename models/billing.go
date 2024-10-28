package models

type AddOn struct {
	Code string `json:"code"`
}

type BillingAddOnsResponse struct {
	Success bool     `json:"success"`
	Data    []AddOn  `json:"data"`
}
